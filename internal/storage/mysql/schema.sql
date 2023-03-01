CREATE TABLE IF NOT EXISTS users (
  name varchar(100) NOT NULL,
  email varchar(100) NOT NULL,
  password varchar(100) NOT NULL,
  role varchar(100) NOT NULL,
  PRIMARY KEY (email)
);
CREATE TABLE IF NOT EXISTS  companies (
  id BINARY(16) NOT NULL,
  name varchar(15) NOT NULL,
  description varchar(3000) DEFAULT NULL,
  amountOfEmployees int NOT NULL,
  registered boolean NOT NULL,
  type varchar(30) NOT NULL,

  PRIMARY KEY (email)
);
DELETE FROM users WHERE email IN ('u1@xm.com','u2@xm.com');

INSERT INTO users VALUES ('u1','u1@xm.com','','admin');
INSERT INTO users VALUES ('u2','u2@xm.com','','user');
