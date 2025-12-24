INSERT INTO employees (first_name, last_name, id, hire_date, termination_date, salary) VALUES
('Bob', 'Smith', 1, '2009-06-20', '2016-01-01', 10000),
('Joe', 'Jarrod', 2, '2010-02-12', NULL, 20000),
('Nancy', 'Soley', 3, '2012-03-14', NULL, 30000),
('Keith', 'Widjaja', 4, '2013-09-10', '2014-01-01', 20000),
('Kelly', 'Smalls', 5, '2013-09-10', NULL, 20000),
('Frank', 'Nguyen', 6, '2015-04-10', '2015-05-01', 60000);

INSERT INTO annual_reviews (id, emp_id, review_date) VALUES
(10, 1, '2016-01-01'),
(20, 2, '2016-04-12'),
(30, 10, '2015-02-13'),
(40, 22, '2010-10-12'),
(50, 11, '2009-01-01'),
(60, 12, '2009-03-03'),
(70, 13, '2008-12-01'),
(80, 1, '2003-04-12'),
(90, 1, '2014-04-30');
