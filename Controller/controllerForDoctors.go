package Controller

import (
	"SP/dataBase"
	"SP/module"
	"SP/responses"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct{}

func (d *Doctor) AddSubjectGrades(c *gin.Context) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var subject module.Subject

	if err := c.BindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	newItem := module.Subject{
		ID:         primitive.NewObjectID(),
		StudentId:  subject.StudentId,
		Name:       subject.Name,
		CourseWork: subject.CourseWork,
	}

	result, err := dataBase.UniSubjects.InsertOne(ctx, newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}

func (d *Doctor) GetSubjectById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	subjectId := c.Param("subjectId")
	var subject module.Subject
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(subjectId)

	err := dataBase.UniSubjects.FindOne(ctx, bson.M{"_id": objId}).Decode(&subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": subject}})
}

func (d *Doctor) AddGradesToTranscript(c *gin.Context) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	subjectId := c.Param("subjectId")
	var subject module.Subject

	if err := c.BindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	objId, _ := primitive.ObjectIDFromHex(subjectId)

	err := dataBase.UniSubjects.FindOne(ctx, bson.M{"_id": objId}).Decode(&subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	filter := bson.M{"studentId": subject.StudentId}
	update := bson.M{
		"$push": bson.M{
			"subjects": subject,
		},
	}

	result := dataBase.StudentTranscript.FindOneAndUpdate(ctx, filter, update)

	if result.Err() != nil {
		fmt.Println(result.Err())
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Subject Not Added"}})
		return
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Subject Added successfully"}})
}
