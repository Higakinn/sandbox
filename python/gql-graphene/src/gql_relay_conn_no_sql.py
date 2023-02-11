from fastapi import FastAPI

import graphene

from starlette_graphene3 import GraphQLApp, make_graphiql_handler

# graphene object 
class User(graphene.ObjectType):
    class Meta:
        interfaces = (graphene.relay.Node,)

    name = graphene.String()

# graphene relay connection object 
class UserConnection(graphene.relay.Connection):
    class Meta:
        node = User

    total_count = graphene.Int()

    def resolve_total_count(root, info):
        print(root)
        return len(root.iterable)

# graphql query object
class Query(graphene.ObjectType):
    me = graphene.relay.ConnectionField(
        UserConnection, description="The ships used by the faction."
    )

    def resolve_me(self, info, **args):
        users = [
            User(id=1, name="hoge"),
            User(id=2, name="fuga"),
            User(id=3, name="piyo")
        ]
        
        return list(filter(lambda x: x.id >= 2, users))

# fastapi にマウント
app = FastAPI()
schema = graphene.Schema(query=Query)
# starlette_graphene3を利用
# 参考: https://github.com/ciscorn/starlette-graphene3#example
app.mount("/", GraphQLApp(schema, on_get=make_graphiql_handler()))  # Graphiql IDE