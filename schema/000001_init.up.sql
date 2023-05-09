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
    creditor_id int not null,
    debtor_id int not null,
    amount float not null,

    primary key (debtor_id, creditor_id),
    foreign key (debtor_id) references users(id),
    foreign key (creditor_id) references users(id)
);

create table purchases(
    id serial primary key,
    group_id int not null,
    amount float not null,
    buyer_id int not null,
    description varchar(255) not null,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    foreign key (group_id) references groups(id),
    foreign key (buyer_id) references users(id)
);