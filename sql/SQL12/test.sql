-- 这种方法是不对的，虽然保证了salary是最大的，但是emp_no不一定对应着最大薪水
select d.dept_no,d.emp_no,max(s.salary)
FROM dept_emp d left join salaries s
on d.emp_no=s.emp_no
group by d.dept_no
order by d.dept_no 

-- 改进
select temp.dept_no,temp2.emp_no,temp.maxSalary
from (
    -- 这里查出每个部门的最高工资
    select d.dept_no dept_no,max(s.salary) maxSalary
    FROM dept_emp d left join salaries s
    on d.emp_no=s.emp_no
    group by d.dept_no
    order by d.dept_no  
) as temp left join (
    -- 这是查出所有的信息
    select d.dept_no dept_no,d.emp_no,s.salary salary
    FROM dept_emp d join salaries s
    on d.emp_no=s.emp_no
)as temp2
on temp.dept_no=temp2.dept_no and temp.maxSalary=temp2.salary   -- 根据两个表的情况继续筛选


