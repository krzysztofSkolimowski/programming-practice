-- +migrate Up
create table if not exists aptTenants (
    id int auto_increment,
    tenantID int,
    aptID int,
    primary key (id),
    foreign key fk_apartments (aptID)
        references apartments(aptID)
        on delete cascade
        on update cascade,
    foreign key fk_tenants (tenantID)
        references tenants(tenantID)
        on delete cascade
        on update cascade
);

-- +migrate Down
drop table aptTenants;