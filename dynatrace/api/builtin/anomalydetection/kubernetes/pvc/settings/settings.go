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

package pvc

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	LowDiskSpaceCritical           *LowDiskSpaceCritical           `json:"lowDiskSpaceCritical"`           // Alerts on low disk space in megabytes for a persistent volume claim.
	LowDiskSpaceCriticalPercentage *LowDiskSpaceCriticalPercentage `json:"lowDiskSpaceCriticalPercentage"` // Alerts on low disk space in % for a persistent volume claim.
	Scope                          *string                         `json:"-" scope:"scope"`                // The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
}

func (me *Settings) Name() string {
	return *me.Scope
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"low_disk_space_critical": {
			Type:        schema.TypeList,
			Description: "Alerts on low disk space in megabytes for a persistent volume claim.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(LowDiskSpaceCritical).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"low_disk_space_critical_percentage": {
			Type:        schema.TypeList,
			Description: "Alerts on low disk space in % for a persistent volume claim.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(LowDiskSpaceCriticalPercentage).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.",
			Optional:    true,
			Default:     "environment",
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"low_disk_space_critical":            me.LowDiskSpaceCritical,
		"low_disk_space_critical_percentage": me.LowDiskSpaceCriticalPercentage,
		"scope":                              me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"low_disk_space_critical":            &me.LowDiskSpaceCritical,
		"low_disk_space_critical_percentage": &me.LowDiskSpaceCriticalPercentage,
		"scope":                              &me.Scope,
	})
}
