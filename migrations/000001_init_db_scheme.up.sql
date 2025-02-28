CREATE TABLE "Songs"
(
    "id"           SERIAL PRIMARY KEY,
    "group"        VARCHAR(255) NOT NULL,
    "song"         VARCHAR(255) NOT NULL,
    "release_date" DATE  NOT NULL,
    "text"         TEXT         NOT NULL,
    "link"         VARCHAR(255) NOT NULL,
    UNIQUE ("group", "song", "link")
);
