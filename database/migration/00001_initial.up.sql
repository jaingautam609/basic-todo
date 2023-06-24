create TABLE if not exists tasks(
    id text primary key,
    title text not null,
    description text not null,
    completed bool not null
)