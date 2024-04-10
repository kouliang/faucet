-- Active: 1710308193473@@127.0.0.1@3306
CREATE DATABASE IF NOT EXISTS faucet
    DEFAULT CHARACTER SET = 'utf8mb4';

USE faucet;


CREATE TABLE IF NOT EXISTS task (
  address VARCHAR(255),
  timestamp INTEGER,
  taskid INTEGER,
  hash VARCHAR(255),
  PRIMARY KEY (address, timestamp, taskid)
);