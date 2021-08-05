/*
查询最晚入职的员工的所有信息
思路1：按照入职时间进行排序，limit 1 取第一个人信息即可

缺陷：如果最晚日期有多个人，就没法筛选出来了
*/


select * 
from employees
order by hire_date desc
limit 1;


SELECT * 
FROM employees 
order by hire_date desc 
limit 0,1

/*
limit语法：
limit n 表示取前n条数据
limit m n 表示取m+1行开始，共n条数据    m表示skip步长，即跳过m条
*/

/*
思路2：找到最大的hire_date
*/
SELECT * FROM employees
WHERE hire_date = (
    SELECT MAX(hire_date)
    FROM employees
);