/*
前年度の年収と比較して、変更なしの場合のみ表示する。


*/
-- テストデータ
with sales as (
  select 
    1990 as year -- 年度
    ,50   as sale -- 年商
  union
  select 1991 as year, 51 as sale
  union
  select 1992 as year, 52 as sale
  union
  select 1993 as year, 52 as sale
  union
  select 1994 as year, 50 as sale
  union
  select 1995 as year, 50 as sale
  union
  select 1996 as year, 49 as sale
  union
  select 1997 as year, 55 as sale

)
select year,sale
from sales s1
where sale = (
  select sale 
  from sales s2 
  where s2.year = s1.year - 1
)
order by year;