create database tiktok;
use tiktok;

create table comments
(
    comment_id  bigint       not null
        primary key,
    video_id    bigint       not null,
    content     varchar(255) not null,
    create_time bigint       not null,
    user_id     bigint       not null
);

INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10305139227234304, 1, 'this is a test', 182857, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10306790612799488, 1, 'this is a test', 183530, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10307517473099776, 1, 'this is a test', 183824, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10307681625575424, 1, 'this is a test', 183903, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10311642411700224, 1, 'this is a test', 185447, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10314650688163840, 1, 'this is a test', 190644, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10316515857403904, 1, '123', 191409, 10246285546229760);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10316552238796800, 1, '123', 191418, 10246285546229760);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10316616797523968, 1, 'test', 191433, 10246285546229760);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10533001981071360, 1, 'this is a test ！！！', 1655084063, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (10533686395015168, 1, 'this is a test ？？', 1655084226, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (19677437084307456, 1, '88888888', 1657264266, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (19700058362482688, 1, 'jhgghhj', 1657269659, 10181697941278720);
INSERT INTO tiktok.comments (comment_id, video_id, content, create_time, user_id) VALUES (19700062862970880, 1, 'jhgghhj', 1657269660, 10181697941278720);


create table favorite
(
    favorite_id bigint not null
        primary key,
    user_id     bigint not null,
    video_id    bigint null
);

INSERT INTO tiktok.favorite (favorite_id, user_id, video_id) VALUES (0, 10181697941278720, 1);
INSERT INTO tiktok.favorite (favorite_id, user_id, video_id) VALUES (12559657105100800, 10246285546229760, 1);


create table relations
(
    relation_id bigint not null
        primary key,
    user_id     bigint not null,
    to_user_id  bigint not null
);

INSERT INTO tiktok.relations (relation_id, user_id, to_user_id) VALUES (10280925019508736, 10246285546229760, 10181697941278720);
INSERT INTO tiktok.relations (relation_id, user_id, to_user_id) VALUES (10281607348883456, 10279700773474304, 10246285546229760);


create table users
(
    user_id  bigint                      not null comment '用户 snowflake id',
    username varchar(30) charset utf8mb4 not null comment '用户名（Email）',
    password varchar(255)                not null,
    constraint ds
        check (`user_id` > 0),
    check (`user_id` is not null)
)
    comment '用户表';

INSERT INTO tiktok.users (user_id, username, password) VALUES (9996985457184768, '1234@qq.com', '313233343536373839a8a9671fc58e16b91f56835a88c28b3f');
INSERT INTO tiktok.users (user_id, username, password) VALUES (10181697941278720, 'konyue', '313131313131a8a9671fc58e16b91f56835a88c28b3f');
INSERT INTO tiktok.users (user_id, username, password) VALUES (10246285546229760, '123@qq.com', '3132333435363738a8a9671fc58e16b91f56835a88c28b3f');
INSERT INTO tiktok.users (user_id, username, password) VALUES (10265307285295104, 'admin12345', '6138613936373166633538653136623931663536383335613838633238623366505de3ad8e991662be2ac79fb3fd111a');
INSERT INTO tiktok.users (user_id, username, password) VALUES (10279411811094528, '12@qq.com', '3132333435363738a8a9671fc58e16b91f56835a88c28b3f');
INSERT INTO tiktok.users (user_id, username, password) VALUES (10279700773474304, 'mrxuexi@qq.com', '3132333435363738a8a9671fc58e16b91f56835a88c28b3f');

create table video
(
    video_id       bigint       not null
        primary key,
    user_id        bigint       null,
    cover_url      varchar(40)  null,
    favorite_count int          null,
    comment_count  int          null,
    is_favorite    tinyint(1)   null,
    date           int          null,
    name           varchar(255) null
)
    charset = utf8;

INSERT INTO tiktok.video (video_id, user_id, cover_url, favorite_count, comment_count, is_favorite, date, name) VALUES (1, 9996985457184768, '', 2, 3, 0, 1654995870, 'firstvideo');
INSERT INTO tiktok.video (video_id, user_id, cover_url, favorite_count, comment_count, is_favorite, date, name) VALUES (12567435622223872, 10246285546229760, '', 0, 0, 0, 1655569110, 'mmexport1642498809694.mp4');
