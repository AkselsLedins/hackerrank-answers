SELECT
    RIGHT(STUDENTS.NAME, 3)
FROM
    STUDENTS
WHERE
    STUDENTS.MARKS > 75
ORDER BY RIGHT(STUDENTS.NAME, 3), STUDENTS.ID ASC
