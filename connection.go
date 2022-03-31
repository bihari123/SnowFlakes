package main

import (
	"fmt"
	"log"

	_ "github.com/snowflakedb/gosnowflake"
)

func main() {

	//db, err := sql.Open("snowflake", "jsmith:mypassword@my_organization-my_account/mydb/testschema?warehouse=mywh")   // defining the schema and warehouse
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database successfully connected")
	defer db.Close()

}
