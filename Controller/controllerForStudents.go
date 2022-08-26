package Controller

import (
	"SP/dataBase"
	"SP/module"
	"SP/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct{}

func (s *Student) CreateTranscript(c *gin.Context) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var transcript module.Transcript

	if err := c.BindJSON(&transcript); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	studentTranscript := module.Transcript{
		ID:        primitive.NewObjectID(),
		StudentId: transcript.StudentId,
		Subjects:  transcript.Subjects,
	}

	result, err := dataBase.StudentTranscript.InsertOne(ctx, studentTranscript)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
}

func (s *Student) GetTranscript(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("TranscriptId")
	var transcript module.Transcript
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := dataBase.StudentTranscript.FindOne(ctx, bson.M{"_id": objId}).Decode(&transcript)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": transcript}})

}
