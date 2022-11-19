/*
  playerテーブルのダミーデータ
  status_idに対する 名前がdb上に存在していない場合に、case文を利用すると便利。

  id: name
  1: alive
  2: recovering
  3: dying
  4: dead
*/ WITH player_data AS
  (SELECT 1 AS player_id
        , 'player1' AS player_name
        , 1 AS status_id
   UNION SELECT 2 AS player_id
              , 'player2' AS player_name
              , 2 AS status_id
   UNION SELECT 3 AS player_id
              , 'player3' AS player_name
              , 3 AS status_id
   UNION SELECT 4 AS player_id
              , 'player4' AS player_name
              , 4 AS status_id)
SELECT *
     ,  CASE
            WHEN status_id=1 THEN 'alive'
            WHEN status_id=2 THEN 'recovering'
            WHEN status_id=3 THEN 'dying'
            ELSE 'dead'
        END status_mame
FROM player_data;
