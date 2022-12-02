package models

import (
	"log"
	"time"
)

type UnityAds struct {
	ID         int       `gorm:"primary_key" json:"id"`
	GameName   string    `json:"game_name"`
	GameId     int       `json:"game_id"`
	Platform   string    `json:"platform"`
	Revenue    float32   `json:"revenue"`
	Impression int       `json:"impression"`
	Ecpm       float32   `json:"ecpm"`
	RegDate    time.Time `json:"regdate"`
}

func AddUnityAdsData() {

	unityAds := UnityAds{
		GameName: "",
		GameId:   1,
		Revenue:  1.2,
	}
	if err := orm.Create(&unityAds).Error; err != nil {
		log.Fatalf("error: %v", err)
	}

	// unitys := []UnityAds{{GameName: "juan"}, {GameName: "pedro"}}

	// orm.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"RegDate"}),
	// }).Create(&unitys).Error
}
