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
	"github.com/go-resty/resty/v2"
)

type UnityAdsResp struct {
	GameName   string    `json:"source_name"`
	GameId     string    `json:"source_game_id"`
	Platform   string    `json:"platform"`
	Revenue    float32   `json:"revenue_sum"`
	Impression int       `json:"view_count"`
	Ecpm       float32   `json:"ecpm"`
	RegDate    time.Time `json:"timestamp"`
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
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Revenue", "revenue", db.Decimal)
	info.AddField("View Count", "view_count", db.Int)
	info.AddField("CreatedAt", "created_at", db.Timestamp)
	info.AddField("UpdatedAt", "updated_at", db.Timestamp).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})

	info.AddButton("GET", icon.Send, action.Ajax("ajax_id",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			client := resty.New()
			resp, err := client.R().
				SetHeader("Accept", "application/json").
				Get("")

			fmt.Println("  Url      :", resp.Request.URL)
			fmt.Println("  Error      :", err)
			fmt.Println("  Body       :\n", resp)

			unityAdsResp := []UnityAdsResp{}
			if err := json.Unmarshal(resp.Body(), &unityAdsResp); err != nil {
				log.Fatalf("error: %v", err)
			}

			for _, item := range unityAdsResp {
				log.Println(item.GameName)
			}
			//models.AddUnityAdsData()

			return true, "success", ""
		}))

	info.AddButton("POST", icon.Send, action.Ajax("ajax_id",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {

			return true, "success", ""
		}))

	info.AddSelectBox("code", types.FieldOptions{
		{Value: "0", Text: "합성소녀"},
		{Value: "1", Text: "대장장이"},
		{Value: "2", Text: "별빛정원"},
		{Value: "3", Text: "마이리틀포레스트"},
		{Value: "4", Text: "아일랜드M(글로벌)"},
		{Value: "5", Text: "프린세스테일"},
		{Value: "6", Text: "무한의기사"},
		{Value: "7", Text: "다크판타지"},
		{Value: "8", Text: "소마스"},
		{Value: "9", Text: "마녀의 화원"},
	}, action.FieldFilter("code"), 200)

	info.SetTable("unity_ads").SetTitle("Unity Ads").SetDescription("Report")

	return
}
