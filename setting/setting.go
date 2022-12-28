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

type Unity struct {
	OrganizationId string
	ApiKey         string
}

var UnitySetting = &Unity{}

type Vungle struct {
	ApiKey string
}

var VungleSetting = &Vungle{}

type Iron struct {
	Secretkey    string
	RefreshToken string
}

var IronSetting = &Iron{}

type Adx struct {
	CompanyId string
	AuthKey   string
}

var AdxSetting = &Adx{}

type Admob struct {
	RefreshToken string
	ClientId     string
	ClientSecret string
	PublisherId  string
}

var AdmobSetting = &Admob{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("adm.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'adm.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
	mapTo("unity", UnitySetting)
	mapTo("vungle", VungleSetting)
	mapTo("iron", IronSetting)
	mapTo("adx", AdxSetting)
	mapTo("admob", AdmobSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
