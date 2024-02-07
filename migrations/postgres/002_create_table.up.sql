create table if not exists branches (
    id uuid primary key not null,
    name varchar(30),
    address text unique,
    phone_number varchar(13) unique
);

insert into branches
    (id, name, address, phone_number)
        values ('aa541fcc-bf74-11ee-ae0b-166244b65504', 'Main Branch', 'Address', '+998123456789');

alter table users
    add column if not exists branch_id uuid references branches(id) default 'aa541fcc-bf74-11ee-ae0b-166244b65504';

alter table products
    add column if not exists branch_id uuid references branches(id) default 'aa541fcc-bf74-11ee-ae0b-166244b65504';