CREATE DATABASE simple_bank;

DROP TABLE accounts;
DROP TABLE entries;
DROP TABLE transfers;

CREATE TABLE accounts (
    id int NOT NULL,
    owner varchar NOT NULL,
    balance float NOT NULL,
    currency varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    PRIMARY KEY (id)
);

CREATE TABLE entries (
    id int NOT NULL,
    account_id int NOT NULL,
    amount float NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE transfers (
    id int NOT NULL,
    from_account_id int NOT NULL,
    to_account_id int NOT NULL,
    amount float NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    PRIMARY KEY (id),
    FOREIGN KEY (from_account_id) REFERENCES accounts(id),
    FOREIGN KEY (to_account_id) REFERENCES accounts(id)
);

CREATE INDEX ON accounts (owner);
CREATE INDEX ON entries (account_id);
CREATE INDEX ON transfers (to_account_id);
CREATE INDEX ON transfers (from_account_id, to_account_id);