package controllers

type SettingsController struct {
	// beego.Controller
	baseController
}

// ToPersonalSetting 转到个人设置页
func (this *SettingsController) ToPersonalSetting() {
	this.Data["menu_lv_1"] = "settings"
	this.Data["menu_lv_2"] = "settings_personal"

	this.TplNames = "admin/settings_personal.html"
}

// ToSysSetting 转到系统设置页
func (this *SettingsController) ToSysSetting() {
	this.Data["menu_lv_1"] = "settings"
	this.Data["menu_lv_2"] = "settings_sys"

	this.TplNames = "admin/settings_sys.html"
}

func (this *SettingsController) ToChangePwd() {
	this.TplNames = "admin/ajax_html/settings_personal_form_change_pwd.html"
}
