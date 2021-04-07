provider "sql" {
  alias = "mysql"
  type = "mysql"
  address = "root:root@tcp(127.0.0.1:3306)/test"
}


data "sql_query" "mariadb_test" {
  provider = sql.mysql
  sql = "select * from accounts;"
}

output "mariadb_query" {
  value = jsondecode(data.sql_query.mariadb_test.data)
}