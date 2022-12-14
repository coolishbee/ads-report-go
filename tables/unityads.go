package tables

import (
	"encoding/json"
	"fmt"
	"log"
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

type UnityAdsResp struct {
	GameName   string    `json:"source_name"`
	GameId     string    `json:"source_game_id"`
	Platform   string    `json:"platform"`
	Revenue    float32   `json:"revenue_sum"`
	Impression int       `json:"view_count"`
	Ecpm       float32   `json:"ecpm"`
	Date       time.Time `json:"timestamp"`
}

func GetUnityAds(ctx *context.Context) (unityadsTable table.Table) {
	unityadsTable = table.NewDefaultTable(table.Config{
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

	info := unityadsTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "game_name", db.Varchar)
	info.AddField("Game ID", "game_id", db.Varchar)
	info.AddField("Revenue", "revenue", db.Decimal)
	info.AddField("Impression", "impression", db.Int)
	info.AddField("Ecpm", "ecpm", db.Decimal)
	info.AddField("Date", "ad_date", db.Varchar).
		FieldFilterable(types.FilterType{FormType: form.Date})

	info.AddButton("GET", icon.Send, action.Ajax("ajax_id",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			endTime := time.Now().UTC()
			startTime := endTime.UTC().AddDate(0, 0, -3)

			reqURL := fmt.Sprintf("https://monetization.api.unity.com/stats/v1/operate/organizations/%s?groupBy=game&fields=revenue_sum,view_count&scale=day&start=%s&end=%s&apikey=%s",
				setting.UnitySetting.OrganizationId,
				startTime.Format("2006-01-02"),
				endTime.Format("2006-01-02"),
				setting.UnitySetting.ApiKey)

			client := resty.New()
			resp, err := client.R().
				SetHeader("Accept", "application/json").
				Get(reqURL)

			fmt.Println("  Url      :", resp.Request.URL)
			fmt.Println("  Error      :", err)
			//fmt.Println("  Body       :\n", resp)

			unityAdsResp := []UnityAdsResp{}
			if err := json.Unmarshal(resp.Body(), &unityAdsResp); err != nil {
				log.Fatalf("error: %v", err)
			}

			db_unityads := []models.UnityAds{}
			for _, item := range unityAdsResp {
				obj := models.UnityAds{
					GameName:   item.GameName,
					GameId:     item.GameId,
					Revenue:    item.Revenue,
					Impression: item.Impression,
					Ecpm:       item.Ecpm,
					AdDate:     item.Date.Format("2006-01-02"),
					RegDate:    time.Now(),
				}
				db_unityads = append(db_unityads, obj)
			}
			models.AddUnityAdsData(db_unityads)

			return true, "success", ""
		}))

	info.SetTable("tb_ad_corp_unity").SetTitle("Unity Ads").SetDescription("Report")

	return
}
