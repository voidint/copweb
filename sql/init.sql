/*
Navicat MySQL Data Transfer

Source Server         : loc_root
Source Server Version : 50614
Source Host           : localhost:3306
Source Database       : cms4go

Target Server Type    : MYSQL
Target Server Version : 50614
File Encoding         : 65001

Date: 2015-03-12 16:15:54
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `t_blog`
-- ----------------------------
DROP TABLE IF EXISTS `t_blog`;
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

-- ----------------------------
-- Records of t_blog
-- ----------------------------
INSERT INTO `t_blog` VALUES ('312b9318-78d6-4442-ae96-acb49944185f', 'golang is good!', null, '<span style=\"font-weight: bold; font-style: italic;\">golang is good</span>', '2', '', '1', '7', '2015-02-21 16:33:09', '2015-03-04 10:26:49');
INSERT INTO `t_blog` VALUES ('4a40e61c-94d7-4bbd-89bc-987aa56551ae', 'unit test', null, '### write in markdown', '1', '', '1', '0', '2015-02-22 14:47:28', '2015-02-22 21:08:58');
INSERT INTO `t_blog` VALUES ('5cab7105-092e-4180-94cd-74d58078aee0', 'Java的语法实在太啰嗦了', null, '##### Java的语法实在太啰嗦了\n```\npublic static void main(string[] args){\n    System.out.print(\"hello world!\");\n}\n```', '1', '', '2', '2', '2015-02-21 16:40:53', '2015-03-03 11:55:37');
INSERT INTO `t_blog` VALUES ('5e2f14ee-517c-4204-9aae-951e1970645e', '你妹', null, '### 你妹\nyour sister!', '1', '', '1', '0', '2015-02-22 22:08:48', '2015-02-22 22:09:06');
INSERT INTO `t_blog` VALUES ('cd96b858-442a-48f5-a60f-97d8cd6abdd2', 'linux is good', null, '<span style=\"font-weight: bold; font-style: italic;\">linux is good</span>', '2', '', '1', '4', '2015-02-21 16:32:34', '2015-03-03 10:48:39');
INSERT INTO `t_blog` VALUES ('e3482660-07df-4fe8-ba8f-0bd55afe2904', '_markdown test', null, '### markdown test', '1', '', '1', '0', '2015-03-03 11:00:27', '2015-03-03 11:11:35');
INSERT INTO `t_blog` VALUES ('f033d61b-2f9e-4d35-b69c-af5c20ea8a8c', 'markdown is good', null, '### markdown is good', '1', '', '1', '3', '2015-02-21 16:31:39', '2015-02-21 21:57:43');

-- ----------------------------
-- Table structure for `t_blog_img`
-- ----------------------------
DROP TABLE IF EXISTS `t_blog_img`;
CREATE TABLE `t_blog_img` (
  `blog_img_id` char(36) NOT NULL COMMENT '主键',
  `blog_img_blog_id` char(36) NOT NULL COMMENT '产品ID',
  `blog_img_path` varchar(512) NOT NULL COMMENT '图片路径（建议使用相对路径）',
  `blog_img_created` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`blog_img_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_blog_img
-- ----------------------------
INSERT INTO `t_blog_img` VALUES ('08d7d4c6-efc9-4953-9029-12393ab452eb', '93b34f2f-021d-40c9-88df-4fff919a6524', '/image/e35ce8bf-d84a-4efe-9c8a-a803f333d464.jpg', '2015-02-22 14:33:33');
INSERT INTO `t_blog_img` VALUES ('8e325777-2bed-4261-b47b-8a240d2ac2cf', 'cc2db631-58c4-4845-8559-5627d460199b', '/image/e35ce8bf-d84a-4efe-9c8a-a803f333d464.jpg', '2015-02-22 14:32:35');
INSERT INTO `t_blog_img` VALUES ('cf0df36e-a012-45a0-a039-dbe986a10ca9', '59d4e102-843a-4973-b24b-b9b029fdee3e', '/image/e35ce8bf-d84a-4efe-9c8a-a803f333d464.jpg', '2015-02-22 14:34:34');
INSERT INTO `t_blog_img` VALUES ('f1f10b4b-37ca-449b-9755-dfd6fe585ce1', '75f7f2e4-6c92-4ee3-ad43-919d47399cdb', '/image/e35ce8bf-d84a-4efe-9c8a-a803f333d464.jpg', '2015-02-22 14:31:12');

-- ----------------------------
-- Table structure for `t_config`
-- ----------------------------
DROP TABLE IF EXISTS `t_config`;
CREATE TABLE `t_config` (
  `cfg_key` varchar(100) NOT NULL COMMENT '配置键',
  `cfg_value` varchar(256) NOT NULL COMMENT '配置内容',
  `cfg_created` datetime NOT NULL,
  `cfg_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cfg_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_config
-- ----------------------------

-- ----------------------------
-- Table structure for `t_contact_msg`
-- ----------------------------
DROP TABLE IF EXISTS `t_contact_msg`;
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

-- ----------------------------
-- Records of t_contact_msg
-- ----------------------------
INSERT INTO `t_contact_msg` VALUES ('171ace84-031c-4590-a0c4-8c354e1a53ac', '张三', 'zhangsan@gmail.com', '110', 'facebook', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Quae repudiandae fugiat illo cupiditate excepturi esse officiis consectetur, laudantium qui voluptatem. Ad necessitatibus velit, accusantium expedita debitis impedit rerum totam id. Lorem ipsum dolor sit amet, consectetur adipisicing elit. Natus quibusdam recusandae illum, nesciunt, architecto, saepe facere, voluptas eum incidunt dolores magni itaque autem neque velit in. At quia quaerat asperiores.', '2', '2015-03-07 22:16:15', '2015-03-09 10:24:50');
INSERT INTO `t_contact_msg` VALUES ('5c1c7d8d-dd65-4e4e-b3ec-1807b13fe713', 'voidint', 'tangtianyun.coder@gmail.com', '15857113047', '', 'Flot is a pure JavaScript plotting library for jQuery, with a focus on simple usage, attractive looks, and interactive features. In SB Admin, we are using the most recent version of Flot along with a few plugins to enhance the user experience. The Flot plugins being used are the tooltip plugin for hoverable tooltips, and the resize plugin for fully responsive charts. The documentation for Flot Charts is available on their website, http://www.flotcharts.org/.', '1', '2014-12-21 14:17:20', '2015-02-24 21:38:28');
INSERT INTO `t_contact_msg` VALUES ('5d5d8893-2ad3-4a5d-8b6c-06917454250d', 'tangtianyun', 'tangtianyun@126.com', '15857113047', 'google', 'fuck u', '1', '2015-01-20 15:26:37', '2015-03-06 17:27:19');
INSERT INTO `t_contact_msg` VALUES ('7152ff3e-24a9-425e-a4ed-d4e9c3f62ba5', 'voidint', 'voidint@126.com', '17091601585', 'google', 'fuck', '2', '2015-01-20 15:22:41', '2015-03-11 15:06:57');
INSERT INTO `t_contact_msg` VALUES ('7b341939-c060-46e6-bfc1-3c6a2c869f7f', 'heyili', 'heyili@love.org', '15857113047', 'google', 'love', '1', '2015-01-22 14:15:04', '2015-03-06 17:27:49');
INSERT INTO `t_contact_msg` VALUES ('7d1439e8-25e1-48f8-b1c8-fe3eb8825cfe', '严飞云', 'asktianyun@yeah.net', '15857113047', '中软国际科技服务有限公司', '几个礼拜之前我被问到：“对于Go语言，最令你惊讶的是什么？”当时我就明确地给出了答案：“虽然我希望C++程序员能够使用Go作为替代拼，但实际上大部分Go程序员都是从Python和Ruby转过来的，其中却少有C++程序员。”', '2', '2015-03-06 11:36:16', '2015-03-06 17:46:03');
INSERT INTO `t_contact_msg` VALUES ('84903dfb-fd10-484f-b32e-7cbfc15a48ea', '2', 'voidint@126.com', '17091601535', 'google', 'hello world', '1', '2015-01-16 07:41:23', '2015-01-16 07:41:23');
INSERT INTO `t_contact_msg` VALUES ('b48fa0db-c21f-417e-9728-f29cd34c5323', '1', 'voidint@126.com', '15857113047', 'google', 'Hi', '1', '2015-01-16 07:36:34', '2015-01-16 07:36:34');
INSERT INTO `t_contact_msg` VALUES ('c4b182e4-fd2a-4be3-8fb6-ce79cd3d2c9c', '123', 'voidint@126.com', '123', '', '', '1', '2014-12-22 11:09:19', '2014-12-22 11:09:19');
INSERT INTO `t_contact_msg` VALUES ('e6d276cd-cfc9-4555-b73c-901e2289f781', 'voidint', 'voidint@126.com', '15857113047', 'voidint', 'test message', '1', '2014-12-21 09:33:11', '2014-12-21 09:33:11');
INSERT INTO `t_contact_msg` VALUES ('e8871f41-de36-475b-ab3b-9509d7c90f58', '汤天云', 'tangtianyun@126.com', '17091601585', '', '这篇文章是Google首席工程师、Go语言之父Rob Pike自己整理的6月21日在旧金山给Go SF的演讲稿。Rob提到：Go语言本是以C为原型，以C++为目标设计，但最终却大相径庭。值得一提的是，这3门语言都曾当选TIOBE年度语言。', '2', '2015-03-06 11:30:51', '2015-03-06 18:01:04');

-- ----------------------------
-- Table structure for `t_home_carousel`
-- ----------------------------
DROP TABLE IF EXISTS `t_home_carousel`;
CREATE TABLE `t_home_carousel` (
  `caro_id` char(36) NOT NULL,
  `caro_img_path` varchar(512) NOT NULL COMMENT '滚动巨幕图片路径',
  `caro_caption` varchar(200) NOT NULL COMMENT '滚动巨幕图片配文',
  `caro_sort_no` int(10) NOT NULL COMMENT '滚动巨幕顺序，序号越大越靠前。',
  `caro_created` datetime NOT NULL,
  `caro_modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`caro_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_home_carousel
-- ----------------------------
INSERT INTO `t_home_carousel` VALUES ('0eb485bf-bc10-4a44-9bc5-edfe0270e6bf', '/image/d0144872-9040-43fa-8f8a-191af43c5775.jpg', '天堂岛', '7', '2015-02-12 18:42:32', '2015-02-19 20:21:10');
INSERT INTO `t_home_carousel` VALUES ('3c560704-690c-4356-baeb-0c8192249195', '/image/4cf7d38c-d9aa-4535-9460-480e66bad43d.jpg', '火星地貌', '9', '2015-02-19 17:08:01', '2015-02-19 20:22:04');
INSERT INTO `t_home_carousel` VALUES ('4b6d3055-5cb2-4051-8c41-b9b62c58e4ce', '/image/aa360364-2e08-48ab-9197-2a7c04778578.jpg', '背山面海', '0', '2015-03-03 17:57:53', '2015-03-03 17:57:53');
INSERT INTO `t_home_carousel` VALUES ('aba4c923-7df3-43c3-9237-c69f09d6983a', '/image/191754dc-7e3d-43e7-b743-aa99f8d9084e.jpg', '北极冰川正在消融', '10', '2015-02-18 14:38:47', '2015-02-19 20:22:21');

-- ----------------------------
-- Table structure for `t_home_flagship_product`
-- ----------------------------
DROP TABLE IF EXISTS `t_home_flagship_product`;
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

-- ----------------------------
-- Records of t_home_flagship_product
-- ----------------------------
INSERT INTO `t_home_flagship_product` VALUES ('4d4979c8-c070-4572-8506-4bbec3fc501c', 'e075c59c-5e56-464f-820e-d108322eaa22', '11', '2015-02-18 16:25:05', '2015-02-19 20:45:37');
INSERT INTO `t_home_flagship_product` VALUES ('52f372da-b054-4226-96dc-d335432798fb', '2834101b-1ae3-4ced-b7fb-48d5bd086de6', '10', '2015-02-18 13:43:46', '2015-02-19 20:45:36');
INSERT INTO `t_home_flagship_product` VALUES ('e89ce7db-71e1-4986-8edf-195f7c4e4d56', 'cfa622f1-e33d-4e29-b0a9-f873529f1b3f', '9', '2015-02-16 18:51:58', '2015-02-19 20:45:32');

-- ----------------------------
-- Table structure for `t_log`
-- ----------------------------
DROP TABLE IF EXISTS `t_log`;
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

-- ----------------------------
-- Records of t_log
-- ----------------------------
INSERT INTO `t_log` VALUES ('ac35f358-cebc-4015-9bf6-727410263993', '127.0.0.1', '127.0.0.1', 'login', 'success', '', '2014-12-18 06:42:30');
INSERT INTO `t_log` VALUES ('e2b47465-4adf-4b7f-9820-0019092ff69f', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 07:23:39');
INSERT INTO `t_log` VALUES ('cfac9e58-6dd4-4097-84bb-bf82378e8dc5', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 07:33:51');
INSERT INTO `t_log` VALUES ('1d30be94-c811-4fed-9f2c-0ee0c913510a', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 07:34:43');
INSERT INTO `t_log` VALUES ('69ebc07f-13fe-4dc0-a8de-5a58305cdce9', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 07:37:51');
INSERT INTO `t_log` VALUES ('9b817270-4058-426a-9578-34acdcc82656', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 08:04:17');
INSERT INTO `t_log` VALUES ('398c9af7-d044-416e-be15-539d06b367b9', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 08:25:42');
INSERT INTO `t_log` VALUES ('335183ae-8f44-4820-83d3-e1bdbaabd2d2', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 09:14:38');
INSERT INTO `t_log` VALUES ('c68b5f2b-279e-43bb-aa39-2c62e11829d1', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 10:40:03');
INSERT INTO `t_log` VALUES ('64a91d3a-f867-4bc0-8bff-516af0da73d4', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 10:58:05');
INSERT INTO `t_log` VALUES ('61aece6e-2c34-4477-bd3c-65cdfe7da1c5', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-19 11:18:28');
INSERT INTO `t_log` VALUES ('8ea2cbcd-26cd-42d0-bc28-8cf16b43a2e2', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: e6d276cd-cfc9-4555-b73c-901e2289f781', '2014-12-21 09:33:11');
INSERT INTO `t_log` VALUES ('d98a2df3-72e7-49a1-a0ff-5405153a85ae', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: 5c1c7d8d-dd65-4e4e-b3ec-1807b13fe713', '2014-12-21 14:17:20');
INSERT INTO `t_log` VALUES ('703318e2-1865-4867-acd3-05cef5e17a67', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: c4b182e4-fd2a-4be3-8fb6-ce79cd3d2c9c', '2014-12-22 11:09:20');
INSERT INTO `t_log` VALUES ('661a83b0-d756-4189-9ea9-f7580bb62907', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-23 08:08:59');
INSERT INTO `t_log` VALUES ('5b1aabfb-087c-4f70-95f4-1e2d0d3c07bf', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-24 08:50:43');
INSERT INTO `t_log` VALUES ('66e1db52-6a5c-4542-96e6-63f2dc18af30', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-25 02:45:43');
INSERT INTO `t_log` VALUES ('8e9e511c-07ef-4dc4-a77b-abfc8d1b6b1e', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2014-12-25 08:49:23');
INSERT INTO `t_log` VALUES ('d3ccd67f-97f7-4470-b86f-82a07938dbcb', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-14 04:19:26');
INSERT INTO `t_log` VALUES ('47daa489-14dc-4d02-b811-0732db3863a6', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-15 09:21:43');
INSERT INTO `t_log` VALUES ('d9c471fa-bdb3-4786-9313-51e2ae78a973', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-16 02:08:58');
INSERT INTO `t_log` VALUES ('f480786b-07f3-42e9-a335-d6200240f327', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: b48fa0db-c21f-417e-9728-f29cd34c5323', '2015-01-16 07:36:35');
INSERT INTO `t_log` VALUES ('af7269aa-b360-4fa3-93d2-cd7d408cc351', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: 84903dfb-fd10-484f-b32e-7cbfc15a48ea', '2015-01-16 07:41:23');
INSERT INTO `t_log` VALUES ('0cb8a9d2-d915-4466-b0eb-140f24ceb87c', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-16 07:56:21');
INSERT INTO `t_log` VALUES ('7cff02a2-e34d-4259-a53f-061ec7ceee52', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-18 03:27:14');
INSERT INTO `t_log` VALUES ('e97b0bd7-98fb-4f0f-b6e2-d17c0baf4d61', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-19 02:58:13');
INSERT INTO `t_log` VALUES ('4540e838-5843-41c2-a538-290fe2a69e67', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-20 02:52:28');
INSERT INTO `t_log` VALUES ('b61b2a83-d3b6-4a57-aef6-37c5d69e879d', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-20 06:36:06');
INSERT INTO `t_log` VALUES ('3b5f1f4a-777b-4143-b849-2332c6440db2', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-20 06:40:26');
INSERT INTO `t_log` VALUES ('79e28169-9627-4cc5-8b8e-265ac62ab954', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-20 06:44:06');
INSERT INTO `t_log` VALUES ('25602d88-92fb-458d-9964-9f8355b74587', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-20 06:48:25');
INSERT INTO `t_log` VALUES ('815bc82a-e422-4883-907d-c305990d3cff', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-20 15:20:57');
INSERT INTO `t_log` VALUES ('e87ced61-12d8-420a-8b1a-6789d6231749', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: 7152ff3e-24a9-425e-a4ed-d4e9c3f62ba5', '2015-01-20 15:22:41');
INSERT INTO `t_log` VALUES ('4d2f4371-4368-452b-8591-7d8cd0152f29', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: 5d5d8893-2ad3-4a5d-8b6c-06917454250d', '2015-01-20 15:26:37');
INSERT INTO `t_log` VALUES ('ebf4786e-be67-4cb7-a6ab-37dad3db3d0a', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-20 15:28:47');
INSERT INTO `t_log` VALUES ('9b1b5012-bfe0-48b7-933a-f2be90b87f64', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-20 19:35:33');
INSERT INTO `t_log` VALUES ('d0304239-661a-4f43-83f3-636fb0c76a8f', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-21 07:50:54');
INSERT INTO `t_log` VALUES ('1d1a2e70-65e5-40f8-887e-5a67576848a0', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-22 09:03:25');
INSERT INTO `t_log` VALUES ('b0244844-fec2-4bb0-a8cc-772a8f03c79e', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: 7b341939-c060-46e6-bfc1-3c6a2c869f7f', '2015-01-22 14:15:04');
INSERT INTO `t_log` VALUES ('2bbc1327-644c-4071-8a45-c0adf665eef8', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-22 14:15:19');
INSERT INTO `t_log` VALUES ('9bd62fd2-a896-4095-8336-fb8b87de83aa', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-22 22:26:37');
INSERT INTO `t_log` VALUES ('b95a4907-87cb-4bcd-95b2-d6386ac8d1fa', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-23 13:51:12');
INSERT INTO `t_log` VALUES ('62e9a724-bb13-4a03-8efa-d0240e04c1ce', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-24 09:57:08');
INSERT INTO `t_log` VALUES ('14b7544c-6718-4de4-a422-a223409abee5', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-24 18:24:14');
INSERT INTO `t_log` VALUES ('a374e3e3-f497-4ae6-96eb-596e2d0f2dcd', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-25 10:34:14');
INSERT INTO `t_log` VALUES ('66ee372b-ed62-415d-ba40-9ff634c8c685', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-25 10:48:22');
INSERT INTO `t_log` VALUES ('920b5c0c-c626-4ec1-8102-9c035d3e2948', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-26 10:31:09');
INSERT INTO `t_log` VALUES ('2ae77a39-ae0f-4393-a71a-8283a61807f7', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-27 10:07:46');
INSERT INTO `t_log` VALUES ('4d6e4ffc-56cf-4b0f-99e3-c73028b554f6', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-28 13:23:15');
INSERT INTO `t_log` VALUES ('cb81132a-1f4e-4347-9bd3-7b3a626d0b05', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-29 19:51:23');
INSERT INTO `t_log` VALUES ('7e5b3a27-21e8-4aa2-a3f1-0f4397dfeece', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-30 14:22:51');
INSERT INTO `t_log` VALUES ('32bcdf51-3b85-4ea9-a608-797b6dae5631', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-30 20:23:15');
INSERT INTO `t_log` VALUES ('20edee73-8506-4fb7-938f-daae4ca20986', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-01-31 14:46:25');
INSERT INTO `t_log` VALUES ('2483f3ad-53ed-4ace-b1b1-ecfff36bb187', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-07 09:45:59');
INSERT INTO `t_log` VALUES ('d152cd66-a4ef-42d0-878d-c234e9d16f96', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-08 12:04:31');
INSERT INTO `t_log` VALUES ('3ded7184-5615-419f-8196-4cbac77940a7', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-09 12:52:53');
INSERT INTO `t_log` VALUES ('c84600ea-b7b6-4fe3-b2ae-3d8a12d699a0', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-10 09:40:51');
INSERT INTO `t_log` VALUES ('3db86323-c7da-45ae-96f9-5ea1d2e94c06', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-10 22:17:18');
INSERT INTO `t_log` VALUES ('ea8e120b-b22b-4e92-9b19-34fa04797457', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-10 22:43:24');
INSERT INTO `t_log` VALUES ('275d0d04-32fb-4946-abb2-50a446ef54d3', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-11 10:18:45');
INSERT INTO `t_log` VALUES ('d3adfb1d-6bda-4352-ad84-d9fcc72e7858', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-12 09:39:29');
INSERT INTO `t_log` VALUES ('bf3f415d-cbdb-41af-affc-f875d9fd6390', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-15 11:16:56');
INSERT INTO `t_log` VALUES ('c23114bc-7189-4a72-8939-97007533a245', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-16 10:41:27');
INSERT INTO `t_log` VALUES ('5ec62b2a-c96d-470d-af66-bdc67cfd82f8', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-18 10:59:32');
INSERT INTO `t_log` VALUES ('f9b5ed0c-572e-4499-ba3d-c63c4b4c9bc9', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-18 19:31:21');
INSERT INTO `t_log` VALUES ('f61f1bb0-aaa1-4057-a304-3f6aafdb9b06', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-19 16:07:19');
INSERT INTO `t_log` VALUES ('e24aba5f-3379-4c7b-bd58-25d34433b023', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-20 13:54:25');
INSERT INTO `t_log` VALUES ('76b36edf-c930-4293-b717-822b5093d9d9', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-21 10:35:00');
INSERT INTO `t_log` VALUES ('21138315-b9ac-4d24-8348-c8faad6bd0ef', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-21 16:35:27');
INSERT INTO `t_log` VALUES ('7619f3e3-c2c4-4607-87e9-5b424c9796be', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-22 14:15:20');
INSERT INTO `t_log` VALUES ('0d4dab34-d414-4f21-90d1-9c0e781e3c0e', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-22 21:00:58');
INSERT INTO `t_log` VALUES ('837bcfa1-0ae2-401e-ab80-c798b3a70519', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-24 14:33:47');
INSERT INTO `t_log` VALUES ('b9f92a5e-b133-4578-ba15-ca3c415cbc1c', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 16:27:56');
INSERT INTO `t_log` VALUES ('7f1aa462-9add-4250-a129-6427ed874432', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 16:30:06');
INSERT INTO `t_log` VALUES ('e079977c-4d61-4b6e-82f9-251e1a9c3130', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 16:35:30');
INSERT INTO `t_log` VALUES ('e53467c9-c4db-4b98-a1f6-ea4e510df4f1', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 16:40:14');
INSERT INTO `t_log` VALUES ('004ef7d0-9137-4070-b65c-cddf5c2b8ec6', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 16:42:33');
INSERT INTO `t_log` VALUES ('c41d8a15-45e3-488b-8c83-df75dfffc562', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 16:45:55');
INSERT INTO `t_log` VALUES ('991b691f-3549-46fc-8a6f-08b3e8e13ee4', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 16:54:21');
INSERT INTO `t_log` VALUES ('3772b120-b9c9-4b23-9a09-00ae5b32c54e', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 16:59:52');
INSERT INTO `t_log` VALUES ('994bbb1d-4279-4418-abc8-e54f77e5ab1f', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 22:23:39');
INSERT INTO `t_log` VALUES ('12164e68-e569-46ef-be87-610785e8d0bf', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 22:25:03');
INSERT INTO `t_log` VALUES ('ee71b710-4c24-4afe-9135-90eb71c7003e', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 22:25:51');
INSERT INTO `t_log` VALUES ('71f6ed9c-f2f8-4518-b867-c9d21c276884', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 22:32:05');
INSERT INTO `t_log` VALUES ('41ceed53-b3e9-42c1-ade5-87304f1b1b13', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 22:33:19');
INSERT INTO `t_log` VALUES ('24b9b42f-91ed-480f-8f8c-af4d08cfb28c', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 22:33:48');
INSERT INTO `t_log` VALUES ('7a294ac5-f5ab-465b-ac07-e2b82ef1ec6d', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-27 22:35:08');
INSERT INTO `t_log` VALUES ('8625b534-9ec1-4e19-aab7-d58cdcde0b9a', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 10:28:34');
INSERT INTO `t_log` VALUES ('316a6182-bc59-4c40-b87c-e5f79d5715ca', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 10:36:55');
INSERT INTO `t_log` VALUES ('cc17f718-ce6d-4bb7-b041-c2ade341522f', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 10:45:26');
INSERT INTO `t_log` VALUES ('75148a5f-6bea-4610-bd2c-6e78e08bfdd7', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 10:51:54');
INSERT INTO `t_log` VALUES ('7dcc18ea-5fe1-4609-9fa8-b6871e0db9eb', '7450b7fb-8145-11e4-beed-005056c00008', '127.0.0.1', 'logout', 'success', '', '2015-02-28 10:52:26');
INSERT INTO `t_log` VALUES ('1bac6308-2128-4067-a9ab-a361c850b0a1', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 11:16:27');
INSERT INTO `t_log` VALUES ('c7a09e14-1cb4-4b56-a5b5-18a239898f1c', '7450b7fb-8145-11e4-beed-005056c00008', '127.0.0.1', 'logout', 'success', '', '2015-02-28 11:16:44');
INSERT INTO `t_log` VALUES ('ac1b76ab-6089-4d67-898b-5597a7a7f497', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 15:19:02');
INSERT INTO `t_log` VALUES ('2f3ce792-4ef0-4a93-b024-f6c47bc1d605', '7450b7fb-8145-11e4-beed-005056c00008', '127.0.0.1', 'logout', 'success', '', '2015-02-28 15:21:49');
INSERT INTO `t_log` VALUES ('b5e5d7bf-9c93-4960-8057-6a996b60b515', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 15:44:20');
INSERT INTO `t_log` VALUES ('c7357247-c48b-400c-b03a-4237c2f06853', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 15:58:33');
INSERT INTO `t_log` VALUES ('ec8b2ca8-8537-4a86-9bdc-a25ebe66b32e', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 18:05:25');
INSERT INTO `t_log` VALUES ('cfb2d5a2-1562-4e16-8fac-1cad7c2dd0da', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-02-28 21:48:48');
INSERT INTO `t_log` VALUES ('8fa9834d-5562-4f71-b2a5-e16c85d6f8e4', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-02 16:56:41');
INSERT INTO `t_log` VALUES ('83887800-5cca-4347-bc2a-914221e1993c', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-02 17:19:33');
INSERT INTO `t_log` VALUES ('9f76917a-f81d-4591-9171-7b7ab21af8c9', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-02 18:44:12');
INSERT INTO `t_log` VALUES ('c5451420-7725-4f1b-9e0b-34c3209e2a92', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-02 19:03:19');
INSERT INTO `t_log` VALUES ('da6d3011-8363-4b94-8bcf-2cf175e31875', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-02 19:13:06');
INSERT INTO `t_log` VALUES ('ebd423bf-ef56-4cb4-bfc8-c838b8fabde9', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-03 10:13:10');
INSERT INTO `t_log` VALUES ('6127c2be-5153-4191-9d55-6ff85e58b618', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-04 10:13:54');
INSERT INTO `t_log` VALUES ('c019fc93-7eba-44ea-ba2d-0c02d0c7dfa3', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 10:46:14');
INSERT INTO `t_log` VALUES ('ec1d24bb-f056-4108-9986-00ced0225b70', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 14:37:16');
INSERT INTO `t_log` VALUES ('b6e6970f-ee79-46f3-9963-df70308bb6b4', '7450b7fb-8145-11e4-beed-005056c00008', '127.0.0.1', 'logout', 'success', '', '2015-03-05 14:37:28');
INSERT INTO `t_log` VALUES ('1b4a3516-bfd7-4b08-bcc8-719fb311edb9', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 14:37:41');
INSERT INTO `t_log` VALUES ('52145d3a-faf4-407e-8e30-500740646632', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 14:47:09');
INSERT INTO `t_log` VALUES ('c7fb75d9-4c78-4f66-91e2-e3d6dce7c503', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 18:17:46');
INSERT INTO `t_log` VALUES ('f123c59f-c877-4df0-ba91-84fdc082fa1c', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 18:23:29');
INSERT INTO `t_log` VALUES ('efe66628-de81-4276-af70-2bc3eb76dd60', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 18:24:46');
INSERT INTO `t_log` VALUES ('526914a5-d0e4-4ac5-804c-af87bc60e136', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 18:52:22');
INSERT INTO `t_log` VALUES ('7d6ab5c7-32c3-4f21-a5c3-7fa44925163c', '7450b7fb-8145-11e4-beed-005056c00008', '127.0.0.1', 'logout', 'success', '', '2015-03-05 18:55:51');
INSERT INTO `t_log` VALUES ('e7060b4c-576c-446f-b323-c8a30abd966d', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-05 18:56:07');
INSERT INTO `t_log` VALUES ('b5ccce54-01d3-463e-a5e1-1df386711d8b', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-06 10:10:07');
INSERT INTO `t_log` VALUES ('5acdd3a2-7c94-490e-a07c-36fbd05a367b', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: e8871f41-de36-475b-ab3b-9509d7c90f58', '2015-03-06 11:30:51');
INSERT INTO `t_log` VALUES ('6cd4970d-1ef1-4bd6-8596-85351bf277af', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: 7d1439e8-25e1-48f8-b1c8-fe3eb8825cfe', '2015-03-06 11:36:16');
INSERT INTO `t_log` VALUES ('20c2d5e8-e788-4eac-8dd6-17dfc7e48699', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-06 14:53:38');
INSERT INTO `t_log` VALUES ('0a793a99-0ea6-42cd-b210-31be0b821f84', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-07 22:01:17');
INSERT INTO `t_log` VALUES ('2bd9f1f4-703a-48b7-95fb-295e28521e0e', 'visitor', '127.0.0.1', 'send_contact_msg', 'success', 'contact message id: 171ace84-031c-4590-a0c4-8c354e1a53ac', '2015-03-07 22:16:16');
INSERT INTO `t_log` VALUES ('cfd2bfaf-9745-463c-bbf9-80322f6ab594', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-07 23:00:13');
INSERT INTO `t_log` VALUES ('14c66fdd-89d3-44e8-9d84-b97ca300f31f', '7450b7fb-8145-11e4-beed-005056c00008', '127.0.0.1', 'logout', 'success', '', '2015-03-07 23:04:31');
INSERT INTO `t_log` VALUES ('4b513037-d470-4c2d-8ed9-cff9f7d57129', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-08 09:53:49');
INSERT INTO `t_log` VALUES ('3410d0cc-a098-4b58-8ff7-05b54738bd3e', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-09 10:20:13');
INSERT INTO `t_log` VALUES ('e4784541-9c14-4267-ad6e-c1c7fe2fac3a', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-10 14:53:57');
INSERT INTO `t_log` VALUES ('95cb3283-d4a8-4b3e-9ce7-db33f0adfc67', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-11 14:33:58');
INSERT INTO `t_log` VALUES ('0c788f4c-a4c5-4a6d-83e7-f9ff5dab40e3', 'voidint@126.com', '127.0.0.1', 'login', 'success', '', '2015-03-12 16:04:58');

-- ----------------------------
-- Table structure for `t_login`
-- ----------------------------
DROP TABLE IF EXISTS `t_login`;
CREATE TABLE `t_login` (
  `login_user_id` char(36) NOT NULL COMMENT '用户ID,关联t_user表中的user_id字段',
  `login_name` varchar(100) NOT NULL COMMENT '登录账号(推荐使用Email)',
  `login_pwd` varchar(50) NOT NULL COMMENT '登录密码=MD5(原密码+login_salt)',
  `login_salt` varchar(20) NOT NULL COMMENT '加盐值',
  PRIMARY KEY (`login_name`),
  KEY `idx_userid_login` (`login_user_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_login
-- ----------------------------
INSERT INTO `t_login` VALUES ('7450b7fb-8145-11e4-beed-005056c00008', 'voidint@126.com', 'd3aec9ef22406323fa726804644b6af9', '616769');

-- ----------------------------
-- Table structure for `t_product`
-- ----------------------------
DROP TABLE IF EXISTS `t_product`;
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

-- ----------------------------
-- Records of t_product
-- ----------------------------
INSERT INTO `t_product` VALUES ('2834101b-1ae3-4ced-b7fb-48d5bd086de6', 'selinux', '超级稳定的操作系统', '<span style=\"font-weight: bold; font-style: italic;\">linux is good</span>', '2', '1', '7', '2015-02-08 20:54:40', '2015-02-20 15:43:29');
INSERT INTO `t_product` VALUES ('730cab15-91ba-45ad-a716-e77cd3e18cdc', 'unit test', 'unit test', 'unit test', '1', '1', '3', '2015-02-19 20:56:46', '2015-02-20 13:56:15');
INSERT INTO `t_product` VALUES ('a41bb236-226d-44b9-9ad1-9081e6d94033', 'product 4', 'product 4', '<h1>product <span style=\"font-weight: bold;\">4</span></h1>', '2', '1', '0', '2015-03-03 11:32:36', '2015-03-03 11:33:20');
INSERT INTO `t_product` VALUES ('b30ff759-5702-4b80-acc6-0dade39de5a8', '产品2', '产品2产品2产品2产品2产品2产品2', '### 产品2', '1', '1', '0', '2015-02-16 16:55:26', '2015-02-16 16:55:26');
INSERT INTO `t_product` VALUES ('cfa622f1-e33d-4e29-b0a9-f873529f1b3f', '产品3', '产品3产品3产品3产品3产品3产品3', '## 产品3', '1', '1', '0', '2015-02-16 16:57:07', '2015-02-16 16:57:07');
INSERT INTO `t_product` VALUES ('d1cae4ab-e298-461c-85d1-bbb836be777f', '产品1', '产品1产品1产品1产品1产品1', '### 产品1', '1', '1', '5', '2015-02-16 16:54:39', '2015-02-20 13:55:38');
INSERT INTO `t_product` VALUES ('e075c59c-5e56-464f-820e-d108322eaa22', 'golang', 'golang', '**golang真好啊**', '1', '1', '6', '2015-02-10 22:18:19', '2015-03-03 11:58:23');

-- ----------------------------
-- Table structure for `t_product_img`
-- ----------------------------
DROP TABLE IF EXISTS `t_product_img`;
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

-- ----------------------------
-- Records of t_product_img
-- ----------------------------
INSERT INTO `t_product_img` VALUES ('0bae6f31-c3a0-412a-8cbd-c21f68bd61bd', null, 'e075c59c-5e56-464f-820e-d108322eaa22', '/image/4b13b5d1-7596-418f-9297-31e6b293e13d.jpg', '2', '2015-03-03 11:58:23');
INSERT INTO `t_product_img` VALUES ('2a3c3816-3ac4-4850-b043-56c80eee9c2b', null, 'd1cae4ab-e298-461c-85d1-bbb836be777f', '/image/6337e6b6-a326-4753-b9b5-6eb686ccb819.jpg', '1', '2015-02-20 13:55:38');
INSERT INTO `t_product_img` VALUES ('2af4fd9f-1eae-4b60-afc3-d8d3878f9edb', null, 'b30ff759-5702-4b80-acc6-0dade39de5a8', '/image/66f259e6-1ed1-4294-bcb6-5479d8eef67c.jpg', '2', '2015-02-16 16:55:26');
INSERT INTO `t_product_img` VALUES ('2d5bae38-1bb4-4dba-8e89-d7ccc555fb73', null, '2834101b-1ae3-4ced-b7fb-48d5bd086de6', '/image/b4f59174-0847-4542-8afa-524a18d94991.jpg', '2', '2015-02-20 15:43:29');
INSERT INTO `t_product_img` VALUES ('2dea1885-2f7e-40d0-a478-59688777dd4a', null, '2834101b-1ae3-4ced-b7fb-48d5bd086de6', '/image/cd6e04ae-768d-48f2-88c2-3585c615cbab.jpg', '1', '2015-02-20 15:43:29');
INSERT INTO `t_product_img` VALUES ('34e4226b-df40-4bc5-81cb-f97e62dfdcea', null, '730cab15-91ba-45ad-a716-e77cd3e18cdc', '/image/c01d652d-7457-49e3-b353-2ae9031d6964.jpg', '2', '2015-02-20 13:56:15');
INSERT INTO `t_product_img` VALUES ('525d4846-5f0f-4c27-baff-9ae43bb14797', null, 'b30ff759-5702-4b80-acc6-0dade39de5a8', '/image/7febea96-1962-425f-9add-b5f6896395c1.jpg', '1', '2015-02-16 16:55:26');
INSERT INTO `t_product_img` VALUES ('5f63a350-2b47-4019-b55f-3e476ec976c3', null, '2834101b-1ae3-4ced-b7fb-48d5bd086de6', '/image/2876e828-1dbe-4ef0-8e55-cf6768931436.jpg', '2', '2015-02-20 15:43:29');
INSERT INTO `t_product_img` VALUES ('6ae57245-af1f-43f6-8c8f-3894eddbbfd6', null, 'd1cae4ab-e298-461c-85d1-bbb836be777f', '/image/ebd8a936-c1a8-4cf6-a9a4-ba389c94334f.jpg', '2', '2015-02-20 13:55:38');
INSERT INTO `t_product_img` VALUES ('79cdd689-0132-452e-a359-428c36a51d06', null, 'a41bb236-226d-44b9-9ad1-9081e6d94033', '/image/1729ab3b-5c12-4c7a-ae5b-f4a94289682a.jpg', '1', '2015-03-03 11:33:20');
INSERT INTO `t_product_img` VALUES ('8c141e4b-7e1c-4c15-8e80-7b826459957c', null, '730cab15-91ba-45ad-a716-e77cd3e18cdc', '/image/acedd39e-f253-47e4-aeb9-0bc2bfd59ecb.jpg', '1', '2015-02-20 13:56:15');
INSERT INTO `t_product_img` VALUES ('99509974-5e4e-424b-b710-3737e9acc06f', null, 'cfa622f1-e33d-4e29-b0a9-f873529f1b3f', '/image/481fe9e7-e186-404a-9cc0-a3d1253ea4af.jpg', '1', '2015-02-16 16:57:06');
INSERT INTO `t_product_img` VALUES ('a26609f6-ea30-4c4e-b1ea-0781d8c980b9', null, '730cab15-91ba-45ad-a716-e77cd3e18cdc', '/image/ba944d7b-31ae-467e-9627-792ba600a3d7.jpg', '2', '2015-02-20 13:56:15');
INSERT INTO `t_product_img` VALUES ('ad372df8-7c57-436d-8bf2-bde33eca5cc2', null, 'e075c59c-5e56-464f-820e-d108322eaa22', '/image/b61fd2e3-76c0-4ce5-ab42-938e423b59d1.jpg', '2', '2015-03-03 11:58:23');
INSERT INTO `t_product_img` VALUES ('cb9e2d58-174b-4845-9381-870cdbf8e2f2', null, '730cab15-91ba-45ad-a716-e77cd3e18cdc', '/image/e21fb8ed-2d46-46f0-a8e1-f03cf6bb86fc.jpg', '2', '2015-02-20 13:56:15');
INSERT INTO `t_product_img` VALUES ('d3a99f54-155d-40b1-ae38-7e8b030d21db', null, 'cfa622f1-e33d-4e29-b0a9-f873529f1b3f', '/image/241e3024-b887-4755-bfaf-f2bcc8f0dd13.jpg', '2', '2015-02-16 16:57:07');
INSERT INTO `t_product_img` VALUES ('d872ab0f-47fd-4e83-ae12-19a2fa873d2b', null, 'e075c59c-5e56-464f-820e-d108322eaa22', '/image/5f49604d-fdeb-42be-ba46-9b1f1fee3574.jpg', '1', '2015-03-03 11:58:23');
INSERT INTO `t_product_img` VALUES ('dd99c3a0-dbec-4c7c-aa66-09652b5972a1', null, '2834101b-1ae3-4ced-b7fb-48d5bd086de6', '/image/95a0ec00-a4e6-424f-bf8b-d5f50e54b65d.jpg', '2', '2015-02-20 15:43:29');
INSERT INTO `t_product_img` VALUES ('ebc045b8-2bea-4056-90f5-df651f5ec341', null, 'b30ff759-5702-4b80-acc6-0dade39de5a8', '/image/4e0ddec6-3dce-4ef4-89ac-bc7f6439e5ed.jpg', '2', '2015-02-16 16:55:26');
INSERT INTO `t_product_img` VALUES ('ee892ddc-ecdc-4a6b-8dd5-21ec9793816f', null, 'a41bb236-226d-44b9-9ad1-9081e6d94033', '/image/56178000-13ca-40f9-8771-e1a4cdde4401.png', '2', '2015-03-03 11:33:21');

-- ----------------------------
-- Table structure for `t_user`
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
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

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES ('7450b7fb-8145-11e4-beed-005056c00008', '汤天云', '1989-02-16', 'male', 'voidint', '2014-12-09 17:15:14', '2014-12-11 22:55:08');
