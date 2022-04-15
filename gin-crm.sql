/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : 127.0.0.1:3306
 Source Schema         : gin-crm

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 15/04/2022 23:29:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for g_access
-- ----------------------------
DROP TABLE IF EXISTS `g_access`;
CREATE TABLE `g_access` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `module_name` varchar(128) NOT NULL DEFAULT '' COMMENT '模块名',
  `type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '类型,1模块,2菜单,3操作',
  `url` varchar(128) NOT NULL DEFAULT '' COMMENT '跳转地址',
  `module_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属模块id',
  `sort` int(11) unsigned NOT NULL DEFAULT '100' COMMENT '排序,默认倒叙',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '0禁用 1启用',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

-- ----------------------------
-- Records of g_access
-- ----------------------------
BEGIN;
INSERT INTO `g_access` VALUES (1, '菜单管理', 1, '', 0, 100, 1, '2022-04-13 22:10:41', '2022-04-13 22:36:28');
INSERT INTO `g_access` VALUES (2, '菜单列表', 2, '/admin/access', 1, 100, 1, '2022-04-13 22:11:13', '2022-04-13 22:36:37');
INSERT INTO `g_access` VALUES (3, '添加权限菜单', 3, '/admin/access/add', 1, 100, 1, '2022-04-13 22:12:22', '2022-04-13 22:12:40');
INSERT INTO `g_access` VALUES (4, '编辑权限菜单', 3, '/admin/access/edit', 1, 100, 1, '2022-04-13 22:13:35', '2022-04-13 22:13:35');
INSERT INTO `g_access` VALUES (5, '删除权限菜单', 3, '/admin/access/delete', 1, 100, 1, '2022-04-13 22:14:14', '2022-04-13 22:14:14');
INSERT INTO `g_access` VALUES (6, '管理员管理', 1, '/admin/manager', 0, 100, 1, '2022-04-13 22:16:13', '2022-04-13 22:37:59');
INSERT INTO `g_access` VALUES (7, '添加管理员', 3, '/admin/manager/add', 6, 100, 1, '2022-04-13 22:17:10', '2022-04-13 22:38:20');
INSERT INTO `g_access` VALUES (8, '编辑管理员', 3, '/admin/manager/edit', 6, 100, 1, '2022-04-13 22:17:26', '2022-04-13 22:38:21');
INSERT INTO `g_access` VALUES (9, '删除管理员', 3, '/admin/manager/delete', 6, 100, 1, '2022-04-13 22:17:46', '2022-04-13 22:38:24');
INSERT INTO `g_access` VALUES (10, '角色管理', 1, '/admin/role', 0, 100, 1, '2022-04-13 22:19:14', '2022-04-13 22:38:31');
INSERT INTO `g_access` VALUES (11, '添加角色', 3, '/admin/role/add', 10, 100, 1, '2022-04-13 22:20:12', '2022-04-13 22:38:37');
INSERT INTO `g_access` VALUES (12, '编辑角色', 3, '/admin/role/edit', 10, 100, 1, '2022-04-13 22:20:27', '2022-04-13 22:38:39');
INSERT INTO `g_access` VALUES (13, '删除角色', 3, '/admin/role/delete', 10, 100, 1, '2022-04-13 22:20:43', '2022-04-13 22:38:40');
INSERT INTO `g_access` VALUES (14, '角色授权', 3, '/admin/auth', 10, 100, 1, '2022-04-13 22:21:09', '2022-04-13 22:38:41');
INSERT INTO `g_access` VALUES (15, '执行授权', 3, '/admin/auth/do', 10, 100, 1, '2022-04-13 22:21:31', '2022-04-13 22:38:43');
INSERT INTO `g_access` VALUES (16, '角色列表', 2, '/admin/role', 10, 100, 1, '2022-04-13 22:37:03', '2022-04-13 22:39:05');
INSERT INTO `g_access` VALUES (17, '管理员列表', 2, '', 6, 100, 1, '2022-04-13 22:39:27', '2022-04-13 22:39:27');
COMMIT;

-- ----------------------------
-- Table structure for g_manager
-- ----------------------------
DROP TABLE IF EXISTS `g_manager`;
CREATE TABLE `g_manager` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `nick_name` varchar(255) NOT NULL DEFAULT '' COMMENT '管理员名称',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '启用状态,1启用,0禁用',
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '性别,1男,2女',
  `is_super` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否超级管理员,1是,0否',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_name_password` (`nick_name`,`password`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Records of g_manager
-- ----------------------------
BEGIN;
INSERT INTO `g_manager` VALUES (1, 1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '18810981789', '18810981789@163.com', 1, 1, 1, '2022-04-14 18:18:05', '2022-04-14 20:43:22');
COMMIT;

-- ----------------------------
-- Table structure for g_role
-- ----------------------------
DROP TABLE IF EXISTS `g_role`;
CREATE TABLE `g_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '角色名称',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '启用状态,1启用,0禁用',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Records of g_role
-- ----------------------------
BEGIN;
INSERT INTO `g_role` VALUES (1, '技术部', 1, '2022-04-14 18:16:43', '2022-04-14 18:16:43');
INSERT INTO `g_role` VALUES (2, '产品部', 1, '2022-04-14 22:29:53', '2022-04-14 22:29:53');
INSERT INTO `g_role` VALUES (3, '运营部', 1, '2022-04-14 22:29:53', '2022-04-14 22:29:53');
COMMIT;

-- ----------------------------
-- Table structure for g_role_access
-- ----------------------------
DROP TABLE IF EXISTS `g_role_access`;
CREATE TABLE `g_role_access` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `access_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限id',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限表';

-- ----------------------------
-- Records of g_role_access
-- ----------------------------
BEGIN;
COMMIT;


SET FOREIGN_KEY_CHECKS = 1;
