package models

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/coolishbee/ads-report-go/setting"
	"github.com/jinzhu/gorm"
)

var (
	orm *gorm.DB
	err error
)

func Init(c db.Connection) {
	orm, err = gorm.Open(setting.DatabaseSetting.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.Username,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Database))

	if err != nil {
		panic("initialize orm failed")
	}
}
