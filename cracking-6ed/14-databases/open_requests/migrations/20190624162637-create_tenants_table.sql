-- +migrate Up
create table if not exists tenants (
    tenantID   int auto_increment,
    tenantName varchar(100),
    primary key (tenantID)
);

-- +migrate Down
drop table tenants;
