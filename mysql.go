package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

import "fmt"
import "os"

func openMySQLWriter() {
	var err error
	MySQLWriter, err = sql.Open("mysql", makeMysqlDataSourceName(appConfig.DB.MySQLWriter))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = MySQLWriter.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func openMySQLReader() {
	var err error
	MySQLReader, err = sql.Open("mysql", makeMysqlDataSourceName(appConfig.DB.MySQLReader))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = MySQLReader.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func makeMysqlDataSourceName(conf confMySQL) string {
	return conf.Username + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.Database
}
