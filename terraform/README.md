# Get Started


## COCKROACHDB
```dotnetcli
docker compose up -d
docker compose exec tf sh
cd cockroachdb
terraform init
terraform apply \
            -auto-approve \
            -var="cluster_name=${COCKROACHDB_CLUSTER_NAME}" \
            -var="sql_user_name=${COCKROACHDB_USER}" \
            -var="sql_user_password=${COCKROACHDB_PASS}"
```


## localstack

```dotnetcli
docker compose up -d
docker compose exec tf sh
cd localstack
terraform init
terraform apply 
```