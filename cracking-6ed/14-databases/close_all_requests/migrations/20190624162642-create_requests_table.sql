-- +migrate Up
create table if not exists requests (
    requestID   int auto_increment,
    status      varchar(100),
    description varchar(100),
    aptID       int,
    primary key (requestID),
    foreign key fk_apartments (aptID)
        references apartments (aptID)
        on delete cascade
        on update cascade
);

-- +migrate Down
drop table requests;
