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

package protoc

import (
	"context"
	"fmt"
	"strings"

	"github.com/element-of-surprise/buf/private/bufpkg/bufimage"
	"github.com/element-of-surprise/buf/private/pkg/app"
	"github.com/element-of-surprise/buf/private/pkg/app/appproto/appprotoexec"
	"github.com/element-of-surprise/buf/private/pkg/command"
	"github.com/element-of-surprise/buf/private/pkg/storage/storageos"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/pluginpb"
)

type pluginInfo struct {
	// Required
	Out string
	// optional
	Opt []string
	// optional
	Path string
}

func newPluginInfo() *pluginInfo {
	return &pluginInfo{}
}

func executePlugin(
	ctx context.Context,
	logger *zap.Logger,
	storageosProvider storageos.Provider,
	runner command.Runner,
	container app.EnvStderrContainer,
	images []bufimage.Image,
	pluginName string,
	pluginInfo *pluginInfo,
) (*pluginpb.CodeGeneratorResponse, error) {
	response, err := appprotoexec.NewGenerator(
		logger,
		storageosProvider,
		runner,
	).Generate(
		ctx,
		container,
		pluginName,
		bufimage.ImagesToCodeGeneratorRequests(
			images,
			strings.Join(pluginInfo.Opt, ","),
			appprotoexec.DefaultVersion,
			false,
			false,
		),
		appprotoexec.GenerateWithPluginPath(pluginInfo.Path),
	)
	if err != nil {
		return nil, fmt.Errorf("--%s_out: %v", pluginName, err)
	}
	return response, nil
}
