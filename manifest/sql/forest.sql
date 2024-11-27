/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 90001 (9.0.1)
 Source Host           : localhost:3306
 Source Schema         : forest

 Target Server Type    : MySQL
 Target Server Version : 90001 (9.0.1)
 File Encoding         : 65001

 Date: 27/11/2024 11:04:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department`  (
  `department_id` int NOT NULL AUTO_INCREMENT COMMENT '部门Id',
  `department_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '部门名称',
  `department_parent_id` int NOT NULL COMMENT '父类部门',
  PRIMARY KEY (`department_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for department_affiche
-- ----------------------------
DROP TABLE IF EXISTS `department_affiche`;
CREATE TABLE `department_affiche`  (
  `department_affiche_id` int NOT NULL AUTO_INCREMENT COMMENT '部门公告主键',
  `department_affiche_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '部门公告标题',
  `department_affiche_content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '部门公告内容',
  `department_id` int NOT NULL COMMENT '发布部门',
  `user_id` int NOT NULL COMMENT '创建用户',
  `department_affiche_create_time` bigint NOT NULL COMMENT '创建时间',
  `department_affiche_update_time` bigint NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`department_affiche_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for forest
-- ----------------------------
DROP TABLE IF EXISTS `forest`;
CREATE TABLE `forest`  (
  `forest_id` int NOT NULL AUTO_INCREMENT COMMENT '林地主键',
  `forest_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '林地名',
  PRIMARY KEY (`forest_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for forest_resource_info
-- ----------------------------
DROP TABLE IF EXISTS `forest_resource_info`;
CREATE TABLE `forest_resource_info`  (
  `forest_resource_info_id` int NOT NULL AUTO_INCREMENT COMMENT '林地资源信息主键',
  `forest_id` int NOT NULL COMMENT '林地id',
  `tree_species_id` int NOT NULL COMMENT '树种id',
  `forest_resource_examine_time` bigint NOT NULL COMMENT '检查日期',
  `forest_resource_examine_type` int NOT NULL COMMENT '检查类型',
  `forest_status` int NOT NULL COMMENT '林地状态',
  `forest_resource_examine_info` json NOT NULL COMMENT '检测数据',
  `problem_des` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '问题描述',
  `treatment_suggestion` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '处理建议',
  `scene_photo` json NOT NULL COMMENT '现场照片',
  `scene_video` json NOT NULL COMMENT '视频记录',
  PRIMARY KEY (`forest_resource_info_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for group_chat
-- ----------------------------
DROP TABLE IF EXISTS `group_chat`;
CREATE TABLE `group_chat`  (
  `group_chat_id` int NOT NULL AUTO_INCREMENT COMMENT '部门聊天主键',
  `group_chat_send_id` int NOT NULL COMMENT '发送方Id',
  `group_chat_receiver_id` int NOT NULL COMMENT '接收方Id',
  `group_chat_create_time` bigint NOT NULL COMMENT '创建时间',
  `group_chat_message` json NOT NULL COMMENT '消息内容',
  PRIMARY KEY (`group_chat_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for identity
-- ----------------------------
DROP TABLE IF EXISTS `identity`;
CREATE TABLE `identity`  (
  `identity_id` int NOT NULL AUTO_INCREMENT COMMENT '身份ID',
  `identity_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '身份名称',
  PRIMARY KEY (`identity_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice`  (
  `notice_id` int NOT NULL AUTO_INCREMENT COMMENT '公告id',
  `notice_publisher` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '发布人',
  `notice_time` bigint NOT NULL COMMENT '发布日期',
  `notice_priority` int NOT NULL COMMENT '优先级',
  `notice_content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公告内容',
  `notice_annex` json NOT NULL COMMENT '附件信息',
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for people_application
-- ----------------------------
DROP TABLE IF EXISTS `people_application`;
CREATE TABLE `people_application`  (
  `people_application_id` int NOT NULL AUTO_INCREMENT COMMENT '民事申请ID',
  `people_application_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '民事申请标题',
  `people_applicant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '申请人',
  `people_application_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '联系电话',
  `people_application_type` int NOT NULL COMMENT '申请类型',
  `people_application_content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '申请内容',
  `people_application_annex` json NOT NULL COMMENT '附件',
  `modify_time` bigint NOT NULL COMMENT '修改时间',
  `people_application_state` int NOT NULL COMMENT '申请状态(0:待审核 1:已处理)',
  `people_application_desc` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '处理描述',
  PRIMARY KEY (`people_application_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for policy
-- ----------------------------
DROP TABLE IF EXISTS `policy`;
CREATE TABLE `policy`  (
  `policy_id` int NOT NULL AUTO_INCREMENT COMMENT '政策Id',
  `policy_headline` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '政策标题',
  `policy_department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '发布部门',
  `policy_time` bigint NOT NULL COMMENT '发布日期',
  `policy_type` int NOT NULL COMMENT '政策类型',
  `policy_content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '政策内容',
  `policy_annex` json NOT NULL COMMENT '附件',
  PRIMARY KEY (`policy_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for private_chat
-- ----------------------------
DROP TABLE IF EXISTS `private_chat`;
CREATE TABLE `private_chat`  (
  `private_chat_id` int NOT NULL AUTO_INCREMENT COMMENT '私聊id',
  `private_chat_send_id` int NOT NULL COMMENT '发送方id',
  `private_chat_receiver_id` int NOT NULL COMMENT '接收方id',
  `private_chat_create_time` bigint NOT NULL COMMENT '发送时间',
  `private_chat_message` json NOT NULL COMMENT '消息内容',
  PRIMARY KEY (`private_chat_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '私聊记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tree_species
-- ----------------------------
DROP TABLE IF EXISTS `tree_species`;
CREATE TABLE `tree_species`  (
  `tree_species_id` int NOT NULL AUTO_INCREMENT COMMENT '树种ID',
  `tree_species_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '树种名称',
  PRIMARY KEY (`tree_species_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '树种表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int NOT NULL AUTO_INCREMENT COMMENT '用户主键',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户密码',
  `department_id` int NOT NULL COMMENT '部门ID',
  `user_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '联系电话',
  `identity_id` int NOT NULL COMMENT '身份ID',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
