import pymongo
 
# client = pymongo.MongoClient("mongodb://localhost:27017/")
client = pymongo.MongoClient(host='localhost',port=27017, username='admin', password='123456')
print(client)

dblist = client.list_database_names()

print(dblist)

# if "runoobdb" in dblist:
#   print("数据库已存在！")