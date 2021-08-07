-- 找出所有员工当前具体的薪水salary情况，对于相同的薪水只显示一次，并按照逆序显示

select salary
from salaries
group by salary         --去重
order by salary desc

select distinct salary  --去重
from salaries   
order by salary desc

/*
说明：
对于distinct与group by的使用：
1.当对系统的性能高并且数据量大时使用group by
2.当对系统的性能不高时或者使用数据量少时两者皆可
3.尽量使用group by
*/