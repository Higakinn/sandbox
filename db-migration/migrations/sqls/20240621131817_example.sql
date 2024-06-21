-- +goose Up
CREATE TABLE post (
    id int NOT NULL comment '投稿ID',
    title text comment '投稿タイトル',
    body text comment '投稿内容',
    PRIMARY KEY(id)
) comment='投稿テーブル';

-- +goose Down
DROP TABLE post;