CREATE TABLE `study_user`
(
    `id`          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    `create_time` INT                     NOT NULL,
    `username`    VARCHAR(100) DEFAULT '' NOT NULL,
    `password`    VARCHAR(100) DEFAULT '' NOT NULL,
    `status`      INT          DEFAULT 1  NOT NULL --'1-normal 2-deleted'
);

INSERT INTO `study_user` (`create_time`, `username`, `password`) VALUES ( 1717294770, 'admin', '21232f297a57a5a743894a0e4a801fc3');

CREATE TABLE `study_question`
(
    `id`          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    `create_time` INT NOT NULL,
    `question`    TEXT DEFAULT '',
    `answer`      TEXT DEFAULT '',
    `test_id`     INT  DEFAULT 0,
    `type`        INT NOT NULL --'1-math 2-english '
);

CREATE TABLE `study_test`
(
    `id`   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    `name` VARCHAR(100) DEFAULT '' NOT NULL
);