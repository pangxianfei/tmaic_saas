/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : tenanta

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2022-03-04 02:17:28
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
INSERT INTO `tmaic_user` VALUES ('1', '1', 'pangfei-A', '421339244@qq.com', '$2a$10$SJTsJ0b50quSwXrmJxBb5uiaHbZF8rSVR3xd8zBVouZWqpDfacYqm', '2021-09-24 17:30:33', '2021-09-24 17:30:33', null);
