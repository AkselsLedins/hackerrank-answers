SELECT
    DISTINCT(STATION.CITY)
FROM
    STATION
WHERE
    STATION.CITY REGEXP "^[^aeiou].*$"
;
