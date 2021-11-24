package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testDataSourceConfigBasic = `
data "dhall" "test" {
  entrypoint = "\"Hello world\""
}

output "simple" {
  value = "${data.dhall.test.result}"
}
`

func TestDataSourceDhall(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceConfigBasic,
				Check: func(s *terraform.State) error {
					_, ok := s.RootModule().Resources["data.dhall.test"]
					if !ok {
						return fmt.Errorf("missing data resource")
					}

					outputs := s.RootModule().Outputs

					if outputs["simple"] == nil {
						return fmt.Errorf("missing 'simple' output")
					}

					if outputs["simple"].Value != "\"Hello world\"" {
						return fmt.Errorf(
							"'simple' output is %q; want '\"Hello World\"'",
							outputs["simple"].Value,
						)
					}

					return nil
				},
			},
		},
	})
}
