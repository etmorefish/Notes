# sqlalchemy 通过映射类对表能很容易进行增删改查
Person.query.filter(Person.id == 1).first()
# 查询的结果是一个ResultProxy对象，可以通过.获取到对应的值，但是怎么转成python字典呢
def row2dict(row):
    d = {}
    for column in row.__table__.columns:
        d[column.name] = str(getattr(row, column.name))
    return d


# 或者匿名函数
row2dict = lambda r: {c.name: str(getattr(r, c.name)) for c in r.__table__.columns}

