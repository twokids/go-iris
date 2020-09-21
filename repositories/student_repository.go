package repositories

import "go-iris/datamodels"

//接口
type StudentRepository interface {
	GetStudent() datamodels.Student
}

//实现类
type StudentManager struct {
}

//构造函数，controller中使用
func NewStudentManager() StudentRepository {
	return &StudentManager{}
}

//实现具体方法
func (c *StudentManager) GetStudent() datamodels.Student {
	stu := datamodels.Student{}
	return stu
}
