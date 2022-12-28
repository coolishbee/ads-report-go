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

type AdxAdsResp struct {
	Code int       `json:"code"`
	Data []AdxData `json:"data"`
}

type AdxData struct {
	Date       string  `json:"date"`
	Revenue    float32 `json:"revenue"`
	Impression int     `json:"impression"`
	GameNameEN string  `json:"app"`
	Flatform   string  `json:"osType"`
	Flag       bool
}

type AdxDB struct {
	GameName  string
	GameId    string
	Revenue   float32
	ViewCount int
	Ecpm      float32
	Date      string
}

func GetAdxAds(ctx *context.Context) (adxTable table.Table) {
	adxTable = table.NewDefaultTable(table.Config{
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

	info := adxTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "game_name", db.Varchar)
	info.AddField("Game ID", "game_id", db.Varchar)
	info.AddField("Revenue", "revenue", db.Decimal)
	info.AddField("Impression", "impression", db.Int)
	info.AddField("Ecpm", "ecpm", db.Decimal)
	info.AddField("Date", "ad_date", db.Varchar).
		FieldFilterable(types.FilterType{FormType: form.Date})

	info.AddButton("GET", icon.Send, action.Ajax("ajax_adx",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			endTime := time.Now().UTC().AddDate(0, 0, -1)
			startTime := endTime.UTC().AddDate(0, 0, -2)

			reqURL := fmt.Sprintf("https://developer.adxcorp.kr/api/v2.1/report?companyId=%s&fromDate=%s&toDate=%s",
				setting.AdxSetting.CompanyId,
				startTime.Format("2006-01-02"),
				endTime.Format("2006-01-02"))

			fmt.Println("startTime :", startTime.Format("2006-01-02"))
			fmt.Println("endTime :", endTime.Format("2006-01-02"))

			client := resty.New()
			resp, err := client.R().
				SetHeader("Authorization", setting.AdxSetting.AuthKey).
				Get(reqURL)
			fmt.Println("  Error      :", err)
			fmt.Println("  Body       :\n", resp)

			adxResp := AdxAdsResp{}
			if err := json.Unmarshal(resp.Body(), &adxResp); err != nil {
				log.Fatalf("error: %v", err)
			}

			resultCode := adxResp.Code
			adxQuery := []AdxDB{}

			if resultCode == 200 {
				for i := 0; i < len(adxResp.Data); i++ {
					tmpRevenue := adxResp.Data[i].Revenue
					tmpImpression := adxResp.Data[i].Impression

					for j := 0; j < len(adxResp.Data); j++ {

						if i != j && adxResp.Data[i].Date == adxResp.Data[j].Date &&
							adxResp.Data[i].GameNameEN == adxResp.Data[j].GameNameEN {
							tmpRevenue += adxResp.Data[j].Revenue
							tmpImpression += adxResp.Data[j].Impression
							adxResp.Data[j].Flag = true
							break
						}
					}

					if !adxResp.Data[i].Flag {
						tmpQuery := AdxDB{
							GameName:  adxResp.Data[i].GameNameEN,
							GameId:    adxResp.Data[i].GameNameEN,
							Revenue:   tmpRevenue,
							ViewCount: tmpImpression,
							Ecpm:      0,
							Date:      adxResp.Data[i].Date,
						}
						adxQuery = append(adxQuery, tmpQuery)
					}
				}

				db_adx := []models.Adx{}
				for _, item := range adxQuery {
					obj := models.Adx{
						GameName:   item.GameName,
						GameId:     item.GameId,
						Revenue:    item.Revenue,
						Impression: item.ViewCount,
						Ecpm:       item.Ecpm,
						AdDate:     item.Date,
						RegDate:    time.Now(),
					}
					db_adx = append(db_adx, obj)
				}
				models.AddAdxData(db_adx)
			}

			return true, "success", ""
		}))

	info.SetTable("tb_ad_corp_adx").SetTitle("Adx Ads").SetDescription("Report")

	return
}
