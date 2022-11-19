/*
for i in range(3):
    for j in range(3):
      print(i,j)
0 0
0 1
0 2
1 0
1 1
1 2
2 0
2 1
2 2
*/ WITH series AS
  (SELECT generate_series v
   FROM generate_series(0, 10))
SELECT s1.v AS i
     , s2.v AS j
FROM series s1
INNER JOIN series s2 ON 1 = 1
