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

type displayService struct {
	logger          *zap.Logger
	client          v1alpha1.DisplayServiceClient
	contextModifier func(context.Context) context.Context
}

// DisplayOrganizationElements returns which organization elements should be displayed to the user.
func (s *displayService) DisplayOrganizationElements(
	ctx context.Context,
	organizationId string,
) (createRepository bool, createPlugin bool, createTemplate bool, settings bool, updateSettings bool, delete bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.DisplayOrganizationElements(
		ctx,
		&v1alpha1.DisplayOrganizationElementsRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, false, false, false, false, false, err
	}
	return response.CreateRepository, response.CreatePlugin, response.CreateTemplate, response.Settings, response.UpdateSettings, response.Delete, nil
}

// DisplayRepositoryElements returns which repository elements should be displayed to the user.
func (s *displayService) DisplayRepositoryElements(
	ctx context.Context,
	repositoryId string,
) (settings bool, delete bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.DisplayRepositoryElements(
		ctx,
		&v1alpha1.DisplayRepositoryElementsRequest{
			RepositoryId: repositoryId,
		},
	)
	if err != nil {
		return false, false, err
	}
	return response.Settings, response.Delete, nil
}

// DisplayPluginElements returns which plugin elements should be displayed to the user.
func (s *displayService) DisplayPluginElements(
	ctx context.Context,
	pluginId string,
) (createVersion bool, settings bool, delete bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.DisplayPluginElements(
		ctx,
		&v1alpha1.DisplayPluginElementsRequest{
			PluginId: pluginId,
		},
	)
	if err != nil {
		return false, false, false, err
	}
	return response.CreateVersion, response.Settings, response.Delete, nil
}

// DisplayTemplateElements returns which template elements should be displayed to the user.
func (s *displayService) DisplayTemplateElements(
	ctx context.Context,
	templateId string,
) (createVersion bool, settings bool, delete bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.DisplayTemplateElements(
		ctx,
		&v1alpha1.DisplayTemplateElementsRequest{
			TemplateId: templateId,
		},
	)
	if err != nil {
		return false, false, false, err
	}
	return response.CreateVersion, response.Settings, response.Delete, nil
}

// DisplayUserElements returns which user elements should be displayed to the user.
func (s *displayService) DisplayUserElements(ctx context.Context) (delete bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.DisplayUserElements(
		ctx,
		&v1alpha1.DisplayUserElementsRequest{},
	)
	if err != nil {
		return false, err
	}
	return response.Delete, nil
}

// DisplayServerElements returns which server elements should be displayed to the user.
func (s *displayService) DisplayServerElements(ctx context.Context) (adminPanel bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.DisplayServerElements(
		ctx,
		&v1alpha1.DisplayServerElementsRequest{},
	)
	if err != nil {
		return false, err
	}
	return response.AdminPanel, nil
}

// ListManageableRepositoryRoles returns which roles should be displayed
// to the user when they are managing contributors on the repository.
func (s *displayService) ListManageableRepositoryRoles(ctx context.Context, repositoryId string) (roles []v1alpha1.RepositoryRole, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListManageableRepositoryRoles(
		ctx,
		&v1alpha1.ListManageableRepositoryRolesRequest{
			RepositoryId: repositoryId,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Roles, nil
}

// ListManageableUserRepositoryRoles returns which roles should be displayed
// to the user when they are managing a specific contributor on the repository.
func (s *displayService) ListManageableUserRepositoryRoles(
	ctx context.Context,
	repositoryId string,
	userId string,
) (roles []v1alpha1.RepositoryRole, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListManageableUserRepositoryRoles(
		ctx,
		&v1alpha1.ListManageableUserRepositoryRolesRequest{
			RepositoryId: repositoryId,
			UserId:       userId,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Roles, nil
}

// ListManageablePluginRoles returns which roles should be displayed
// to the user when they are managing contributors on the plugin.
func (s *displayService) ListManageablePluginRoles(ctx context.Context, pluginId string) (roles []v1alpha1.PluginRole, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListManageablePluginRoles(
		ctx,
		&v1alpha1.ListManageablePluginRolesRequest{
			PluginId: pluginId,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Roles, nil
}

// ListManageableUserPluginRoles returns which roles should be displayed
// to the user when they are managing a specific contributor on the plugin.
func (s *displayService) ListManageableUserPluginRoles(
	ctx context.Context,
	pluginId string,
	userId string,
) (roles []v1alpha1.PluginRole, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListManageableUserPluginRoles(
		ctx,
		&v1alpha1.ListManageableUserPluginRolesRequest{
			PluginId: pluginId,
			UserId:   userId,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Roles, nil
}

// ListManageableTemplateRoles returns which roles should be displayed
// to the user when they are managing contributors on the template.
func (s *displayService) ListManageableTemplateRoles(ctx context.Context, templateId string) (roles []v1alpha1.TemplateRole, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListManageableTemplateRoles(
		ctx,
		&v1alpha1.ListManageableTemplateRolesRequest{
			TemplateId: templateId,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Roles, nil
}

// ListManageableUserTemplateRoles returns which roles should be displayed
// to the user when they are managing a specific contributor on the template.
func (s *displayService) ListManageableUserTemplateRoles(
	ctx context.Context,
	templateId string,
	userId string,
) (roles []v1alpha1.TemplateRole, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListManageableUserTemplateRoles(
		ctx,
		&v1alpha1.ListManageableUserTemplateRolesRequest{
			TemplateId: templateId,
			UserId:     userId,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Roles, nil
}
