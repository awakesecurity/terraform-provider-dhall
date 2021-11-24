package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// New creates a new instance of the dhall data provider
func New() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"dhall": dataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{},
	}
}
