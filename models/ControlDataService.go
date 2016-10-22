package models

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

//get course catalog list by company ids
func GetCourseCatalog(companyIds []string) map[int]string  {
	catalogs := make(map[int]string)
	
	// ToDo call rest interface
	return catalogs
}
