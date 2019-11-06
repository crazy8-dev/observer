create table if not exists records
(
    key bytea not null
        constraint records_id_pk
            primary key,
    value bytea not null,
    pulse bigint,
    key_debug text,
    value_type text
);

create table if not exists raw_requests
(
    request_id varchar(256) not null constraint raw_requests_pk
        primary key,
    reason_id varchar(256) not null,
    request_body bytea not null
);

create table if not exists raw_results
(
    request_id varchar(256) not null
        constraint raw_results_pk
            primary key,
    result_body bytea not null
);

create table if not exists raw_side_effects
(
    id varchar(256) not null
        constraint raw_side_effects_pk
            primary key,
    request_id varchar(256) not null,
    side_effect_body bytea not null
);

create index if not exists idx_raw_side_effects_by_request_id
    on raw_requests (request_id);

create table if not exists objects
(
    object_id varchar(256) not null
        constraint objects_pkey
            primary key,
    domain varchar(256),
    request varchar(256),
    memory text,
    image varchar(256),
    parent varchar(256),
    prev_state varchar(256),
    type varchar(256)
);

create table if not exists requests
(
    request_id varchar(256) not null
        constraint requests_pkey
            primary key,
    caller varchar(256),
    return_mode varchar(256),
    base varchar(256),
    object varchar(256),
    prototype varchar(256),
    method text,
    arguments text,
    reason varchar(256)
);

create table if not exists results
(
    result_id varchar(256) not null
        constraint results_pkey
            primary key,
    request varchar(256),
    payload text
);

create table if not exists fees
(
    id bigint generated by default as identity
        constraint fees_pkey
            primary key,
    start_sum varchar(256),
    fin_sum varchar(256),
    percent varchar(128),
    min_amount varchar(256)
);

create table if not exists pulses
(
    pulse bigint not null
        constraint pulses_pkey
            primary key,
    pulse_date bigint,
    entropy varchar(256),
    requests_count integer
);

create table if not exists migration_addresses
(
    addr varchar(256),
    timestamp bigint, -- time when the address was ADDED (not assigned). API uses this field for sorting
    wasted boolean, -- if `true` the address was assigned, `false` otherwise
    id bigint generated by default as identity
);

create index if not exists idx_migration_addresses_addr
    on migration_addresses (addr);

create table if not exists members
(
    member_ref bytea not null
        constraint members_pkey
            primary key,
    balance varchar(256),
    migration_address varchar(256)
        constraint members_migration_address_key
            unique,
    status varchar(256),
    wallet_ref bytea,
    account_state bytea,
    account_ref bytea
);

create table if not exists deposits
(
    eth_hash varchar(256) not null,
    deposit_ref bytea,
    member_ref bytea not null,
    transfer_date bigint,
    hold_release_date bigint,
    amount varchar(256),
    balance varchar(256),
    deposit_state bytea,
    vesting bigint,
    vesting_step bigint,
    constraint deposits_pk
        primary key (member_ref, eth_hash)
);

create table if not exists transactions
(
    id bigserial not null
        constraint transactions_pkey
            primary key,
    tx_id bytea
        constraint transactions_tx_id_key
            unique,
    amount varchar(256),
    fee varchar(256),
    transfer_date bigint,
    pulse_num bigint,
    member_from_ref bytea,
    member_to_ref bytea,
    wallet_from_ref bytea,
    wallet_to_ref bytea,
    status varchar(256),
    eth_hash varchar(256),
    transfer_request_member bytea,
    transfer_request_wallet bytea,
    transfer_request_account bytea,
    accept_request_member bytea,
    accept_request_wallet bytea,
    accept_request_account bytea,
    calc_fee_request bytea,
    fee_member_request bytea,
    cost_center_ref bytea,
    fee_member_ref bytea
);

create index if not exists idx_transactions_eth_hash
    on transactions (eth_hash);

create index if not exists idx_transactions_member_from_ref
    on transactions (member_from_ref);

create index if not exists idx_transactions_member_to_ref
    on transactions (member_to_ref);

create table if not exists blockchain_stats
(
    pulse_num bigint not null
        constraint blockchain_stats_pkey
            primary key,
    total_transactions bigint,
    total_accounts bigint,
    nodes bigint,
    count_transactions bigint,
    max_transactions bigint,
    last_month_transactions bigint
);


