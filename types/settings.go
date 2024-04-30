package types

type SettingsForm struct {
	Amount   string `form:"amount" json:"amount"`
	SearchOn bool   `form:"searchOn" json:"searchOn"`
	AddNew   bool   `form:"addNew" json:"addNew"`
}
