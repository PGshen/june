/*
 Navicat Premium Data Transfer

 Source Server         : docker-mysql
 Source Server Type    : MySQL
 Source Server Version : 50710
 Source Host           : localhost:3306
 Source Schema         : june

 Target Server Type    : MySQL
 Target Server Version : 50710
 File Encoding         : 65001

 Date: 29/08/2020 14:06:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_sys_api
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_api`;
CREATE TABLE `t_sys_api` (
  `api_id` int(11) DEFAULT NULL,
  `parent_api_id` int(11) DEFAULT NULL,
  `cascade_path` tinytext COLLATE utf8mb4_unicode_ci,
  `type` int(11) DEFAULT NULL,
  `name` tinytext COLLATE utf8mb4_unicode_ci,
  `description` tinytext COLLATE utf8mb4_unicode_ci,
  `method` tinytext COLLATE utf8mb4_unicode_ci,
  `uri` tinytext COLLATE utf8mb4_unicode_ci,
  `is_del` int(11) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `tenant_id` tinytext COLLATE utf8mb4_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='接口定义表';

-- ----------------------------
-- Records of t_sys_api
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_api` VALUES (1, 0, '[]', 0, '根', '根目录', 'GET', '', 0, '2019-07-18 05:43:09', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (2, 1, '[1]', 0, '系统管理', '系统管理', '', '', 0, '2019-07-18 05:43:09', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (3, 5, '[1,2,5]', 1, 'API接口新增', 'API接口新增', 'POST', '/api', 0, '2019-07-18 05:43:09', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (4, 5, '[1,2,5]', 1, 'API接口列表', 'API接口列表', 'POST', '/api/list', 0, '2019-07-18 05:43:09', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (5, 2, '[1,2]', 0, 'API接口', 'API接口管理', '', '', 0, '2019-07-31 05:28:52', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (6, 5, '[1,2,5]', 1, 'API接口修改', 'API接口修改', 'PUT', '/api/{id}', 0, '2019-07-31 05:33:18', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (7, 5, '[1,2,5]', 1, 'API接口删除', 'API接口删除', 'DELETE', '/api/{id}', 0, '2019-07-31 05:34:29', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (8, 5, '[1,2,5]', 1, 'API接口树', 'API接口完整树', 'GET', '/api/tree', 0, '2019-07-31 05:35:27', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (9, 5, '[1,2,5]', 1, 'API接口查找', '按ID查找API接口', 'GET', '/api/{id}', 0, '2019-07-31 05:36:22', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (10, 5, '[1,2,5]', 1, 'API接口树[ID]', '按ID返回API树', 'GET', '/api/tree/{apiId}', 0, '2019-07-31 05:37:33', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (11, 2, '[1,2]', 0, '菜单', '菜单', '', '', 0, '2019-07-31 05:38:15', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (12, 11, '[1,2,11]', 1, '菜单列表', '菜单列表', 'POST', '/api/list', 0, '2019-07-31 05:38:45', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (13, 2, '[1,2]', 0, '客户端', '客户端', '', '', 0, '2019-07-31 05:39:19', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (14, 2, '[1,2]', 0, '角色', '角色', '', '', 0, '2019-07-31 05:39:32', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (15, 2, '[1,2]', 0, '用户', '用户', '', '', 0, '2019-07-31 05:39:46', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (16, 2, '[1,2]', 0, '组织机构', '组织机构', '', '', 0, '2019-07-31 05:40:14', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (17, 2, '[1,2]', 0, '租户', '租户', '', '', 0, '2019-07-31 05:40:31', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (18, 13, '[1,2,13]', 1, '客户端列表', '客户端列表', 'POST', '/client/list', 0, '2019-07-31 05:41:09', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (19, 14, '[1,2,14]', 1, '角色列表', '角色列表', 'POST', '/role/list', 0, '2019-07-31 05:41:36', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (20, 15, '[1,2,15]', 1, '用户列表', '用户列表', 'POST', '/user/list', 0, '2019-07-31 05:42:03', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (21, 16, '[1,2,16]', 1, '组织机构列表', '组织机构列表', 'POST', '/org/list', 0, '2019-07-31 05:42:32', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (22, 17, '[1,2,17]', 1, '租户列表', '租户列表', 'POST', '/tenant/list', 0, '2019-07-31 05:42:59', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (23, 11, '[1,2,11]', 1, '菜单新增', '菜单新增', 'POST', '/menu', 0, '2019-07-31 05:43:46', '2019-08-04 03:36:50', '1');
INSERT INTO `t_sys_api` VALUES (24, 1, '[1]', 0, '运维监控', '运维监控', '', '', 0, '2019-07-31 11:59:01', '2019-08-04 03:36:50', NULL);
INSERT INTO `t_sys_api` VALUES (25, 24, '[1,24]', 0, 'API文档', 'API接口文档', '', '', 0, '2019-07-31 11:59:21', '2019-08-04 03:36:50', NULL);
INSERT INTO `t_sys_api` VALUES (26, 24, '[1,24]', 0, '链路追踪', '链路追踪', '', '', 0, '2019-07-31 11:59:49', '2019-08-04 03:36:50', NULL);
INSERT INTO `t_sys_api` VALUES (27, 24, '[1,24]', 1, '测试', '测试', 'POST', '/test', 1, '2019-07-31 12:00:22', '2019-08-04 03:36:50', NULL);
INSERT INTO `t_sys_api` VALUES (28, 24, '[1,24]', 0, 'ceshi', 'test', '', '', 1, '2019-07-31 12:11:27', '2019-08-04 03:36:50', NULL);
INSERT INTO `t_sys_api` VALUES (29, 24, '[1,24]', 0, '应用监控', '应用监控', '', '', 0, '2019-07-31 12:17:46', '2019-08-04 03:36:50', NULL);
INSERT INTO `t_sys_api` VALUES (30, 24, '[1,24]', 0, '操作日志', '操作日志', '', '', 0, '2019-08-04 03:41:33', '2019-08-04 03:43:29', NULL);
INSERT INTO `t_sys_api` VALUES (31, 30, '[1,24,30]', 1, '日志查询', '日志查询', 'POST', '/log', 1, '2019-08-04 03:46:16', '2019-08-04 03:47:55', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_sys_client
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_client`;
CREATE TABLE `t_sys_client` (
  `client_id` int(11) DEFAULT NULL,
  `app_id` tinytext COLLATE utf8mb4_unicode_ci,
  `app_secret` tinytext COLLATE utf8mb4_unicode_ci,
  `description` tinytext COLLATE utf8mb4_unicode_ci,
  `resource_ids` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `scope` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `authorized_grant_types` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `web_server_redirect_uri` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `authorities` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `access_token_validity` bigint(20) DEFAULT NULL,
  `refresh_token_validity` bigint(20) DEFAULT NULL,
  `additional_information` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `autoapprove` tinytext COLLATE utf8mb4_unicode_ci,
  `is_del` int(11) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `tenant_id` tinytext COLLATE utf8mb4_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='外部客户端表';

-- ----------------------------
-- Records of t_sys_client
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_client` VALUES (1, '2209899660', 'sdfghjyuiore7893', '测试的', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, '2019-07-31 13:08:32', '2019-07-31 13:11:55', '1');
INSERT INTO `t_sys_client` VALUES (2, '3202725393', 'gfddvbfre464532', 'Testing', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 1, '2019-07-31 13:11:42', '2019-07-31 13:15:11', NULL);
INSERT INTO `t_sys_client` VALUES (3, 'test', 'qwertyuiop', '测试客户端', '', 'server', 'password', '', '', 0, 0, '', 'false', 0, '2019-09-04 02:34:46', '2019-09-04 02:34:59', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_sys_client_api
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_client_api`;
CREATE TABLE `t_sys_client_api` (
  `client_id` int(11) DEFAULT NULL,
  `ip` tinytext COLLATE utf8mb4_unicode_ci,
  `api_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='客户端接口关联表';

-- ----------------------------
-- Records of t_sys_client_api
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_client_api` VALUES (1, '10.18.97.1', 0);
INSERT INTO `t_sys_client_api` VALUES (1, '10.18.97.1', 12);
INSERT INTO `t_sys_client_api` VALUES (1, '10.18.97.1', 23);
INSERT INTO `t_sys_client_api` VALUES (1, '102.2.2.2', 0);
INSERT INTO `t_sys_client_api` VALUES (1, '102.2.2.2', 21);
INSERT INTO `t_sys_client_api` VALUES (1, '102.2.2.2', 22);
INSERT INTO `t_sys_client_api` VALUES (1, '127.0.0.1', 0);
INSERT INTO `t_sys_client_api` VALUES (1, '127.0.0.1', 3);
INSERT INTO `t_sys_client_api` VALUES (1, '127.0.0.1', 4);
INSERT INTO `t_sys_client_api` VALUES (1, '127.0.0.1', 6);
INSERT INTO `t_sys_client_api` VALUES (1, '127.0.0.1', 7);
INSERT INTO `t_sys_client_api` VALUES (1, '127.0.0.1', 8);
INSERT INTO `t_sys_client_api` VALUES (1, '127.0.0.1', 9);
INSERT INTO `t_sys_client_api` VALUES (1, '127.0.0.1', 10);
COMMIT;

-- ----------------------------
-- Table structure for t_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_menu`;
CREATE TABLE `t_sys_menu` (
  `menu_id` int(11) DEFAULT NULL,
  `parent_menu_id` int(11) DEFAULT NULL,
  `cascade_path` tinytext COLLATE utf8mb4_unicode_ci,
  `menu_name` tinytext COLLATE utf8mb4_unicode_ci,
  `title` tinytext COLLATE utf8mb4_unicode_ci,
  `icon` tinytext COLLATE utf8mb4_unicode_ci,
  `perm` tinytext COLLATE utf8mb4_unicode_ci,
  `type` int(11) DEFAULT NULL,
  `order_num` int(11) DEFAULT NULL,
  `hidden` int(11) DEFAULT NULL,
  `always_show` int(11) DEFAULT NULL,
  `component` tinytext COLLATE utf8mb4_unicode_ci,
  `path` tinytext COLLATE utf8mb4_unicode_ci,
  `redirect` tinytext COLLATE utf8mb4_unicode_ci,
  `is_del` int(11) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `tenant_id` tinytext COLLATE utf8mb4_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单表';

-- ----------------------------
-- Records of t_sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_menu` VALUES (0, -1, '[]', 'root', '菜单树根', '', '', 0, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:19', '1');
INSERT INTO `t_sys_menu` VALUES (1, 0, '[0]', 'system', '系统管理', 'component', 'sys:system', 0, 0, 0, 0, 'layout/Layout', '/system', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (2, 1, '[0,1]', 'user', '用户管理', 'user', 'sys:user', 0, 0, 0, 0, 'system/user', 'user', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (3, 1, '[0,1]', 'role', '角色管理', 'role', 'sys:role', 0, 0, 0, 0, 'system/role', 'role', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (4, 1, '[0,1]', 'menu', '菜单管理', 'menu', 'sys:menu', 0, 0, 0, 0, 'system/menu', 'menu', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (5, 1, '[0,1]', 'api', '接口管理', 'api', 'sys:api', 0, 0, 0, 0, 'system/api', 'api', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (6, 1, '[0,1]', 'client', '客户端管理', 'client', 'sys:client', 0, 0, 0, 0, 'system/client', 'client', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (7, 2, '[0,1,2]', 'user_add', '新增用户', 'user_add', 'sys:user:add', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (8, 2, '[0,1,2]', 'user_edit', '编辑用户', 'user_edit', 'sys:user:edit', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (9, 2, '[0,1,2]', 'user_delete', '删除用户', 'user_delete', 'sys:user:delete', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (10, 3, '[0,1,3]', 'role_add', '新增角色', 'role_add', 'sys:role:add', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (11, 3, '[0,1,3]', 'role_edit', '编辑角色', 'role_edit', 'sys:role:edit', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (12, 3, '[0,1,3]', 'role_delete', '删除角色', 'role_delete', 'sys:role:delete', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (13, 4, '[0,1,4]', 'menu_add', '新增菜单', 'menu_add', 'sys:menu:add', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (14, 4, '[0,1,4]', 'menu_edit', '编辑菜单', 'menu_edit', 'sys:menu:edit', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (15, 4, '[0,1,4]', 'menu_delete', '删除菜单', 'menu_delete', 'sys:menu:delete', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (16, 4, '[0,1,4]', 'menu_ban', '禁用菜单', 'menu_ban', 'sys:menu:ban', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (17, 0, '[0]', 'maintain', '运维监控', 'example', 'maintain', 0, 0, 0, 0, 'layout/Layout', '/maintain', '', 0, '2019-07-28 13:56:58', '2019-08-03 09:21:02', '1');
INSERT INTO `t_sys_menu` VALUES (18, 17, '[0,17]', 'swagger', 'API文档', 'documentation', 'maintain:doc', 0, 0, 0, 0, 'maintain/swagger', 'swagger', '', 0, '2019-07-28 13:59:27', '2019-08-03 09:15:44', '1');
INSERT INTO `t_sys_menu` VALUES (19, 2, '[0,1,2]', 'user_ban', '禁用用户', 'user_ban', 'sys:user:ban', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (20, 3, '[0,1,3]', 'role_auth', '授权角色', 'role_auth', 'sys:role:auth', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (21, 5, '[0,1,5]', 'api_add', '新增接口', 'api', 'sys:api:add', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (22, 5, '[0,1,5]', 'api_edit', '编辑接口', 'api', 'sys:api:edit', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (23, 5, '[0,1,5]', 'api_delete', '删除接口', 'api', 'sys:api:delete', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (24, 5, '[0,1,5]', 'api_ban', '禁用接口', 'api', 'sys:api:ban', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (25, 5, '[0,1,5]', 'api_list', '接口列表', 'api', 'sys:api:list', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (26, 6, '[0,1,6]', 'client_add', '新增客户端', 'client', 'sys:client:add', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (27, 6, '[0,1,6]', 'client_edit', '编辑客户端', 'client', 'sys:client:edit', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (28, 6, '[0,1,6]', 'client_delete', '删除客户端', 'client', 'sys:client:delete', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (29, 6, '[0,1,6]', 'client_ban', '禁用客户端', 'client', 'sys:client:ban', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (30, 6, '[0,1,6]', 'client_list', '客户端列表', 'client', 'sys:client:list', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (31, 2, '[0,1,2]', 'user_list', '用户列表', 'user', 'sys:user:list', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (32, 3, '[0,1,3]', 'role_list', '角色列表', 'role', 'sys:role:list', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (33, 3, '[0,1,3]', 'role_auth', '角色授权', 'role', 'sys:role:auth', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (34, 4, '[0,1,4]', 'menu_api', '菜单关联API', 'menu', 'sys:menu:api', 2, 0, 0, 0, '', '', '', 0, '2019-07-26 00:46:58', '2019-08-03 09:27:10', '1');
INSERT INTO `t_sys_menu` VALUES (35, 17, '[0,17]', 'zipkin', '链路追踪', 'documentation', 'maintain:zipkin', 0, 0, 1, 0, 'maintain/zipkin', 'zipkin', '', 0, '2019-08-03 07:55:45', '2020-04-05 03:07:38', '1');
INSERT INTO `t_sys_menu` VALUES (36, 17, '[0,17]', 'operateLog', '操作日志', 'log', 'maintain:log', 1, 0, 0, 0, 'maintain/log', 'log', '', 0, '2019-08-03 08:09:36', '2020-04-05 03:07:38', '1');
INSERT INTO `t_sys_menu` VALUES (37, 0, '[0]', 'Test', '测试目录', 'guide', 'test', 0, 0, 0, 0, 'layout/Layout', '/test', '', 0, '2019-08-03 08:12:13', '2020-04-05 03:07:38', '1');
INSERT INTO `t_sys_menu` VALUES (38, 37, '[0,37]', 'Testing', '测试菜单', 'guide', 'test:test', 1, 0, 0, 0, 'test/test', 'test', '', 0, '2019-08-03 08:14:09', '2020-04-05 03:07:38', '1');
INSERT INTO `t_sys_menu` VALUES (39, 38, '[0,37,38]', 'testButton', '测试按钮', '', 'test:test:btn', 2, 0, 0, 0, '', '', '', 0, '2019-08-03 08:15:07', '2020-04-05 03:13:24', '1');
INSERT INTO `t_sys_menu` VALUES (40, 38, '[0,37,38]', 'testmm', 'ceshi', '', 'test:test:testmm', 2, 0, 0, 0, '', '', '', 1, '2019-08-03 09:34:40', '2020-04-05 03:13:24', '1');
INSERT INTO `t_sys_menu` VALUES (41, 36, '[0,17,36]', 'logBtn', '日志按钮', '', 'maintain:log:btn', 2, 0, 0, 0, '', '', '', 0, '2019-08-03 09:39:10', '2020-04-05 03:13:24', '1');
INSERT INTO `t_sys_menu` VALUES (42, 37, '[0,37]', 'Temp', '临时菜单', '', 'test:temp', 2, 0, 0, 0, '', 'temp', '', 0, '2019-08-03 09:44:15', '2020-04-05 03:13:24', '1');
INSERT INTO `t_sys_menu` VALUES (43, 37, '[0,37]', 'test2', 'test2', 'guide', 'test:test2', 1, 0, 0, 0, 'test/test2', 'test2', '', 0, '2019-08-03 09:56:14', '2020-04-05 03:07:38', '1');
INSERT INTO `t_sys_menu` VALUES (44, 6, '[0,1,6]', 'client_api', '客户端关联API', '', 'sys:client:api', 2, 0, 1, 0, '', '', '', 0, '2019-08-04 04:35:23', '2020-04-05 03:07:38', '1');
INSERT INTO `t_sys_menu` VALUES (45, 6, '[0,1,6]', 'client_ip_add', '客户端新增IP', '', 'sys:client:ip:add', 2, 0, 1, 0, '', '', '', 0, '2019-08-04 05:11:14', '2020-04-05 03:07:38', '1');
INSERT INTO `t_sys_menu` VALUES (46, 6, '[0,1,6]', 'client_ip_del', '客户端删除IP', '', 'sys:client:ip:delete', 2, 0, 1, 0, '', '', '', 0, '2019-08-04 05:11:51', '2020-04-05 03:07:38', '1');
COMMIT;

-- ----------------------------
-- Table structure for t_sys_menu_api
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_menu_api`;
CREATE TABLE `t_sys_menu_api` (
  `menu_id` int(11) DEFAULT NULL,
  `api_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单API接口关联表';

-- ----------------------------
-- Records of t_sys_menu_api
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_menu_api` VALUES (1, 3);
INSERT INTO `t_sys_menu_api` VALUES (1, 4);
INSERT INTO `t_sys_menu_api` VALUES (1, 6);
INSERT INTO `t_sys_menu_api` VALUES (1, 8);
INSERT INTO `t_sys_menu_api` VALUES (1, 10);
INSERT INTO `t_sys_menu_api` VALUES (37, 12);
INSERT INTO `t_sys_menu_api` VALUES (37, 23);
COMMIT;

-- ----------------------------
-- Table structure for t_sys_org
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_org`;
CREATE TABLE `t_sys_org` (
  `org_id` int(11) DEFAULT NULL,
  `parent_org_id` int(11) DEFAULT NULL,
  `org_name` tinytext COLLATE utf8mb4_unicode_ci,
  `remark` tinytext COLLATE utf8mb4_unicode_ci,
  `order_num` int(11) DEFAULT NULL,
  `is_del` int(11) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `tenant_id` tinytext COLLATE utf8mb4_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='组织机构表';

-- ----------------------------
-- Table structure for t_sys_role
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_role`;
CREATE TABLE `t_sys_role` (
  `role_id` int(11) DEFAULT NULL,
  `role_name` tinytext COLLATE utf8mb4_unicode_ci,
  `alias` tinytext COLLATE utf8mb4_unicode_ci,
  `description` tinytext COLLATE utf8mb4_unicode_ci,
  `is_del` int(11) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `tenant_id` tinytext COLLATE utf8mb4_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- ----------------------------
-- Records of t_sys_role
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_role` VALUES (1, 'admin', 'superAdmin', '超级管理员', 0, '2019-07-28 05:47:57', '2019-07-28 05:47:57', '1');
INSERT INTO `t_sys_role` VALUES (2, 'test', 'superTest', '超级测试', 0, '2019-07-28 05:47:57', '2019-07-28 05:47:57', '1');
INSERT INTO `t_sys_role` VALUES (5, 'temp', '临时', '临时角色', 1, '2019-08-03 04:08:30', '2019-08-03 04:08:35', NULL);
INSERT INTO `t_sys_role` VALUES (6, 'temp', '临时', '临时的', 0, '2019-08-03 04:17:48', '2019-08-03 04:17:48', NULL);
INSERT INTO `t_sys_role` VALUES (7, 'dev', '开发', '开发的', 1, '2019-08-03 04:18:56', '2019-08-03 04:19:01', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_role_menu`;
CREATE TABLE `t_sys_role_menu` (
  `role_id` int(11) DEFAULT NULL,
  `menu_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色菜单关联表';

-- ----------------------------
-- Records of t_sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_role_menu` VALUES (1, 1);
INSERT INTO `t_sys_role_menu` VALUES (1, 2);
INSERT INTO `t_sys_role_menu` VALUES (1, 3);
INSERT INTO `t_sys_role_menu` VALUES (1, 4);
INSERT INTO `t_sys_role_menu` VALUES (1, 5);
INSERT INTO `t_sys_role_menu` VALUES (1, 6);
INSERT INTO `t_sys_role_menu` VALUES (1, 7);
INSERT INTO `t_sys_role_menu` VALUES (1, 8);
INSERT INTO `t_sys_role_menu` VALUES (1, 9);
INSERT INTO `t_sys_role_menu` VALUES (1, 10);
INSERT INTO `t_sys_role_menu` VALUES (1, 11);
INSERT INTO `t_sys_role_menu` VALUES (1, 12);
INSERT INTO `t_sys_role_menu` VALUES (1, 13);
INSERT INTO `t_sys_role_menu` VALUES (1, 14);
INSERT INTO `t_sys_role_menu` VALUES (1, 15);
INSERT INTO `t_sys_role_menu` VALUES (1, 16);
INSERT INTO `t_sys_role_menu` VALUES (1, 18);
INSERT INTO `t_sys_role_menu` VALUES (1, 19);
INSERT INTO `t_sys_role_menu` VALUES (1, 20);
INSERT INTO `t_sys_role_menu` VALUES (1, 21);
INSERT INTO `t_sys_role_menu` VALUES (1, 22);
INSERT INTO `t_sys_role_menu` VALUES (1, 23);
INSERT INTO `t_sys_role_menu` VALUES (1, 24);
INSERT INTO `t_sys_role_menu` VALUES (1, 25);
INSERT INTO `t_sys_role_menu` VALUES (1, 26);
INSERT INTO `t_sys_role_menu` VALUES (1, 27);
INSERT INTO `t_sys_role_menu` VALUES (1, 28);
INSERT INTO `t_sys_role_menu` VALUES (1, 29);
INSERT INTO `t_sys_role_menu` VALUES (1, 30);
INSERT INTO `t_sys_role_menu` VALUES (1, 31);
INSERT INTO `t_sys_role_menu` VALUES (1, 32);
INSERT INTO `t_sys_role_menu` VALUES (1, 33);
INSERT INTO `t_sys_role_menu` VALUES (1, 34);
INSERT INTO `t_sys_role_menu` VALUES (1, 35);
INSERT INTO `t_sys_role_menu` VALUES (1, 36);
INSERT INTO `t_sys_role_menu` VALUES (1, 41);
INSERT INTO `t_sys_role_menu` VALUES (1, 44);
INSERT INTO `t_sys_role_menu` VALUES (1, 45);
INSERT INTO `t_sys_role_menu` VALUES (1, 46);
INSERT INTO `t_sys_role_menu` VALUES (1, 47);
INSERT INTO `t_sys_role_menu` VALUES (1, 48);
INSERT INTO `t_sys_role_menu` VALUES (1, 49);
INSERT INTO `t_sys_role_menu` VALUES (1, 50);
INSERT INTO `t_sys_role_menu` VALUES (1, 51);
INSERT INTO `t_sys_role_menu` VALUES (1, 52);
INSERT INTO `t_sys_role_menu` VALUES (1, 53);
INSERT INTO `t_sys_role_menu` VALUES (1, 54);
INSERT INTO `t_sys_role_menu` VALUES (1, 55);
INSERT INTO `t_sys_role_menu` VALUES (1, 56);
INSERT INTO `t_sys_role_menu` VALUES (1, 57);
INSERT INTO `t_sys_role_menu` VALUES (1, 58);
INSERT INTO `t_sys_role_menu` VALUES (1, 59);
INSERT INTO `t_sys_role_menu` VALUES (1, 60);
INSERT INTO `t_sys_role_menu` VALUES (1, 61);
INSERT INTO `t_sys_role_menu` VALUES (1, 62);
INSERT INTO `t_sys_role_menu` VALUES (1, 63);
INSERT INTO `t_sys_role_menu` VALUES (1, 64);
INSERT INTO `t_sys_role_menu` VALUES (2, 1);
INSERT INTO `t_sys_role_menu` VALUES (2, 2);
INSERT INTO `t_sys_role_menu` VALUES (2, 3);
INSERT INTO `t_sys_role_menu` VALUES (2, 17);
INSERT INTO `t_sys_role_menu` VALUES (2, 18);
INSERT INTO `t_sys_role_menu` VALUES (6, 1);
INSERT INTO `t_sys_role_menu` VALUES (6, 17);
COMMIT;

-- ----------------------------
-- Table structure for t_sys_tenant
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_tenant`;
CREATE TABLE `t_sys_tenant` (
  `tenant_id` int(11) DEFAULT NULL,
  `tenant_name` tinytext COLLATE utf8mb4_unicode_ci,
  `tenant_code` tinytext COLLATE utf8mb4_unicode_ci,
  `description` tinytext COLLATE utf8mb4_unicode_ci,
  `is_enable` int(11) DEFAULT NULL,
  `is_del` int(11) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='租户表';

-- ----------------------------
-- Records of t_sys_tenant
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_tenant` VALUES (1, 'Testing', '1', '测试测试呀', 0, 0, '2019-07-23 11:00:25', '2019-07-23 11:02:40');
INSERT INTO `t_sys_tenant` VALUES (2, 'Niubi Company', '2', '牛逼哄哄', 0, 1, '2019-07-24 11:00:25', '2019-07-24 08:34:49');
COMMIT;

-- ----------------------------
-- Table structure for t_sys_user
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_user`;
CREATE TABLE `t_sys_user` (
  `user_id` int(11) DEFAULT NULL,
  `login_name` tinytext COLLATE utf8mb4_unicode_ci,
  `password` tinytext COLLATE utf8mb4_unicode_ci,
  `name` tinytext COLLATE utf8mb4_unicode_ci,
  `email` tinytext COLLATE utf8mb4_unicode_ci,
  `phone` tinytext COLLATE utf8mb4_unicode_ci,
  `avatar` varchar(512) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remark` tinytext COLLATE utf8mb4_unicode_ci,
  `is_enable` int(11) DEFAULT NULL,
  `is_del` int(11) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `tenant_id` tinytext COLLATE utf8mb4_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- ----------------------------
-- Records of t_sys_user
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_user` VALUES (1, '18899009900', '123456', 'admin', 'admin@september.space', '18800990088', 'https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=422175229,193945588&fm=11&gp=0.jpg', '备注备注。。。', 1, 0, '2019-07-27 07:22:51', '2020-04-02 06:22:45', '1');
INSERT INTO `t_sys_user` VALUES (2, '18810102020', '12', '测试', 'test@test.com', '19927710022', 'https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=422175229,193945588&fm=11&gp=0.jpg', '测试测试', 1, 1, '2019-08-01 15:02:47', '2020-04-09 03:02:24', NULL);
INSERT INTO `t_sys_user` VALUES (3, 'test', '123456', 'test', '123123@126.com', '12314234235', 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=1970997771,2057738770&fm=26&gp=0.jpg', NULL, 1, 1, '2019-08-01 15:06:51', '2020-04-09 03:02:14', NULL);
INSERT INTO `t_sys_user` VALUES (4, '19929293939', '12', '123123', '123@123.cn', '19900002222', '', '淡定淡定', 0, 1, '2019-08-01 15:14:07', '2020-04-09 03:02:26', NULL);
INSERT INTO `t_sys_user` VALUES (5, '18090909090', '12', '132', '123123@88.com', '123123', '', '4321', 0, 1, '2019-08-01 15:20:34', '2020-04-02 06:26:08', NULL);
INSERT INTO `t_sys_user` VALUES (6, '18899009999', '123456', '432', '321', '2334', '', '', 0, 1, '2019-08-01 15:21:59', '2019-08-02 05:07:41', NULL);
INSERT INTO `t_sys_user` VALUES (7, 'root', '2B54867A33C94AA773F9A21DDA47424C', 'admin', 'admin@exhi.com', NULL, '19919192929', '超级账号', 0, 0, '2019-09-05 06:52:40', '2019-09-08 15:33:46', NULL);
INSERT INTO `t_sys_user` VALUES (8, '19929293939', '200820E3227815ED1756A6B531E7E0D2', '测试', '19929293939@189.cn', '19929293939', 'https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=422175229,193945588&fm=11&gp=0.jpg', '', 1, 0, '2020-04-09 03:03:28', '2020-04-09 03:04:13', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_user_role`;
CREATE TABLE `t_sys_user_role` (
  `user_id` int(11) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- ----------------------------
-- Records of t_sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `t_sys_user_role` VALUES (1, 1);
INSERT INTO `t_sys_user_role` VALUES (2, 2);
INSERT INTO `t_sys_user_role` VALUES (3, 2);
INSERT INTO `t_sys_user_role` VALUES (4, 2);
INSERT INTO `t_sys_user_role` VALUES (4, 6);
INSERT INTO `t_sys_user_role` VALUES (8, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
