create table if not exists dealer (
    id uuid primary key not null ,
    name varchar(30) ,
    sum integer
);

insert into dealer values ('1cfd84e6-72cb-4135-a802-85d10e4183ea', 'Dealer 1', 0);