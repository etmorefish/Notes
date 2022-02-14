from threading import local
import requests
from lxml import etree
import pymysql
import time

def get_mysql_connection(user, password, host, port, database):
    """[summary]
    """
    connect = pymysql.connect(host=host,port=port, user=user, password=password, database=database)
    print("连接数据库成功")
    return connect

def get_items(path):
    """获取爬取的产品名称

    Args:
        path ([type]): [description]

    Returns:
        [List]: [description]
    """
    tmp = None
    with open(path, 'r') as f:
        tmp = f.readlines()
    res = [i.strip() for i in tmp]
    return res

def get_item_url(item):
    url2 = 'https://www.jeebei.com/pingce/s.html?act=search&type=1&keywords='
    # title = 'AHC The Pure 第五代全效多功能眼霜'
    headers = {"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.16 Safari/537.36 Edg/99.0.1150.7"}
    headers.update({'referer': 'https://www.jeebei.com/pingce/search.php'})
    resp2 = requests.get(url=url2 + item, headers=headers, timeout=200)
    html2 = resp2.content.decode('UTF-8')
    res2 = etree.HTML(html2)
    item_url = res2.xpath('//div[@class="news_list_title"]/p/a/@href')[0]
    return item_url
    
        
if __name__ == '__main__':
    path = 'item.py'
    items = get_items(path)
    
    user = 'root'
    password = '123'
    host = '0.0.0.0'
    port=3306
    database='demo'
    ty = 'charset=utf8mb4'
    
    db = get_mysql_connection(user, password, host, port, database)
    
    urls = []
    for item in items:
        url = get_item_url(item)
        print(url)
    
