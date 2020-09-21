package services

import (
	"go-iris/datamodels"
	"go-iris/repositories"
)

type StudentService interface {
	GetStudent() datamodels.Student
}

type StudentServiceManager struct {
	repo repositories.StudentRepository
}

func NewStudentServiceManager(repo repositories.StudentRepository) StudentService {
	return &StudentServiceManager{repo: repo}
}

func (s *StudentServiceManager) GetStudent() datamodels.Student {
	result := datamodels.Student{Name: "lisa"}
	return result
}
