package common

import (
	"fmt"
	"io/ioutil"
)

type Environment struct {
	PostgreDatabaseUrl   	string
	OracleDatabaseUrl			string
	WriteValue	  				string
	WriteHeaders					string
	Production						string
	MainRoot							string
	MovePath							string
}

func ParseEnv(key string) string {
	data,_ := ioutil.ReadFile("../prod.json")
	// _ = godotenv.Load("../prod.txt")
	fmt.Println(data)
	return "value"
}

func GetEnvironment() *Environment {
	return &Environment{
		PostgreDatabaseUrl:            ParseEnv("POSTGRE_DATABASE_URL"),
		OracleDatabaseUrl:   					 ParseEnv("ORACLE_DATABASE_URL"),
		WriteValue:	  								 ParseEnv("WRITE_VALUES"),
		WriteHeaders:								   ParseEnv("WRITE_HEADERS"),
		Production:				             ParseEnv("PRODUCTION"),
		MainRoot: 										 ParseEnv("MAIN_ROOT"),
		MovePath: 										 ParseEnv("MOVE_PATH"),
	}
}