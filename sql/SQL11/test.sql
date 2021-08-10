-- 获取所有员工当前的manager

-- 方法1：按照dept_no左连接查询，
--        如果员工本身是经理的话，having对结果再次筛选
select d1.emp_no,d2.emp_no
from dept_emp d1 left join dept_manager d2
on d1.dept_no=d2.dept_no
having d1.emp_no!=d2.emp_no

-- 方法2 提前约束d1.emp_no不在dept_manager表中出现
select d1.emp_no,d2.emp_no
from dept_emp d1 left join dept_manager d2
on d1.dept_no=d2.dept_no
where d1.emp_no not in (
    select DISTINCT emp_no
    from dept_manager
)

-- 方法3：inner内连接  内连接相当于一种并集
select d1.emp_no,d2.emp_no
from dept_emp d1 inner join dept_manager d2
on d1.dept_no=d2.dept_no and d1.emp_no!=d2.emp_no 