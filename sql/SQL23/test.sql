-- 对员工的工资进行排序
-- 很简单的需求，但需要好好考虑

-- 得到两张一样的表，员工num和工资的表
-- 从第一张表取出一个salary，统计另一张表大于他的工资数有多少（并用distinct去重）
select temp1.emp_no, temp1.salary , count(DISTINCT temp2.salary)
from 
    (
        select emp_no,salary
        from salaries 
    ) temp1 , (
        select emp_no,salary
        from salaries 
    ) temp2
where 
    temp1.salary<=temp2.salary
GROUP BY
    temp1.emp_no
order by 
    salary desc,emp_no asc