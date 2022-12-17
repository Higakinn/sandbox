import graphene
from graphene import relay
from graphene_sqlalchemy import SQLAlchemyConnectionField
from ..models.genres import Genres as GenresModel
from ..models.books import Books as BooksModel
from ..types.books import Books
from ..types.genres import Genres
from ..types.genres import mock_genres
from graphql_relay.node.node import from_global_id
from ..utils import input_to_dictionary
from ..database.db_session import db_session
from sqlalchemy.sql.functions import concat
import sqlalchemy
import inspect
import base64
class m_query(graphene.ObjectType):


    mock_books_by_name = graphene.List(Genres, name=graphene.String())
    @staticmethod
    def resolve_mock_books_by_name(parent, info, **args):
        print(info)
        book_id=base64.b64encode('0'.encode('utf-8'))
        print(book_id)
        print(GenresModel.__dict__)
        result = [
            Genres(
                id = '0',
                name='higahiga'
            )
        ]
        # result = input_to_dictionary(result)
        return result

class Query(graphene.ObjectType):
    node = relay.Node.Field()
    test = graphene.String()
    def resolve_test(parent,info):
        model = Books._meta.model
        from sqlalchemy.sql.expression import cast
        print(db_session.query(cast(model.id,sqlalchemy.String) + model.name).all())
        return "test"
    books_by_name = graphene.List(Books, name=graphene.String(),sort=graphene.Argument(Books.sort_enum()))
    books_by_genre = graphene.List(Books, name=graphene.String())
    all_books = SQLAlchemyConnectionField(Books.connection, name=graphene.String())
    @staticmethod
    def resolve_books_by_name(parent, info, **args):
        q = args.get('name')
        # print(GenresModel.__dict__)
        # print(GenresModel.__table__.columns)
        l = [[1,1],[2,2]]
        li = []
        # print(Books.__builtins__)
        print(vars(GenresModel))
        # print(__builtins__)
        # for j in l:
        #     for i, column in enumerate(GenresModel.__table__.columns):
        #         g = GenresModel()
        #         setattr(g, column.name, j[i])
        #         print(g.id)
        #     li.append(g)
        # print(li)
                # setattr(g,co)
            # arg = [setattr(g, Column.name, j[i]) for i, column in enumerate(GenresModel.__table__.columns)]
            # print(arg[0])
            # setattr(column,i)
            # print(arg)
        books_query = Books.get_query(info)
        result = books_query.filter(BooksModel.name.contains(q)).all()
        return result

    @staticmethod
    def resolve_books_by_genre(parent, info, **args):
        q = args.get('name')

        books_query = Books.get_query(info)

        return books_query.join(GenresModel).filter(GenresModel.name == q).all()

    # @staticmethod
    # def resolve_books(parent,info):
    #     print(parent)
    #     books_query = Books.get_query(info)
    #     return books_query.all()