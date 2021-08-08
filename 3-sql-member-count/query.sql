SELECT
  ReportsTo,
  COUNT(ReportsTo) AS "Members",
  CAST(AVG(Age) AS DECIMAL(3)) AS "Average Age"
FROM
  maintable_D9O9V
WHERE
  ReportsTo IS NOT NULL
GROUP BY
  ReportsTo
ORDER BY
  ReportsTo ASC;
