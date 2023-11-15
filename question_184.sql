SELECT Employee.name AS Employee, Employee.salary as Salary, t.dname AS Department
FROM Employee
INNER JOIN
(SELECT MAX(Employee.salary) AS dmax, Employee.departmentId AS did, Department.name AS dname FROM Employee
INNER JOIN Department ON Employee.departmentId=Department.id 
GROUP BY Employee.departmentId) AS t 
ON Employee.departmentId=t.did AND Employee.salary=t.dmax;
