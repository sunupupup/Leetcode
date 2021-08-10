/*
查询要求：
1. last_name不能是Mary
2. id为奇数
3. 按照hire_date逆序排列
*/

select * from employees
where last_name!='Mary' and emp_no%2=1
order by hire_date desc

/*
结果
10005|1953-11-07|Mary   |Sluis    |F| 1990-01-22    //时间倒叙，
10001|1953-09-02|Georgi |Facello  |M| 1986-06-26
*/