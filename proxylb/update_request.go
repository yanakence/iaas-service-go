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

package proxylb

import (
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-service-go/serviceutil"
	"github.com/sacloud/packages-go/validate"
)

type UpdateRequest struct {
	ID types.ID `service:"-" validate:"required"`

	Name                 *string                           `service:",omitempty" validate:"omitempty,min=1"`
	Description          *string                           `service:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags                 *types.Tags                       `service:",omitempty"`
	IconID               *types.ID                         `service:",omitempty"`
	Plan                 *types.EProxyLBPlan               `service:",omitempty" validate:"omitempty,oneof=100 500 1000 5000 10000 50000 100000 400000"`
	HealthCheck          *iaas.ProxyLBHealthCheck          `service:",omitempty"`
	SorryServer          *iaas.ProxyLBSorryServer          `service:",omitempty"`
	BindPorts            *[]*iaas.ProxyLBBindPort          `service:",omitempty"`
	Servers              *[]*iaas.ProxyLBServer            `service:",omitempty"`
	Rules                *[]*iaas.ProxyLBRule              `service:",omitempty"`
	LetsEncrypt          *iaas.ProxyLBACMESetting          `service:",omitempty"`
	StickySession        *iaas.ProxyLBStickySession        `service:",omitempty"`
	Gzip                 *iaas.ProxyLBGzip                 `service:",omitempty"`
	BackendHttpKeepAlive *iaas.ProxyLBBackendHttpKeepAlive `service:",omitempty"`
	ProxyProtocol        *iaas.ProxyLBProxyProtocol        `service:",omitempty"`
	Syslog               *iaas.ProxyLBSyslog               `service:",omitempty"`
	Timeout              *iaas.ProxyLBTimeout              `service:",omitempty"`
	SettingsHash         string
}

func (req *UpdateRequest) Validate() error {
	return validate.New().Struct(req)
}

func (req *UpdateRequest) ToRequestParameter(current *iaas.ProxyLB) (*iaas.ProxyLBUpdateRequest, error) {
	r := &iaas.ProxyLBUpdateRequest{}
	if err := serviceutil.RequestConvertTo(current, r); err != nil {
		return nil, err
	}
	if err := serviceutil.RequestConvertTo(req, r); err != nil {
		return nil, err
	}
	return r, nil
}
