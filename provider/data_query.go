package provider

import (
	"database/sql"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataQuery() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sql": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Query string",
				ForceNew:    true,
			},
			"data": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Read: dataReadQuery,
	}
}

func dataReadQuery(d *schema.ResourceData, m interface{}) error {
	db := m.(*sql.DB)

	query := d.Get("sql").(string)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	data, err := RowsToJSONArray(rows)
	if err != nil {
		return err
	}

	d.SetId(query)
	d.Set("sql", query)
	d.Set("data", data)
	return nil
}
