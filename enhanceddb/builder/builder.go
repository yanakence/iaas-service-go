// Copyright 2022-2023 The sacloud/iaas-service-go Authors
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

package builder

import (
	"context"
	"errors"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Builder エンハンスドデータベースのビルダー
type Builder struct {
	ID types.ID

	Name         string
	Description  string
	Tags         types.Tags
	IconID       types.ID
	DatabaseName string
	Password     string
	SettingsHash string
	Client       iaas.EnhancedDBAPI
}

func (b *Builder) Build(ctx context.Context) (*EnhancedDB, error) {
	if b.ID.IsEmpty() {
		return b.create(ctx)
	}
	return b.update(ctx)
}

func (b *Builder) create(ctx context.Context) (*EnhancedDB, error) {
	created, err := b.Client.Create(ctx, &iaas.EnhancedDBCreateRequest{
		Name:         b.Name,
		Description:  b.Description,
		Tags:         b.Tags,
		IconID:       b.IconID,
		DatabaseName: b.DatabaseName,
	})
	if err != nil {
		return nil, err
	}

	err = b.Client.SetPassword(ctx, created.ID, &iaas.EnhancedDBSetPasswordRequest{
		Password: b.Password,
	})
	if err != nil {
		return nil, err
	}
	return Read(ctx, b.Client, created.ID)
}

func (b *Builder) update(ctx context.Context) (*EnhancedDB, error) {
	current, err := b.Client.Read(ctx, b.ID)
	if err != nil {
		return nil, err
	}
	if current.DatabaseName != b.DatabaseName {
		return nil, errors.New("DatabaseName cannot be changed")
	}

	updated, err := b.Client.Update(ctx, b.ID, &iaas.EnhancedDBUpdateRequest{
		Name:         b.Name,
		Description:  b.Description,
		Tags:         b.Tags,
		IconID:       b.IconID,
		SettingsHash: b.SettingsHash,
	})
	if err != nil {
		return nil, err
	}

	if b.Password != "" {
		err := b.Client.SetPassword(ctx, updated.ID, &iaas.EnhancedDBSetPasswordRequest{
			Password: b.Password,
		})
		if err != nil {
			return nil, err
		}
	}

	return Read(ctx, b.Client, b.ID)
}
