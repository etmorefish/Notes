import pymongo
from datetime import datetime
client = pymongo.MongoClient("mongodb://admin:123456@localhost:27017/")
# client = pymongo.MongoClient(host='localhost',port=27017, username='admin', password='123456')
# print(client)
# dblist = client.list_database_names()
# print(dblist)
db_test = client['test']
test = db_test['test']
col = test.find_one()

data1 = {"name": "data1", "createDatetime": datetime.now(), "uid": 1.0}
data2 = {"name": "data2", "createDatetime": datetime.now(), "uid": 2.0}

res = test.insert_one(data1)
print(res.inserted_id)
# with client.start_session() as session:
#     session.start_transaction()
#     test.insert_one(data1, session=session)
#     session.commit_transaction()

# client.close()
