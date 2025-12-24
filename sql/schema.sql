DROP TABLE IF EXISTS annual_reviews;
DROP TABLE IF EXISTS employees;

CREATE TABLE employees (
    id INT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    hire_date DATE NOT NULL,
    termination_date DATE,
    salary INT NOT NULL
);

CREATE TABLE annual_reviews (
    id INT PRIMARY KEY,
    emp_id INT NOT NULL,
    review_date DATE NOT NULL
);
