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
*/

with series as (
  select generate_series v from generate_series(0,10)
)
select s1.v as i, s2.v as j 
from series s1
inner join series s2 on 1 = 1
