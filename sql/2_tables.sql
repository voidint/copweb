/*选择目标数据库*/
USE corpweb;

-- MySQL dump 10.13  Distrib 5.6.14, for Win64 (x86_64)
--
-- Host: localhost    Database: corpweb
-- ------------------------------------------------------
-- Server version 5.6.14

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `t_blog`
--

DROP TABLE IF EXISTS `t_blog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_blog` (
  `blog_id` char(36) NOT NULL COMMENT '主键',
  `blog_title` varchar(100) NOT NULL COMMENT '博客标题',
  `blog_intro` varchar(250) DEFAULT NULL,
  `blog_body` text NOT NULL COMMENT '博客正文',
  `blog_body_use_markdown` tinyint(1) unsigned NOT NULL COMMENT '是否使用了markdown来书写。1-是；2-否；',
  `blog_tags` varchar(250) DEFAULT '' COMMENT '标签（多个标签可用,分隔）',
  `blog_is_public` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否公开。1-公开;2-不公开。',
  `blog_sort_no` int(10) NOT NULL DEFAULT '0' COMMENT '排序序号，序号越大越靠前。',
  `blog_created` datetime NOT NULL COMMENT '博客创建时间',
  `blog_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`blog_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_blog_img`
--

DROP TABLE IF EXISTS `t_blog_img`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_blog_img` (
  `blog_img_id` char(36) NOT NULL COMMENT '主键',
  `blog_img_blog_id` char(36) NOT NULL COMMENT '产品ID',
  `blog_img_path` varchar(512) NOT NULL COMMENT '图片路径（建议使用相对路径）',
  `blog_img_created` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`blog_img_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_config`
--

DROP TABLE IF EXISTS `t_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_config` (
  `cfg_key` varchar(100) NOT NULL COMMENT '配置键',
  `cfg_value` varchar(256) NOT NULL COMMENT '配置内容',
  `cfg_created` datetime NOT NULL,
  `cfg_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cfg_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_contact_msg`
--

DROP TABLE IF EXISTS `t_contact_msg`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_contact_msg` (
  `msg_id` char(36) NOT NULL COMMENT '主键id',
  `msg_name` varchar(50) NOT NULL COMMENT '联系人姓名',
  `msg_email` varchar(100) NOT NULL COMMENT '联系人邮箱',
  `msg_phone` varchar(20) NOT NULL COMMENT '联系人电话',
  `msg_company` varchar(150) DEFAULT '' COMMENT '联系人公司',
  `msg_text` varchar(1024) DEFAULT '' COMMENT '联系人留言',
  `msg_state` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '处理状态。1-未处理;2-已处理;',
  `msg_created` datetime NOT NULL COMMENT '创建时间',
  `msg_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`msg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='联系人消息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_home_carousel`
--

DROP TABLE IF EXISTS `t_home_carousel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_home_carousel` (
  `caro_id` char(36) NOT NULL,
  `caro_img_path` varchar(512) NOT NULL COMMENT '滚动巨幕图片路径',
  `caro_caption` varchar(200) NOT NULL COMMENT '滚动巨幕图片配文',
  `caro_sort_no` int(10) NOT NULL COMMENT '滚动巨幕顺序，序号越大越靠前。',
  `caro_created` datetime NOT NULL,
  `caro_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`caro_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_home_flagship_product`
--

DROP TABLE IF EXISTS `t_home_flagship_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_home_flagship_product` (
  `fp_id` char(36) NOT NULL,
  `fp_prod_id` char(36) NOT NULL,
  `fp_sort_no` int(10) NOT NULL,
  `fp_created` datetime NOT NULL,
  `fp_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`fp_id`),
  UNIQUE KEY `uk_prodid_flagship_prod` (`fp_prod_id`) USING BTREE,
  CONSTRAINT `fk_prod_id_flagship_prod` FOREIGN KEY (`fp_prod_id`) REFERENCES `t_product` (`prod_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_log`
--

DROP TABLE IF EXISTS `t_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_log` (
  `log_id` char(36) NOT NULL COMMENT '日志ID',
  `log_sponsor` varchar(100) NOT NULL DEFAULT '' COMMENT '发起者，建议记录IP。',
  `log_terminal` varchar(50) DEFAULT NULL COMMENT '发起者，建议记录IP。',
  `log_action` enum('send_contact_msg','send_email','logout','login') NOT NULL COMMENT '动作',
  `log_result` enum('success','fail') NOT NULL COMMENT '操作结果',
  `log_msg` varchar(1024) DEFAULT NULL COMMENT '详细日志',
  `log_created` datetime NOT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`log_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_login`
--

DROP TABLE IF EXISTS `t_login`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_login` (
  `login_user_id` char(36) NOT NULL COMMENT '用户ID,关联t_user表中的user_id字段',
  `login_name` varchar(100) NOT NULL COMMENT '登录账号(推荐使用Email)',
  `login_pwd` varchar(50) NOT NULL COMMENT '登录密码=MD5(原密码+login_salt)',
  `login_salt` varchar(20) NOT NULL COMMENT '加盐值',
  PRIMARY KEY (`login_name`),
  KEY `idx_userid_login` (`login_user_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_mail`
--

DROP TABLE IF EXISTS `t_mail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_mail` (
  `mail_id` char(32) NOT NULL,
  `mail_from` varchar(256) NOT NULL COMMENT '邮件发送人',
  `mail_to` varchar(1024) NOT NULL COMMENT '邮件接收人列表',
  `mail_cc` varchar(1024) DEFAULT NULL COMMENT '邮件抄送列表',
  `mail_bcc` varchar(255) DEFAULT NULL COMMENT '邮件密送列表',
  `mail_subject` varchar(512) NOT NULL COMMENT '邮件主题',
  `mail_body` text NOT NULL COMMENT '邮件正文',
  `mail_date` datetime NOT NULL COMMENT '邮件发送时间',
  `mail_owner` char(36) NOT NULL COMMENT '所有者（系统用户ID）',
  PRIMARY KEY (`mail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_product`
--

DROP TABLE IF EXISTS `t_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_product` (
  `prod_id` char(36) NOT NULL,
  `prod_title` varchar(100) NOT NULL COMMENT '产品标题',
  `prod_intro` varchar(200) NOT NULL COMMENT '简介',
  `prod_desc` varchar(5120) NOT NULL COMMENT '产品描述（可含有HTML或者markdown语法内容）',
  `prod_desc_use_markdown` tinyint(1) unsigned NOT NULL COMMENT '产品描述是否使用了markdown来书写。1-是；2-否；',
  `prod_is_public` tinyint(1) unsigned NOT NULL COMMENT '是否公开。1-公开（前后台均可见）;2-非公开（仅后台可见）;',
  `prod_sort_no` int(10) NOT NULL DEFAULT '0' COMMENT '排序序号，序号越大越靠前。',
  `prod_created` datetime NOT NULL,
  `prod_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`prod_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='产品信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_product_img`
--

DROP TABLE IF EXISTS `t_product_img`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_product_img` (
  `prod_img_id` char(36) NOT NULL COMMENT '主键',
  `prod_img_pid` char(36) DEFAULT NULL COMMENT '源图片ID',
  `prod_img_prod_id` char(36) NOT NULL COMMENT '产品ID',
  `prod_img_path` varchar(512) NOT NULL COMMENT '图片路径（建议使用相对路径）',
  `prod_img_place_at` tinyint(1) unsigned NOT NULL COMMENT '图片放置位置。1-封面；2-详情；',
  `prod_img_created` datetime NOT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`prod_img_id`),
  KEY `idx_prodid_prodimg` (`prod_img_prod_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_settings_mail`
--

DROP TABLE IF EXISTS `t_settings_mail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_settings_mail` (
  `mailset_id` char(36) NOT NULL COMMENT '主键ID',
  `mailset_account` varchar(128) NOT NULL COMMENT '邮件账号',
  `mailset_pwd` varchar(50) NOT NULL COMMENT '邮箱密码（加密）',
  `mailset_outgoing` varchar(128) NOT NULL COMMENT '邮件发件服务器，如smtp.126.com',
  `mailset_outgoing_port` int(5) unsigned NOT NULL DEFAULT '25' COMMENT '发件服务器端口号，默认值25',
  `mailset_created` datetime NOT NULL,
  `mailset_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `mailset_owner` char(36) NOT NULL COMMENT '所有者（系统用户ID）',
  PRIMARY KEY (`mailset_id`),
  UNIQUE KEY `uk_userid_t_settings_mail` (`mailset_owner`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user` (
  `user_id` char(36) NOT NULL COMMENT '用户ID',
  `user_full_name` varchar(100) DEFAULT NULL COMMENT '姓名',
  `user_birthday` date DEFAULT NULL COMMENT '用户出生年月',
  `user_gender` enum('male','female','other') DEFAULT 'male' COMMENT '性别',
  `user_nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
  `user_created` datetime NOT NULL COMMENT '创建时间',
  `user_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户基本信息表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-03-24 17:38:06