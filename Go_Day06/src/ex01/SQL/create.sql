
CREATE TYPE role AS ENUM ('User', 'Admin', 'Moderator', 'Root');

create table Users (
                       ID bigserial primary key unique,
                       Login varchar,
                       Password varchar,
                       Role role DEFAULT 'User'
);

create table Articles (
                          ID bigserial primary key unique,
                          IDAuthor int,
                          Title varchar,
                          Content varchar,
                          Created timestamp,
                          foreign key (IDAuthor) REFERENCES Users(ID)
);

insert into users (Login, Password, Role)
values ('admin','827ccb0eea8a706c4c34a16891f84e7b', 'Admin');