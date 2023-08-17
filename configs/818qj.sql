-- Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file. The original repo for
-- this file is https://github.com/auroralzdf/apis.

/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : 192.168.36.2:3306
 Source Schema         : 818qj

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 08/08/2023 15:14:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_accesses
-- ----------------------------
DROP TABLE IF EXISTS `admin_accesses`;
CREATE TABLE `admin_accesses` (
  `uid` int(11) unsigned NOT NULL,
  `group_id` int(11) unsigned NOT NULL,
  UNIQUE KEY `uid_group_id` (`uid`,`group_id`) USING BTREE,
  KEY `uid` (`uid`) USING BTREE,
  KEY `group_id` (`group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限组与用户对应关系表';

-- ----------------------------
-- Table structure for admin_groups
-- ----------------------------
DROP TABLE IF EXISTS `admin_groups`;
CREATE TABLE `admin_groups` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` char(100) NOT NULL DEFAULT '' COMMENT '名称',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0 启用 1 禁用',
  `rules` text NOT NULL COMMENT '授权相关的规则组合',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `title` (`title`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员权限组表';

-- ----------------------------
-- Table structure for admin_logs
-- ----------------------------
DROP TABLE IF EXISTS `admin_logs`;
CREATE TABLE `admin_logs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '操作者UID',
  `ip` varchar(16) NOT NULL DEFAULT '' COMMENT '操作者IP',
  `content` text NOT NULL COMMENT '日志内容',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '操作的URL',
  `request` text NOT NULL COMMENT '请求信息内容 post、get等数据',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid` (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统管理员操作日志表';

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员用户ID',
  `name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '管理员姓名',
  `avatar` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '管理员头像地址',
  `phone` char(11) COLLATE utf8mb4_bin NOT NULL COMMENT '管理员手机号码、同时也是登录账号',
  `password` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '登录密码',
  `remember_token` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '登录记住账户token',
  `status` tinyint(1) unsigned NOT NULL COMMENT '账号状态 0 正常 1 被禁用',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for areas
-- ----------------------------
DROP TABLE IF EXISTS `areas`;
CREATE TABLE `areas` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '省份、城市、区县、街道等名称',
  `pid` int(11) unsigned NOT NULL COMMENT '父级ID',
  `level` tinyint(2) unsigned NOT NULL COMMENT '层级',
  `status` tinyint(2) unsigned NOT NULL COMMENT '是否启用 0 启用 1 禁用',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE,
  KEY `level` (`level`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区域配置表';

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '所属类目名称',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
  `is_hidden` tinyint(11) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `icon_active` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '图标选中的样式',
  `icon` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '图标',
  `cover` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '封面',
  `cover_mobile` varchar(200) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机端分类图片地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='众包平台 - 所属类目表';

-- ----------------------------
-- Table structure for configs
-- ----------------------------
DROP TABLE IF EXISTS `configs`;
CREATE TABLE `configs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `k` varchar(100) NOT NULL COMMENT '变量',
  `v` varchar(255) NOT NULL COMMENT '值',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0系统，1自定义',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '说明',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `k` (`k`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='站点配置表';

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级菜单ID',
  `url` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '菜单URL',
  `title` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '菜单名称',
  `component` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '组件路径',
  `icon` varchar(30) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '菜单ICON图标',
  `is_show` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否在操作菜单列表显示 0 显示 1 不显示',
  `level` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '菜单层级',
  `sort` int(10) unsigned NOT NULL DEFAULT '50' COMMENT '排序 越低的越排前面',
  `order` varchar(200) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '菜单节点排序',
  `type` varchar(45) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `tips` text COLLATE utf8mb4_bin COMMENT '提示',
  `status` tinyint(3) unsigned DEFAULT '0' COMMENT '菜单是否可用 0 可用 1不可用',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `pid_title` (`pid`,`title`) USING BTREE,
  KEY `is_show` (`is_show`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE,
  KEY `sort` (`sort`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for places
-- ----------------------------
DROP TABLE IF EXISTS `places`;
CREATE TABLE `places` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for general_configs
-- ----------------------------
DROP TABLE IF EXISTS `general_configs`;
CREATE TABLE `general_configs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '键',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态：0=》启用；1=》禁用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='818站点通用配置';
