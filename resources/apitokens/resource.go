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

package apitokens

import (
	"context"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/apitokens"
	settings "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/apitokens/settings"
	cfg "github.com/dynatrace-oss/terraform-provider-dynatrace/provider/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/provider/logging"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/confighcl"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource produces terraform resource definition for API tokens
func Resource() *schema.Resource {
	return &schema.Resource{
		Schema:        new(settings.APIToken).Schema(),
		CreateContext: logging.Enable(Create),
		UpdateContext: logging.Enable(Update),
		ReadContext:   logging.Enable(Read),
		DeleteContext: logging.Enable(Delete),
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

// Create expects the configuration within the given ResourceData and sends it to the Dynatrace Server in order to create that resource
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := new(settings.APIToken)
	if err := config.UnmarshalHCL(confighcl.DecoderFrom(d, Resource())); err != nil {
		return diag.FromErr(err)
	}

	objStub, err := apitokens.Service(cfg.Credentials(m)).Create(config)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(objStub.ID)
	apiToken := objStub.Value.(settings.APIToken)
	marshalled := hcl.Properties{}
	err = (&apiToken).MarshalHCL(marshalled)
	if err != nil {
		return diag.FromErr(err)
	}

	for k, v := range marshalled {
		d.Set(k, v)
	}

	return diag.Diagnostics{}
}

// Update expects the configuration within the given ResourceData and send them to the Dynatrace Server in order to update that resource
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := new(settings.APIToken)
	if err := config.UnmarshalHCL(confighcl.DecoderFrom(d, Resource())); err != nil {
		return diag.FromErr(err)
	}

	if err := apitokens.Service(cfg.Credentials(m)).Update(d.Id(), config); err != nil {
		return diag.FromErr(err)
	}

	return Read(ctx, d, m)
}

// Read queries the Dynatrace Server for the configuration
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := new(settings.APIToken)
	err := apitokens.Service(cfg.Credentials(m)).Get(d.Id(), config)
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
func Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := apitokens.Service(cfg.Credentials(m)).Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}
