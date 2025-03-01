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

// Code generated by protoc-gen-go-api. DO NOT EDIT.

package registryv1alpha1api

import (
	context "context"
	v1alpha1 "github.com/element-of-surprise/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
)

// UserService is the User service.
type UserService interface {
	// CreateUser creates a new user with the given username.
	CreateUser(ctx context.Context, username string) (user *v1alpha1.User, err error)
	// GetUser gets a user by ID.
	GetUser(ctx context.Context, id string) (user *v1alpha1.User, err error)
	// GetUserByUsername gets a user by username.
	GetUserByUsername(ctx context.Context, username string) (user *v1alpha1.User, err error)
	// ListUsers lists all users.
	ListUsers(
		ctx context.Context,
		pageSize uint32,
		pageToken string,
		reverse bool,
		userStateFilter v1alpha1.UserState,
	) (users []*v1alpha1.User, nextPageToken string, err error)
	// ListOrganizationUsers lists all users for an organization.
	// TODO: #663 move this to organization service
	ListOrganizationUsers(
		ctx context.Context,
		organizationId string,
		pageSize uint32,
		pageToken string,
		reverse bool,
	) (users []*v1alpha1.OrganizationUser, nextPageToken string, err error)
	// DeleteUser deletes a user.
	DeleteUser(ctx context.Context) (err error)
	// Deactivate user deactivates a user.
	DeactivateUser(ctx context.Context, id string) (err error)
	// UpdateUserServerRole update the role of an user in the server.
	UpdateUserServerRole(
		ctx context.Context,
		userId string,
		serverRole v1alpha1.ServerRole,
	) (err error)
	// CountUsers returns the number of users in the server by the user state provided.
	CountUsers(ctx context.Context, userStateFilter v1alpha1.UserState) (totalCount uint32, err error)
}
