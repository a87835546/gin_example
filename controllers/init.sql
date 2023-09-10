create table actor
(
    id         int auto_increment
        primary key,
    name       varchar(1000)                           null,
    created_at timestamp     default CURRENT_TIMESTAMP not null,
    avatar_url varchar(200)                            null,
    name_en    varchar(1000) default ''                not null,
    constraint actor_id_uindex
        unique (id)
);

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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


create table admin
(
    id         int auto_increment
        primary key,
    username   varchar(255)                        null,
    password   varchar(255)                        null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null,
    role       int       default 1                 null
);
create table favorite
(
    id         int auto_increment
        primary key,
    user_id    int                                 not null,
    video_id   bigint    default (0)               not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    constraint FK_favorite_billboard
        foreign key (video_id) references billboard (id),
    constraint FK_favorite_user
        foreign key (user_id) references user (id)
)
    comment '收藏的电影';


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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `video_type` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `title` varchar(200) NOT NULL,
                              `author` varchar(200) DEFAULT NULL,
                              `title_en` varchar(255) DEFAULT NULL,
                              `created_at` bigint DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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