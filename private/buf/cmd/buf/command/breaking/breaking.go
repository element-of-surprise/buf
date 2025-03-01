// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package breaking

import (
	"context"
	"errors"
	"fmt"

	"github.com/element-of-surprise/buf/private/buf/bufcli"
	"github.com/element-of-surprise/buf/private/buf/buffetch"
	"github.com/element-of-surprise/buf/private/buf/bufwire"
	"github.com/element-of-surprise/buf/private/bufpkg/bufanalysis"
	"github.com/element-of-surprise/buf/private/bufpkg/bufcheck/bufbreaking"
	"github.com/element-of-surprise/buf/private/bufpkg/bufimage"
	"github.com/element-of-surprise/buf/private/pkg/app/appcmd"
	"github.com/element-of-surprise/buf/private/pkg/app/appflag"
	"github.com/element-of-surprise/buf/private/pkg/command"
	"github.com/element-of-surprise/buf/private/pkg/stringutil"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	errorFormatFlagName       = "error-format"
	excludeImportsFlagName    = "exclude-imports"
	pathsFlagName             = "path"
	limitToInputFilesFlagName = "limit-to-input-files"
	configFlagName            = "config"
	againstFlagName           = "against"
	againstConfigFlagName     = "against-config"
	excludePathsFlagName      = "exclude-path"
	disableSymlinksFlagName   = "disable-symlinks"
)

// NewCommand returns a new Command.
func NewCommand(
	name string,
	builder appflag.Builder,
) *appcmd.Command {
	flags := newFlags()
	return &appcmd.Command{
		Use:   name + " --against against-input <input>",
		Short: "Verify that the input location has no breaking changes compared to the against location.",
		Long:  bufcli.GetInputLong(`the source, module, or Image to check for breaking changes`),
		Args:  cobra.MaximumNArgs(1),
		Run: builder.NewRunFunc(
			func(ctx context.Context, container appflag.Container) error {
				return run(ctx, container, flags)
			},
			bufcli.NewErrorInterceptor(),
		),
		BindFlags: flags.Bind,
	}
}

type flags struct {
	ErrorFormat       string
	ExcludeImports    bool
	LimitToInputFiles bool
	Paths             []string
	Config            string
	Against           string
	AgainstConfig     string
	ExcludePaths      []string
	DisableSymlinks   bool
	// special
	InputHashtag string
}

func newFlags() *flags {
	return &flags{}
}

func (f *flags) Bind(flagSet *pflag.FlagSet) {
	bufcli.BindPaths(flagSet, &f.Paths, pathsFlagName)
	bufcli.BindInputHashtag(flagSet, &f.InputHashtag)
	bufcli.BindExcludePaths(flagSet, &f.ExcludePaths, excludePathsFlagName)
	bufcli.BindDisableSymlinks(flagSet, &f.DisableSymlinks, disableSymlinksFlagName)
	flagSet.StringVar(
		&f.ErrorFormat,
		errorFormatFlagName,
		"text",
		fmt.Sprintf(
			"The format for build errors or check violations printed to stdout. Must be one of %s.",
			stringutil.SliceToString(bufanalysis.AllFormatStrings),
		),
	)
	flagSet.BoolVar(
		&f.ExcludeImports,
		excludeImportsFlagName,
		false,
		"Exclude imports from breaking change detection.",
	)
	flagSet.BoolVar(
		&f.LimitToInputFiles,
		limitToInputFilesFlagName,
		false,
		fmt.Sprintf(
			`Only run breaking checks against the files in the input.
When set, the against input contains only the files in the input.
Overrides --%s.`,
			pathsFlagName,
		),
	)
	flagSet.StringVar(
		&f.Config,
		configFlagName,
		"",
		`The file or data to use for configuration.`,
	)
	flagSet.StringVar(
		&f.Against,
		againstFlagName,
		"",
		fmt.Sprintf(
			`Required. The source, module, or Image to check against. Must be one of format %s.`,
			buffetch.AllFormatsString,
		),
	)
	flagSet.StringVar(
		&f.AgainstConfig,
		againstConfigFlagName,
		"",
		`The file or data to use to configure the against source, module, or image.`,
	)
}

func run(
	ctx context.Context,
	container appflag.Container,
	flags *flags,
) error {
	if flags.Against == "" {
		return appcmd.NewInvalidArgumentErrorf("required flag %q not set", againstFlagName)
	}
	if err := bufcli.ValidateErrorFormatFlag(flags.ErrorFormat, errorFormatFlagName); err != nil {
		return err
	}
	input, err := bufcli.GetInputValue(container, flags.InputHashtag, ".")
	if err != nil {
		return err
	}
	ref, err := buffetch.NewRefParser(container.Logger(), buffetch.RefParserWithProtoFileRefAllowed()).GetRef(ctx, input)
	if err != nil {
		return err
	}
	storageosProvider := bufcli.NewStorageosProvider(flags.DisableSymlinks)
	runner := command.NewRunner()
	registryProvider, err := bufcli.NewRegistryProvider(ctx, container)
	if err != nil {
		return err
	}
	imageConfigReader, err := bufcli.NewWireImageConfigReader(
		container,
		storageosProvider,
		runner,
		registryProvider,
	)
	if err != nil {
		return err
	}
	imageConfigs, fileAnnotations, err := imageConfigReader.GetImageConfigs(
		ctx,
		container,
		ref,
		flags.Config,
		flags.Paths,        // we filter checks for files
		flags.ExcludePaths, // we exclude these paths
		false,              // files specified must exist on the main input
		false,              // we must include source info for this side of the check
	)
	if err != nil {
		return err
	}
	if len(fileAnnotations) > 0 {
		if err := bufanalysis.PrintFileAnnotations(
			container.Stdout(),
			fileAnnotations,
			flags.ErrorFormat,
		); err != nil {
			return err
		}
		return errors.New("")
	}
	// TODO: this doesn't actually work because we're using the same file paths for both sides
	// if the roots change, then we're torched
	externalPaths := flags.Paths
	if flags.LimitToInputFiles {
		externalPaths, err = getExternalPathsForImages(imageConfigs, flags.ExcludeImports)
		if err != nil {
			return err
		}
	}
	againstRef, err := buffetch.NewRefParser(container.Logger(), buffetch.RefParserWithProtoFileRefAllowed()).GetRef(ctx, flags.Against)
	if err != nil {
		return err
	}
	againstImageConfigs, fileAnnotations, err := imageConfigReader.GetImageConfigs(
		ctx,
		container,
		againstRef,
		flags.AgainstConfig,
		externalPaths,      // we filter checks for files
		flags.ExcludePaths, // we exclude these paths
		true,               // files are allowed to not exist on the against input
		true,               // no need to include source info for against
	)
	if err != nil {
		return err
	}
	if len(fileAnnotations) > 0 {
		if err := bufanalysis.PrintFileAnnotations(
			container.Stdout(),
			fileAnnotations,
			flags.ErrorFormat,
		); err != nil {
			return err
		}
		return bufcli.ErrFileAnnotation
	}
	if len(imageConfigs) != len(againstImageConfigs) {
		// If workspaces are being used as input, the number
		// of images MUST match. Otherwise the results will
		// be meaningless and yield false positives.
		//
		// And similar to the note above, if the roots change,
		// we're torched.
		return fmt.Errorf("input contained %d images, whereas against contained %d images", len(imageConfigs), len(againstImageConfigs))
	}
	var allFileAnnotations []bufanalysis.FileAnnotation
	for i, imageConfig := range imageConfigs {
		fileAnnotations, err := breakingForImage(
			ctx,
			container,
			imageConfig,
			againstImageConfigs[i],
			flags.ExcludeImports,
			flags.ErrorFormat,
		)
		if err != nil {
			return err
		}
		allFileAnnotations = append(allFileAnnotations, fileAnnotations...)
	}
	if len(allFileAnnotations) > 0 {
		if err := bufanalysis.PrintFileAnnotations(
			container.Stdout(),
			bufanalysis.DeduplicateAndSortFileAnnotations(allFileAnnotations),
			flags.ErrorFormat,
		); err != nil {
			return err
		}
		return bufcli.ErrFileAnnotation
	}
	return nil
}

func breakingForImage(
	ctx context.Context,
	container appflag.Container,
	imageConfig bufwire.ImageConfig,
	againstImageConfig bufwire.ImageConfig,
	excludeImports bool,
	errorFormat string,
) ([]bufanalysis.FileAnnotation, error) {
	image := imageConfig.Image()
	if excludeImports {
		image = bufimage.ImageWithoutImports(image)
	}
	againstImage := againstImageConfig.Image()
	if excludeImports {
		againstImage = bufimage.ImageWithoutImports(againstImage)
	}
	return bufbreaking.NewHandler(container.Logger()).Check(
		ctx,
		imageConfig.Config().Breaking,
		againstImage,
		image,
	)
}

func getExternalPathsForImages(imageConfigs []bufwire.ImageConfig, excludeImports bool) ([]string, error) {
	externalPaths := make(map[string]struct{})
	for _, imageConfig := range imageConfigs {
		image := imageConfig.Image()
		if excludeImports {
			image = bufimage.ImageWithoutImports(image)
		}
		for _, imageFile := range image.Files() {
			externalPaths[imageFile.ExternalPath()] = struct{}{}
		}
	}
	return stringutil.MapToSlice(externalPaths), nil
}
