from pymongo import MongoClient, collection
'''
MongoDB 可以创建很多 db
db 内包含很多个集合，有点类似 mysql 这类关系型数据库中的表
插入一条数据，MongoDB 每条记录都有一个唯一标识。返回一个 InsertOneResult 对象，若需要获取唯一标识，找到 InsertOneResult 对象的属性 inserted_id 即可
插入多条数据，使用 insert_many 批量插入

1）常规查询

 find_one ：查询单条记录，返回一个字典。
 find：查询多条记录 ，返回一个游标对象。

'''
class mongodb:
    def __init__(self,host,username, password,db,port = 27017):
        '''
        :param host: str mongodb地址
        :param db: str 数据库
        :param port: int 端口，默认为27017
        '''
        self.host = host
        self.username = username
        self.password = password
        self.port = port
        client = MongoClient(host=host,username=username,password=password,port=port)
        self.db = client[db]

    def insert_one(self,table,dic):
        '''
        :param table: str 数据库中的集合
        :param dic: dict 要插入的字典
        :return: 返回一个包含ObjectId类型的对象
        '''
        collection = self.db[table]
        rep = collection.insert_one(dic)

        return rep
    
    def insert_many(self,table,lists):
        '''
        :param table: str 数据库中的集合
        :param dic: dict 要插入的列表，列表中的元素为字典
        :return: 返回包含多个ObjectId类型的列表对象
        '''
        collection = self.db[table]
        rep = collection.insert_many(lists)

        return rep
    
    def find_one(self,table,dic):
        '''
        :param table: str 数据库中的集合
        :param dic: dict 查询条件
        :return: dict 返回单条记录的字典
        '''
        collection = self.db[table]
        rep = collection.find_one(dic)

        return rep

    def find(self,table,dic):
        '''
        :param table: str 数据库中的集合
        :param dic: dict 查询条件
        :return: list 返回查询到记录的列表
        '''
        collection = self.db[table]
        rep = list(collection.find(dic))

        return rep

if __name__=='__main__':
    db = mongodb(host='localhost',db = 'test', username='admin', password='123456')

    
    # dic = {'姓名':'小明','English':100,'math':90}
    # rep = db.insert_one('test',dic)
    # print(rep.inserted_id)
    
    # lists = [{'姓名':'小hei','English':100,'math':90},
    #          {'姓名':'小华','English':90,'math':100}]
    # rep = db.insert_many('test',lists)
    # for i in rep.inserted_ids:
    #     print(i)

    # 查询 English 成绩为 100 的所有记录
    # dic = {'English':100}
    # rep = db.find('test',dic)
    # print(rep, rep[0]['_id'])
    
    # 模糊搜索key为"姓名"，value包含"明"的记录
    # dic = {'姓名':{'$regex':'明'}}
    # rep = list(db.find('test', dic))
    # print(rep)
    
    # 模糊搜索key为"姓名"，value包含"明"的记录
    
    # projection字典参数。1显示，0不显示
    # 以下查询结果只返回key为English的相关字段
    # rep = list(db.find(dic,projection={'English':1,'_id':0})) XXX
    
    # 计数查询只需要在普通查询后加上 count() 即可
    # count = db.find().count()  
    # count = db.find('test', {'English':{'$gt':90}}).count()
    # print(count)