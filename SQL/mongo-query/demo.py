import pymongo
from pymongo import collection
 
# client = pymongo.MongoClient("mongodb://localhost:27017/")
client = pymongo.MongoClient(host='localhost',port=27017, username='admin', password='123456')
print(client)

dblist = client.list_database_names()

print(dblist)

# if "runoobdb" in dblist:
#   print("数据库已存在！")

db = client['zq']
db.facedata.count()
print(db.get_collection('zq'))


dbml = client['maolei']
if 'face' in dbml.list_collection_names():
    print('face ')