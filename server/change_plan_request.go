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

package server

import (
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/validate"
)

type ChangePlanRequest struct {
	Zone string   `service:"-" validate:"required"`
	ID   types.ID `service:"-" validate:"required"`

	CPU                  int
	MemoryMB             int
	ServerPlanCommitment types.ECommitment
	ServerPlanGeneration types.EPlanGeneration
}

func (req *ChangePlanRequest) Validate() error {
	return validate.New().Struct(req)
}
