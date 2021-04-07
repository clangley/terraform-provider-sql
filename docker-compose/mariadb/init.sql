CREATE TABLE IF NOT EXISTS test.accounts (
  email varchar(255),
  age int,
  PRIMARY KEY(email)
);
insert into test.accounts values ('bob@email.com', 55);