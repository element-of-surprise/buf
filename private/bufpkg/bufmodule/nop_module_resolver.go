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
	"context"

	"github.com/element-of-surprise/buf/private/bufpkg/bufmodule/bufmoduleref"
	"github.com/element-of-surprise/buf/private/pkg/storage"
)

type nopModuleResolver struct{}

func newNopModuleResolver() *nopModuleResolver {
	return &nopModuleResolver{}
}

func (*nopModuleResolver) GetModulePin(_ context.Context, moduleReference bufmoduleref.ModuleReference) (bufmoduleref.ModulePin, error) {
	return nil, storage.NewErrNotExist(moduleReference.String())
}
