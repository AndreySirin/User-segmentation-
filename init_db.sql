
CREATE TABLE users (
                       id_user SERIAL PRIMARY KEY
);
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;


CREATE TABLE segments (
                          id_segment SERIAL PRIMARY KEY,
                          title TEXT NOT NULL UNIQUE
);

INSERT INTO segments (title) VALUES
                                 ('Segment A'),
                                 ('Segment B'),
                                 ('Segment C');


CREATE TABLE subscriptions (
                               PRIMARY KEY (id_user, id_segment),
                               id_user INTEGER NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
                               id_segment INTEGER NOT NULL REFERENCES segments(id_segment) ON DELETE CASCADE,
                               created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               expires_at TIMESTAMP NOT NULL
);
