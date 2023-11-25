-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE owners (
    id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY
);

CREATE TABLE blobs(
    id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    owner_id UUID NOT NULL,
    value jsonb NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES owners (id)
);

CREATE INDEX blobs_owner_id_idx ON blobs (owner_id);

INSERT INTO owners (id) VALUES
(uuid_generate_v4()),
(uuid_generate_v4()),
(uuid_generate_v4()),
(uuid_generate_v4()),
(uuid_generate_v4()),
(uuid_generate_v4()),
(uuid_generate_v4()),
(uuid_generate_v4()),
(uuid_generate_v4()),
(uuid_generate_v4());

INSERT INTO blobs VALUES
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_1"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_2"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_3"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_4"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_5"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_6"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_7"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_8"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_9"}'),
(uuid_generate_v4(), (SELECT id FROM owners ORDER BY random() LIMIT 1), '{"Val": "random_value_10"}');

-- +migrate Down
DROP TABLE blobs;
DROP TABLE owners;
