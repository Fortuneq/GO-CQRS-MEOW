Drop table if exists meows;
Create table meows(
    id Varchar(32) Primary key,
    body text not null,
    created_at timestamp with time zone not null
)