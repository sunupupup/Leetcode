-- 查找所有已经分配部门的员工的last_name和first_name以及dept_no

-- 方法1：使用左连接
select e.last_name,e.first_name,d.dept_no
from employees e right join dept_emp d
on e.emp_no = d.emp_no
where d.dept_no not null;   -- 注意：因为左连接会自动未分配部门的员工的dept_no字段补null，所以要用not null剔除。


-- 方法2：使用内连接
select last_name, first_name, dept_no
from dept_emp as d
inner join employees as e
on d.emp_no = e.emp_no;