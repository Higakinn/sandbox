/*
  playerテーブルのダミーデータ
  status_idに対する 名前がdb上に存在していない場合に、case文を利用すると便利。

  id: name
  1: alive
  2: recovering
  3: dying
  4: dead
*/
with player_data as (
  select 1 as player_id, 'player1' as player_name, 1 as status_id
  union
  select 2 as player_id, 'player2' as player_name, 2 as status_id
  union
  select 3 as player_id, 'player3' as player_name, 3 as status_id
  union
  select 4 as player_id, 'player4' as player_name, 4 as status_id

)
select 
  *,
  case
    when status_id=1 then 'alive'
    when status_id=2 then 'recovering'
    when status_id=3 then 'dying'
    else 'dead'
  end status_mame
from player_data;