# Tik tok 数据库设计

## Table

### user

| 数据库字段 | 服务端结构体字段 | Description     |
| ---------- | ---------------- | --------------- |
| user_id    | UserId           | 主键（加密）    |
| username   | Username         | 用户名（email） |
| password   | Password         | 密码            |

### vedio

| 数据库字段 | 服务端结构体字段 | Description  |
| ---------- | ---------------- | ------------ |
| vedio_id   | VedioId          | 主键（加密） |
| user_id    | UserID           | 视频作者     |
| title      | Title            | 视频标题     |
| play_url   | PlayUrl          | 视频播放地址 |
| cover_url  | CoverUrl         | 视频封面地址 |
|            |                  |              |

### favorite

| 数据库字段  | 服务端结构体字段 | Description  |
| ----------- | ---------------- | ------------ |
| favorite_id | Favorite_id      | 主键         |
| user_id     | UserID           | 用户id       |
| video_id    | video_id         | 视频id       |

### relation
| 数据库字段  | 服务端结构体字段 | Description  |
| ----------- | ---------------- | ------------ |
|relation_id  |RelationID        |主键
| user_id     | UserID           | 用户         |
| to_user_id  | ToUserID         | 喜欢用户     |
