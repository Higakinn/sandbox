from fastapi import FastAPI

import graphene

from starlette_graphene3 import GraphQLApp, make_graphiql_handler
import sqlalchemy
from sqlalchemy import create_engine, MetaData, Table

dsn="mysql+pymysql://root:pass@db/user?charset=utf8"
engine = create_engine(dsn)
metadata_obj = MetaData()
users_table = Table("users", metadata_obj, autoload_with=engine)


from sqlalchemy import select

print(users_table)

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
        UserConnection, description="The ships used by the faction.",
        name=graphene.String(default_value=None)
    )

    def resolve_me(self, info, **args):
        name = args.get('name',None)
        stmt = select(
            users_table.c.id.label('id'),
            users_table.c.username.label('name')
        ).where(users_table.c.username.like(f'%{name}%'))

        print(stmt)
        result = None
        with engine.connect() as conn:

            result = list(conn.execute(stmt))
    
        return result

# fastapi にマウント
app = FastAPI()
schema = graphene.Schema(query=Query)
# starlette_graphene3を利用
# 参考: https://github.com/ciscorn/starlette-graphene3#example
app.mount("/", GraphQLApp(schema, on_get=make_graphiql_handler()))  # Graphiql IDE