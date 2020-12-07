/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50644
Source Host           : localhost:3306
Source Database       : www_demo_dev

Target Server Type    : MYSQL
Target Server Version : 50644
File Encoding         : 65001

Date: 2020-12-07 17:17:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tmaic_article
-- ----------------------------
DROP TABLE IF EXISTS `tmaic_article`;
CREATE TABLE `tmaic_article` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT NULL,
  `body` varchar(100) NOT NULL,
  `slug` varchar(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tmaic_article
-- ----------------------------
INSERT INTO `tmaic_article` VALUES ('12', '这是标题', '这是内容', '这是标题', '2020-12-02 22:29:03', '2020-12-02 22:29:03', null);
INSERT INTO `tmaic_article` VALUES ('13', '这是标题', '这是内容', '这是标题', '2020-12-02 23:20:58', '2020-12-02 23:20:58', null);
INSERT INTO `tmaic_article` VALUES ('14', '这是标题', '这是内容', '这是标题', '2020-12-02 23:21:01', '2020-12-02 23:21:01', null);
INSERT INTO `tmaic_article` VALUES ('15', '这是标题', '这是内容', '这是标题', '2020-12-03 10:40:55', '2020-12-03 10:40:55', null);
INSERT INTO `tmaic_article` VALUES ('16', '这是标题', '这是内容', '这是标题', '2020-12-03 10:40:57', '2020-12-03 10:40:57', null);
INSERT INTO `tmaic_article` VALUES ('17', '这是标题', '这是内容', '这是标题', '2020-12-03 10:40:59', '2020-12-03 10:40:59', null);
INSERT INTO `tmaic_article` VALUES ('18', '这是标题', '这是内容', '这是标题', '2020-12-03 10:41:23', '2020-12-03 10:41:23', null);

-- ----------------------------
-- Table structure for tmaic_failed_queue
-- ----------------------------
DROP TABLE IF EXISTS `tmaic_failed_queue`;
CREATE TABLE `tmaic_failed_queue` (
  `failed_queue_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `failed_queue_hash` varchar(100) NOT NULL,
  `failed_queue_topic_name` varchar(100) NOT NULL,
  `failed_queue_channel_name` varchar(100) NOT NULL,
  `failed_queue_data` varbinary(2048) DEFAULT NULL,
  `failed_queue_pushed_at` datetime NOT NULL,
  `failed_queue_delay` bigint(20) unsigned NOT NULL,
  `failed_queue_retries` int(10) unsigned NOT NULL,
  `failed_queue_tried` int(10) unsigned NOT NULL,
  `failed_queue_err` longtext,
  `failed_queue_created_at` datetime DEFAULT NULL,
  `failed_queue_updated_at` datetime DEFAULT NULL,
  `failed_queue_deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`failed_queue_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tmaic_failed_queue
-- ----------------------------

-- ----------------------------
-- Table structure for tmaic_migrations
-- ----------------------------
DROP TABLE IF EXISTS `tmaic_migrations`;
CREATE TABLE `tmaic_migrations` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) DEFAULT NULL,
  `batch` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tmaic_migrations
-- ----------------------------
INSERT INTO `tmaic_migrations` VALUES ('10', 'CreateUserTable1548750742', '2');
INSERT INTO `tmaic_migrations` VALUES ('11', 'CreateUserAffiliationTable1553678539', '2');
INSERT INTO `tmaic_migrations` VALUES ('12', 'CreateFailedQueueTable1556612225', '2');

-- ----------------------------
-- Table structure for tmaic_user
-- ----------------------------
DROP TABLE IF EXISTS `tmaic_user`;
CREATE TABLE `tmaic_user` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) DEFAULT NULL,
  `user_email` varchar(100) NOT NULL,
  `user_password` varchar(100) NOT NULL,
  `user_created_at` datetime DEFAULT NULL,
  `user_updated_at` datetime DEFAULT NULL,
  `user_deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tmaic_user
-- ----------------------------
INSERT INTO `tmaic_user` VALUES ('4', '421339258@qq.com', '421339258@qq.com', '$2a$10$bFnd/KDXs6cufSFvWVxkye1NSiy95ji5TTNNfQm4gCNXkAsf8vM6S', '2020-12-01 20:29:20', '2020-12-01 20:29:20', null);
