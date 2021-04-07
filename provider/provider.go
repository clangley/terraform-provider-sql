package provider

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	_ "github.com/lib/pq"
	"github.com/tidwall/sjson"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SQL_TYPE", ""),
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SQL_ADDRESS", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"sql_query": dataQuery(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	t := d.Get("type").(string)
	return sql.Open(t, address)
}

//RowsToJSONArray will convert rows into proper json array as a string
func RowsToJSONArray(rows *sql.Rows) (string, error) {
	var ret string
	var err error
	ret = "[]"

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	//Scan requires pointers and []*interface does not work for rows.Scan so this is a workaround
	//Since interface can be anything, we create a pointer to another interface in another slice to pass type-check
	//https://stackoverflow.com/questions/29102725/go-sql-driver-get-interface-column-values
	colPointers := make([]interface{}, len(columns))
	cols := make([]interface{}, len(columns))
	for i := range colPointers {
		colPointers[i] = &cols[i]
	}

	counter := 0
	for rows.Next() {
		err := rows.Scan(colPointers...)
		if err != nil {
			return "", err
		}
		for i, v := range cols {
			path := fmt.Sprintf("%d.%s", counter, columns[i])
			ret, err = sjson.Set(ret, path, v)
			if err != nil {
				return "", err
			}
		}
		counter++
	}
	return ret, nil
}
