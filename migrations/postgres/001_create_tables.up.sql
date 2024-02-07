create type user_role_enum as enum ('admin', 'customer');

create table users (
    id uuid primary key not null,
    full_name varchar(30),
    phone varchar(30) unique not null,
    password varchar(30) not null,
    user_role user_role_enum not null,
    cash int
);

create table categories (
    id uuid primary key not null,
    name varchar(30) not null
);

create table products (
    id uuid primary key not null,
    name varchar(30) not null,
    price int default 0,
    original_price int default 0,
    quantity int default 0,
    category_id uuid references categories(id)
);

create table baskets (
    id uuid primary key not null,
    customer_id uuid references users(id) not null,
    total_sum integer default 0
);

create table basket_products (
    id uuid primary key not null,
    basket_id uuid references baskets(id),
    product_id uuid references products(id),
    quantity int
);