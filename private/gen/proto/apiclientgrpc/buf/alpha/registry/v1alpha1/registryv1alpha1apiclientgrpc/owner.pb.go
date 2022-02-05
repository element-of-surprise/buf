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

type ownerService struct {
	logger          *zap.Logger
	client          v1alpha1.OwnerServiceClient
	contextModifier func(context.Context) context.Context
}

// GetOwnerByName takes an owner name and returns the owner as
// either a user or organization.
func (s *ownerService) GetOwnerByName(ctx context.Context, name string) (owner *v1alpha1.Owner, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetOwnerByName(
		ctx,
		&v1alpha1.GetOwnerByNameRequest{
			Name: name,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Owner, nil
}
