-- 先查出所有源给对应的部门号、员工号、工资
-- 排除是领导的emp_no即可
select 
    d.dept_no,d.emp_no,s.salary
from 
    dept_emp d join salaries s on d.emp_no=s.emp_no
where 
    d.emp_no not in (
        select emp_no
        from dept_manager
    )
order by 
    emp_no
