CREATE TABLE IF NOT EXISTS accounts (
  email varchar(255),
  age int,
  PRIMARY KEY(email)
);

insert into accounts values ('bob@email.com', 55);