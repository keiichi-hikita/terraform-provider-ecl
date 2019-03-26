package ecl

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIdentityAuthScopeV3DataSource_basic(t *testing.T) {
	userName := os.Getenv("OS_USERNAME")
	projectName := os.Getenv("OS_PROJECT_NAME")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIdentityAuthScopeV3DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIdentityAuthScopeV3DataSourceID("data.ecl_identity_auth_scope_v3.token"),
					resource.TestCheckResourceAttr(
						"data.ecl_identity_auth_scope_v3.token", "user_name", userName),
					resource.TestCheckResourceAttr(
						"data.ecl_identity_auth_scope_v3.token", "project_name", projectName),
				),
			},
		},
	})
}

func testAccCheckIdentityAuthScopeV3DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find token data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Token data source ID not set")
		}

		return nil
	}
}

const testAccIdentityAuthScopeV3DataSource_basic = `
data "ecl_identity_auth_scope_v3" "token" {
	name = "my_token"
}
`