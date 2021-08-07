-- 查找薪水记录超过15次的员工号emp_no以及其对应的记录次数t

select emp_no,count(emp_no) t
from salaries
group by emp_no
having t > 15

/*
注意这几点：
1、用COUNT()函数和GROUP BY语句可以统计同一emp_no值的记录条数
2、根据题意，输出的变动次数为t，故用AS语句将COUNT(emp_no)的值转换为t
3、由于COUNT()函数不可用于WHERE语句中，故使用HAVING语句来限定t>15的条件

group by 和 having用于筛选where之后的数据
*/