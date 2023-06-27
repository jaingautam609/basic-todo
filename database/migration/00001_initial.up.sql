create TABLE if not exists users(
                                    id serial primary key not null,
                                    username text not null unique,
                                    password text not null
);

create TABLE if not exists tasks(
                                    id serial primary key not null,
                                    title text not null,
                                    description text not null,
                                    completed bool not null,
                                    due_date timestamp not null,
                                    created_at timestamp default current_timestamp,
                                    userid int REFERENCES users(id)
    );

create table if not exists sessions(
                                       logedInAt timestamp default current_timestamp,
                                       token text primary key not null,
                                       user_id int REFERENCES users(id) not null
    );