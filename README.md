# corpweb
corpweb是一个企业站点，站点包含了首页、产品、博客、关于我们、联系我们等通用的页面，还有针对这些页面的管理后台。

由于我一个做beanbag外贸生意的朋友需要一个企业站点，正好我也在学习golang，可以利用这次机会真正写一个完整的程序，于是便有了这个项目。本着`取之于民，用之于民`的原则，我把项目源代码推送到了这里，不管最终是否帮到别人，我都会感到因开源而得到的内心的一份荣耀。


# 安装
- 安装[golang](http://golang.org/)
- 设置`GOROOT`和`GOPATH`环境变量
- 安装依赖
```
go get github.com/astaxie/beego
go get github.com/beego/i18n
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm
go get github.com/nfnt/resize
go get github.com/oliamb/cutter
go get github.com/satori/go.uuid
go get github.com/slene/blackfriday
```
- 进入`$GOPATH/src/corpweb`目录并执行`go build`命令。
- 安装[MySQL](http://www.mysql.com/)数据库服务器。创建一个数据库如`corpweb`并在其中执行初始化SQL脚本`sql/init.sql`。
- 根据实际情况修改配置文件`conf/app.conf`。


# 运行
安装完成后进入corpweb目录并运行。
```
cd $GOPATH/src/corpweb
./corpweb
```
生产环境中的部署请使用`Supervisord`工具。

# 实际效果截图
![](https://github.com/voidint/corpweb/raw/master/screenshots/home.PNG)
![](https://github.com/voidint/corpweb/raw/master/screenshots/product.PNG)
![](https://github.com/voidint/corpweb/raw/master/screenshots/product_item.PNG)
![](https://github.com/voidint/corpweb/raw/master/screenshots/contact.PNG)


