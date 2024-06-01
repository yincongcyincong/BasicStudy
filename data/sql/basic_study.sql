CREATE TABLE `study_user`
(
    `id`          INT PRIMARY KEY NOT NULL,
    `create_time` INT             NOT NULL,
    `username`    TEXT DEFAULT '' NOT NULL,
    `password`    TEXT DEFAULT '' NOT NULL,
    `status`      INT  DEFAULT 1  NOT NULL --'1-normal 2-deleted'
);

CREATE TABLE `study_question`
(
    `id`          INT PRIMARY KEY NOT NULL,
    `create_time` INT             NOT NULL,
    `question`    TEXT            NOT NULL,
    `answer`      TEXT            NOT NULL,
    `test_id`     INT             NOT NULL,
    `type`        INT             NOT NULL --'1-math 2-english '
);

CREATE TABLE `study_test`
(
    `id`   INT PRIMARY KEY NOT NULL,
    `name` TEXT DEFAULT '' NOT NULL
);