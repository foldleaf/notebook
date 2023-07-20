# 创建表、修改表、添加列
```sql
-- 创建表 Student
CREATE TABLE Student(
    Sno char(5) not null unique,
    Sname char(20),
    Ssex char(2),
    Sage int,
    Sdept char(15)
);

-- 修改表 Student
    -- 添加新列
    ALTER TABLE Student add Scome DATE
    -- 修改数据类型
    ALTER TABLE Student modify Sage smallint
    -- 删除学号的唯一标识
    ALTER TABLE Student drop unique(Sno)

-- 删除表
drop table Student;
-- 创建索引
CREATE CLUSTER INDEX Stusname ON Student(Sname);    
-- 删除索引
drop index Stusname;
```
# 基本查询
```sql

select Sname,Sage from Student;
 
select Sname,Sage,Sdept from Student;
 
select * from Student;
 
select Sname,1996-Sage from Student;
 
select Sname,'Year of Birth:',1996-SagLOWER(Sdept) from Student;
 
select Sname NAME,'Year of Birth:'BIRT1996-Sage BIRTHDAY,LOWER(SdeptDEPARTMENT FROM Student;    /*为列创建别名*/
```
# 复杂查询
```sql
--消除重复行
 select DISTINCT Sno from Student;
--查询满足条件的元组
 select Sname from Student where Sdept="CD";
 
 select Sname,Sage from Student where Sage>=20;
 
 alter table Student add Grade int;
 
 select DISTINCT Sno from Student where Grade<60;
--确定范围-----------between and 
 select Sname,Sage,Sdept from Student where Sage between 20 and 23;
--确定集合-----------IN('','','')
 select Sname,Ssex from Student where Sdept IN('IS','MA','CS');
--字符匹配-----------LIKE+++( %：代表任意长度)( _ ：代表单个字符)
 select Sname,Sage from Student where Sname LIKE 'A%';
 
 SELECT Sname ,Sage from Student where Sname LIKE 'A_';
 
--字符匹配-------里面的转义字符-----ESCAPE'\'：表示\为转义字符
--查询条件是“A_”此时这里的_不是代表一个字符，只是单纯的表示下划线而已。因为语句前面有转义字符。
 select Sname from Student where Sage LIKE 'A\__' ESCAPE '\';
 
--涉及空值的查询
 select Sname,Sage from Student where Grade IS NULL;  /*查询成绩为空的学生*/
```

# 更复杂查询
```sql
--多重条件查询
   select Sname from Student where Sdept='cd' and Sage>20;
 
   select Sname from Student where Sdept='cd' or Sage>20;
 
--对查询结果进行排序
 
   select Sname Grade from Student where Sage>20 order by Sage DESC;
 
   select * from Student order by Sage DESC;
 
--使用集函数
   select count(*) from Student;    /*求总个数*/
 
   select count(distinct Sno) from Student;
 
   select avg(Sage) from Student where Sname='ahui';     /*avg：求平均值*/
 
   select Sname,count(Sage) from Student group by Sname;
 
   select MAX(Sage) from Student where Sno='1';    /*最大值*/
```

# 连接查询
```sql
/*连接查询*/ 
   /*等值的查询*/
      select Student.*,SC.* from Student,SC where Student.Sno=SC.Sno;
 
     /*自然连接两个表*/
     select Student.Sno,Sname,Ssex,Sage,Sdept,Cno,SC.Grade from Student,SC where Student.Sno=SC.Sno;
     /*将表Student进行了重新命名，为两个名字，从而进行对自己的查询*/
     select FIRS.Sage,SECO.Sno from Student FIRS,Student SECO where FIRS.Sno=SECO.Sage; 
 
    /*外连接*/
 
    select Student.Sno,Sname,Sage,Sdept,Ssex,Cno,SC.Grade from Student,SC where Student.Sno=SC.Sno(*);
 
    /*复合条件连接--------就是利用and来进行操作*/
 
     select Student.Sno,Sname from Student,SC where Student.Sno=SC.Sno and SC.Cno='2' and Student.Sage=2;
```

# 嵌套查询
```sql
/*嵌套查询---就是把一个查询的结果当做另一个查询的条件来进行查询*/
 
   select Sname from Student where Sno 
         IN(
            select Sno from SC where Cno='2'
         );
         /*--01：带有IN的子查询*/
   select Sno,Sname,Sdept from Student where Sdept IN(  select Sdept from Student where Sname='ahui');
        /*--02：带有比较运算符的子查询*/
   select Sno,Sname,Sdept from Student where Sdept=(  select Sdept from Student where Sname='ahui');
```
