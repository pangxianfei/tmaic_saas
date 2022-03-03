/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : tenantc

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2022-03-04 02:06:20
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tmaic_user
-- ----------------------------
DROP TABLE IF EXISTS `tmaic_user`;
CREATE TABLE `tmaic_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tenants_id` int(11) DEFAULT NULL,
  `user_name` varchar(100) DEFAULT NULL,
  `user_email` varchar(100) NOT NULL,
  `user_password` varchar(100) NOT NULL,
  `user_created_at` datetime DEFAULT NULL,
  `user_updated_at` datetime DEFAULT NULL,
  `user_deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of tmaic_user
-- ----------------------------
INSERT INTO `tmaic_user` VALUES ('3', '3', 'pangxianfei', '421339246@qq.com', '$2a$10$DWDZ9SCO14kuXmVFz2euI.PWzU79QXPVgQut9fAHKq9WAGYfW.Q/G', '2021-12-31 15:11:34', '2021-12-31 15:11:34', null);
