import pymongo
from datetime import datetime 
# client = pymongo.MongoClient("mongodb://localhost:27017/")
# client = pymongo.MongoClient(host='localhost',port=27017, username='admin', password='123456')

import random
mongodbUri = 'mongodb://admin:123456@localhost:27017/admin'
mongodbUri = 'mongodb://admin:123456@localhost:27017'

client = pymongo.MongoClient(mongodbUri)
db = client.maolei
collection = db.face

start = "2019-01-04 05:00:00"

query = {"create_time": {"$gte": start}}
cursor = collection.find(query)
list_cur = list(cursor)
print(len(list_cur))