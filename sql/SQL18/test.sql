--  获取当前薪水第二多的员工的emp_no以及其对应的薪水salary

-- 使用order by
select e.emp_no,s.salary,s.last_name,s.first_name
from employees e left join salaries s
on e.emp_no=s.emp_no
order by s.salary
limit 1,1

-- 不使用order by，用聚合函数
-- 主要思想为多层SELECT嵌套与MAX()函数结合
select e.emp_no,max(s.salary),e.last_name,e.first_name
from employees e left join salaries s
on e.emp_no=s.emp_no
where s.salary = (
    select max(salary)      -- 这是找第二大的salary
    from salaries   
    where salary<(
        select max(salary)  -- 这是找第一大的salary
        from salaries
    )
)

