/*创建数据库*/
CREATE DATABASE IF NOT EXISTS corpweb DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci;


/*创建数据库管理员用户*/
CREATE USER 'panda_corpweb'@'%' IDENTIFIED BY 'abc#123';


/*赋予当前数据库所有操作权限*/
GRANT ALL PRIVILEGES ON corpweb.* TO 'panda_corpweb'@'%';


/*选择目标数据库*/
USE corpweb;