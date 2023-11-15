SELECT t.name AS Employee, t.salary AS Salary, t.Department AS Department FROM (
SELECT
  e1.name,e1.salary,D.name AS Department,
  (
    SELECT
      COUNT(DISTINCT e2.salary)
    FROM
      Employee e2
    WHERE
      e2.salary >= e1.salary and e1.departmentId = e2.departmentId
  ) AS 'rank'
FROM
  Employee e1
INNER JOIN Department AS D
ON e1.departmentId=D.id
ORDER BY e1.salary DESC) t
WHERE t.rank <= 3;
