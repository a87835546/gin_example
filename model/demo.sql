-- -------------------------------------------------------------
-- TablePlus 5.4.0(504)
--
-- https://tableplus.com/
--
-- Database: demo
-- Generation Time: 2023-09-10 19:53:52.2730
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `actor` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(1000) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `avatar_url` varchar(200) DEFAULT NULL,
  `name_en` varchar(1000) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `actor_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `role` int DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `banner` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` int DEFAULT NULL,
  `video_id` bigint DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `created_at` bigint DEFAULT NULL,
  `updated_at` bigint DEFAULT NULL,
  `operation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作者',
  `video_url` varchar(255) DEFAULT NULL,
  `video_theme_url` varchar(255) DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `video_id` (`video_id`),
  KEY `type` (`menu_id`),
  CONSTRAINT `banner_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `billboard` (`id`),
  CONSTRAINT `banner_ibfk_2` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `billboard` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `category_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `author` varchar(100) NOT NULL DEFAULT '',
  `theme_url` varchar(255) DEFAULT NULL,
  `types` varchar(200) DEFAULT NULL,
  `actor` varchar(1000) DEFAULT NULL,
  `rate` varchar(100) NOT NULL DEFAULT '0',
  `years` int NOT NULL DEFAULT '1990',
  `duration` int NOT NULL DEFAULT '0' COMMENT '视频时长，分钟',
  `menu_title` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `favorite` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` bigint NOT NULL DEFAULT (0),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_favorite_billboard` (`video_id`),
  KEY `FK_favorite_user` (`user_id`),
  CONSTRAINT `FK_favorite_billboard` FOREIGN KEY (`video_id`) REFERENCES `billboard` (`id`),
  CONSTRAINT `FK_favorite_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='收藏的电影';

CREATE TABLE `history` (
  `id` int NOT NULL AUTO_INCREMENT,
  `video_id` bigint NOT NULL,
  `user_id` int NOT NULL,
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `video_id` (`video_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `history_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `billboard` (`id`),
  CONSTRAINT `history_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `desc` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `title_en` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `role` int NOT NULL,
  `position` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `status` int DEFAULT NULL COMMENT '0 显示 1 隐藏',
  `desc_en` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `menu_category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `title_en` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `index` int NOT NULL DEFAULT '1',
  `menu_id` int DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` bigint NOT NULL DEFAULT '0' COMMENT '0 显示 1 隐藏',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `menu_category_menu_id_title_idx` (`menu_id`,`title`) USING BTREE,
  CONSTRAINT `menu_category_ibfk_1` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `ip` varchar(100) NOT NULL COMMENT 'last login ip address',
  `device_type` tinyint NOT NULL DEFAULT '0' COMMENT '设备类型 1 -- ios 2 --- android 0 --- others',
  `email` varchar(255) NOT NULL DEFAULT '',
  `gender` tinyint(1) NOT NULL DEFAULT '1',
  `location` varchar(255) NOT NULL DEFAULT '',
  `birthday` varchar(255) NOT NULL DEFAULT '',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `phone_number` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `video_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL,
  `author` varchar(200) DEFAULT NULL,
  `title_en` varchar(255) DEFAULT NULL,
  `created_at` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `actor` (`id`, `name`, `created_at`, `avatar_url`, `name_en`) VALUES
(1, '刘德华', '2023-09-08 14:54:01', '', 'Andy Liu'),
(2, '张三', '2023-09-08 14:54:19', '', 'Zhan San');

INSERT INTO `admin` (`id`, `username`, `password`, `created_at`, `updated_at`, `role`) VALUES
(1, 'admin', '123456', '2023-09-08 11:54:32', '2023-09-08 11:54:32', 0);

INSERT INTO `banner` (`id`, `menu_id`, `video_id`, `title`, `desc`, `created_at`, `updated_at`, `operation`, `video_url`, `video_theme_url`, `status`) VALUES
(4, 2, 4, '11', '23', 1694250996, 1694250996643, '', 'https://s5.bfzycdn.com/video/xinshengzhongsheng/HD/index.m3u8', 'https://img.bfzypic.com/upload/vod/20230908-1/f882147dbba5306ce52c035a7efe400d.jpg', 1),
(5, 1, 7, '小姨子的梦', '小姨子的梦', 1694328956, 1694328956058, '', 'https://s1.bfzyll.com/video/xiaoyizidemeng/HD/index.m3u8', 'https://img.bfzypic.com/upload/vod/20230909-1/c54ae0a2feb39ec74a478f84b34d4673.jpg', 1),
(6, 1, 5, '狄仁杰之幽兵借路', '狄仁杰之幽兵借路', 1694329220, 1694329220015, '', 'https://s5.bfzycdn.com/video/direnjiezhiyoubingjielu/HD/index.m3u8', 'https://img.bfzypic.comhttps//pic0.iqiyipic.com/image/20230909/ec/f0/v_173795306_m_601_m4_260_360.jpg?caplist=jpg,webp,avif#err2023-09-10', 1);

INSERT INTO `billboard` (`id`, `url`, `title`, `desc`, `category_id`, `created_at`, `updated_at`, `author`, `theme_url`, `types`, `actor`, `rate`, `years`, `duration`, `menu_title`) VALUES
(1, 'https://s5.bfzycdn.com/video/kesouyanzhinvdi23ji/第01集/index.m3u8', '科搜研之女第23季', '1', 10, '2023-09-08 15:34:11', '2023-09-09 14:54:03', 'admin', 'https://img.bfzypic.com/upload/vod/20230908-1/f882147dbba5306ce52c035a7efe400d.jpg', '剧情,动作', '泽口靖子,内藤刚志,小池彻平,若村麻由美,风间彻,齐藤晓', '5.6', 2023, 0, NULL),
(2, 'https://s5.bfzycdn.com/video/kesouyanzhinvdi23ji/第02集/index.m3u8', '科搜研之女第23季', '科搜研之女第23季', 10, '2023-09-08 19:27:29', '2023-09-09 14:54:03', 'admin', 'https://img.bfzypic.com/upload/vod/20230908-1/f882147dbba5306ce52c035a7efe400d.jpg', '剧情', '泽口靖子,内藤刚志,小池彻平,若村麻由美,风间彻,齐藤晓', '5.0', 2023, 0, NULL),
(3, 'hhttps://s6.bfzycdn.com/video/babi/HD中字/index.m3u8', '芭比2023', '芭比娃娃 芭比真人版 芭比娃娃真人版', 10, '2023-09-08 19:40:56', '2023-09-09 14:54:03', 'admin', 'https://img.bfzypic.com/upload/vod/20230722-1/8c3f3794990321c95d20e14dd7a20dda.jpg', '剧情,喜剧', '玛格特·罗比', '2.0', 2020, 0, NULL),
(4, 'https://s5.bfzycdn.com/video/xinshengzhongsheng/HD/index.m3u8', '新生，重生', '1', 11, '2023-09-09 15:50:04', '2023-09-09 15:57:33', 'admin', 'https://img.bfzypic.com/upload/vod/20230824-1/04f11723f7a02314b284d08fc77af295.jpg', '恐怖', '朱迪·雷耶斯,布瑞达·伍尔,马琳·爱尔兰,莫尼克·加布里埃拉·库尔内', '3.0', 2023, 0, '电视剧'),
(5, 'https://s5.bfzycdn.com/video/direnjiezhiyoubingjielu/HD/index.m3u8', '狄仁杰之幽兵借路', '狄仁杰之幽兵借路', 9, '2023-09-10 14:03:55', '2023-09-10 14:49:23', 'admin', 'https://img.bfzypic.comhttps//pic0.iqiyipic.com/image/20230909/ec/f0/v_173795306_m_601_m4_260_360.jpg?caplist=jpg,webp,avif#err2023-09-10', '2,1,4', '杜奕衡李若希祝昕愿', '2.0', 2023, 0, NULL),
(6, 'https://img.bfzypic.com/upload/vod/20230910-1/7c2e2716621367eb8994e212c8ce89fa.webp', '天坑寻龙', '天坑寻龙', 9, '2023-09-10 14:50:39', '2023-09-10 14:50:39', 'admin', 'https://img.bfzypic.com/upload/vod/20230910-1/7c2e2716621367eb8994e212c8ce89fa.webp', '恐怖', '林枫烨,胡雪儿,郜玄铭,恩璟,郑冬,赵亮', '2.0', 2023, 0, NULL),
(7, 'https://s1.bfzyll.com/video/xiaoyizidemeng/HD/index.m3u8', '小姨子的梦', '小姨子的梦', 8, '2023-09-10 14:53:03', '2023-09-10 14:53:03', 'admin', 'https://img.bfzypic.com/upload/vod/20230909-1/c54ae0a2feb39ec74a478f84b34d4673.jpg', '喜剧,动作,科幻', '정향 하빈 위지웅', '0.0', 2023, 0, '电影'),
(8, 'https://s5.bfzycdn.com/video/weixiaodong/HD/index.m3u8', '喂！小咚', '喂！小咚', 14, '2023-09-10 18:12:00', '2023-09-10 18:12:00', 'admin', 'https://img.bfzypic.com/upload/vod/20230908-1/39929179fdbf88b6c90b8e8020fdf6b7.webp', '剧情,喜剧', '坂口辰平 大塚弘太 遠藤隆太 师冈广明 宫部纯子', '0.0', 2022, 0, '电影'),
(9, 'https://s5.bfzycdn.com/video/chayoumen/中字/index.m3u8', '茶友们', '茶友们', 13, '2023-09-10 18:13:00', '2023-09-10 18:13:00', 'admin', 'https://img.bfzypic.com/upload/vod/20230908-1/4e2ada31bfcdf9889577ced2113a113e.webp', '剧情', '冈本玲,渡边哲,伊藤庆德', '2.0', 2023, 0, '电影');

INSERT INTO `menu` (`id`, `desc`, `title`, `title_en`, `role`, `position`, `created_at`, `updated_at`, `status`, `desc_en`) VALUES
(1, '电影', '电影', 'Movies', 0, 0, '2023-09-08 14:57:53', '2023-09-08 14:57:53', 0, NULL),
(2, '电视剧', '电视剧', 'Serious', 0, 0, '2023-09-08 14:58:17', '2023-09-08 14:58:17', 0, NULL),
(7, '动漫', '动漫', 'Animation', 0, 0, '2023-09-08 19:22:07', '2023-09-08 19:22:07', 0, NULL),
(8, '综艺', '综艺', 'TV Show', 0, 0, '2023-09-08 19:22:56', '2023-09-08 19:22:56', 0, NULL),
(9, '真人', '午夜', 'Pron', 0, 0, '2023-09-08 19:23:25', '2023-09-08 19:23:25', 0, NULL),
(10, '其他', '其他', 'Other', 0, 0, '2023-09-09 13:45:35', '2023-09-09 13:45:35', 0, NULL);

INSERT INTO `menu_category` (`id`, `title`, `title_en`, `desc`, `index`, `menu_id`, `status`, `created_at`, `updated_at`) VALUES
(8, '热门', 'hot', 'hot', 0, 1, 1, 1694239901105, '2023-09-09 14:11:41'),
(9, '测试分类', 'test', 'test', 1, 1, 1, 1694241371059, '2023-09-09 14:36:11'),
(10, '热门', 'hot', 'hot', 0, 2, 1, 1694241747265, '2023-09-09 14:42:27'),
(11, '电视剧测试分类', 'test 1', 'test', 0, 2, 1, 1694242363363, '2023-09-09 14:52:43'),
(12, '动作片', 'Action Film', 'Action Film', 0, 1, 1, 1694340571778, '2023-09-10 18:09:32'),
(13, 'Top100', 'Top100', 'Top100', 0, 1, 1, 1694340583395, '2023-09-10 18:09:43'),
(14, '喜剧片', 'Comedy Film', 'Comedy Film', 0, 1, 1, 1694340623868, '2023-09-10 18:10:24'),
(15, '热门', 'hot', 'hot', 0, 7, 1, 1694343039083, '2023-09-10 18:50:39'),
(16, '热门', 'hot', '综艺 热门', 0, 8, 1, 1694343061547, '2023-09-10 18:51:02'),
(17, '热门', 'Hot', '成人 热门', 0, 9, 1, 1694343094902, '2023-09-10 18:51:35');

INSERT INTO `user` (`id`, `username`, `password`, `created_at`, `updated_at`, `ip`, `device_type`, `email`, `gender`, `location`, `birthday`, `avatar`, `phone_number`) VALUES
(3, 'zhansan', '123456', '2023-09-08 11:50:27', '2023-09-08 11:50:27', '123456', 0, '', 0, '', '', '', '');

INSERT INTO `video_type` (`id`, `title`, `author`, `title_en`, `created_at`) VALUES
(1, '动作', 'admin', 'Action', 1694244016),
(2, '剧情', 'admin', 'drama film', 1694245021),
(3, '喜剧', 'admin', 'Comedy', 1694245037),
(4, '恐怖', 'admin', 'Horror', 1694245100),
(5, '科幻', 'admin', 'Science Fiction Film', 1694325963),
(6, '爱情片', 'admin', 'Lover Film', 1694325995);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;