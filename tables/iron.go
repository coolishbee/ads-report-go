package tables

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/coolishbee/ads-report-go/models"
	"github.com/coolishbee/ads-report-go/setting"

	"github.com/go-resty/resty/v2"
)

type IronAdsResp struct {
	GameName string `json:"appName"`
	GameId   string `json:"appKey"`
	Flag     bool
	Date     string     `json:"date"`
	IronData []IronData `json:"data"`
}

type IronData struct {
	Revenue   float32 `json:"revenue"`
	ViewCount int     `json:"impressions"`
	Ecpm      float32 `json:"eCPM"`
}

type IronAdsDB struct {
	GameName  string
	GameId    string
	Revenue   float32
	ViewCount int
	Ecpm      float32
	Date      string
}

func GetIronSourceAds(ctx *context.Context) (ironTable table.Table) {
	ironTable = table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := ironTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "game_name", db.Varchar)
	info.AddField("Game ID", "game_id", db.Varchar)
	info.AddField("Revenue", "revenue", db.Decimal)
	info.AddField("Impression", "impression", db.Int)
	info.AddField("Ecpm", "ecpm", db.Decimal)
	info.AddField("Date", "ad_date", db.Varchar).
		FieldFilterable(types.FilterType{FormType: form.Date})

	info.AddButton("GET", icon.Send, action.Ajax("ajax_iron",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			client := resty.New()
			resp, err := client.R().
				SetHeader("secretkey", setting.IronSetting.Secretkey).
				SetHeader("refreshToken", setting.IronSetting.RefreshToken).
				Get("https://platform.ironsrc.com/partners/publisher/auth")

			authKey := resp.String()

			//fmt.Println("  Url      :", resp.Request.URL)
			fmt.Println("  Error      :", err)
			authKey = strings.ReplaceAll(authKey, "\"", "")
			//fmt.Println("  Key       :\n", authKey)

			endTime := time.Now().UTC().AddDate(0, 0, -1)
			//fmt.Println(" endTime :", endTime.Format("2006-01-02"))
			startTime := endTime.UTC().AddDate(0, 0, -2)
			//fmt.Println(" startTime :", startTime.Format("2006-01-02"))

			reqURL := fmt.Sprintf("https://platform.ironsrc.com/partners/publisher/mediation/applications/v6/stats?startDate=%s&endDate=%s&adSource=ironSource&metrics=revenue,impressions,eCPM&breakdowns=app",
				startTime.Format("2006-01-02"),
				endTime.Format("2006-01-02"))

			resp, err = client.R().
				SetHeader("Accept", "application/json").
				SetAuthToken(authKey).
				Get(reqURL)

			fmt.Println("  Url      :", resp.Request.URL)
			//fmt.Println("  Header      :", resp.Request.Header)
			fmt.Println("  Error      :", err)
			//fmt.Println("  Body       :\n", resp)

			ironResp := []IronAdsResp{}
			ironDest := []IronAdsDB{}
			if err := json.Unmarshal(resp.Body(), &ironResp); err != nil {
				log.Fatalf("error: %v", err)
			}

			count := 0
		TEST:
			for i := range ironResp {
				ironResp[i].Flag = false

				for j := 0; j < len(ironResp); j++ {
					if i != j {
						if ironResp[i].Date == ironResp[j].Date &&
							ironResp[i].GameId == ironResp[j].GameId {

							if !ironResp[i].Flag && !ironResp[j].Flag {
								count++
								ironResp[i].Flag = true
								continue TEST
							} else {
								ironAdsQuery := IronAdsDB{
									GameName:  ironResp[i].GameName,
									GameId:    ironResp[i].GameId,
									Revenue:   ironResp[i].IronData[0].Revenue + ironResp[j].IronData[0].Revenue,
									ViewCount: ironResp[i].IronData[0].ViewCount + ironResp[j].IronData[0].ViewCount,
									Ecpm:      ironResp[i].IronData[0].Ecpm + ironResp[j].IronData[0].Ecpm,
									Date:      ironResp[i].Date,
								}
								ironDest = append(ironDest, ironAdsQuery)
								continue TEST
							}
						}
					}
				}
				ironAdsDB := IronAdsDB{
					GameName:  ironResp[i].GameName,
					GameId:    ironResp[i].GameId,
					Revenue:   ironResp[i].IronData[0].Revenue,
					ViewCount: ironResp[i].IronData[0].ViewCount,
					Ecpm:      ironResp[i].IronData[0].Ecpm,
					Date:      ironResp[i].Date,
				}
				ironDest = append(ironDest, ironAdsDB)
			}

			db_ironAds := []models.IronSource{}
			for _, item := range ironDest {
				obj := models.IronSource{
					GameName:   item.GameName,
					GameId:     item.GameId,
					Revenue:    item.Revenue,
					Impression: item.ViewCount,
					Ecpm:       item.Ecpm,
					AdDate:     item.Date,
					RegDate:    time.Now(),
				}
				db_ironAds = append(db_ironAds, obj)
			}
			models.AddIronSourceData(db_ironAds)
			return true, "success", ""
		}))

	info.SetTable("tb_ad_corp_iron").SetTitle("IronSource Ads").SetDescription("Report")

	return
}
