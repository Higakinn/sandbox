


## memo

```
 uvicorn gql_relay_conn_no_sql:app --reload --host 0.0.0.0
```

server access

http://127.0.0.1:8000/


### graphql query

```graphql
query {
  me(first:1) {
    totalCount
    edges {
      node {
        id
        name
      }
    }
  }
}
```