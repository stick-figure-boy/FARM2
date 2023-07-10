CREATE TABLE IF NOT EXISTS users (
    account_id VARCHAR(16) NOT NULL PRIMARY KEY COMMENT 'account id',
    name VARCHAR(32) NOT NULL COMMENT 'user name',
    created_at DATETIME default CURRENT_TIMESTAMP,
    updated_at DATETIME default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
) COMMENT = 'user info table'