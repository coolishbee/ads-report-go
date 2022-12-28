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

type VungleAdsResp struct {
	GameName  string  `json:"application name"`
	GameId    string  `json:"application id"`
	Revenue   float32 `json:"revenue"`
	ViewCount int     `json:"views"`
	Ecpm      float32 `json:"ecpm"`
	Date      string  `json:"date"`
}

func GetVungleAds(ctx *context.Context) (vungleTable table.Table) {
	vungleTable = table.NewDefaultTable(table.Config{
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

	info := vungleTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "game_name", db.Varchar)
	info.AddField("Game ID", "game_id", db.Varchar)
	info.AddField("Revenue", "revenue", db.Decimal)
	info.AddField("Impression", "impression", db.Int)
	info.AddField("Ecpm", "ecpm", db.Decimal)
	info.AddField("Date", "ad_date", db.Varchar).
		FieldFilterable(types.FilterType{FormType: form.Date})

	info.AddButton("GET", icon.Send, action.Ajax("ajax_vungle",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			endTime := time.Now().UTC().AddDate(0, 0, -1)
			startTime := endTime.UTC().AddDate(0, 0, -2)

			reqURL := fmt.Sprintf("https://report.api.vungle.com/ext/pub/reports/performance?start=%s&end=%s&dimensions=application,date&aggregates=views,revenue,ecpm",
				startTime.Format("2006-01-02"),
				endTime.Format("2006-01-02"))

			client := resty.New()
			resp, err := client.R().
				SetHeader("Accept", "application/json").
				SetAuthToken(setting.VungleSetting.ApiKey).
				Get(reqURL)

			fmt.Println("  Url      :", resp.Request.URL)
			fmt.Println("  Error      :", err)
			//fmt.Println("  Body       :\n", resp)

			vungleResp := []VungleAdsResp{}
			if err := json.Unmarshal(resp.Body(), &vungleResp); err != nil {
				log.Fatalf("error: %v", err)
			}

			db_vungleAds := []models.VungleAds{}
			for _, item := range vungleResp {
				obj := models.VungleAds{
					GameName:   item.GameName,
					GameId:     item.GameId,
					Revenue:    item.Revenue,
					Impression: item.ViewCount,
					Ecpm:       item.Ecpm,
					AdDate:     item.Date,
					RegDate:    time.Now(),
				}
				db_vungleAds = append(db_vungleAds, obj)
			}
			models.AddVungleAdsData(db_vungleAds)
			return true, "success", ""
		}))

	info.SetTable("tb_ad_corp_vungle").SetTitle("Vungle Ads").SetDescription("Report")

	return
}
