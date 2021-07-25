package Controllers

import (
	"awesomeProject1/LearnGo/Day3/Exercises/Exercise2/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetStudents ... Get all students
func GetStudents(c *gin.Context) {
	var s []Models.Student
	err := Models.GetAllStudents(&s)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, s)
	}
}

//CreateStudent ... Create student
func CreateStudent(c *gin.Context) {
	var s Models.Student
	c.BindJSON(&s)
	fmt.Println(s)
	err := Models.CreateStudent(&s)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, s)
	}
}

//GetStudentByID ... Get the student by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var s Models.Student
	err := Models.GetStudentByID(&s, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, s)
	}
}

//UpdateStudent ... Update the student information
func UpdateStudent(c *gin.Context) {
	var s Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&s, id)
	if err != nil {
		c.JSON(http.StatusNotFound, s)
	}
	c.BindJSON(&s)
	err = Models.UpdateStudent(&s, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, s)
	}
}

//DeleteStudent ... Delete the student
func DeleteStudent(c *gin.Context) {
	var s Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&s, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
