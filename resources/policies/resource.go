/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package policies

import (
	"context"
	"fmt"
	"strings"

	policies_service "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/cluster/v1/policies"
	policies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/cluster/v1/policies/settings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/provider/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/provider/logging"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource produces terraform resource definition for Management Zones
func Resource() *schema.Resource {
	return &schema.Resource{
		Schema:        new(policies.Policy).Schema(),
		CreateContext: logging.Enable(Create),
		UpdateContext: logging.Enable(Update),
		ReadContext:   logging.Enable(Read),
		DeleteContext: logging.Enable(Delete),
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func NewService(m any) *policies_service.PolicyServiceClient {
	conf := m.(*config.ProviderConfiguration)
	apiService := policies_service.NewPolicyService(fmt.Sprintf("%s%s", conf.ClusterAPIV2URL, "/api/cluster/v2"), conf.ClusterAPIToken)
	return apiService
}

// Create expects the configuration within the given ResourceData and sends it to the Dynatrace Server in order to create that resource
func Create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	config := new(policies.Policy)
	if err := config.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	objStub, err := NewService(m).Create(config)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(objStub.ID)
	return Read(ctx, d, m)
}

// Update expects the configuration within the given ResourceData and send them to the Dynatrace Server in order to update that resource
func Update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	config := new(policies.Policy)
	if err := config.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	if err := NewService(m).Update(d.Id(), config); err != nil {
		return diag.FromErr(err)
	}
	return Read(ctx, d, m)
}

// Read queries the Dynatrace Server for the configuration
func Read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	config := new(policies.Policy)
	err := NewService(m).Get(d.Id(), config)
	if err != nil {
		return diag.FromErr(err)
	}
	marshalled := hcl.Properties{}
	err = config.MarshalHCL(marshalled)
	if err != nil {
		return diag.FromErr(err)
	}
	for k, v := range marshalled {
		d.Set(k, v)
	}
	return diag.Diagnostics{}
}

// Delete the configuration
func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := NewService(m).Delete(d.Id()); err != nil {
		if strings.HasSuffix(err.Error(), " doesn't exist") {
			return diag.Diagnostics{}
		}
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
