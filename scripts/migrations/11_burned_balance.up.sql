create table if not exists burned_balance
(
    id bigint generated by default as identity primary key,
    balance varchar(256),
    account_state bytea
);
