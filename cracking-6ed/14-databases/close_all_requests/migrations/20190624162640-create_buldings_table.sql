-- +migrate Up
create table if not exists buildings (
    buildingID   int auto_increment,
    complexID    int,
    buildingName varchar(100),
    address      varchar(500),
    primary key (buildingID),
    foreign key fk_complexes (complexID)
        references complexes (complexID)
        on delete cascade
        on update cascade
);

-- +migrate Down
drop table buildings;