package setting

var SQLView map[string]string  = map[string]string{
	"task_data_export_view": `(select a.task_user_id,
                        a.task_id,
                        b.task_name,
                        case when b.task_type= 1 then '特殊任务' else '每日任务' end task_type,
                        c.user_id,
                        c.employee_id,c.realname,c.department,
                          case when b.expired_time=0 then '' else date_format(from_unixtime(b.expired_time),'%Y%m%d') end as task_expired_time,
                        date_format(from_unixtime(b.create_time), '%Y-%m-%d %H:%i:%s') as task_create_time,
                       date_format( from_unixtime(b.update_time), '%Y-%m-%d %H:%i:%s') as task_update_time ,
                        b.gold_num,
                        b.diamond_num,
                        b.description,
                        b.condition_type,
                       a.user_task_create_time,
                       a.user_task_update_time
                       from (select * from (select a.*,rank() over(partition by task_user_id order by user_task_update_time desc) rank from (
                 select distinct a.task_user_id,a.task_id, a.user_id, a.task_status, a.sub_task_total_num, a.sub_task_finish_num, a.task_day, date_format(from_unixtime(a.task_create_time), '%Y-%m-%d %H:%i%s') user_task_create_time, date_format(from_unixtime(a.task_update_time), '%Y-%m-%d %H:%i%s') user_task_update_time
                   from task a
                  where p between date_format(current_date -interval '30' day,'%Y%m%d')
                    and date_format(current_date,'%Y%m%d')
                    and company_id in (?))a) where rank=1 )a,
                          task_dim b,
                          (
                 select *
                   from(
                 select a.*, rank() over(partition by company_id, employee_id
                  order by update_time desc) rank
                   from user_meta a
                  where company_id in (?))
                  where is_test!= 1
                    and rank= 1) c
                  where a.task_id= b.task_id
                    and a.user_id= c.user_id)`,
}
