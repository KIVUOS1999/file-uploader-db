create database uploader_file_details

create table chunk_details (
    chunk_id       VARCHAR(255) PRIMARY KEY,
    file_id   VARCHAR(255),
    check_sum VARCHAR(255),
    chunk_order  INT
    created_at INT
);

create table file_details (
    file_id VARCHAR(255) PRIMARY KEY,
    file_name VARCHAR(255),
    file_size INT,
    total_chunks INT,
    created_at INT
);

\l -> show database
\c db_name -> connect to a db
\dt -> show table
\d table_name -> describe table

goose postgres postgresql://postgres:password@localhost:5432/uploader_file_details up