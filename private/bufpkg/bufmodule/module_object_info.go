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

package bufmodule

import (
	"github.com/element-of-surprise/buf/private/bufpkg/bufmodule/bufmoduleref"
	"github.com/element-of-surprise/buf/private/pkg/storage"
)

// moduleObjectInfo is used in moduleReadBucket.
type moduleObjectInfo struct {
	storage.ObjectInfo

	moduleIdentity bufmoduleref.ModuleIdentity
	commit         string
}

func newModuleObjectInfo(
	storageObjectInfo storage.ObjectInfo,
	moduleIdentity bufmoduleref.ModuleIdentity,
	commit string,
) *moduleObjectInfo {
	return &moduleObjectInfo{
		ObjectInfo:     storageObjectInfo,
		moduleIdentity: moduleIdentity,
		commit:         commit,
	}
}

func (o *moduleObjectInfo) ModuleIdentity() bufmoduleref.ModuleIdentity {
	return o.moduleIdentity
}

func (o *moduleObjectInfo) Commit() string {
	return o.commit
}
