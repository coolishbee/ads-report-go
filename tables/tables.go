// This file is generated by GoAdmin CLI adm.
package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// Generators is a map of table models.
//
// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// example end
var Generators = map[string]table.Generator{
	"tb_ad_corp_unity":    GetUnityAds,
	"tb_ad_corp_admob":    GetGoogleAdmob,
	"tb_ad_corp_facebook": GetFacebookAds,
	"tb_ad_corp_iron":     GetIronSourceAds,
	"tb_ad_corp_vungle":   GetVungleAds,
	"tb_ad_corp_adx":      GetAdxAds,
	// generators end
}
