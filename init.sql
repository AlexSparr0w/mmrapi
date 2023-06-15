create table Users (
"user_id" varchar(60) primary key not null,
"email" varchar(60) not null,
"password" varchar(255) not null,
"total_matches" int,
"total_mmr" varchar(255),
"created_at" date not null
);

CREATE USER apiuser WITH PASSWORD 'dota2govno';
GRANT ALL PRIVILEGES ON DATABASE api TO apiuser;