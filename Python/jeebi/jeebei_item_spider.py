from threading import local
import requests
from lxml import etree
import pymysql
import time
import re
import json

def get_mysql_connection(user, password, host, port, database):
    """[summary]
    """
    connect = pymysql.connect(host=host,port=port, user=user, password=password, database=database)
    print("连接数据库成功")
    return connect

def get_items(path):
    """读取文件中的 产品名称 or url

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

        
def get_item_html(url):
    headers = {"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.16 Safari/537.36 Edg/99.0.1150.7"}
    resp = requests.get(url=url, headers=headers, timeout=200)
    html = resp.content.decode('UTF-8')
    return html
    
def item_spider(html):
    htmls = etree.HTML(html)
    target1 = htmls.xpath("//table[@class='table']")[0]
    trs = target1.xpath("./tr")
    data = {}
    ingredient_list = []
    for i in range(1, len(trs)-1):
        ingredient_id = trs[i].xpath("./td[1]/@data")[0]
        ingredient_name = trs[i].xpath("./td[1]/span/text()")[0]
        if trs[i].xpath("./td[3]/span/text()"):
            safe = trs[i].xpath("./td[3]/span/text()")[0]
        else:
            safe = 0
        hxcf = len(trs[i].xpath("./td[4]/span/img")) or 0
        acne = trs[i].xpath("./td[5]/span/text()") or 0
        uvan = trs[i].xpath("./td[6]/img/@src") or None
        # print(uvan)
        uv = None
        if uvan is not None:
            uva = re.findall("icon/(.*).gif", uvan[0])
            uvb = re.findall("icon/(.*).gif", uvan[1])
            uv = "|".join(uva + uvb)
        uv = uv or uvan
        tmp = {
            "ingredient_id": ingredient_id,
            "ingredient_name": ingredient_name,
            "safe": safe,
            "hxcf": hxcf,
            "acne": acne,
            "uv": uv,
        }
        ingredient_list.append(tmp)
        
    atttions = htmls.xpath("//div[@class='safe_detail']/p/text()")
    xj = atttions[0].split("：")[1]
    ffj = atttions[1].split("：")[1]
    fxcf = atttions[2].split("：")[1]
    yfsy = atttions[3].split("：")[1]
    atttion = {
        "xj": xj,
        "ffj": ffj,
        "fxcf": fxcf,
        "yfsy": yfsy,  
    }
    data['ingredient_list'] = ingredient_list
    data['atttion'] = atttion
    
    return data

def save_to_file(data: dict):
    res  = json.dumps(data, ensure_ascii=False)
    with open("yunqi_out.txt", 'a+', encoding='utf-8') as f:
        f.writelines(res)
        f.write('\n')

def main():
    item_path = 'yunqi_item'
    items = get_items(item_path)
    
    user = 'root'
    password = '123'
    host = '0.0.0.0'
    port=3306
    database='demo'
    ty = 'charset=utf8mb4'
    DOMAIN = "https://www.jeebei.com" 
    
    db = get_mysql_connection(user, password, host, port, database)
    
    # 1 url write in file
    # urls = []
    # for item in items:
    #     url = DOMAIN + get_item_url(item)
    #     print(url)
    #     urls.append(url)
        
    # # urls 写入文件    
    # with open('yunqi_item_urls', 'w') as f:
    #     f.writelines(urls)
    
    # # 2 read url and spider
    urls = get_items("yunqi_item_urls")
    a = 1
    for url in urls:
        print(f" {a} {url}")
        html = get_item_html(url)
        data = item_spider(html)
        save_to_file(data)
        time.sleep(0.3)
        a += 1
        
    
        
if __name__ == '__main__':
    main()