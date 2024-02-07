drop table if exists store;

alter table users
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table categories
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table products
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table baskets
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table basket_products
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table branches
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;