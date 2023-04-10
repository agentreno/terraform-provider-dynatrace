package export

import "fmt"

type DataSource struct {
	ID   string
	Type string
	Name string
}

func AsDataSource(resource *Resource) string {
	if resource == nil {
		return ""
	}
	if !resource.IsReferencedAsDataSource() {
		return ""
	}
	switch resource.Type {
	case ResourceTypes.ManagementZoneV2:
		return fmt.Sprintf(`data "dynatrace_management_zone" "%s" {
			name = "%s"
		}`, resource.UniqueName, resource.Name)
	case ResourceTypes.Alerting:
		return fmt.Sprintf(`data "dynatrace_alerting_profile" "%s" {
			name = "%s"
		}`, resource.UniqueName, resource.Name)
	case ResourceTypes.RequestAttribute:
		return fmt.Sprintf(`data "dynatrace_request_attribute" "%s" {
			name = "%s"
		}`, resource.UniqueName, resource.Name)
	case ResourceTypes.WebApplication:
		return fmt.Sprintf(`data "dynatrace_application" "%s" {
			name = "%s"
		}`, resource.UniqueName, resource.Name)
	case ResourceTypes.RequestNaming:
		return fmt.Sprintf(`data "dynatrace_request_naming" "%s" {
			name = "%s"
		}`, resource.UniqueName, resource.Name)
	default:
		return ""
	}
}
