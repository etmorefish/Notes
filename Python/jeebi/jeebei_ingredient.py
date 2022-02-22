from threading import local
import requests
from lxml import etree
import pymysql
import time

user = 'root'
password = '123'
host = '0.0.0.0'
port=3306
database='demo'
ty = 'charset=utf8mb4'

db = pymysql.connect(host=host,port=port, user=user, password=password, database=database)
print("连接数据库成功")
cursor = db.cursor()
print(cursor)
for i in range(3469, 3533):
    url = "https://www.jeebei.com/pingce/chengfen/{}.html".format(i)
    headers = {"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.16 Safari/537.36 Edg/99.0.1150.7"}

    resp = requests.get(url=url, headers=headers, timeout=200)
    if resp.status_code != 200:
        number = i
        name = description = safe = recommend = effects = None

        with db.cursor() as cursor:
            sql = "INSERT INTO ingredient_copy1 (name, effects, recommend, safe, description, number) VALUES (%s, %s, %s, %s, %s, %s);"
            cursor.execute(sql, (name, effects, recommend, safe, description, number))
        db.commit()
        
    else:
        html = resp.content.decode('UTF-8')
        # print(html)
        res = etree.HTML(html)
        number = i

        target = res.xpath("//div[@class='component met-cons']")[0]
        # if target:
        name = target.xpath("//h1/text()")[0] # 成分名称 str
        description = target.xpath("//p[@class='content']/text()")[0].strip("参考资料【The beauty bible】jeebei.com']") # 成分描述 str
        safe = target.xpath("//p/span[@class='safe-color']/text()")[0].strip() or 0  # 安全等级 int
        if target.xpath("//p/img[contains(@src,'cestart')]"):
            recommend = len(target.xpath("//p/img[contains(@src,'cestart')]")) or 0 # 推荐指数 int
        else:
            recommend = 0
        effects = target.xpath("//p/text()")[0].split("：")[-1].strip()  # 功效 str
        # if i <= 3704:
        #     with db.cursor() as cursor:
        #         sql = "UPDATE ingredient_copy1 SET name=%s, effects=%s, recommend=%s, safe=%s, description=%s, number=%s WHERE id=%s;"
        #         cursor.execute(sql, (name, effects, int(recommend), safe, description, int(number), i))
        #     db.commit()
        # else:
        with db.cursor() as cursor:
            sql = "INSERT INTO ingredient_copy1 (name, effects, recommend, safe, description, number) VALUES (%s, %s, %s, %s, %s, %s);"
            cursor.execute(sql, (name, effects, int(recommend), safe, description, int(number)))
        db.commit()
            
        # else:
        #     name = description = safe = recommend = effects = None

        #     with db.cursor() as cursor:
        #         sql = "INSERT INTO ingredient_copy1 (name, effects, recommend, safe, description, number) VALUES (%s, %s, %s, %s, %s, %s);"
        #         cursor.execute(sql, (name, effects, recommend, safe, description, number))
        #     db.commit()
        
        print(f"{i}   {name} ")
        time.sleep(0.05)