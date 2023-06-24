create TABLE if not exists todo(
    if text primary key,
    title text not null,
    description text not null,
    completed bool not null
)