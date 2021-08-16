-- 首先找到每个员工、员工的薪水和他领导
-- 再找到领导的薪水
-- 筛选出对的信息

select temp.no1,temp.no2,temp.salary,s.salary
from (
    SELECT d1.emp_no no1,s.salary,d2.emp_no no2        -- 得到员工、员工薪水、员工领导
    from dept_emp d1
    left join salaries s on d1.emp_no=s.emp_no
    left join dept_manager d2 on d1.dept_no=d2.dept_no
) temp 
left join salaries s
on temp.no2=s.emp_no
where temp.salary > s.salary and temp.no1!=temp.no2