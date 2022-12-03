CSV=files/csvkit_tutorial/data.csv
SQL="select * from data"
TIME=07140000
FILE=hoge.txt
.PHONY: list
	
list:
	@LC_ALL=C $(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/(^|\n)# Files(\n|$$)/,/(^|\n)# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

up:
	docker compose up -d
build:
	docker compose up -d --build
bash:
	docker compose exec alpine bash
setup_k8s:
	mkdir -p ~/.kube/
	cat ./okteto-kube.config > ~/.kube/config
	kubectl config get-contexts
create_dummy_k8s_pod:
	kubectl run busybox --image=busybox --command --restart=Never -- tail -f /dev/null
# ---------------------------------------------------------------------------
csv_col:	# csvファイルのカラム名を取得
	csvcut -n ${CSV}
csv_sql:	# csvファイルの中身をSQLで取得
	csvsql --query "${SQL}" ${CSV}
check_unix_kernel:
	uname -a
create_specified_time_file: #ファイル作成時刻を指定してファイルを作成
	touch -t ${TIME} ${FILE}
