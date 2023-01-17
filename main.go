package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	UID       string `json:"uid"`
	Name      string `json: "name"`
	Age       int32  `json: "age"`
	Gender    string `json: "gender"`
	Faculty   string `json: "faculty"`
	Semester  int32  `json: "semester"`
	Graduated bool   `json: "graduated"`
}

var students = []Student{
	{UID: "111", Name: "Muhammad Ali", Age: 20, Gender: "male", Faculty: "Computer Science", Semester: 1, Graduated: false},
	{UID: "112", Name: "David", Age: 21, Gender: "male", Faculty: "Data Science", Semester: 2, Graduated: false},
	{UID: "113", Name: "Fitri Aulia", Age: 19, Faculty: "Architecture", Semester: 3, Graduated: false},
}

func getStudents(Context *gin.Context) {
	Context.JSON(http.StatusOK, students)
}

func getStudentByID(Context *gin.Context) {
	uid := Context.Param("uid")

	for _, student := range students {
		if student.UID == uid {
			Context.JSON(http.StatusOK, student)
			return
		}
	}

	Context.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

func addStudent(Context *gin.Context) {
	var newStudent Student

	if err := Context.ShouldBindJSON(&newStudent); err != nil {
		Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	students = append(students, newStudent)
	Context.JSON(http.StatusOK, newStudent)
}

func editStudent(Context *gin.Context) {
	uid := Context.Param("uid")

	for index, student := range students {
		if student.UID == uid {
			var newStudent Student

			if err := Context.ShouldBindJSON(&newStudent); err != nil {
				Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			students[index] = newStudent
			Context.JSON(http.StatusOK, newStudent)
			return
		}
	}

	Context.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

func deleteStudent(Context *gin.Context) {
	uid := Context.Param("uid")

	for index, student := range students {
		if student.UID == uid {
			students = append(students[:index], students[index+1:]...)
			Context.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
			return
		}
	}

	Context.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

func main() {
	router := gin.Default()
	router.GET("/", func(Context *gin.Context) {
		Context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	router.GET("/students", getStudents)
	router.GET("/students/:uid", getStudentByID)
	router.PUT("/students/edit/:uid", editStudent)
	router.DELETE("/students/delete/:uid", deleteStudent)
	router.POST("/students", addStudent)
	router.Run("localhost:8080")
}
