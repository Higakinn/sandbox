/*
 
 */
CREATE OR REPLACE VIEW test_data_view AS
SELECT '{"gender": "male", "prefecture": "Tokyo"}' meta
UNION
SELECT '{"gender": "female", "prefecture": "Osaka"}' meta
UNION
SELECT '{"gender": "male", "prefecture": "Nagoya"}' meta ;


SELECT JSON_PRETTY(meta) meta
FROM test_data_view
WHERE JSON_EXTRACT(meta, '$.gender') = 'male';
