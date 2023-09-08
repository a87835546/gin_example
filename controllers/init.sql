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

create table billboard
(
    id          bigint auto_increment
        primary key,
    url         varchar(255)                           not null,
    title       varchar(255)                           null,
    `desc`      varchar(255)                           null,
    category_id varchar(200)                           null,
    created_at  timestamp    default CURRENT_TIMESTAMP null,
    updated_at  timestamp    default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    author      varchar(100) default ''                not null,
    theme_url   varchar(255)                           null,
    types       varchar(200)                           null,
    actor       varchar(1000)                          null,
    rate        varchar(100) default '0'               not null,
    years       int          default 1990              not null,
    duration    int          default 0                 not null comment '视频时长，分钟'
);

create table menu
(
    id         int auto_increment
        primary key,
    `desc`     varchar(255)                        null,
    title      varchar(255)                        null,
    role       int                                 not null,
    position   int       default 0                 null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null,
    status     int                                 null,
    title_en   varchar(200)                        not null,
    desc_en    text                                null
);

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


create table video_category
(
    id          int auto_increment
        primary key,
    title       varchar(255)                           not null,
    `desc`      varchar(255)                           null,
    `index`     int          default 1                 not null,
    status      tinyint(1)   default 1                 not null,
    created_at  bigint       default 0                 not null,
    updated_at  timestamp    default CURRENT_TIMESTAMP null,
    super_title varchar(100) default '0'               null,
    title_en    varchar(100)                           not null,
    constraint title
        unique (title)
);

create table video_type
(
    id     int auto_increment
        primary key,
    title  varchar(200) not null,
    author varchar(200) null,
    constraint title
        unique (title)
);

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
                          `type` int DEFAULT NULL,
                          `video_id` bigint DEFAULT NULL,
                          `title` varchar(255) DEFAULT NULL,
                          `desc` varchar(255) DEFAULT NULL,
                          `created_at` bigint DEFAULT NULL,
                          `updated_at` bigint DEFAULT NULL,
                          `operation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作者',
                          PRIMARY KEY (`id`),
                          KEY `video_id` (`video_id`),
                          KEY `type` (`type`),
                          CONSTRAINT `banner_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `billboard` (`id`),
                          CONSTRAINT `banner_ibfk_2` FOREIGN KEY (`type`) REFERENCES `menu` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;