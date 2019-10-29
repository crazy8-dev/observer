create table simple_transactions
(
    id bigserial not null constraint simple_transactions_pkey primary key,
    tx_id bytea not null constraint simple_transactions_tx_id_key unique,

    status_registered bool,
    pulse_record bigint[2] unique,
    member_from_ref bytea,
    member_to_ref bytea,
    deposit_to_ref bytea,
    deposit_from_ref bytea,
    amount varchar(256),
    fee varchar(256),

    status_sent bool,

    status_finished bool,
    finish_success bool,
    finish_pulse_record bigint[2] unique
);
