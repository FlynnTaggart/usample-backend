SET TIME ZONE 'Europe/Moscow';

ALTER TABLESPACE pg_global
    OWNER TO postgres;
ALTER TABLESPACE pg_default
    OWNER TO postgres;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create TABLE if not exists users (
    id uuid PRIMARY KEY NOT NULL,
    nickname varchar(64) NOT NULL,
    first_name varchar(64),
    second_name varchar(64),
    default_access_type smallint NOT NULL,
    user_type smallint NOT NULL,
    bio text
);

create TABLE if not exists user_links(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    link_type varchar(64) NOT NULL,
    link_url text,
    user_id uuid NOT NULL references users(id)
);

create TABLE if not exists friend_requests (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    sender_id uuid NOT NULL REFERENCES users(id),
    receiver_id uuid NOT NULL REFERENCES users(id),
    is_accepted bit NOT NULL
);

create TABLE if not exists samples (
    id uuid PRIMARY KEY NOT NULL,
    sample_name varchar(64) NOT NULL,
    file_path text NOT NULL,
    cover_file_path text,
    listens int NOT NULL, 
    private_comment text,
    descr text,
    access_type varchar(64),
    author_id uuid REFERENCES users(id)
);

create TABLE if not exists comments (
    id uuid PRIMARY KEY NOT NULL,
    author_id uuid NOT NULL REFERENCES users(id),
    comment_text text NOT NULL,
    sending_time TIMESTAMP NOT NULL,
    sample_id uuid REFERENCES samples(id),
    comment_id uuid REFERENCES comments(id)
);  

create TABLE if not exists likes (
    id uuid PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id),
    sending_time TIMESTAMP NOT NULL,
    comment_id uuid REFERENCES comments(id),
    sample_id uuid REFERENCES samples(id)
);

create TABLE if not exists sample_usages (
    id uuid PRIMARY KEY NOT NULL,
    usage_type varchar(64) NOT NULL,
    user_id uuid NOT NULL references users(id),
    sample_id uuid NOT NULL REFERENCES samples(id)
);

