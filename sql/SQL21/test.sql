-- 查找在职员工自入职以来的薪水涨幅情况
-- 思路：
-- 1.先找出所有在职的员工的现在工资
-- 2.找出所有员工入职时的工资
-- 3.上面对应emp_no的工资相减即可得到growth

/*
-- 找出每个员工当前的工资
select emp_no,salary as maxSalary
from salaries 
where salaries.to_date='9999-01-01'

-- 再找出入职时的工资
select emp_no,salary as minSalary
from employees e left join salaries s
on e.hire_date=s.hire_date
*/
select temp1.emp_no,temp1.maxSalary-temp2.minSalary as growth
from (
    select emp_no,salary as maxSalary
    from salaries 
    where salaries.to_date='9999-01-01'
) temp1 left join (
    select e.emp_no,s.salary as minSalary
    from employees e left join salaries s
    on e.hire_date=s.from_date
) temp2 
on temp1.emp_no=temp2.emp_no
order by growth asc;
