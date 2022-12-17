from pymongo import MongoClient
from datetime import datetime
import pandas as pd
class TestMongo(object):

     def __init__(self):
         self.clint = MongoClient('mongo',username='root',password='example')
         self.db = self.clint['test']

     def add_one(self):
        """データ挿入"""
        post = {
            'title': 'ハリネズミ',
            'content': 'ハリネズミ可愛い~',
            'created_at': datetime.now()
        }
        return self.db.test.insert_one(post)
     def insert(self,data):
        return self.db.test.insert_many(data)
def main():
    obj = TestMongo()
    rest = obj.add_one()
    df = pd.DataFrame({
    'city': ['osaka', 'osaka', 'osaka', 'osaka', 'tokyo', 'tokyo', 'tokyo'],
    'food': ['apple', 'orange', 'banana', 'banana', 'apple', 'apple', 'banana'],
    'price': [100, 200, 250, 300, 150, 200, 400],
    'quantity': [1, 2, 3, 4, 5, 6, 7]
    })
    final_dict = [dic for index, dic in df.to_dict(orient="index").items() if index!=0]
    # obj.insert(df.to_dict()
    print(df.to_json())
    for i  in final_dict: 
        obj.insert(i)
    obj.insert(df.to_json())
if __name__ == '__main__':
    main()  