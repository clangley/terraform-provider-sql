.PHONY: test
test:
	#Setup databases and verify they work correctly
	cd docker-compose && docker-compose up wait_for_postgres wait_for_mariadb
	psql --host localhost -U postgres -c "select * from accounts;"
	echo "select * from test.accounts" | mysql --host=127.0.0.1 --port=3306 --user=root --password=root test
	go build 

	#Move file to correct location
	mkdir -p test/.terraform/plugins/test.io/clangley/sql/1.0.0/linux_amd64/
	mv terraform-provider-sql test/.terraform/plugins/test.io/clangley/sql/1.0.0/linux_amd64/terraform-provider-sql_1.0 
	chmod +x test/.terraform/plugins/test.io/clangley/sql/1.0.0/linux_amd64/terraform-provider-sql_1.0 

	#Run the test
	cd test/ && terraform init && terraform apply --auto-approve
