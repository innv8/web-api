create table user (
    id int(11) primary key auto_increment,
    first_name varchar(20) not null,
    last_name varchar(20) not null,
    phone_number varchar(20) not null unique,
    date_of_birth date not null,
    created datetime not null default current_timestamp,
    updated datetime null on update current_timestamp
);
