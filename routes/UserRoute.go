package routes

import (
	"SP/Controller"
	"SP/Middlewares"

	"github.com/gin-gonic/gin"
)

var auth Middlewares.Auth
var Doc Controller.Doctor
var Student Controller.Student

func UserRoute(router *gin.Engine) {
	router.POST("/addSubjectGrades", auth.AuthForDoctors, Doc.AddSubjectGrades)
	router.GET("/getById/:subjectId", auth.AuthForDoctors, Doc.GetSubjectById)
	router.POST("/AddGradesToTranscript/:subjectId", auth.AuthForDoctors, Doc.AddGradesToTranscript)
	router.GET("/GetTranscript/:TranscriptId", auth.AuthForStudents, Student.GetTranscript)
	router.POST("/createTranscript", auth.AuthForStudents, Student.CreateTranscript)
}
