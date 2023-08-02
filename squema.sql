create database test;
use test;

create table users
(
    id        int         not null auto_increment,
    user_name varchar(50) not null,
    email     varchar(50) not null,
    password  varchar(50) not null,
    primary key (id)
);


create table roles(
    id int not null auto_increment,
    role_name varchar(50) not null,
    primary key (id)
);

create table user_roles(
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) references users,
    foreign key (role_id) references roles
);