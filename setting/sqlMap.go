package setting

var SQLView map[string]string  = map[string]string{
	"task_data_export_view": `(
						   select a.task_user_id,
                        a.task_id,
                        b.task_name,
                        case when b.task_type= 1 then '每日任务' when b.task_type=0 then '特殊任务' when b.task_type=2 then '专区任务' else '未知' end task_type,
                        c.user_id,
                        c.employee_id,c.realname,c.department,
                        case when b.expired_time=0 then '' else date_format(from_unixtime(b.expired_time),'%Y%m%d') end as task_expired_time,
                        date_format(from_unixtime(b.create_time), '%Y-%m-%d %H:%i:%s') as task_create_time,
                        date_format( from_unixtime(b.update_time), '%Y-%m-%d %H:%i:%s') as task_update_time ,
                        b.gold_num,
                        b.diamond_num,
                        b.description,
                        case when b.condition_type=0 then '按照条件' when b.condition_type=1 then '指定人' else '未知' end as condition_type,
                        a.task_status, a.sub_task_total_num, a.sub_task_finish_num, a.task_day, date_format(from_unixtime(a.task_create_time), '%Y-%m-%d %H:%i%s') user_task_create_time, date_format(from_unixtime(a.task_update_time), '%Y-%m-%d %H:%i%s') user_task_update_time,
                        a.p
                        from task a,
                        task_dim b,
                        employee c
                        where a.task_id= b.task_id
                        and a.user_id= c.user_id and a.company_id=c.company_id and a.company_id in (?)
                    )`,
}
