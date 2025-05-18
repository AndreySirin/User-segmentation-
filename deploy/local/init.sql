CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid PRIMARY KEY default uuid_generate_v4(),

    full_name VARCHAR NOT NULL,
    gender VARCHAR NOT NULL,
    date_of_birth DATE NOT NULL,

    create_at TIMESTAMP NOT NULL DEFAULT now(),
    update_at TIMESTAMP NOT NULL DEFAULT now(),
    delete_at TIMESTAMP
);

CREATE TABLE segments (
    id uuid PRIMARY KEY default uuid_generate_v4(),

    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL ,
    auto_user_prc SMALLINT NOT NULL,

    create_at TIMESTAMP NOT NULL DEFAULT now(),
    update_at TIMESTAMP NOT NULL DEFAULT now(),
    delete_at TIMESTAMP
);

CREATE TABLE subscriptions (
    PRIMARY KEY (user_id, segment_id),
    user_id uuid NOT NULL REFERENCES users(id),
    segment_id uuid NOT NULL REFERENCES segments(id),

    ttl TIMESTAMP,
    is_auto_add BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    update_at TIMESTAMP NOT NULL DEFAULT now(),
    delete_at TIMESTAMP
);
