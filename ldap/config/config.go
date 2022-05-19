package config

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

var (
	MyHost   string
	MyPort   string
	MyDB     string
	MyUser   string
	MyPass   string
	MyConf   string
	MyPath   = MyHost + MyPort
	MyBaseDn string
) /*  */

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg.Section("").Key("app_mode"))
	MyHost = cfg.Section("ldap").Key("host").String()
	MyPort = cfg.Section("ldap").Key("port").String()
	MyDB = cfg.Section("ldap").Key("dbname").String()
	MyUser = cfg.Section("mysql").Key("username").String()
	MyPass = cfg.Section("ldap").Key("password").String()
	MyConf = cfg.Section("mysql").Key("conf").String()
	MyPath = MyHost + ":" + MyPort
	MyBaseDn = cfg.Section("ldap").Key("basedn").String()
}
