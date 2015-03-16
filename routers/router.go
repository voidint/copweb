package routers

import (
	"corpweb/controllers"
	"corpweb/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	portalContr := &controllers.PortalController{}
	beego.Router("/", portalContr, "get:ToHome")
	beego.Router("/home", portalContr, "get:ToHome")
	beego.Router("/products", portalContr, "get:ToProducts")
	beego.Router("/products.json", portalContr, "get:GetProducts")
	beego.Router("/product/item/:prodId", portalContr, "get:ToProductItem")
	beego.Router("/blog", portalContr, "get:ToBlog")
	beego.Router("/blog/page", portalContr, "get:AjaxGetBlogPage")
	beego.Router("/blog/post/:blogId", portalContr, "get:ToBlogPost")
	beego.Router("/about", portalContr, "get:ToAbout")
	beego.Router("/contact", portalContr, "get:ToContact")
	beego.Router("/contact/msg/add", portalContr, "post:AddContactMsg")

	userContr := &controllers.UserController{}
	beego.Router("/login", userContr, "get:ToLogin")
	beego.Router("/login", userContr, "post:Login")
	beego.Router("/logout", userContr, "get:ToLogout")
	beego.Router("/admin/user/changepwd", userContr, "post:ChangePwd")

	prodContr := &controllers.ProductController{}
	beego.Router("/admin/product/add", prodContr, "get:ToProductAdd")
	beego.Router("/admin/product/add.json", prodContr, "post:AddProduct")
	beego.Router("/admin/product/mod", prodContr, "get:ToProductMod")
	beego.Router("/admin/product/mod.json", prodContr, "post:ModProduct")
	beego.Router("/admin/product", prodContr, "get:ToProducts")
	beego.Router("/admin/products.json", prodContr, "get:AjaxGetProductList")
	beego.Router("/admin/products/remove.json", prodContr, "post:RmProducts")
	beego.Router("/admin/products/pushpin.json", prodContr, "post:PushPin")

	caroContr := &controllers.CarouselController{}
	beego.Router("/admin/home/carousel", caroContr, "get:ToCarousel")
	beego.Router("/admin/home/carousel/add", caroContr, "get:ToAddCarousel")
	beego.Router("/admin/home/carousel/mod", caroContr, "get:ToModCarousel")
	beego.Router("/admin/home/carousel.json", caroContr, "get:GetCarousels")
	beego.Router("/admin/home/carousel/add.json", caroContr, "post:AddCarousel")
	beego.Router("/admin/home/carousel/mod.json", caroContr, "post:ModCarousel")
	beego.Router("/admin/home/carousel/remove.json", caroContr, "post:RmCarousel")
	beego.Router("/admin/home/carousel/pushpin.json", caroContr, "post:PushPin")

	flagshipContr := &controllers.FlagshipProductController{}
	beego.Router("/admin/home/products/flagship", flagshipContr, "get:ToFlagshipProducts")
	beego.Router("/admin/home/products/unflagship.json", flagshipContr, "get:AjaxGetProductsButFlagships")
	beego.Router("/admin/home/products/flagship.json", flagshipContr, "get:GetFlagshipProducts")
	beego.Router("/admin/home/products/flagship/add.json", flagshipContr, "post:AddFlagshipProducts")
	beego.Router("/admin/home/products/flagship/remove.json", flagshipContr, "post:RmFlagshipProduct")
	beego.Router("/admin/home/products/flagship/pushpin.json", flagshipContr, "post:PushPin")

	blogContr := &controllers.BlogController{}
	beego.Router("/admin/blog", blogContr, "get:ToBlog")
	beego.Router("/admin/blog/write", blogContr, "get:ToAddBlog")
	beego.Router("/admin/blog/edit/:blogId", blogContr, "get:ToEditBlog")
	beego.Router("/admin/blog/add", blogContr, "post:AjaxAddBlog")
	beego.Router("/admin/blog/mod/:blogId", blogContr, "post:AjaxModBlog")
	beego.Router("/admin/blog/remove/:blogId", blogContr, "post:AjaxRmBlog")
	beego.Router("/admin/blog/list", blogContr, "get:AjaxGetBlogList")
	beego.Router("/admin/blog/pushpin/:blogId", blogContr, "post:AjaxPushPin")

	msgContr := &controllers.MessageController{}
	beego.Router("/admin/message/contact", msgContr, "get:ToContactMsg")
	beego.Router("/admin/message/contact/page", msgContr, "get:AjaxGetContactMsgPage")
	beego.Router("/admin/message/contact/mark", msgContr, "post:AjaxMarkMsg")
	beego.Router("/admin/message/contact/search", msgContr, "post:SearchContactMsgPage")

	adminContr := &controllers.AdminController{}
	beego.Router("/admin/index", adminContr, "*:Index")
	beego.Router("/admin/img/upload", adminContr, "post:ImgUpload")
	beego.Router("/admin/img/crop", adminContr, "post:ImgCrop")
	beego.Router("/admin/markdown2html.json", adminContr, "post:Markdown2html")

	settingContr := &controllers.SettingsController{}
	beego.Router("/admin/settings/personal", settingContr, "get:ToPersonalSetting")
	beego.Router("/admin/settings/sys", settingContr, "get:ToSysSetting")
	beego.Router("/admin/settings/changepwd", settingContr, "get:ToChangePwd")

	// 登录过滤器
	beego.InsertFilter("/admin/*", beego.BeforeRouter, func(ctx *context.Context) {
		user, ok := ctx.Input.Session("UserInfo").(models.User)

		if !ok || len(user.UserId) <= 0 || len(user.LoginName) <= 0 {
			ctx.Redirect(302, "/login")
		}
	})

}
