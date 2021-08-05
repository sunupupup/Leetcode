
/*SQL2 查找入职员工时间排名倒数第三的员工所有信息*/

/*
思路1: 按照最晚日期倒叙排，利用limit取第三行数据即可
不足：倒三入职的员工可能不止一个
*/
SELECT * 
from employees
order by hire_date desc 
limit 2,1;

/*
思路2： 找到倒三的日期，用 = 号进行筛选
*/

SELECT *
from employees
where hire_date = (
    select hire_date
    FROM employees
    group by hire_date
    order by hire_date DESC
    limit 2,1              
);

SELECT *
FROM employees
WHERE hire_date = (
    SELECT DISTINCT hire_date     -- 去重
    FROM employees
    ORDER BY hire_date DESC       -- 倒序
    LIMIT 1 OFFSET 2              -- 去掉排名倒数第一第二的时间，取倒数第三
);

-- 关于 DISTINCT 的用法，一般用于统计某个字段不重复的记录，如果是多个字段，不太好
-- 可以用 group by hire_date 替代