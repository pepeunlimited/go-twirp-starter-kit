CREATE DATABASE IF NOT EXISTS hello_world CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE hello_world;

CREATE TABLE hello (
    id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE world (
    id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    hello_id BIGINT NOT NULL
    FOREIGN KEY (hello_id) REFERENCES hello (id),
    PRIMARY KEY (id)
);

-- add columns to a table using ADD COLUMN Statement
-- ALTER TABLE your_table_name
-- ADD COLUMN column_name INT NOT NULL;

-- rename existing column
-- ALTER TABLE your_table_name
-- RENAME COLUMN original_column_name TO new_column_name;

-- drop column
-- ALTER TABLE your_table_name
-- DROP COLUMN column_name;

-- ALTER TABLE your_table_name
-- MODIFY column_name tinyint(4) NOT NULL;