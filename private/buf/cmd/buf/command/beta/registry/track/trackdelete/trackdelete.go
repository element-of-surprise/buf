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

package trackdelete

import (
	"context"

	"github.com/element-of-surprise/buf/private/buf/bufcli"
	"github.com/element-of-surprise/buf/private/bufpkg/bufmodule/bufmoduleref"
	"github.com/element-of-surprise/buf/private/pkg/app/appcmd"
	"github.com/element-of-surprise/buf/private/pkg/app/appflag"
	"github.com/element-of-surprise/buf/private/pkg/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const forceFlagName = "force"

// NewCommand returns a new Command
func NewCommand(
	name string,
	builder appflag.Builder,
) *appcmd.Command {
	flags := newFlags()
	return &appcmd.Command{
		Use:   name + " <buf.build/owner/repository:track>",
		Short: "Delete a track",
		Args:  cobra.ExactArgs(1),
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
	Force bool
}

func newFlags() *flags {
	return &flags{}
}

func (f *flags) Bind(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(
		&f.Force,
		forceFlagName,
		false,
		"Force deletion without confirming. Use with caution.",
	)
}

func run(
	ctx context.Context,
	container appflag.Container,
	flags *flags,
) error {
	bufcli.WarnBetaCommand(ctx, container)
	moduleReference, err := bufmoduleref.ModuleReferenceForString(container.Arg(0))
	if err != nil {
		return appcmd.NewInvalidArgumentError(err.Error())
	}
	registryProvider, err := bufcli.NewRegistryProvider(ctx, container)
	if err != nil {
		return err
	}
	repositoryTrackService, err := registryProvider.NewRepositoryTrackService(ctx, moduleReference.Remote())
	if err != nil {
		return err
	}
	if !flags.Force {
		if err := bufcli.PromptUserForDelete(container, "track", moduleReference.String()); err != nil {
			return err
		}
	}
	if err := repositoryTrackService.DeleteRepositoryTrackByName(
		ctx,
		moduleReference.Owner(),
		moduleReference.Repository(),
		moduleReference.Reference(),
	); err != nil {
		if rpc.GetErrorCode(err) == rpc.ErrorCodeNotFound {
			return bufcli.NewModuleReferenceNotFoundError(moduleReference)
		}
		return err
	}
	return nil
}
