package tables

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
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
	"github.com/shopspring/decimal"
)

type AdmobReq struct {
	ReportSpec struct {
		DateRange struct {
			StartDate struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"startDate"`
			EndDate struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"endDate"`
		} `json:"dateRange"`
		Metrics              []string `json:"metrics"`
		LocalizationSettings struct {
			CurrencyCode string `json:"currencyCode"`
			LanguageCode string `json:"languageCode"`
		} `json:"localizationSettings"`
		Dimensions []string `json:"dimensions"`
	} `json:"reportSpec"`
}

type Authorization struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type AdmobResp struct {
	Row struct {
		DimensionValues struct {
			Date struct {
				Value string `json:"value"`
			} `json:"DATE"`
			App struct {
				GameId   string `json:"value"`
				GameName string `json:"displayLabel"`
			} `json:"APP"`
		} `json:"dimensionValues"`
		MetricValues struct {
			Revenue struct {
				MicrosValue string `json:"microsValue"`
			} `json:"ESTIMATED_EARNINGS"`
			ViewCount struct {
				IntegerValue string `json:"integerValue"`
			} `json:"IMPRESSIONS"`
			Ecpm struct {
				DoubleValue float64 `json:"doubleValue"`
			} `json:"IMPRESSION_RPM"`
		} `json:"metricValues"`
	} `json:"row"`
}

func GetGoogleAdmob(ctx *context.Context) (admobTable table.Table) {
	admobTable = table.NewDefaultTable(table.Config{
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

	info := admobTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "game_name", db.Varchar)
	info.AddField("Game ID", "game_id", db.Varchar)
	info.AddField("Revenue", "revenue", db.Decimal)
	info.AddField("Impression", "impression", db.Int)
	info.AddField("Ecpm", "ecpm", db.Decimal)
	info.AddField("Date", "ad_date", db.Varchar).
		FieldFilterable(types.FilterType{FormType: form.Date})

	info.AddButton("GET", icon.Send, action.Ajax("ajax_google",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			client := resty.New()
			resp, err := client.R().
				SetHeader("Content-Type", "application/x-www-form-urlencoded").
				SetFormData(map[string]string{
					"refresh_token": setting.AdmobSetting.RefreshToken,
					"client_id":     setting.AdmobSetting.ClientId,
					"client_secret": setting.AdmobSetting.ClientSecret,
					"grant_type":    "refresh_token",
				}).
				Post("https://accounts.google.com/o/oauth2/token")
			if err != nil {
				log.Println(err.Error())
			}

			authResp := Authorization{}
			if err := json.Unmarshal(resp.Body(), &authResp); err != nil {
				log.Fatalf("error: %v", err)
			}

			endTime := time.Now().UTC().AddDate(0, 0, -1)
			startTime := endTime.UTC().AddDate(0, 0, -2)
			startDateSlice := strings.Split(startTime.Format("2006-01-02"), "-")
			startYear, _ := strconv.Atoi(startDateSlice[0])
			startMonth, _ := strconv.Atoi(startDateSlice[1])
			startDay, _ := strconv.Atoi(startDateSlice[2])

			endDateSlice := strings.Split(endTime.Format("2006-01-02"), "-")
			endYear, _ := strconv.Atoi(endDateSlice[0])
			endMonth, _ := strconv.Atoi(endDateSlice[1])
			endDay, _ := strconv.Atoi(endDateSlice[2])

			var admobReq AdmobReq
			admobReq.ReportSpec.DateRange.StartDate.Year = startYear
			admobReq.ReportSpec.DateRange.StartDate.Month = startMonth
			admobReq.ReportSpec.DateRange.StartDate.Day = startDay
			admobReq.ReportSpec.DateRange.EndDate.Year = endYear
			admobReq.ReportSpec.DateRange.EndDate.Month = endMonth
			admobReq.ReportSpec.DateRange.EndDate.Day = endDay
			admobReq.ReportSpec.Metrics = []string{"ESTIMATED_EARNINGS", "IMPRESSIONS", "IMPRESSION_RPM"}
			admobReq.ReportSpec.LocalizationSettings.CurrencyCode = "USD"
			admobReq.ReportSpec.LocalizationSettings.LanguageCode = "en-US"
			admobReq.ReportSpec.Dimensions = []string{"DATE", "APP"}

			reportResp, err := client.R().
				SetAuthToken(authResp.AccessToken).
				SetHeader("ContentType", "application/json").
				SetBody(admobReq).
				Post(fmt.Sprintf("https://admob.googleapis.com/v1/accounts/pub-%s/networkReport:generate",
					setting.AdmobSetting.PublisherId))

			fmt.Println("  Url      :", reportResp.Request.Body)
			fmt.Println("  Error      :", err)
			//fmt.Println("  Body       :\n", reportResp)

			admobResp := []AdmobResp{}
			if err := json.Unmarshal(reportResp.Body(), &admobResp); err != nil {
				log.Fatalf("error: %v", err)
			}
			admobResp = admobResp[1:]                //앞부분 제거
			admobResp = admobResp[:len(admobResp)-1] //뒷부분 제거

			db_admob := []models.Admob{}

			for _, item := range admobResp {

				dateParse, _ := time.Parse("20060102", item.Row.DimensionValues.Date.Value)
				viewCountConv, _ := strconv.Atoi(item.Row.MetricValues.ViewCount.IntegerValue)
				revenueConv, _ := strconv.ParseInt(item.Row.MetricValues.Revenue.MicrosValue, 10, 64)
				revenue := decimal.New(revenueConv, -6).Round(2)
				revenueFloat, _ := strconv.ParseFloat(revenue.String(), 32)

				obj := models.Admob{
					GameName:   item.Row.DimensionValues.App.GameName,
					GameId:     item.Row.DimensionValues.App.GameId,
					Revenue:    float32(revenueFloat),
					Impression: viewCountConv,
					Ecpm:       float32(item.Row.MetricValues.Ecpm.DoubleValue),
					AdDate:     dateParse.Format("2006-01-02"),
					RegDate:    time.Now(),
				}
				db_admob = append(db_admob, obj)
			}
			models.AddAdmobData(db_admob)
			return true, "success", ""
		}))

	info.SetTable("tb_ad_corp_admob").SetTitle("Google Admob").SetDescription("Report")

	return
}
