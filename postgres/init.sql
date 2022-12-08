drop DATABASE if exists usampleDB;
create DATABASE usampleDB;
use usampleDB;
set names 'utf8';

create TABLE users (
    id_users bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    nickname varchar(64) NOT NULL,
    first_name varchar(64),
    second_name varchar(64),
    email varchar(64) NOT NULL,
    password varchar(64) NOT NULL,
    default_access_type varchar(64) NOT NULL,
    user_type bit NOT NULL,
    bio text
);

create TABLE user_links(
    id_user_links bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    link_type varchar(64) NOT NULL,
    link_url text,
    user_id bigint NOT NULL references users(id_users)
);

create TABLE chats(
    id_chats bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_a bigint NOT NULL references users(id_users),
    user_b bigint NOT NULL references users(id_users),
    constraint chk_diff_users CHECK (user_a <> user_b)
);

create TABLE messages (
    id_messages bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    sender_id bigint NOT NULL REFERENCES users(id_users),
    message_text text NOT NULL,
    sending_time TIMESTAMP NOT NULL,
    chat_id bigint references chats(id_chats)
);

create TABLE chat_participations (
    id_chat_participations bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_id bigint NOT NULL REFERENCES users(id_users),
    chat_id bigint NOT NULL REFERENCES chats(id_chats)
);

create TABLE friend_requests (
    id_friend_requests bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    sender_id bigint NOT NULL REFERENCES users(id_users),
    reciever_id bigint NOT NULL REFERENCES users(id_users),
    is_accepted bit NOT NULL
);

create TABLE samples (
    id_samples bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    sample_name varchar(64) NOT NULL,
    file_path text NOT NULL,
    cover_file_path text,
    listens int NOT NULL, 
    private_comment text,
    descr text,
    access_type varchar(64),
    author_id bigint REFERENCES users(id_users)
);

create TABLE comments (
    id_comments bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    author_id bigint NOT NULL REFERENCES users(id_users),
    comment_text text NOT NULL,
    sending_time TIMESTAMP NOT NULL,
    sample_id bigint REFERENCES samples(id_samples),
    comment_id bigint REFERENCES comments(id_comments)
);  

create TABLE likes (
    id_likes bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_id bigint NOT NULL REFERENCES users(id_users),
    sending_time TIMESTAMP NOT NULL,
    comment_id bigint REFERENCES comments(id_comments),
    sample_id bigint REFERENCES samples(id_samples)
);

create TABLE sample_usages (
    id_sample_usages bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    usage_type varchar(64) NOT NULL,
    user_id bigint NOT NULL references users(id_users),
    sample_id bigint NOT NULL REFERENCES samples(id_samples)
);

