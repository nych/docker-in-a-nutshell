CREATE DATABASE IF NOT EXISTS app_db;

USE app_db;

CREATE TABLE IF NOT EXISTS hit (
    id          INT     NOT NULL,
    hit_count   INT,
    PRIMARY KEY (id)
);

INSERT IGNORE INTO hit(id, hit_count)
VALUES(0, 0) ON DUPLICATE KEY
UPDATE
  id = 0,
  hit_count = 0;
