-- +migrate Up
create table if not exists apartments (
    aptID      int auto_increment,
    unitNumber varchar(10),
    buildingID int,
    primary key (aptID),
    foreign key fk_buildings (buildingID)
        references buildings (buildingID)
        on delete cascade
        on update cascade
);

-- +migrate Down
drop table apartments;
