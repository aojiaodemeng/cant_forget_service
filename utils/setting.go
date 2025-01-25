package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	JwtKey     string
	Issuer     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AirtableToken string
	AirtableDBId  string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	} else {
		fmt.Println("配置文件读取成功", err)
	}
	LoadServer(file)
	//LoadData(file)
	LoadAirtable(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").String()
	Issuer = file.Section("server").Key("Issuer").String()
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root1234")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadAirtable(file *ini.File) {
	AirtableToken = file.Section("airtable").Key("AirtableToken").String()
	AirtableDBId = file.Section("airtable").Key("AirtableDBId").String()
}
