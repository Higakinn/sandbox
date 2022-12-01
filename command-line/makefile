CSV=files/csvkit_tutorial/data.csv
SQL="select * from data"

csvcols:	# csvファイルのカラム名を取得
	csvcut -n ${CSV}
csvsqls:	# csvファイルの中身をSQLで取得
	csvsql --query "${SQL}" ${CSV}
