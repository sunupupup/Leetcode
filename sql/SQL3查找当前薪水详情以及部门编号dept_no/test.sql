
select s.* , d.dept_no
from salaries as s right join dept_manager as d
on s.emp_no=d.emp_no
order by emp_no

-- join   join等价于inner join内连接抄，是返回两个表中都有的符合条件的行。相当于取交集
-- right join右连接，是返回右表中所有的行及左表中符合条件的行。

select s.* ,d.dept_no
from salaries as s join dept_manager as d 
on s.emp_no=d.emp_no
order by emp_no
