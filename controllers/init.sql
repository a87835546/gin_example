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

create table user
(
    id          int auto_increment
        primary key,
    username    varchar(255)                        null,
    password    varchar(255)                        null,
    created_at  timestamp default CURRENT_TIMESTAMP null,
    updated_at  timestamp default CURRENT_TIMESTAMP null,
    ip          varchar(100)                        not null comment 'last login ip address',
    device_type tinyint   default 0                 not null comment '设备类型 1 -- ios 2 --- android 0 --- others'
);



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

