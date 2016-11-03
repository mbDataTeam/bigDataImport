package util

import "fmt"

//radio button or checkbox for yes or no data
func InitYesOrNo() map[int]string {
	return map[int]string{
		1 : "是",
		0 : "否",
	}
}

//radio button or checkbox for gender data
func InitGenders() map[int]string {
	return map[int]string{
		1 : "男",
		0 : "女",
	}
}

//radio button or checkbox for condtionType data
func InitConditionType() map[string]string {
	return map[string]string{
		"指定人" : "指定人",
		"按照条件" : "按照条件",
	}
}

//radio button or checkbox for task type data
func InitTaskType() map[string]string {
	return map[string]string{
		"特殊任务" : "特殊任务",
		"每日任务" : "每日任务",
		"专区任务" : "专区任务",
		"未知" : "未知",
	}
}

//get course catalog list by company ids
func GetParentCatalogs(companyIds string) *[]SelectSchema  {
	sql := `select distinct Parent_Category_name, Parent_Category_Id
			from mb.view.company_course
			where company_id in(`+companyIds+`)`
	queryResult := QueryData(sql)
	catalogs := []SelectSchema{}
	for _,row := range queryResult.Rows{
		catalogs = append(catalogs,SelectSchema{
			OptionText: fmt.Sprintf("%v",row[0]), OptionValue:fmt.Sprintf("%v",row[1]),
		})
	}
	return &catalogs
}

func GetCatlogs(companyIds string,parentCalaIds string) *[]SelectSchema  {
	sql := `select distinct Category_name, Category_Id,Parent_Category_name
			from mb.view.company_course
			where company_id in(`+companyIds+`) and Parent_Category_id in (`+parentCalaIds+`)`
	queryResult := QueryData(sql)
	catalogs := []SelectSchema{}
	for _,row := range queryResult.Rows{
		catalogs = append(catalogs,SelectSchema{
			OptionText: fmt.Sprintf("%v",row[0]), OptionValue:fmt.Sprintf("%v",row[1]), GroupName:fmt.Sprintf("%v",row[2]),
		})
	}
	return &catalogs
}

func GetCourseList(companyIds string,parentCalaIds string) *[]SelectSchema  {
	sql := `select distinct title,course_id,type
			from mb.view.company_course
			where company_id in(`+companyIds+`) and category_id in (`+parentCalaIds+`)`
	queryResult := QueryData(sql)
	courses := []SelectSchema{}
	for _,row := range queryResult.Rows{
		courses = append(courses,SelectSchema{
			OptionText: fmt.Sprintf("%v",row[0]), OptionValue:fmt.Sprintf("%v",row[1]), GroupName:fmt.Sprintf("%v",row[2]),
		})
	}
	return &courses
}
