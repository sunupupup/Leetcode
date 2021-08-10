-- 请你统计出各个title类型对应的员工薪水对应的平均工资avg。
-- 结果给出title以及平均工资avg，并且以avg升序排序，以上例子输出如下
select t.title,AVG(s.salary)
from titles t left join salaries s
on t.emp_no=s.emp_no
group by t.title
order by AVG(s.salary) 