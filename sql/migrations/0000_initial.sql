-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE api_audit (
    ts bigint PRIMARY KEY NOT NULL,
    ip_address varchar(255) NOT NULL,
    method varchar(10) NOT NULL,
    request_path varchar(255) NOT NULL,
    status int NOT NULL,
    user_agent varchar(255)
);

CREATE TABLE example (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    occupation varchar(255) NOT NULL,
    telephone varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    notes varchar(255) NOT NULL
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE api_audit;

DROP TABLE example;

