/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : www_demo_dev

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2022-03-04 02:06:26
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tmaic_migrations
-- ----------------------------
DROP TABLE IF EXISTS `tmaic_migrations`;
CREATE TABLE `tmaic_migrations` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) DEFAULT NULL,
  `batch` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of tmaic_migrations
-- ----------------------------
INSERT INTO `tmaic_migrations` VALUES ('1', 'CreateUserTable1548750742', '1');
INSERT INTO `tmaic_migrations` VALUES ('2', 'CreateUserAffiliationTable1553678539', '1');
INSERT INTO `tmaic_migrations` VALUES ('3', 'CreateFailedQueueTable1556612225', '1');

-- ----------------------------
-- Table structure for tmaic_tenants_db
-- ----------------------------
DROP TABLE IF EXISTS `tmaic_tenants_db`;
CREATE TABLE `tmaic_tenants_db` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `tenants_id` int(10) unsigned NOT NULL,
  `host` varchar(255) NOT NULL,
  `drivername` varchar(255) NOT NULL,
  `port` int(10) unsigned NOT NULL,
  `prefix` varchar(255) NOT NULL,
  `dbname` varchar(255) NOT NULL,
  `dbuser` varchar(255) NOT NULL,
  `charset` varchar(255) NOT NULL,
  `collation` varchar(255) NOT NULL,
  `setmaxIdleconns` int(10) unsigned NOT NULL,
  `setmaxopenconns` int(10) unsigned NOT NULL,
  `setconnmaxlifetime` int(10) unsigned NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of tmaic_tenants_db
-- ----------------------------
INSERT INTO `tmaic_tenants_db` VALUES ('1', '1', '1', '127.0.0.1', 'mysql', '3306', 'tmaic_', 'tenantb', 'tenantb', 'utf8mb4', 'utf8mb4_general_ci', '1', '1', '60', 'AtfLzLhTGwk8mfpD', '2022-03-03 15:12:44', '2022-03-03 15:13:26', null);
INSERT INTO `tmaic_tenants_db` VALUES ('3', '3', '3', '127.0.0.1', 'mysql', '3306', 'tmaic_', 'tenantc', 'tenantc', 'utf8mb4', 'utf8mb4_general_ci', '1', '1', '60', 'PbxRsHeGYdABB3BG', '2022-03-03 15:12:48', '2022-03-03 15:12:48', null);

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
INSERT INTO `tmaic_user` VALUES ('1', '1', 'b2319ab8665eff3a0c1e001ecc271fcf', '421339244@qq.com', '$2a$10$SJTsJ0b50quSwXrmJxBb5uiaHbZF8rSVR3xd8zBVouZWqpDfacYqm', '2021-09-24 17:30:33', '2021-09-24 17:30:33', null);
INSERT INTO `tmaic_user` VALUES ('2', '2', 'c05f766373d9be68bbdfc8627a6dd1f5', '421339245@qq.com', '$2a$10$UF3YgjCtVJy7x7XL7C5WZudmJldNPEyhtkqOK7UDEPfm39dmpb0pK', '2021-12-07 16:51:27', '2021-12-07 16:51:27', null);
INSERT INTO `tmaic_user` VALUES ('3', '3', '1c9db57a335c9181a46f4cb6c393ada6', '421339246@qq.com', '$2a$10$DWDZ9SCO14kuXmVFz2euI.PWzU79QXPVgQut9fAHKq9WAGYfW.Q/G', '2021-12-31 15:11:34', '2021-12-31 15:11:34', null);

-- ----------------------------
-- Table structure for tmaic_user_affiliation
-- ----------------------------
DROP TABLE IF EXISTS `tmaic_user_affiliation`;
CREATE TABLE `tmaic_user_affiliation` (
  `user_id` int(10) unsigned NOT NULL,
  `uaff_code` varchar(32) NOT NULL,
  `uaff_from_code` varchar(32) DEFAULT NULL,
  `uaff_root_id` int(10) unsigned DEFAULT NULL,
  `uaff_parent_id` int(10) unsigned DEFAULT NULL,
  `uaff_left_id` int(10) unsigned NOT NULL,
  `uaff_right_id` int(10) unsigned NOT NULL,
  `uaff_level` int(10) unsigned NOT NULL,
  `user_created_at` datetime DEFAULT NULL,
  `user_updated_at` datetime DEFAULT NULL,
  `user_deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of tmaic_user_affiliation
-- ----------------------------
