create table if not exists incomes (
    id uuid primary key,
    external_id varchar(6) unique not null,
    total_sum integer default 0,
    created_at timestamp default now(),
    deleted_at integer default 0
);

create table if not exists income_products (
    id uuid primary key,
    income_id uuid references incomes(id),
    product_id uuid references products(id),
    quantity int not null,
    price int not null,
    created_at timestamp default now(),
    updated_at timestamp,
    deleted_at integer default 0
);