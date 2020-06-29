package config

import (
	"inidulu/utils"
)

type dbConf struct {
	DbUser               string
	DbPass               string
	DbHost               string
	DbPort               string
	DbSchema             string
	AllowNativePasswords bool
}

type Conf struct {
	Db dbConf
}

func NewAppConfig() *Conf {
	return &Conf{dbConf{
		DbUser:               utils.GetCustomConf("DbUser", "u0qelr8zarez91tv"),
		DbPass:               utils.GetCustomConf("DbPass", "vO91or76vlFPsSZJuYbc"),
		DbHost:               utils.GetCustomConf("DbHost", "bhkgwplznepnz2yokzjl-mysql.services.clever-cloud.com"),
		DbPort:               utils.GetCustomConf("DbPort", "3306"),
		DbSchema:             utils.GetCustomConf("DbSchema", "bhkgwplznepnz2yokzjl"),
		AllowNativePasswords: true,
	}}
}
