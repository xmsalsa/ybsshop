-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: 2021-04-13 16:37:53
-- 服务器版本： 5.7.19
-- PHP Version: 7.2.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `shop`
--

-- --------------------------------------------------------

--
-- 表的结构 `ybs_casbin_rule`
--

CREATE TABLE `ybs_casbin_rule` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `p_type` varchar(40) DEFAULT NULL,
  `v0` varchar(40) DEFAULT NULL,
  `v1` varchar(40) DEFAULT NULL,
  `v2` varchar(40) DEFAULT NULL,
  `v3` varchar(40) DEFAULT NULL,
  `v4` varchar(40) DEFAULT NULL,
  `v5` varchar(40) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `ybs_casbin_rule`
--

INSERT INTO `ybs_casbin_rule` (`id`, `p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
(23, 'g', '1', '1', '', '', '', ''),
(4, 'p', '1', '/api/v1/admin/change_avatar', 'POST', '', '', ''),
(2, 'p', '1', '/api/v1/admin/clear', 'GET', '', '', ''),
(6, 'p', '1', '/api/v1/admin/dashboard', 'GET', '', '', ''),
(3, 'p', '1', '/api/v1/admin/expire', 'GET', '', '', ''),
(1, 'p', '1', '/api/v1/admin/logout', 'GET', '', '', ''),
(17, 'p', '1', '/api/v1/admin/perms', 'GET', '', '', ''),
(19, 'p', '1', '/api/v1/admin/perms', 'POST', '', '', ''),
(21, 'p', '1', '/api/v1/admin/perms/{id:uint}', 'DELETE', '', '', ''),
(18, 'p', '1', '/api/v1/admin/perms/{id:uint}', 'GET', '', '', ''),
(20, 'p', '1', '/api/v1/admin/perms/{id:uint}', 'POST', '', '', ''),
(5, 'p', '1', '/api/v1/admin/profile', 'GET', '', '', ''),
(12, 'p', '1', '/api/v1/admin/roles', 'GET', '', '', ''),
(14, 'p', '1', '/api/v1/admin/roles', 'POST', '', '', ''),
(16, 'p', '1', '/api/v1/admin/roles/{id:uint}', 'DELETE', '', '', ''),
(13, 'p', '1', '/api/v1/admin/roles/{id:uint}', 'GET', '', '', ''),
(15, 'p', '1', '/api/v1/admin/roles/{id:uint}', 'POST', '', '', ''),
(22, 'p', '1', '/api/v1/admin/upload_file', 'POST', '', '', ''),
(7, 'p', '1', '/api/v1/admin/users', 'GET', '', '', ''),
(9, 'p', '1', '/api/v1/admin/users', 'POST', '', '', ''),
(11, 'p', '1', '/api/v1/admin/users/{id:uint}', 'DELETE', '', '', ''),
(8, 'p', '1', '/api/v1/admin/users/{id:uint}', 'GET', '', '', ''),
(10, 'p', '1', '/api/v1/admin/users/{id:uint}', 'POST', '', '', '');

-- --------------------------------------------------------

--
-- 表的结构 `ybs_clients`
--

CREATE TABLE `ybs_clients` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'id',
  `merch_id` int(11) NOT NULL DEFAULT '0',
  `uid` int(11) NOT NULL DEFAULT '0',
  `real_name` varchar(25) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `phone` char(15) DEFAULT '' COMMENT '手机号码',
  `birthday` int(11) NOT NULL DEFAULT '0' COMMENT '生日',
  `card_id` varchar(20) NOT NULL DEFAULT '' COMMENT '身份证号码',
  `mark` varchar(255) NOT NULL DEFAULT '' COMMENT '用户备注',
  `avatar` varchar(256) NOT NULL DEFAULT '' COMMENT '用户头像',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '详细地址',
  `partner_id` int(11) NOT NULL DEFAULT '0' COMMENT '合伙人id',
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户分组id',
  `spread_uid` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '推广员id',
  `spread_time` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '推广员关联时间',
  `spread_count` int(11) DEFAULT '0' COMMENT '下级人数',
  `is_promoter` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否为推广员',
  `source` tinyint(1) DEFAULT '1' COMMENT '来源1pc,2app,3h5,4小程序',
  `type` varchar(32) DEFAULT '1' COMMENT '用户类型',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1为正常，0为禁止',
  `level` tinyint(2) UNSIGNED NOT NULL DEFAULT '0' COMMENT '等级',
  `clean_time` int(11) DEFAULT '0' COMMENT '清理会员时间',
  `is_money_level` tinyint(1) NOT NULL DEFAULT '0' COMMENT '会员来源  0: 购买商品升级   1：花钱购买的会员2: 会员卡领取',
  `is_ever_level` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否永久性会员  0: 非永久会员  1：永久会员',
  `overdue_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '会员到期时间',
  `balance` int(11) NOT NULL DEFAULT '0' COMMENT '用户余额',
  `brokerage` int(11) NOT NULL DEFAULT '0' COMMENT '佣金金额',
  `exp` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '会员经验',
  `integral` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户剩余积分',
  `pay_count` int(10) UNSIGNED DEFAULT '0' COMMENT '用户购买次数',
  `add_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(16) NOT NULL DEFAULT '' COMMENT '添加ip',
  `last_time` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '最后一次登录时间',
  `last_ip` varchar(16) NOT NULL DEFAULT '' COMMENT '最后一次登录ip',
  `login_type` varchar(36) NOT NULL DEFAULT '' COMMENT '用户登陆类型，h5,wechat,routine',
  `record_phone` varchar(11) NOT NULL DEFAULT '0' COMMENT '记录临时电话',
  `sign_num` int(11) NOT NULL DEFAULT '0' COMMENT '连续签到天数',
  `created_at` int(11) NOT NULL DEFAULT '0',
  `created_uid` int(11) NOT NULL DEFAULT '0',
  `updated_at` int(11) NOT NULL DEFAULT '0',
  `updated_uid` int(11) NOT NULL DEFAULT '0',
  `effect` tinyint(4) NOT NULL DEFAULT '1',
  `memo` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表' ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `ybs_client_groups`
--

CREATE TABLE `ybs_client_groups` (
  `id` int(11) NOT NULL,
  `name` varchar(24) NOT NULL DEFAULT '' COMMENT '用户分组名称',
  `pid` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类父级 id',
  `merch_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '商户 id',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序DESC',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `created_uid` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_uid` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `effect` tinyint(1) UNSIGNED NOT NULL DEFAULT '1',
  `memo` varchar(255) NOT NULL DEFAULT ''
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='用户分组表';

--
-- 转存表中的数据 `ybs_client_groups`
--

INSERT INTO `ybs_client_groups` (`id`, `name`, `pid`, `merch_id`, `sort`, `created_at`, `created_uid`, `updated_at`, `updated_uid`, `effect`, `memo`) VALUES
(1, 'UI', 0, 1, 1, 1618193009, 1, 1618198059, 1, 1, ''),
(2, 'UI3232', 0, 1, 1, 1618196284, 1, 1618217639, 1, 1, ''),
(3, 'dklsdjf', 0, 1, 1, 1618196290, 1, 1618233582, 1, 0, ''),
(4, '老板', 0, 1, 1, 1618196369, 1, 1618198479, 1, 0, ''),
(5, '前台', 0, 1, 1, 1618196862, 1, 1618196862, 1, 1, ''),
(6, '程序员', 0, 1, 1, 1618217567, 1, 1618217661, 1, 0, ''),
(7, '程序员', 0, 1, 1, 1618233284, 1, 1618233284, 1, 1, '');

-- --------------------------------------------------------

--
-- 表的结构 `ybs_client_label_relations`
--

CREATE TABLE `ybs_client_label_relations` (
  `id` int(11) NOT NULL,
  `merch_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '商户id',
  `client_id` int(11) NOT NULL DEFAULT '0' COMMENT '标签ID',
  `label_id` int(11) NOT NULL DEFAULT '0' COMMENT '标签ID',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `created_uid` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_uid` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `effect` tinyint(1) UNSIGNED NOT NULL DEFAULT '1',
  `memo` varchar(255) NOT NULL DEFAULT ''
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='客户标签关联表';

--
-- 转存表中的数据 `ybs_client_label_relations`
--

INSERT INTO `ybs_client_label_relations` (`id`, `merch_id`, `client_id`, `label_id`, `created_at`, `created_uid`, `updated_at`, `updated_uid`, `effect`, `memo`) VALUES
(1, 1, 2, 1, 1618306525, 1, 1618306525, 1, 1, ''),
(2, 1, 2, 2, 1618306525, 1, 1618306525, 1, 1, ''),
(3, 1, 2, 19, 1618306525, 1, 1618306525, 1, 1, ''),
(4, 1, 2, 20, 1618306525, 1, 1618306525, 1, 1, ''),
(5, 1, 5, 1, 1618306525, 1, 1618306525, 1, 1, ''),
(6, 1, 5, 2, 1618306525, 1, 1618306525, 1, 1, ''),
(7, 1, 23, 1, 1618306525, 1, 1618306525, 1, 1, ''),
(8, 1, 23, 2, 1618306525, 1, 1618306525, 1, 1, ''),
(9, 1, 32, 1, 1618306525, 1, 1618306525, 1, 1, ''),
(11, 1, 8, 1, 1618327117, 1, 1618327117, 1, 1, ''),
(12, 1, 8, 2, 1618327117, 1, 1618327117, 1, 1, ''),
(13, 1, 32, 2, 1618327117, 1, 1618327117, 1, 1, ''),
(14, 1, 9, 1, 1618327700, 1, 1618327700, 1, 1, ''),
(15, 1, 9, 2, 1618327700, 1, 1618327700, 1, 1, '');

-- --------------------------------------------------------

--
-- 表的结构 `ybs_labels`
--

CREATE TABLE `ybs_labels` (
  `id` int(11) NOT NULL,
  `name` varchar(24) NOT NULL DEFAULT '' COMMENT '标签名称',
  `merch_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '商户 id',
  `labelcate_id` int(11) NOT NULL DEFAULT '0' COMMENT '标签分类',
  `type` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '1客户标签',
  `source` int(11) NOT NULL DEFAULT '1' COMMENT '0=手动标签 1=自动标签	',
  `sort` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序DESC',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `created_uid` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_uid` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `effect` tinyint(1) UNSIGNED NOT NULL DEFAULT '1',
  `memo` varchar(255) NOT NULL DEFAULT ''
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='标签';

--
-- 转存表中的数据 `ybs_labels`
--

INSERT INTO `ybs_labels` (`id`, `name`, `merch_id`, `labelcate_id`, `type`, `source`, `sort`, `created_at`, `created_uid`, `updated_at`, `updated_uid`, `effect`, `memo`) VALUES
(1, '家庭', 1, 1, 0, 1, 0, 1618047529, 1, 1618047529, 1, 1, ''),
(2, '喜剧', 1, 1, 0, 1, 0, 1618047550, 1, 1618047550, 1, 1, ''),
(21, '民谣', 1, 2, 0, 1, 0, 1618047550, 1, 1618047550, 1, 1, ''),
(19, '流行', 1, 1, 0, 1, 0, 1618047550, 1, 1618047550, 1, 1, ''),
(20, '摇滚', 1, 2, 0, 1, 0, 1618047550, 1, 1618047550, 1, 1, ''),
(18, '科幻', 1, 2, 0, 1, 0, 1618047550, 1, 1618047550, 1, 1, '');

-- --------------------------------------------------------

--
-- 表的结构 `ybs_label_categories`
--

CREATE TABLE `ybs_label_categories` (
  `id` int(11) NOT NULL,
  `name` varchar(24) NOT NULL DEFAULT '' COMMENT '标签分类名称',
  `pid` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类父级 id',
  `merch_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '商户 id',
  `type` int(11) UNSIGNED NOT NULL DEFAULT '1' COMMENT '1=用户;',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序DESC',
  `owner_id` int(11) NOT NULL DEFAULT '0' COMMENT '所有人, 0为全部',
  `other` text NOT NULL COMMENT '其他参数',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `created_uid` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_uid` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `effect` tinyint(1) UNSIGNED NOT NULL DEFAULT '1',
  `memo` varchar(255) NOT NULL DEFAULT ''
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='标签分类表';

--
-- 转存表中的数据 `ybs_label_categories`
--

INSERT INTO `ybs_label_categories` (`id`, `name`, `pid`, `merch_id`, `type`, `sort`, `owner_id`, `other`, `created_at`, `created_uid`, `updated_at`, `updated_uid`, `effect`, `memo`) VALUES
(1, '电影', 0, 1, 1, 99, 0, '', 1618329599, 0, 1618152389, 1, 1, ''),
(2, '歌曲', 0, 1, 1, 1, 0, '', 1617975867, 1, 1617975867, 1, 1, '');

-- --------------------------------------------------------

--
-- 表的结构 `ybs_oplogs`
--

CREATE TABLE `ybs_oplogs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `model_name` varchar(256) DEFAULT NULL,
  `action_name` varchar(256) DEFAULT NULL,
  `content` longtext,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `ybs_oplogs`
--

INSERT INTO `ybs_oplogs` (`id`, `created_at`, `updated_at`, `deleted_at`, `model_name`, `action_name`, `content`, `user_id`) VALUES
(1, '2021-04-07 16:54:24.119', '2021-04-07 16:54:24.119', NULL, '认证', '登录', '', 1),
(2, '2021-04-07 16:54:27.212', '2021-04-07 16:54:27.212', NULL, '认证', '登录', '', 1),
(3, '2021-04-07 16:54:28.315', '2021-04-07 16:54:28.315', NULL, '认证', '登录', '', 1),
(4, '2021-04-07 16:54:33.714', '2021-04-07 16:54:33.714', NULL, '认证', '登录', '', 1),
(5, '2021-04-07 16:54:34.722', '2021-04-07 16:54:34.722', NULL, '认证', '登录', '', 1),
(6, '2021-04-07 16:54:35.578', '2021-04-07 16:54:35.578', NULL, '认证', '登录', '', 1),
(7, '2021-04-07 16:54:36.525', '2021-04-07 16:54:36.525', NULL, '认证', '登录', '', 1),
(8, '2021-04-07 16:54:37.222', '2021-04-07 16:54:37.222', NULL, '认证', '登录', '', 1),
(9, '2021-04-07 16:54:38.132', '2021-04-07 16:54:38.132', NULL, '认证', '登录', '', 1),
(10, '2021-04-07 16:54:39.242', '2021-04-07 16:54:39.242', NULL, '认证', '登录', '', 1),
(11, '2021-04-07 17:01:20.116', '2021-04-07 17:01:20.116', NULL, '认证', '登录', '', 1),
(12, '2021-04-07 17:08:40.450', '2021-04-07 17:08:40.450', NULL, '认证', '登录', '', 1),
(13, '2021-04-07 17:08:41.720', '2021-04-07 17:08:41.720', NULL, '认证', '登录', '', 1),
(14, '2021-04-07 17:09:28.580', '2021-04-07 17:09:28.580', NULL, '认证', '登录', '', 1),
(15, '2021-04-07 17:42:03.343', '2021-04-07 17:42:03.343', NULL, '认证', '登录', '', 1),
(16, '2021-04-07 17:53:44.791', '2021-04-07 17:53:44.791', NULL, '认证', '登录', '', 1),
(17, '2021-04-07 17:56:23.113', '2021-04-07 17:56:23.113', NULL, '认证', '登录', '', 1),
(18, '2021-04-07 17:57:20.810', '2021-04-07 17:57:20.810', NULL, '认证', '登录', '', 1),
(19, '2021-04-08 10:56:53.057', '2021-04-08 10:56:53.057', NULL, '认证', '登录', '', 1),
(20, '2021-04-08 10:58:31.231', '2021-04-08 10:58:31.231', NULL, '认证', '登录', '', 1),
(21, '2021-04-08 11:01:01.877', '2021-04-08 11:01:01.877', NULL, '认证', '登录', '', 1),
(22, '2021-04-08 15:45:28.815', '2021-04-08 15:45:28.815', NULL, '认证', '登录', '', 1),
(23, '2021-04-08 16:12:57.206', '2021-04-08 16:12:57.206', NULL, '认证', '登录', '', 1),
(24, '2021-04-08 16:21:10.895', '2021-04-08 16:21:10.895', NULL, '认证', '登录', '', 1),
(25, '2021-04-08 16:21:15.114', '2021-04-08 16:21:15.114', NULL, '认证', '登录', '', 1),
(26, '2021-04-08 16:22:02.273', '2021-04-08 16:22:02.273', NULL, '认证', '登录', '', 1),
(27, '2021-04-08 16:22:03.082', '2021-04-08 16:22:03.082', NULL, '认证', '登录', '', 1),
(28, '2021-04-08 16:22:47.434', '2021-04-08 16:22:47.434', NULL, '认证', '登录', '', 1),
(29, '2021-04-08 16:24:28.543', '2021-04-08 16:24:28.543', NULL, '认证', '登录', '', 1),
(30, '2021-04-08 16:28:26.632', '2021-04-08 16:28:26.632', NULL, '认证', '登录', '', 1),
(31, '2021-04-08 16:31:43.942', '2021-04-08 16:31:43.942', NULL, '认证', '登录', '', 1),
(32, '2021-04-08 16:33:24.028', '2021-04-08 16:33:24.028', NULL, '认证', '登录', '', 1),
(33, '2021-04-08 17:01:22.992', '2021-04-08 17:01:22.992', NULL, '认证', '登录', '', 1),
(34, '2021-04-09 21:54:13.519', '2021-04-09 21:54:13.519', NULL, '认证', '登录', '', 1),
(35, '2021-04-10 11:18:45.559', '2021-04-10 11:18:45.559', NULL, '认证', '登录', '', 1),
(36, '2021-04-12 11:07:30.329', '2021-04-12 11:07:30.329', NULL, '认证', '登录', '', 1),
(37, '2021-04-12 11:07:31.943', '2021-04-12 11:07:31.943', NULL, '认证', '登录', '', 1),
(38, '2021-04-12 11:07:32.886', '2021-04-12 11:07:32.886', NULL, '认证', '登录', '', 1),
(39, '2021-04-12 11:09:50.361', '2021-04-12 11:09:50.361', NULL, '认证', '登录', '', 1),
(40, '2021-04-12 11:09:51.367', '2021-04-12 11:09:51.367', NULL, '认证', '登录', '', 1),
(41, '2021-04-12 11:09:52.770', '2021-04-12 11:09:52.770', NULL, '认证', '登录', '', 1),
(42, '2021-04-12 11:09:53.780', '2021-04-12 11:09:53.780', NULL, '认证', '登录', '', 1),
(43, '2021-04-12 11:10:06.562', '2021-04-12 11:10:06.562', NULL, '认证', '登录', '', 1),
(44, '2021-04-12 11:10:07.270', '2021-04-12 11:10:07.270', NULL, '认证', '登录', '', 1),
(45, '2021-04-12 11:10:07.875', '2021-04-12 11:10:07.875', NULL, '认证', '登录', '', 1),
(46, '2021-04-12 11:34:36.900', '2021-04-12 11:34:36.900', NULL, '认证', '登录', '', 1);

-- --------------------------------------------------------

--
-- 表的结构 `ybs_roles`
--

CREATE TABLE `ybs_roles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(256) NOT NULL,
  `display_name` varchar(256) DEFAULT NULL,
  `description` varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `ybs_users_system_admin`
--

CREATE TABLE `ybs_users_system_admin` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(60) NOT NULL,
  `username` varchar(60) NOT NULL,
  `password` varchar(100) DEFAULT NULL,
  `intro` varchar(512) NOT NULL,
  `avatar` varchar(1024) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `ybs_users_system_admin`
--

INSERT INTO `ybs_users_system_admin` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `username`, `password`, `intro`, `avatar`) VALUES
(1, '2021-04-07 16:34:05.958', '2021-04-07 16:34:05.958', NULL, 'name', 'username', '$2a$10$35u0p4B6ESwCsKrxSopiCuFz8iAGUvja0CJ67sYgyqBTfp1k7slkK', '超级弱鸡程序猿一枚！！！！', 'https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIPbZRufW9zPiaGpfdXgU7icRL1licKEicYyOiace8QQsYVKvAgCrsJx1vggLAD2zJMeSXYcvMSkw9f4pw/132');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `ybs_casbin_rule`
--
ALTER TABLE `ybs_casbin_rule`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `unique_index` (`p_type`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`);

--
-- Indexes for table `ybs_client_groups`
--
ALTER TABLE `ybs_client_groups`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ybs_client_label_relations`
--
ALTER TABLE `ybs_client_label_relations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `client_id` (`client_id`),
  ADD KEY `merch_id` (`merch_id`,`effect`);

--
-- Indexes for table `ybs_labels`
--
ALTER TABLE `ybs_labels`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- Indexes for table `ybs_label_categories`
--
ALTER TABLE `ybs_label_categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ybs_oplogs`
--
ALTER TABLE `ybs_oplogs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_ybs_oplogs_deleted_at` (`deleted_at`),
  ADD KEY `fk_ybs_users_oplogs` (`user_id`);

--
-- Indexes for table `ybs_roles`
--
ALTER TABLE `ybs_roles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_ybs_roles_name` (`name`),
  ADD KEY `idx_ybs_roles_deleted_at` (`deleted_at`);

--
-- Indexes for table `ybs_users_system_admin`
--
ALTER TABLE `ybs_users_system_admin`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_ybs_users_username` (`username`),
  ADD KEY `idx_ybs_users_deleted_at` (`deleted_at`),
  ADD KEY `idx_ybs_users_name` (`name`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `ybs_casbin_rule`
--
ALTER TABLE `ybs_casbin_rule`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- 使用表AUTO_INCREMENT `ybs_client_groups`
--
ALTER TABLE `ybs_client_groups`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- 使用表AUTO_INCREMENT `ybs_client_label_relations`
--
ALTER TABLE `ybs_client_label_relations`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- 使用表AUTO_INCREMENT `ybs_labels`
--
ALTER TABLE `ybs_labels`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- 使用表AUTO_INCREMENT `ybs_label_categories`
--
ALTER TABLE `ybs_label_categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- 使用表AUTO_INCREMENT `ybs_oplogs`
--
ALTER TABLE `ybs_oplogs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=47;

--
-- 使用表AUTO_INCREMENT `ybs_roles`
--
ALTER TABLE `ybs_roles`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `ybs_users_system_admin`
--
ALTER TABLE `ybs_users_system_admin`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 限制导出的表
--

--
-- 限制表 `ybs_oplogs`
--
ALTER TABLE `ybs_oplogs`
  ADD CONSTRAINT `fk_ybs_users_oplogs` FOREIGN KEY (`user_id`) REFERENCES `ybs_users_system_admin` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
