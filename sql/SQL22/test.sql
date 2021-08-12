-- 先得到所有员工的记录，包含上部门号
/*
    select d.dept_no,depart.dept_name
    from salaries s 
    left join dept_emp d on s.emp_no=d.emp_no
    left join departments depart on depart.dept_no=d.dept_no
*/

select dept_no,dept_name,count(*)
from (
    select d.dept_no dept_no,depart.dept_name dept_name
    from salaries s 
    left join dept_emp d on s.emp_no=d.emp_no
    left join departments depart on depart.dept_no=d.dept_no
) temp
group by dept_no
order by dept_no asc