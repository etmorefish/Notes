import pymongo
 
# client = pymongo.MongoClient("mongodb://localhost:27017/")
# client = pymongo.MongoClient(host='localhost',port=27017, username='admin', password='123456')

import random
mongodbUri = 'mongodb://admin:123456@localhost:27017/admin'
mongodbUri = 'mongodb://admin:123456@localhost:27017'

client = pymongo.MongoClient(mongodbUri)
db = client.somedb
# db.user.drop()
# element_num=10
# for id in range(element_num):
#    name = random.choice(['R9','cat','owen','lee','J'])
#    sex = random.choice(['male','female'])
#    db.user.insert_one({'id':id, 'name':name, 'sex':sex})
# content = db.user.find()
# for i in content:
#    print(i)
   
inserted_id = db.somecoll.insert_one({"somekey":"yiqihapi"}).inserted_id
print(inserted_id)
for doc in db.somecoll.find(dict(_id=inserted_id)):
       print (doc)
for doc in db.somecoll.find({"somekey":"yiqihapi"}):
       print (doc)