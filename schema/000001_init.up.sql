create table users(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    email varchar(255) not null,
    password_hash varchar(255) not null
);

create table groups(
    id serial primary key,
    name varchar(255) not null
);

create table users_groups(
    user_id int not null,
    group_id int not null,
    
    primary key (user_id, group_id),
    foreign key (user_id) references users(id),
    foreign key (group_id) references groups(id)
);

create table debts(
    debtor int not null,
    creditor int not null,
    amount int not null,

    primary key (debtor, creditor),
    foreign key (debtor) references users(id),
    foreign key (creditor) references users(id)
);

create table purchases(
    id serial primary key,
    group_id int not null,
    cost int not null,
    buyer int not null,
    description varchar(255) not null,

    foreign key (group_id) references groups(id),
    foreign key (buyer) references users(id)
);