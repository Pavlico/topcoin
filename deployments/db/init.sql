CREATE DATABASE IF NOT EXISTS coins;
USE coins;
CREATE TABLE IF NOT EXISTS topcoins (
  id int(11) NOT NULL AUTO_INCREMENT,
  symbol varchar(255),
  rank varchar(255),
  score varchar(255),
  CONSTRAINT topcoins PRIMARY KEY (id)
);