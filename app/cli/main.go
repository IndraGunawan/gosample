package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/IndraGunawan/gosample"
	"github.com/IndraGunawan/gosample/database"
)

func main() {
	createDatabaseCommand := flag.NewFlagSet("create-table", flag.ExitOnError)

	if len(os.Args) < 1 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-table":
		createDatabaseCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	databaseOpt := database.Option{
		Host:     gosample.GetEnvWithDefault("MYSQL_HOST", "127.0.0.1"),
		Port:     gosample.GetEnvWithDefault("MYSQL_PORT", "3306"),
		User:     gosample.GetEnv("MYSQL_USER"),
		Password: gosample.GetEnv("MYSQL_PASSWORD"),
		Database: gosample.GetEnv("MYSQL_DATABASE"),
		Charset:  gosample.GetEnvWithDefault("MYSQL_CHARSET", "utf8"),
	}

	mysql, _ := database.New(databaseOpt)

	if createDatabaseCommand.Parsed() {
		schemaSQL, err := ioutil.ReadFile("./database/schema.sql")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(mysql.Db.Exec(string(schemaSQL)))
	}
}
