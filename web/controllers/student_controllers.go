package controllers

import (
	"github.com/kataras/iris/mvc"
	"go-iris/repositories"
	"go-iris/services"
)

type StudentController struct {

}

func (c *StudentController) Get() mvc.View {
	studentRepository:=repositories.NewStudentManager()
	studentService:=services.NewStudentServiceManager(studentRepository)

	StudentResult := studentService.GetStudent()

	return mvc.View{
		Name: "student/index.html",
		Data: StudentResult,
	}
}
