create table if not exists store (
    id uuid primary key,
    branch_id uuid references branches(id),
    profit numeric(100, 2),
    budget numeric(100, 2),
    created_at timestamp default now(),
    updated_at timestamp,
    deleted_at integer default 0
);

alter table users
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table categories
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table products
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table baskets
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table basket_products
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table branches
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;