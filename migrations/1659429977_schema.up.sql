BEGIN;
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
                       user_id       INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
                       name          TEXT NOT NULL,
                       username      TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS audios (
                        audio_id    INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
                        user_id     INTEGER REFERENCES users(user_id) NOT NULL,
                        title       TEXT NOT NULL,
                        duration    INTEGER NOT NULL,
                        file_path   TEXT UNIQUE NOT NULL,
                        UNIQUE (audio_id, user_id)
);

CREATE TABLE IF NOT EXISTS shares (
                        audio_id    INTEGER REFERENCES audios(audio_id) ON DELETE CASCADE NOT NULL,
                        user_id     INTEGER REFERENCES users(user_id) ON DELETE CASCADE NOT NULL,
                        UNIQUE(audio_id, user_id)
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
                        user_id     INTEGER REFERENCES users(user_id) ON DELETE CASCADE NOT NULL UNIQUE,
                        refresh_token uuid NOT NULL,
                        expires_in  timestamp  with time zone NOT NULL
);
COMMIT;