
-- 请你查找所有员工的last_name和first_name以及对应的dept_name，也包括暂时没有分配部门的员
select e.last_name,e.first_name,temp.dept_name  -- 再联查员工的number
from employees e left join (
    select d1.emp_no,d2.dept_name   -- 这select先获取员工number对应的部门名字
    from dept_emp d1 left join departments d2
    on d1.dept_no=d2.dept_no
) temp
on e.emp_no=temp.emp_no

-- 方法2：直接两次left join
select e.last_name,e.first_name,d2.dept_name
from employees e 
left join dept_emp d1 on d1.emp_no=e.emp_no
left join departments d2 on d1.dept_no=d2.dept_no
