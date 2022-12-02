package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

type Database struct {
	Driver   string
	Host     string
	Username string
	Port     string
	Password string
	Database string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("adm.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'adm.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
