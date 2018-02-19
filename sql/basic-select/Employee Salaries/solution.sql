SELECT
    employee.name
FROM
    employee
WHERE
    employee.salary > 2000
    AND employee.months < 10
ORDER BY employee.employee_id ASC
