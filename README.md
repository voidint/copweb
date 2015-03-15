# corpweb
corpweb是一个基于[beego](http://beego.me/)和[bootstrap3](http://getbootstrap.com/)构建的企业站点，站点包含了首页、产品、博客、关于我们、联系我们等通用的页面，还有针对这些页面的管理后台。

由于我一个做beanbag外贸生意的朋友需要一个企业站点，正好我也在学习golang，可以利用这次机会真正写一个完整的程序，于是便有了这个项目。本着`取之于民，用之于民`的原则，我把项目源代码推送到了这里，不管最终是否帮到别人，我都会感到因开源而得到的内心的一份荣耀。


## 安装
- 安装[golang](http://golang.org/)
- 设置`GOROOT`和`GOPATH`环境变量
- 下载源代码
```
cd $GOPATH/src
git clone https://github.com/voidint/corpweb.git
```
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
- 编译
```
cd $GOPATH/src/corpweb
go build 
```
- 安装[MySQL](http://www.mysql.com/)数据库服务器。使用root账号执行初始化SQL脚本`sql/init.sql`，以下内容包含在脚本中，强烈建议用户修改数据库名称、数据库登录账号名、数据库登录账号密码等参数以保障数据库安全，修改完后需要一并修改`conf/app.conf`配置文件中的`dbname`、`dbusername`、`dbuserpwd`三个参数。
```sql
/*创建名为“corpweb”的数据库*/
CREATE DATABASE IF NOT EXISTS corpweb DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci;


/*创建数据库管理员用户，账号名“panda_corpweb”，密码“abc#123”*/
CREATE USER 'panda_corpweb'@'%' IDENTIFIED BY 'abc#123';


/*赋予panda_corpweb用户corpweb数据库的所有操作权限*/
GRANT ALL PRIVILEGES ON corpweb.* TO 'panda_corpweb'@'%';


/*选择目标数据库*/
USE corpweb;

/*...省略以下内容...*/
```


## 配置
- 修改配置文件`conf/app.conf`

本项目基于beego搭建，因此多数参数都可以参照[beego参数配置说明](http://beego.me/docs/mvc/controller/config.md)，以下配置项为项目自定义的配置项说明：
> dbname: 数据库名称
> dbusername: 数据库登录用户账号
> dbuserpwd: 数据库登录用户密码

- 设置环境变量

> RESOURCE_HOME: 用于指定资源（当前主要是用户上传的图片）的根路径，若未设置该环境变量，默认的资源根路径会是项目目录下的resource目录。

## 运行
安装完成后进入corpweb目录并运行。
```
cd $GOPATH/src/corpweb
./corpweb
```
生产环境中的部署请使用`Supervisord`工具。

## 效果图
- 首页
![](https://github.com/voidint/corpweb/raw/master/screenshots/home.PNG)
- 产品列表
![](https://github.com/voidint/corpweb/raw/master/screenshots/product.PNG)
- 产品明细
![](https://github.com/voidint/corpweb/raw/master/screenshots/product_item.PNG)
- 联系我们
![](https://github.com/voidint/corpweb/raw/master/screenshots/contact.PNG)


