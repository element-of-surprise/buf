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

package buf

import (
	"context"
	"time"

	"github.com/element-of-surprise/buf/private/buf/bufcli"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/alpha/protoc"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/alpha/registry/token/tokencreate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/alpha/registry/token/tokendelete"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/alpha/registry/token/tokenget"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/alpha/registry/token/tokenlist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/commit/commitget"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/commit/commitlist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/organization/organizationcreate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/organization/organizationdelete"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/organization/organizationget"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/plugin/plugincreate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/plugin/plugindelete"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/plugin/plugindeprecate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/plugin/pluginlist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/plugin/pluginundeprecate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/plugin/pluginversion/pluginversionlist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/repository/repositorycreate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/repository/repositorydelete"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/repository/repositorydeprecate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/repository/repositoryget"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/repository/repositorylist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/repository/repositoryundeprecate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/tag/tagcreate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/tag/taglist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/template/templatecreate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/template/templatedelete"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/template/templatedeprecate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/template/templatelist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/template/templateundeprecate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/template/templateversion/templateversioncreate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/template/templateversion/templateversionlist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/track/trackdelete"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/beta/registry/track/tracklist"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/breaking"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/build"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/config/configinit"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/config/configlsbreakingrules"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/config/configlslintrules"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/config/configmigratev1beta1"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/export"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/generate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/lint"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/lsfiles"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/mod/modclearcache"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/mod/modopen"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/mod/modprune"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/mod/modupdate"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/push"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/registry/registrylogin"
	"github.com/element-of-surprise/buf/private/buf/cmd/buf/command/registry/registrylogout"
	"github.com/element-of-surprise/buf/private/pkg/app/appcmd"
	"github.com/element-of-surprise/buf/private/pkg/app/appflag"
)

// Main is the entrypoint to the buf CLI.
func Main(name string) {
	appcmd.Main(context.Background(), NewRootCommand(name))
}

// NewRootCommand returns a new root command.
//
// This is public for use in testing.
func NewRootCommand(name string) *appcmd.Command {
	builder := appflag.NewBuilder(
		name,
		appflag.BuilderWithTimeout(120*time.Second),
		appflag.BuilderWithTracing(),
	)
	globalFlags := bufcli.NewGlobalFlags()
	return &appcmd.Command{
		Use:                 name,
		Short:               "The Buf CLI",
		Long:                "A tool for working with Protocol Buffers and managing resources on the Buf Schema Registry (BSR).",
		Version:             bufcli.Version,
		BindPersistentFlags: appcmd.BindMultiple(builder.BindRoot, globalFlags.BindRoot),
		SubCommands: []*appcmd.Command{
			build.NewCommand("build", builder),
			export.NewCommand("export", builder),
			lint.NewCommand("lint", builder),
			breaking.NewCommand("breaking", builder),
			generate.NewCommand("generate", builder),
			lsfiles.NewCommand("ls-files", builder),
			push.NewCommand("push", builder),
			{
				Use:   "mod",
				Short: "Manage Buf modules.",
				SubCommands: []*appcmd.Command{
					modprune.NewCommand("prune", builder),
					modupdate.NewCommand("update", builder),
					modopen.NewCommand("open", builder),
					modclearcache.NewCommand("clear-cache", builder, "cc"),
				},
			},
			{
				Use:   "config",
				Short: "Manage Buf module configuration.",
				SubCommands: []*appcmd.Command{
					configinit.NewCommand("init", builder),
					configlslintrules.NewCommand("ls-lint-rules", builder),
					configlsbreakingrules.NewCommand("ls-breaking-rules", builder),
					configmigratev1beta1.NewCommand("migrate-v1beta1", builder),
				},
			},
			{
				Use:   "registry",
				Short: "Manage assets on the Buf Schema Registry.",
				SubCommands: []*appcmd.Command{
					registrylogin.NewCommand("login", builder),
					registrylogout.NewCommand("logout", builder),
				},
			},
			{
				Use:   "beta",
				Short: "Beta commands. Unstable and likely to change.",
				SubCommands: []*appcmd.Command{
					{
						Use:   "registry",
						Short: "Manage assets on the Buf Schema Registry.",
						SubCommands: []*appcmd.Command{
							{
								Use:   "organization",
								Short: "Manage organizations.",
								SubCommands: []*appcmd.Command{
									organizationcreate.NewCommand("create", builder),
									organizationget.NewCommand("get", builder),
									organizationdelete.NewCommand("delete", builder),
								},
							},
							{
								Use:   "repository",
								Short: "Manage repositories.",
								SubCommands: []*appcmd.Command{
									repositorycreate.NewCommand("create", builder),
									repositoryget.NewCommand("get", builder),
									repositorylist.NewCommand("list", builder),
									repositorydelete.NewCommand("delete", builder),
									repositorydeprecate.NewCommand("deprecate", builder),
									repositoryundeprecate.NewCommand("undeprecate", builder),
								},
							},
							{
								Use:   "track",
								Short: "Manage a repository's tracks.",
								SubCommands: []*appcmd.Command{
									tracklist.NewCommand("list", builder),
									trackdelete.NewCommand("delete", builder),
								},
							},
							{
								Use:   "tag",
								Short: "Manage a repository's tags.",
								SubCommands: []*appcmd.Command{
									tagcreate.NewCommand("create", builder),
									taglist.NewCommand("list", builder),
								},
							},
							{
								Use:   "commit",
								Short: "Manage a repository's commits.",
								SubCommands: []*appcmd.Command{
									commitget.NewCommand("get", builder),
									commitlist.NewCommand("list", builder),
								},
							},
							{
								Use:   "plugin",
								Short: "Manage Protobuf plugins.",
								SubCommands: []*appcmd.Command{
									plugincreate.NewCommand("create", builder),
									pluginlist.NewCommand("list", builder),
									plugindelete.NewCommand("delete", builder),
									plugindeprecate.NewCommand("deprecate", builder),
									pluginundeprecate.NewCommand("undeprecate", builder),
									{
										Use:   "version",
										Short: "Manage Protobuf plugin versions.",
										SubCommands: []*appcmd.Command{
											pluginversionlist.NewCommand("list", builder),
										},
									},
								},
							},
							{
								Use:   "template",
								Short: "Manage Protobuf templates on the Buf Schema Registry.",
								SubCommands: []*appcmd.Command{
									templatecreate.NewCommand("create", builder),
									templatelist.NewCommand("list", builder),
									templatedelete.NewCommand("delete", builder),
									templatedeprecate.NewCommand("deprecate", builder),
									templateundeprecate.NewCommand("undeprecate", builder),
									{
										Use:   "version",
										Short: "Manage Protobuf template versions.",
										SubCommands: []*appcmd.Command{
											templateversioncreate.NewCommand("create", builder),
											templateversionlist.NewCommand("list", builder),
										},
									},
								},
							},
						},
					},
				},
			},
			{
				Use:    "alpha",
				Short:  "Alpha commands. Unstable and recommended only for experimentation. These may be deleted.",
				Hidden: true,
				SubCommands: []*appcmd.Command{
					protoc.NewCommand("protoc", builder),
					{
						Use:   "registry",
						Short: "Manage assets on the Buf Schema Registry.",
						SubCommands: []*appcmd.Command{
							{
								Use:   "token",
								Short: "Manage user tokens.",
								SubCommands: []*appcmd.Command{
									tokencreate.NewCommand("create", builder),
									tokenget.NewCommand("get", builder),
									tokenlist.NewCommand("list", builder),
									tokendelete.NewCommand("delete", builder),
								},
							},
						},
					},
				},
			},
		},
	}
}
