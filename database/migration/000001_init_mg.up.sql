create table storages
(
    id   uuid not null,
    name text not null,
    constraint storages_pk primary key (id)
);

create table products
(
    id   uuid not null,
    name text not null ,
    presentation varchar(10),
    constraint products_pk primary key (id)
);

create table items (
    storage_id uuid not null ,
    product_id uuid not null ,
    quantity int not null default 0,
    constraint items_pk primary key (storage_id, product_id),
    constraint items_storages_fk foreign key (storage_id) references storages(id),
    constraint items_products_fk foreign key (product_id) references products(id)
);

create table transactions(
    id serial not null ,
    type varchar(10) not null ,
    storage_id uuid not null ,
    product_id uuid not null ,
    quantity int not null ,
    user_id text not null ,
    created_at timestamptz not null default now()
);