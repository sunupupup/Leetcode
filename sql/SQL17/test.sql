
-- 输出工资第二多的员工信心，但这个情况如果多个员工一样多就不对了
select emp_no,salary
from salaries
order by salary DESC
limit 1,1

-- 改进，先找到第二多的工资是多少，再用 = 操作
select emp_no,salary
from salaries
where salary = (
    select distinct salary
    from salaries
    order by salary desc 
    limit 1,1
)