alter table products
    drop column if exists branch_id;

alter table users
    drop column if exists branch_id;

drop table if exists branches;