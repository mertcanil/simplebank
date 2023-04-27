CREATE TABLE accounts (
    id serial NOT NULL,
    owner varchar NOT NULL,
    balance bigint NOT NULL,
    currency varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    PRIMARY KEY (id)
);

CREATE TABLE entries (
    id serial NOT NULL,
    account_id int NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE transfers (
    id serial NOT NULL,
    from_account_id int NOT NULL,
    to_account_id int NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    PRIMARY KEY (id),
    FOREIGN KEY (from_account_id) REFERENCES accounts(id),
    FOREIGN KEY (to_account_id) REFERENCES accounts(id)
);

CREATE INDEX ON accounts (owner);
CREATE INDEX ON entries (account_id);
CREATE INDEX ON transfers (to_account_id);
CREATE INDEX ON transfers (from_account_id, to_account_id);