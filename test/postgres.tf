provider "sql" {
  alias = "postgres"
  type = "postgres"
  address = "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"
}

data "sql_query" "postgres_test" {
  provider = sql.postgres
  sql = "select * from public.accounts;"
}

output "postgres_query" {
  value = jsondecode(data.sql_query.postgres_test.data)
}