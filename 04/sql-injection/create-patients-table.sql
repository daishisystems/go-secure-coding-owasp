CREATE DATABASE IF NOT EXISTS Globomantics;

USE Globomantics;

CREATE TABLE patients (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO patients (name, surname, age, gender)
VALUES
  ('Alice', 'Smith', 27, 'Female'),
  ('Bob', 'Johnson', 42, 'Male'),
  ('Charlie', 'Williams', 19, 'Male'),
  ('David', 'Brown', 34, 'Male'),
  ('Eve', 'Jones', 53, 'Female'),
  ('Frank', 'Davis', 44, 'Male');

select * from patients;