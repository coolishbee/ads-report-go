package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

type FacebookAdsResp struct {
	GameName  string  `json:"application name"`
	GameId    string  `json:"application id"`
	Revenue   float32 `json:"revenue"`
	ViewCount int     `json:"views"`
	Ecpm      float32 `json:"ecpm"`
	Date      string  `json:"date"`
}

func GetFacebookAds(ctx *context.Context) (fbTable table.Table) {
	fbTable = table.NewDefaultTable(table.Config{
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

	info := fbTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "game_name", db.Varchar)
	info.AddField("Game ID", "game_id", db.Varchar)
	info.AddField("Revenue", "revenue", db.Decimal)
	info.AddField("Impression", "impression", db.Int)
	info.AddField("Ecpm", "ecpm", db.Decimal)
	info.AddField("Date", "ad_date", db.Varchar).
		FieldFilterable(types.FilterType{FormType: form.Date})

	info.AddButton("GET", icon.Send, action.Ajax("ajax_facebook",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {

			return true, "success", ""
		}))

	info.SetTable("tb_ad_corp_facebook").SetTitle("Facebook Ads").SetDescription("Report")

	return
}
