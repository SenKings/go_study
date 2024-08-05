/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3306
 Source Schema         : go_study

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 02/08/2024 19:59:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'kkkk', '123@qq.com');
INSERT INTO `users` VALUES (2, 'test1', 'lalala@126.com');
INSERT INTO `users` VALUES (3, '老王', '6666@gmail.com');
INSERT INTO `users` VALUES (4, 'sen', '11868@gmail.com');
INSERT INTO `users` VALUES (12, 'test6', 'test6@foxconn.com');
INSERT INTO `users` VALUES (13, 'test7', 'test7@foxconn.com');
INSERT INTO `users` VALUES (14, '桀哥', '没有邮箱');

SET FOREIGN_KEY_CHECKS = 1;
