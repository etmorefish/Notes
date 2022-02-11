from threading import local
import requests
from lxml import etree
import pymysql

url = "https://www.jeebei.com/pingce/chengfen/1.html"
headers = {"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.16 Safari/537.36 Edg/99.0.1150.7"}

resp = requests.get(url=url, headers=headers, timeout=200)
html = resp.content.decode('UTF-8')
res = etree.HTML(html)

target = res.xpath("//div[@class='component met-cons']")[0]
number = ''
name = target.xpath("//h1/text()")[0]
description = target.xpath("//p[@class='content']/text()")[0].strip("参考资料【The beauty bible】jeebei.com']")
safe = target.xpath("//p/span[@class='safe-color']/text()")[0].strip()
recommend = len(target.xpath("//p/img[contains(@src,'cestart')]"))
effects = target.xpath("//p/text()")[0].split("：")[-1].strip()

print(name)