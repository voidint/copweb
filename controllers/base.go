package controllers

import (
	"errors"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
	RESULT_RESP_SUCC string = "success"
	RESULT_RESP_FAIL string = "fail"
)

var (
	mobilePattern = regexp.MustCompile("^((\\+86)|(86))?(1(([35][0-9])|(47)|[8][012356789]))\\d{8}$")
	telPattern    = regexp.MustCompile("^(0\\d{2,3}(\\-)?)?\\d{7,8}$")
	emailPattern  = regexp.MustCompile("^(\\w)+(\\.\\w+)*@(\\w)+((\\.\\w+)+)$")
)

type AjaxFormResp struct {
	Result string
	Msg    string
	Fields map[string]string      `json:",omitempty"`
	ExtMap map[string]interface{} `json:",omitempty"`
	ExtObj interface{}            `json:",omitempty"`
}

var langTypes []*langType // Languages are supported.

// langType represents a language type.
type langType struct {
	Lang, Name string
}

func initLocales() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

func init() {
	initLocales()
}

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	// Redirect to make URL clean.
	if this.setLangVer() {
		i := strings.Index(this.Ctx.Request.RequestURI, "?")
		this.Redirect(this.Ctx.Request.RequestURI[:i], 302)
		return
	}
}

// setLangVer sets site language version.
func (this *baseController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := this.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		this.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	this.Lang = lang
	this.Data["Lang"] = curLang.Lang
	this.Data["CurLang"] = curLang.Name
	this.Data["RestLangs"] = restLangs

	return isNeedRedir
}

// CalcTotalPages 计算总页数
func (this *baseController) CalcTotalPages(pageSize, totalRecords int64) (int64, error) {
	var totalPages int64
	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return totalPages, nil
}

// CalcStartRecordNo 计算分页起始记录号
func (this *baseController) CalcStartRecordNo(curPageNo, pageSize, totalRecords int64) (int64, error) {
	if curPageNo <= 0 && pageSize <= 0 {
		return 0, errors.New("Invalid agruments.")
	}

	totalPages, err := this.CalcTotalPages(pageSize, totalRecords)
	if err != nil {
		return 0, err
	}

	if totalPages > 0 && curPageNo > totalPages {
		curPageNo = totalPages
	}
	return (curPageNo - 1) * pageSize, nil
}
