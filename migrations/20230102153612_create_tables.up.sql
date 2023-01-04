-- Active: 1672673123285@@127.0.0.1@5432@application

create table
    todos (
        id SERIAL primary key not null,
        todo_text VARCHAR(200) not null,
        is_complete BOOLEAN default false
    );
