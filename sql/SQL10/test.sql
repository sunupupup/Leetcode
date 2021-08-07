--找出所有非部门领导的员工emp_no

--题中只有主管才有dept_no

--方法1：not in + 子查询
select emp_no
from employees
where emp_no not in (   
    select distinct emp_no      -- 这里筛选出领导的emp_no
    from dept_manager
)

--方法2：left join左连接 + null
select e.emp_no
from employees e left join dept_manager d
on e.emp_no=d.emp_no    
where dept_no is null       -- 不是部门领导的人，左连接之后dept_no肯定是 null