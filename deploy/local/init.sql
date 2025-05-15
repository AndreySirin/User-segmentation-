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

INSERT INTO users (full_name, gender, date_of_birth)
VALUES
    ('Alice Ivanova', 'female', '1990-05-01'),
    ('Bob Petrov', 'male', '1985-09-15'),
    ('Charlie Smirnov', 'male', '1992-12-25');

-- 🔹 Вставка сегментов
INSERT INTO segments (title, description, auto_user_prc)
VALUES
    ('Segment A', 'First segment for testing', 10),
    ('Segment B', 'Second segment for testing', 25),
    ('Segment C', 'Third segment for testing', 50);

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
