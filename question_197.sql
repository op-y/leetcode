SELECT T.id FROM Weather AS S
INNER JOIN Weather AS T
ON S.recordDate = DATE_SUB(T.recordDate, INTERVAL 1 DAY) AND S.Temperature < T.Temperature;
