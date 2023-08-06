CREATE TABLE IF NOT EXISTS scores(
    id serial PRIMARY KEY,
    name VARCHAR (255) NOT NULL COMMENT 'ユーザー名',
    score DECIMAL(7, 2) NOT NULL COMMENT 'スコア'
);
