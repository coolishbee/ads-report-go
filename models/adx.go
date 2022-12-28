package models

import (
	"log"
	"time"

	"gorm.io/gorm/clause"
)

type Adx struct {
	ID         int       `gorm:"primary_key" json:"id"`
	GameName   string    `json:"game_name"`
	GameId     string    `json:"game_id"`
	Revenue    float32   `json:"revenue"`
	Impression int       `json:"impression"`
	Ecpm       float32   `json:"ecpm"`
	AdDate     string    `json:"ad_date"`
	RegDate    time.Time `json:"reg_date"`
}

func AddAdxData(tb_adx []Adx) {

	err := orm.Table("tb_ad_corp_adx").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "game_id"}, {Name: "addate"}},
		DoUpdates: clause.AssignmentColumns([]string{"revenue", "impression", "ecpm"}),
	}).Create(&tb_adx).Error

	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
