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

create table products
(
    id           int         not null auto_increment,
    product_name varchar(50) not null,
    description  varchar(50) not null,
    price        float       not null,
    created_by   id          not null,
    primary key (id),
    foreign key (created_by) references users (id)
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