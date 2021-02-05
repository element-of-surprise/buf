// Copyright 2020-2021 Buf Technologies, Inc.
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
	registryv1alpha1api "github.com/bufbuild/buf/internal/gen/proto/api/buf/alpha/registry/v1alpha1/registryv1alpha1api"
	registryv1alpha1apiclient "github.com/bufbuild/buf/internal/gen/proto/apiclient/buf/alpha/registry/v1alpha1/registryv1alpha1apiclient"
	v1alpha1 "github.com/bufbuild/buf/internal/gen/proto/go/buf/alpha/registry/v1alpha1"
	grpcclient "github.com/bufbuild/buf/internal/pkg/transport/grpc/grpcclient"
	zap "go.uber.org/zap"
)

// NewProvider returns a new Provider.
func NewProvider(
	logger *zap.Logger,
	clientConnProvider grpcclient.ClientConnProvider,
) registryv1alpha1apiclient.Provider {
	return &provider{
		logger:             logger,
		clientConnProvider: clientConnProvider,
	}
}

type provider struct {
	logger             *zap.Logger
	clientConnProvider grpcclient.ClientConnProvider
}

func (p *provider) NewDownloadService(ctx context.Context, address string) (registryv1alpha1api.DownloadService, error) {
	clientConn, err := p.clientConnProvider.NewClientConn(ctx, address)
	if err != nil {
		return nil, err
	}
	return &downloadService{
		logger: p.logger,
		client: v1alpha1.NewDownloadServiceClient(clientConn),
	}, nil
}

func (p *provider) NewOrganizationService(ctx context.Context, address string) (registryv1alpha1api.OrganizationService, error) {
	clientConn, err := p.clientConnProvider.NewClientConn(ctx, address)
	if err != nil {
		return nil, err
	}
	return &organizationService{
		logger: p.logger,
		client: v1alpha1.NewOrganizationServiceClient(clientConn),
	}, nil
}

func (p *provider) NewPushService(ctx context.Context, address string) (registryv1alpha1api.PushService, error) {
	clientConn, err := p.clientConnProvider.NewClientConn(ctx, address)
	if err != nil {
		return nil, err
	}
	return &pushService{
		logger: p.logger,
		client: v1alpha1.NewPushServiceClient(clientConn),
	}, nil
}

func (p *provider) NewRepositoryBranchService(ctx context.Context, address string) (registryv1alpha1api.RepositoryBranchService, error) {
	clientConn, err := p.clientConnProvider.NewClientConn(ctx, address)
	if err != nil {
		return nil, err
	}
	return &repositoryBranchService{
		logger: p.logger,
		client: v1alpha1.NewRepositoryBranchServiceClient(clientConn),
	}, nil
}

func (p *provider) NewRepositoryCommitService(ctx context.Context, address string) (registryv1alpha1api.RepositoryCommitService, error) {
	clientConn, err := p.clientConnProvider.NewClientConn(ctx, address)
	if err != nil {
		return nil, err
	}
	return &repositoryCommitService{
		logger: p.logger,
		client: v1alpha1.NewRepositoryCommitServiceClient(clientConn),
	}, nil
}

func (p *provider) NewRepositoryService(ctx context.Context, address string) (registryv1alpha1api.RepositoryService, error) {
	clientConn, err := p.clientConnProvider.NewClientConn(ctx, address)
	if err != nil {
		return nil, err
	}
	return &repositoryService{
		logger: p.logger,
		client: v1alpha1.NewRepositoryServiceClient(clientConn),
	}, nil
}

func (p *provider) NewResolveService(ctx context.Context, address string) (registryv1alpha1api.ResolveService, error) {
	clientConn, err := p.clientConnProvider.NewClientConn(ctx, address)
	if err != nil {
		return nil, err
	}
	return &resolveService{
		logger: p.logger,
		client: v1alpha1.NewResolveServiceClient(clientConn),
	}, nil
}

func (p *provider) NewUserService(ctx context.Context, address string) (registryv1alpha1api.UserService, error) {
	clientConn, err := p.clientConnProvider.NewClientConn(ctx, address)
	if err != nil {
		return nil, err
	}
	return &userService{
		logger: p.logger,
		client: v1alpha1.NewUserServiceClient(clientConn),
	}, nil
}
