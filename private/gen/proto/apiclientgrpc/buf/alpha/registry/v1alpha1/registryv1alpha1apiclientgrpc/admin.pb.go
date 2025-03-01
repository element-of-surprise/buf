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

// Code generated by protoc-gen-go-apiclientgrpc. DO NOT EDIT.

package registryv1alpha1apiclientgrpc

import (
	context "context"
	v1alpha1 "github.com/element-of-surprise/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	zap "go.uber.org/zap"
)

type adminService struct {
	logger          *zap.Logger
	client          v1alpha1.AdminServiceClient
	contextModifier func(context.Context) context.Context
}

// ForceDeleteUser forces to delete a user. Resources and organizations that are
// solely owned by the user will also be deleted.
func (s *adminService) ForceDeleteUser(
	ctx context.Context,
	userId string,
) (user *v1alpha1.User, organizations []*v1alpha1.Organization, repositories []*v1alpha1.Repository, plugins []*v1alpha1.Plugin, templates []*v1alpha1.Template, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ForceDeleteUser(
		ctx,
		&v1alpha1.ForceDeleteUserRequest{
			UserId: userId,
		},
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	return response.User, response.Organizations, response.Repositories, response.Plugins, response.Templates, nil
}
