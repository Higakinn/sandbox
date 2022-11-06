with
-- testデータ
test_data as (
    select gen_random_uuid() as id, 'test1' as name
    union
        select gen_random_uuid() as id, 'test2' as name
    union
        select gen_random_uuid() as id, 'test3' as name
    union
        select gen_random_uuid() as id, 'test4' as name
    union
        select gen_random_uuid() as id, 'test5' as name
)
select * from test_data;


-- with句使えない場合はviewで代用
create or replace view test_data_view as 
    select gen_random_uuid() as id, 'test1' as name
    union
        select gen_random_uuid() as id, 'test2' as name
    union
        select gen_random_uuid() as id, 'test3' as name
    union
        select gen_random_uuid() as id, 'test4' as name
    union
        select gen_random_uuid() as id, 'test5' as name
;
select * from test_data_view;