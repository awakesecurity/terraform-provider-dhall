package provider

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/philandstuff/dhall-golang/v6"
)

func dataSource() *schema.Resource {
	return &schema.Resource{
		Description: "The `dhall` data source allows to provide data from a dhall-lang file",

		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"entrypoint": {
				Description: "The dhall expression to be used as entrypoint",
				Type:        schema.TypeString,
				Required:    true,
			},

			"working_dir": {
				Description: "Working directory of the dhall evaluation. If not supplied, the dhall expression will run " +
					"in the current directory.",
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"result": {
				Description: "The return of the dhall evaluation encoded as json.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	entryP := d.Get("entrypoint").(string)
	workingDir := d.Get("working_dir").(string)

	entryPoint := []byte(entryP)

	oldDir, err := os.Getwd()
	log.Println("[DEBUG] olddir: ", oldDir)
	if err != nil {
		return fmt.Errorf("failed to get current directory: %s", err)
	}
	if workingDir != "" {
		defer func() {
			err := os.Chdir(oldDir)
			_ = err
		}()

		err = os.Chdir(workingDir)
		if err != nil {
			return fmt.Errorf("failed to change current directory: %s", err)
		}
	}

	log.Println("[DEBUG] newdir: ", workingDir)

	var data interface{}
	log.Println("[DEBUG] input: ", entryP)
	err = dhall.Unmarshal(entryPoint, &data)
	if err != nil {
		return fmt.Errorf("failed to evaluate dhall: %s", err)
	}

	log.Println("[DEBUG] data: ", data)

	outJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = d.Set("result", string(outJSON))
	if err != nil {
		return fmt.Errorf("failed to export result: %s", err)
	}

	d.SetId("terraform-provider-dhall")
	return nil
}
