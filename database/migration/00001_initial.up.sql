create TABLE if not exists tasks(
    id text primary key,
    title text not null,
    description text not null,
    completed bool not null,
    due_date timestamp not null,
    created_at timestamp default current_timestamp
)