SELECT
  S.score,
  COUNT(DISTINCT T.score) AS 'rank'
FROM
  scores S
  INNER JOIN scores T ON S.score <= T.score
GROUP BY
  S.id,
  S.score
ORDER BY
  S.score DESC;
