package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/snowflakedb/gosnowflake"
)

func main() {
	userName := os.Getenv("SNOWFLAKES_USERNAME")
	password := os.Getenv("SNOWFLAKES_PASSWORD")
	database := "testdb"
	connectionString := fmt.Sprintf("%s:%s@my_organization-my_account/%s", userName, password, database)
	
	db, err := sql.Open("snowflake", connectionString)
	//db, err := sql.Open("snowflake", "jsmith:mypassword@my_organization-my_account/mydb/testschema?warehouse=mywh")   // defining the schema and warehouse
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database successfully connected")
	defer db.Close()

}
