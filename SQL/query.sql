-- 基本查询
SELECT * FROM students;

SELECT * FROM classes;

SELECT 10-29;

-- 条件查询
SELECT * FROM students WHERE score >= 80;

SELECT * FROM students WHERE score >= 80 AND gender = 'M';

SELECT * FROM students WHERE score >= 80 OR gender = 'M';

SELECT * FROM students WHERE NOT class_id = 2;

SELECT * FROM students WHERE (score < 80 OR score > 90) AND gender = 'M';
-- 如果不加括号，条件运算按照NOT、AND、OR的优先级进行，即NOT优先级最高，其次是AND，最后是OR。加上括号可以改变优先级。

-- 投影查询
SELECT id, score, name FROM students;

SELECT id, score points, name FROM students;

SELECT id, score points, name FROM students WHERE gender = 'M';

-- 排序
SELECT id, name, gender, score FROM students ORDER BY score;

SELECT id, name, gender, score FROM students ORDER BY score DESC;

SELECT id, name, gender, score FROM students ORDER BY score DESC, gender;


-- 分页
-- 使用LIMIT <M> OFFSET <N>可以对结果集进行分页，每次查询返回结果集的一部分；
SELECT id, name, gender, score FROM students ORDER BY score DESC LIMIT 3 OFFSET 9;


-- 聚合查询
SELECT COUNT(*) FROM students;
SELECT COUNT(*) num FROM students;

-- SUM	计算某一列的合计值，该列必须为数值类型
-- AVG	计算某一列的平均值，该列必须为数值类型
-- MAX	计算某一列的最大值
-- MIN	计算某一列的最小值

-- 统计出有多少男生、多少女生、多少80分以上的学生等：
SELECT COUNT(*) boys FROM students WHERE gender = 'M' AND score >= 80;
-- 使用聚合查询计算男生平均成绩:
SELECT AVG(score) average FROM students WHERE gender = 'M';
-- 如果聚合查询的WHERE条件没有匹配到任何行，COUNT()会返回0，而SUM()、AVG()、MAX()和MIN()会返回NULL：
-- WHERE条件gender = 'X'匹配不到任何行:
SELECT AVG(score) average FROM students WHERE gender = 'X';
-- 每页3条记录，如何通过聚合查询获得总页数？
SELECT CEILING(COUNT(*)/3) FROM students;
-- 统计每个班的学生数量 
-- 按class_id分组:
SELECT class_id, COUNT(*) num FROM students GROUP BY class_id;
-- 按class_id, gender分组:
SELECT class_id,gender, COUNT(*) num FROM students GROUP BY class_id, gender;

-- 练习
-- 请使用一条SELECT查询查出每个班级的平均分：
SELECT class_id, AVG(score) average FROM students GROUP  BY class_id;
-- 请使用一条SELECT查询查出每个班级男生和女生的平均分：
SELECT class_id, gender, AVG(score) average FROM students GROUP BY class_id, gender;


-- 多表查询  
-- 查询多张表的语法是：SELECT * FROM <表1> <表2>。
SELECT * FROM students, classes;

-- set alias:
SELECT
    students.id sid,
    students.name,
    students.gender,
    students.score,
    classes.id cid,
    classes.name cname
FROM students, classes;

-- set table alias:
SELECT
    s.id sid,
    s.name,
    s.gender,
    s.score,
    c.id cid,
    c.name cname
FROM students s, classes c;
-- WHERE s.class_id =c.id;


SELECT
    s.id sid,
    s.name,
    s.gender,
    s.score,
    c.id cid,
    c.name cname
FROM students s, classes c
WHERE s.gender = 'M' AND c.id = '1';

-- 这种方式的多表查询是个笛卡尔积查询，但真正有意义的多表查询需要将多表的主键外键相关联才行，此处为students的外键class_id和classes表的主键id需要关联相等，即查询时要添加WHERE students.class_id =classes.id。

-- 连接查询
SELECT s.id, s.name, s.class_id, s.gender, s.score FROM students s;

-- 内连接——INNER JOIN
-- 选出所有学生，同时返回班级名称
SELECT s.id, s.name, s.class_id, s.gender, s.score, c.name class_name
FROM students s
INNER JOIN classes c
ON s.class_id = c.id;
-- ORDER BY s.score;

-- 注意INNER JOIN查询的写法是：
-- 
-- 先确定主表，仍然使用FROM <表1>的语法；
-- 再确定需要连接的表，使用INNER JOIN <表2>的语法；
-- 然后确定连接条件，使用ON <条件...>，这里的条件是s.class_id = c.id，表示students表的class_id列与classes表的id列相同的行需要连接；
-- 可选：加上WHERE子句、ORDER BY等子句。
-- 

-- 外连接（OUTER JOIN)
SELECT s.id, s.name, s.class_id, s.gender, s.score, c.name class_name
FROM students s
RIGHT OUTER JOIN  classes c
ON s.class_id = c.id;

SELECT s.id, s.name, s.class_id, s.gender, s.score, c.name class_name
FROM students s
LEFT OUTER JOIN  classes c
ON s.class_id = c.id;

-- 
-- 有RIGHT OUTER JOIN，就有LEFT OUTER JOIN，以及FULL OUTER JOIN。它们的区别是：
-- 
-- INNER JOIN只返回同时存在于两张表的行数据，
-- 由于students表的class_id包含1，2，3，
-- classes表的id包含1，2，3，4，
-- 所以，INNER JOIN根据条件s.class_id = c.id返回的结果集仅包含1，2，3。
-- 
-- RIGHT OUTER JOIN返回右表都存在的行。
-- 如果某一行仅在右表存在，那么结果集就会以NULL填充剩下的字段。
-- 
-- LEFT OUTER JOIN则返回左表都存在的行。
-- 如果我们给students表增加一行，并添加class_id=5，
-- 由于classes表并不存在id=5的行，
-- 所以，LEFT OUTER JOIN的结果会增加一行，对应的class_name是NULL：

-- 先增加一列class_id=5:
INSERT INTO students (class_id, name, gender, score) values (5, '新生', 'M', 88);
-- 使用LEFT OUTER JOIN
SELECT s.id, s.name, s.class_id, s.gender, s.score, c.name class_name
FROM students s
LEFT OUTER JOIN  classes c
ON s.class_id = c.id;


-- MySQL 不支持全连接，但可以通过左外连接 + UNION + 右外连接实现：
SELECT s.id, s.name, s.class_id, s.gender, s.score, c.name class_name
FROM students s
FULL OUTER JOIN  classes c
ON s.class_id = c.id;

SELECT s.id, s.name, s.class_id, s.gender, s.score, c.name class_name FROM students s LEFT OUTER JOIN  classes c ON s.class_id = c.id
UNION
SELECT s.id, s.name, s.class_id, s.gender, s.score, c.name class_name FROM students s RIGHT OUTER JOIN  classes c ON s.class_id = c.id;



-- SELECT col1_name,col2_name,(COUNT(*),(AVG(col3_name)) average) #基本查询方式 聚合查询——聚合函数 投影查询
-- FROM excel1_name replace1_name (,excel2_name replace2_name) #多表查询 同表名字替换
-- INNER(/FULL OUTER/RIGHT OUTER/LEFT OUTER)  JOIN classes c #连接查询
-- (GROUP BY colx_name) #聚合查询——分组查询
-- WHERE (EXPRESSION) #条件查询
-- ORDER BY coly_name (ASC/DESC) #排序
-- LIMIT m OFFSET n; #分页查询


-- 修改数据
-- CRUD：Create、Retrieve、Update、Delete

-- INSERT语句的基本语法是：
-- INSERT INTO <表名> (字段1, 字段2, ...) VALUES (值1, 值2, ...);

INSERT INTO students (class_id, name, gender, score) VALUES (2, '大牛', 'M', 80);
-- 查询并观察结果:
SELECT * FROM students;

INSERT INTO students (class_id, name, gender, score) VALUES
  (1, '大宝', 'M', 87),
  (2, '二宝', 'M', 81);
	
-- UPDATE语句的基本语法是：
-- UPDATE <表名> SET 字段1=值1, 字段2=值2, ... WHERE ...;	
	
UPDATE students SET name='大牛', score=66 WHERE id=1;
-- 查询并观察结果:
SELECT * FROM students WHERE id=1;	
	
-- 更新id=5,6,7的记录	
UPDATE students SET name='小牛', score=77 WHERE id>=5 AND id<=7;
-- 查询并观察结果:
SELECT * FROM students;
	
-- 更新score<80的记录
UPDATE students SET score = score+10 WHERE score < 80;

-- 更新id=999的记录
-- 如果WHERE条件没有匹配到任何记录，UPDATE语句不会报错，也不会有任何记录被更新
UPDATE students SET score=100 WHERE id=999;

UPDATE students SET name='大宝' WHERE id=1;

-- DELETE语句的基本语法是：
-- DELETE FROM <表名> WHERE ...;

DELETE FROM students WHERE id=1;
-- 查询并观察结果:
SELECT * FROM students;

DESC students;

