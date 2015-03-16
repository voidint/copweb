package controllers

import (
	"fmt"
	"image"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"corpweb/conf"
	"corpweb/utils"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

type AdminController struct {
	// beego.Controller
	baseController
}

func (this *AdminController) Index() {
	this.Data["menu_lv_1"] = "dashboard"
	this.Data["menu_lv_2"] = ""

	this.TplNames = "admin/index.html"
}

// ImgUpload 图片上传
func (this *AdminController) ImgUpload() {
	beego.Debug("AdminController.ImgUpload()...")

	ajaxResp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
	}
	defer func() {
		this.Data["json"] = &ajaxResp
		this.ServeJson(true)
	}()

	mFile, fHeader, err := this.GetFile("file")
	if err != nil {
		beego.Error(err)
		ajaxResp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	// 判断上传的是否是支持的图片文件
	data := make([]byte, 512)
	_, err = mFile.Read(data)
	if err != nil {
		beego.Error(err)
		ajaxResp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}
	mimeType := strings.TrimSpace(http.DetectContentType(data))
	if mimeType != "image/jpeg" && mimeType != "image/png" {
		ajaxResp.Msg = this.Tr("tips_invalid_file_type")
		return
	}

	// 转储图片文件
	newFileName := uuid.NewV4().String() + strings.ToLower(filepath.Ext(fHeader.Filename))
	err = this.SaveToFile("file", conf.Resource_Home+"/image/source/"+newFileName)
	if err != nil {
		beego.Error(err)
		ajaxResp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	extension := make(map[string]interface{})
	extension["filepath"] = "/image/source/" + newFileName
	ajaxResp.ExtMap = extension
	ajaxResp.Result = RESULT_RESP_SUCC
}

// CropCoord 图片裁剪坐标信息
type CropCoord struct {
	Imgpath    string `form:"imgpath"`
	X          uint   `form:"x"`
	Y          uint   `form:"y"`
	Width      uint   `form:"width"`
	Height     uint   `form:"height"`
	SizeWidth  uint   `form:"sizeWidth"`
	SizeHeight uint   `form:"sizeHeight"`
}

// ImgCrop 裁剪图片指定区域并进行缩放
func (this *AdminController) ImgCrop() {
	beego.Debug("AdminController.ImgCrop()...")

	ajaxResp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
	}
	defer func() {
		this.Data["json"] = &ajaxResp
		this.ServeJson(true)
	}()

	coord := CropCoord{}
	if err := this.ParseForm(&coord); err != nil {
		beego.Error(err)
		ajaxResp.Msg = fmt.Sprintf("图片裁剪参数错误: %s\n", err.Error())
		return
	}

	fileExt := strings.ToLower(filepath.Ext(coord.Imgpath))
	imgType, err := utils.Convert2ImgType(fileExt)
	if err != nil {
		ajaxResp.Msg = fmt.Sprintf("系统当前不支持该种文件类型的裁剪操作: %s\n", fileExt)
		return
	}

	srcImg, err := os.Open(strings.Replace(coord.Imgpath, "/res/", conf.Resource_Home+"/", 1))
	if err != nil {
		beego.Error(err)
		ajaxResp.Msg = "源图片或已不存在，请重新上传图片。"
		return
	}

	dstImgName := uuid.NewV4().String() + fileExt
	dstImg, err := os.Create(conf.Resource_Home + "/image/" + dstImgName)
	if err != nil {
		beego.Error(err)
		ajaxResp.Msg = "无法在目标目录下创建文件，请联系技术人员。"
		return
	}

	rect := image.Rect(int(coord.X), int(coord.Y), int(coord.X+coord.Width), int(coord.Y+coord.Height))
	err = utils.Thumbnails(dstImg, srcImg, rect, coord.SizeWidth, coord.SizeHeight, imgType)
	if err != nil {
		beego.Error(err)
		ajaxResp.Msg = "源图片或已不存在，请重新上传图片。"
		return
	}

	ajaxResp.Msg = "图片裁剪成功"
	ajaxResp.Result = RESULT_RESP_SUCC
	extension := make(map[string]interface{}, 1)
	extension["newImgPath"] = "/image/" + dstImgName
	ajaxResp.ExtMap = extension
}

// Markdown2html 将Markdown内容转化为html
func (this *AdminController) Markdown2html() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	md := this.GetString("markdown")

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = utils.Markdown2html(md)
	// resp.ExtMap = make(map[string]interface{})
	// resp.ExtMap["htmlContent"] = utils.Markdown2html(md)

}
