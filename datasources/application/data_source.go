package application

import (
	"fmt"

	applicationapi "github.com/dtcookie/dynatrace/api/config/applications/web"
	"github.com/dtcookie/hcl"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/hcl2sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceRead,
		Schema: hcl2sdk.Convert(map[string]*hcl.Schema{
			"name": {
				Type:     hcl.TypeString,
				Required: true,
			},
		}),
	}
}

func DataSourceRead(d *schema.ResourceData, m interface{}) error {
	var name string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	conf := m.(*config.ProviderConfiguration)
	apiService := applicationapi.NewService(conf.DTenvURL, conf.APIToken)
	stubList, err := apiService.List()
	if err != nil {
		return err
	}
	if len(stubList.Values) > 0 {
		for _, stub := range stubList.Values {
			if name == stub.Name {
				d.SetId(stub.ID)
				return nil
			}
		}
	}

	return fmt.Errorf("no application with name '%s'", name)
}
