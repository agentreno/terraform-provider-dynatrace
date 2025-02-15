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

package azure

import (
	"fmt"
	"sync"

	azure "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/azure/settings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/rest"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/settings"
)

const SchemaID = "v1:config:credentials:azure"
const BasePath = "/api/config/v1/azure/credentials"

var mu sync.Mutex

func Service(credentials *settings.Credentials) settings.CRUDService[*azure.AzureCredentials] {
	return settings.NewCRUDService(
		credentials,
		SchemaID,
		settings.DefaultServiceOptions[*azure.AzureCredentials](BasePath).
			WithMutex(mu.Lock, mu.Unlock).
			WithOnBeforeUpdate(func(id string, v *azure.AzureCredentials) error {
				if v.SupportingServicesManagedInDynatrace {
					var creds azure.AzureCredentials
					client := rest.DefaultClient(credentials.URL, credentials.Token)

					if err := client.Get(fmt.Sprintf("%s/%s", BasePath, "id"), 200).Finish(&creds); err != nil {
						return err
					}
					v.SupportingServices = creds.SupportingServices
				}
				return nil
			}),
	)
}
