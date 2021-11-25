
```mysql
-- 查看一个表的结构，使用命令：
DESC students;
SHOW CREATE TABLE students;
```

创建表使用`CREATE TABLE`语句，而删除表使用`DROP TABLE`语句：

```mysql
mysql> DROP TABLE students;
Query OK, 0 rows affected (0.01 sec)
```

修改表就比较复杂。如果要给`students`表新增一列`birth`，使用：

```mysql
ALTER TABLE students ADD COLUMN birth VARCHAR(10) NOT NULL;
```

要修改`birth`列，例如把列名改为`birthday`，类型改为`VARCHAR(20)`：

```mysql
ALTER TABLE students CHANGE COLUMN birth birthday VARCHAR(20) NOT NULL;
```

要删除列，使用：

```mysql
ALTER TABLE students DROP COLUMN birthday;
```
