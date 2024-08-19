create table usermeta
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)  null,
    updated_at datetime(3)  null,
    deleted_at datetime(3)  null,
    uuid       char(36)     not null,
    meta_key   varchar(255) null,
    meta_value text         null,
    meta_json  json         null,
    user_uuid  longtext     null
);

create index idx_usermeta_deleted_at
    on usermeta (deleted_at);

create index idx_usermeta_uuid
    on usermeta (uuid);

create table users
(
    id         bigint unsigned auto_increment,
    created_at datetime(3)      null,
    updated_at datetime(3)      null,
    deleted_at datetime(3)      null,
    uuid       char(36)         not null,
    email      varchar(255)     not null,
    username   varchar(255)     not null,
    password   varchar(255)     not null,
    status     bigint default 1 null,
    primary key (id, uuid),
    constraint uni_users_email
        unique (email),
    constraint uni_users_username
        unique (username)
);

create index idx_users_deleted_at
    on users (deleted_at);

create index idx_users_uuid
    on users (uuid);
