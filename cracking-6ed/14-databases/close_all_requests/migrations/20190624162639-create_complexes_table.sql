-- +migrate Up
create table if not exists complexes (
    complexID   int auto_increment,
    complexName varchar(100),
    primary key (complexID)
);
-- +migrate Down

drop table complexes;
