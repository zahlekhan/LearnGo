package Models

import (
	"awesomeProject1/LearnGo/Day3/Exercises/Exercise2/Config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllStudents Fetch all student data
func GetAllStudents(s *[]Student) (err error) {
	if err = Config.DB.Find(s).Error; err != nil {
		return err
	}
	return nil
}

//CreateStudent ... Insert New data
func CreateStudent(s *Student) (err error) {
	if err = Config.DB.Create(s).Error; err != nil {
		return err
	}
	return nil
}

//GetStudentByID ... Fetch only one user by Id
func GetStudentByID(s *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(s).Error; err != nil {
		return err
	}
	return nil
}

//UpdateStudent ... Update user
func UpdateStudent(s *Student, id string) (err error) {
	Config.DB.Save(s)
	return nil
}

//DeleteStudent ... Delete user
func DeleteStudent(s *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(s)
	return nil
}
