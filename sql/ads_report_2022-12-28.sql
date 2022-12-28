# ************************************************************
# Sequel Ace SQL dump
# Version 20044
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.5.17-MariaDB-1:10.5.17+maria~ubu2004)
# Database: ads_report
# Generation Time: 2022-12-28 03:06:57 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table goadmin_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_menu`;

CREATE TABLE `goadmin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT 0,
  `type` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `order` int(11) unsigned NOT NULL DEFAULT 0,
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(3000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `header` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `plugin_name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `uuid` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_menu` DISABLE KEYS */;

INSERT INTO `goadmin_menu` (`id`, `parent_id`, `type`, `order`, `title`, `icon`, `uri`, `header`, `plugin_name`, `uuid`, `created_at`, `updated_at`)
VALUES
	(1,0,1,2,'Admin','fa-tasks','',NULL,'',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,1,1,2,'Users','fa-users','/info/manager',NULL,'',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(3,1,1,3,'Roles','fa-user','/info/roles',NULL,'',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(4,1,1,4,'Permission','fa-ban','/info/permission',NULL,'',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(5,1,1,5,'Menu','fa-bars','/menu',NULL,'',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(6,1,1,6,'Operation log','fa-history','/info/op',NULL,'',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(7,0,1,1,'Dashboard','fa-bar-chart','/',NULL,'',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(11,0,0,7,'Google Admob','fa-google','/info/tb_ad_corp_admob','','',NULL,'2022-12-08 03:09:45','2022-12-08 03:09:45'),
	(12,0,0,8,'Facebook Ads','fa-facebook','/info/tb_ad_corp_facebook','','',NULL,'2022-12-08 03:10:58','2022-12-08 03:10:58'),
	(13,0,0,9,'Unity Ads','fa-table','/info/tb_ad_corp_unity','','',NULL,'2022-12-08 03:13:21','2022-12-08 03:13:21'),
	(14,0,0,10,'Vungle Ads','fa-vimeo','/info/tb_ad_corp_vungle','','',NULL,'2022-12-08 03:13:56','2022-12-08 03:13:56'),
	(15,0,0,11,'Iron Source Ads','fa-info','/info/tb_ad_corp_iron','','',NULL,'2022-12-08 03:14:39','2022-12-08 03:14:39'),
	(16,0,0,12,'Adx Ads','fa-audio-description','/info/tb_ad_corp_adx','','',NULL,'2022-12-16 02:24:30','2022-12-16 02:24:30');

/*!40000 ALTER TABLE `goadmin_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_operation_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_operation_log`;

CREATE TABLE `goadmin_operation_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


# Dump of table goadmin_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_permissions`;

CREATE TABLE `goadmin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_permissions` (`id`, `name`, `slug`, `http_method`, `http_path`, `created_at`, `updated_at`)
VALUES
	(1,'All permission','*','','*','2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,'Dashboard','dashboard','GET,PUT,POST,DELETE','/','2019-09-10 00:00:00','2019-09-10 00:00:00');

/*!40000 ALTER TABLE `goadmin_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_menu`;

CREATE TABLE `goadmin_role_menu` (
  `role_id` int(11) unsigned NOT NULL,
  `menu_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_role_menu` DISABLE KEYS */;

INSERT INTO `goadmin_role_menu` (`role_id`, `menu_id`, `created_at`, `updated_at`)
VALUES
	(1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(1,7,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,7,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(1,11,'2022-12-08 03:09:45','2022-12-08 03:09:45'),
	(1,12,'2022-12-08 03:10:58','2022-12-08 03:10:58'),
	(1,13,'2022-12-08 03:13:21','2022-12-08 03:13:21'),
	(1,14,'2022-12-08 03:13:56','2022-12-08 03:13:56'),
	(1,15,'2022-12-08 03:14:39','2022-12-08 03:14:39'),
	(1,16,'2022-12-16 02:24:30','2022-12-16 02:24:30');

/*!40000 ALTER TABLE `goadmin_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_permissions`;

CREATE TABLE `goadmin_role_permissions` (
  `role_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_role_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_role_permissions` (`role_id`, `permission_id`, `created_at`, `updated_at`)
VALUES
	(1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(1,2,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,2,'2019-09-10 00:00:00','2019-09-10 00:00:00');

/*!40000 ALTER TABLE `goadmin_role_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_users`;

CREATE TABLE `goadmin_role_users` (
  `role_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_users` WRITE;
/*!40000 ALTER TABLE `goadmin_role_users` DISABLE KEYS */;

INSERT INTO `goadmin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`)
VALUES
	(1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,2,'2019-09-10 00:00:00','2019-09-10 00:00:00');

/*!40000 ALTER TABLE `goadmin_role_users` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_roles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_roles`;

CREATE TABLE `goadmin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_roles` WRITE;
/*!40000 ALTER TABLE `goadmin_roles` DISABLE KEYS */;

INSERT INTO `goadmin_roles` (`id`, `name`, `slug`, `created_at`, `updated_at`)
VALUES
	(1,'Administrator','administrator','2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,'Operator','operator','2019-09-10 00:00:00','2019-09-10 00:00:00');

/*!40000 ALTER TABLE `goadmin_roles` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_session
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_session`;

CREATE TABLE `goadmin_session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `values` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `goadmin_session` WRITE;
/*!40000 ALTER TABLE `goadmin_session` DISABLE KEYS */;

INSERT INTO `goadmin_session` (`id`, `sid`, `values`, `created_at`, `updated_at`)
VALUES
	(103,'aae2a5af-3747-48bf-846a-7ef2419a8f71','{\"user_id\":1}','2022-12-16 05:10:34','2022-12-16 05:10:34');

/*!40000 ALTER TABLE `goadmin_session` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_site
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_site`;

CREATE TABLE `goadmin_site` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(3000) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `state` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_site` WRITE;
/*!40000 ALTER TABLE `goadmin_site` DISABLE KEYS */;

INSERT INTO `goadmin_site` (`id`, `key`, `value`, `description`, `state`, `created_at`, `updated_at`)
VALUES
	(1,'logger_encoder_name_key','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(2,'logger_encoder_message_key','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(3,'logger_encoder_caller','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(4,'hide_visitor_user_center_entrance','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(5,'access_assets_log_off','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(6,'error_log_off','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(7,'custom_404_html','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(8,'theme','adminlte',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(9,'logger_level','0',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(10,'logger_rotate_compress','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(11,'session_life_time','7200',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(12,'language','en',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(13,'logger_rotate_max_age','0',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(14,'info_log_path','/Users/gamepub/Documents/go/src/github.com/coolishbee/ads-report-go/logs/info.log',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(15,'info_log_off','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(16,'logger_rotate_max_size','0',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(17,'logger_encoder_level','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(18,'hide_config_center_entrance','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(19,'title','GoAdsReport',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(20,'index_url','/',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(21,'login_title','GoAdmin',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(22,'operation_log_off','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(23,'env','local',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(24,'logger_encoder_caller_key','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(25,'go_mod_file_path','/Users/gamepub/Documents/go/src/github.com/coolishbee/ads-report-go/go.mod',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(26,'animation_type','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(27,'hide_app_info_entrance','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(28,'hide_tool_entrance','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(29,'custom_403_html','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(30,'domain','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(31,'footer_info','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(32,'animation_delay','0.00',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(33,'custom_foot_html','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(34,'auth_user_table','goadmin_users',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(35,'no_limit_login_ip','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(36,'custom_500_html','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(37,'open_admin_api','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(38,'exclude_theme_components','null',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(39,'access_log_path','/Users/gamepub/Documents/go/src/github.com/coolishbee/ads-report-go/logs/access.log',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(40,'logger_encoder_duration','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(41,'asset_url','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(42,'prohibit_config_modification','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(43,'debug','true',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(44,'logger_encoder_time','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(45,'access_log_off','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(46,'logger_encoder_level_key','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(47,'file_upload_engine','{\"name\":\"local\"}',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(48,'login_logo','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(49,'site_off','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(50,'allow_del_operation_log','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(51,'logger_encoder_stacktrace_key','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(52,'color_scheme','skin-black',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(53,'asset_root_path','./public/',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(54,'error_log_path','/Users/gamepub/Documents/go/src/github.com/coolishbee/ads-report-go/logs/error.log',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(55,'bootstrap_file_path','/Users/gamepub/Documents/go/src/github.com/coolishbee/ads-report-go/bootstrap.go',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(56,'hide_plugin_entrance','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(57,'logger_encoder_encoding','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(58,'custom_head_html','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(59,'logo','GoAdsReport',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(60,'extra','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(61,'app_id','bcv86F3RO3NG',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(62,'url_prefix','admin',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(63,'sql_log','false',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(64,'logger_rotate_max_backups','0',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(65,'logger_encoder_time_key','',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(66,'animation_duration','0.00',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(67,'mini_logo','GA',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13'),
	(68,'login_url','/login',NULL,1,'2022-11-29 06:41:13','2022-11-29 06:41:13');

/*!40000 ALTER TABLE `goadmin_site` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_user_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_user_permissions`;

CREATE TABLE `goadmin_user_permissions` (
  `user_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_user_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_user_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_user_permissions` (`user_id`, `permission_id`, `created_at`, `updated_at`)
VALUES
	(1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,2,'2019-09-10 00:00:00','2019-09-10 00:00:00');

/*!40000 ALTER TABLE `goadmin_user_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_users`;

CREATE TABLE `goadmin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_users` WRITE;
/*!40000 ALTER TABLE `goadmin_users` DISABLE KEYS */;

INSERT INTO `goadmin_users` (`id`, `username`, `password`, `name`, `avatar`, `remember_token`, `created_at`, `updated_at`)
VALUES
	(1,'admin','$2a$10$bIKCvw7GzsHZyvDUklnZUebgVD4Efe0rclpvMTO.QcfidXCheRNOm','admin','','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,'operator','$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.','Operator','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');

/*!40000 ALTER TABLE `goadmin_users` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tb_ad_corp_admob
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_ad_corp_admob`;

CREATE TABLE `tb_ad_corp_admob` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_name` varchar(50) DEFAULT NULL,
  `game_id` varchar(50) DEFAULT NULL,
  `revenue` decimal(10,2) DEFAULT 0.00,
  `impression` int(11) DEFAULT NULL,
  `ecpm` decimal(10,2) DEFAULT 0.00,
  `ad_date` varchar(10) DEFAULT NULL,
  `reg_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_tb_ad_corp_admob` (`game_id`,`ad_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tb_ad_corp_adx
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_ad_corp_adx`;

CREATE TABLE `tb_ad_corp_adx` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_name` varchar(50) DEFAULT NULL,
  `game_id` varchar(50) DEFAULT NULL,
  `revenue` decimal(10,2) DEFAULT 0.00,
  `impression` int(11) DEFAULT NULL,
  `ecpm` decimal(10,2) DEFAULT 0.00,
  `ad_date` varchar(10) DEFAULT NULL,
  `reg_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_tb_ad_corp_adx` (`game_id`,`ad_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tb_ad_corp_facebook
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_ad_corp_facebook`;

CREATE TABLE `tb_ad_corp_facebook` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_name` varchar(50) DEFAULT NULL,
  `game_id` varchar(50) DEFAULT NULL,
  `revenue` decimal(10,2) DEFAULT 0.00,
  `impression` int(11) DEFAULT NULL,
  `ecpm` decimal(10,2) DEFAULT 0.00,
  `ad_date` varchar(10) DEFAULT NULL,
  `reg_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_tb_ad_corp_facebook` (`game_id`,`ad_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tb_ad_corp_iron
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_ad_corp_iron`;

CREATE TABLE `tb_ad_corp_iron` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_name` varchar(50) DEFAULT NULL,
  `game_id` varchar(50) DEFAULT NULL,
  `revenue` decimal(10,2) DEFAULT 0.00,
  `impression` int(11) DEFAULT NULL,
  `ecpm` decimal(10,2) DEFAULT 0.00,
  `ad_date` varchar(10) DEFAULT NULL,
  `reg_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_tb_ad_corp_iron` (`game_id`,`ad_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tb_ad_corp_unity
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_ad_corp_unity`;

CREATE TABLE `tb_ad_corp_unity` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_name` varchar(50) DEFAULT NULL,
  `game_id` varchar(50) DEFAULT NULL,
  `revenue` decimal(10,2) DEFAULT 0.00,
  `impression` int(11) DEFAULT NULL,
  `ecpm` decimal(10,2) DEFAULT 0.00,
  `ad_date` varchar(10) DEFAULT NULL,
  `reg_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_tb_ad_corp_unity` (`game_id`,`ad_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tb_ad_corp_vungle
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_ad_corp_vungle`;

CREATE TABLE `tb_ad_corp_vungle` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_name` varchar(50) DEFAULT NULL,
  `game_id` varchar(50) DEFAULT NULL,
  `revenue` decimal(10,2) DEFAULT 0.00,
  `impression` int(11) DEFAULT NULL,
  `ecpm` decimal(10,2) DEFAULT 0.00,
  `ad_date` varchar(10) DEFAULT NULL,
  `reg_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_tb_ad_corp_vungle` (`game_id`,`ad_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
