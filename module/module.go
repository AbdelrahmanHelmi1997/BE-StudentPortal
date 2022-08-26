package module

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	RoleType  string             `bson:"roleType"`
	Token     string             `bson:"token"`
}
type LoginData struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	RoleType  string             `bson: "roleType"`
	Token     string             `bson:"token"`
}

type Subject struct {
	ID         primitive.ObjectID `bson:"_id"`
	StudentId  string             `bson:"studentId"`
	Name       string             `bson : "name"`
	CourseWork CourseWork         `bson : "courseWork"`
}

type CourseWork struct {
	Midterm    int `bson : "midterm"`
	Final      int `bson : "final"`
	Quizzes    int `bson : "quizzes"`
	Assigments int `bson : "assigments"`
}

type Transcript struct {
	ID        primitive.ObjectID `bson:"_id"`
	StudentId string             `bson:"studentId"`
	Subjects  []Subject          `bson:"subjects"`
}

type AddSubjectToStudent struct {
	StudentId   string `bson:"studentId"`
	SubjectName string `bson:"subjectName"`
}

type AddGrades struct {
	StudentId   string     `bson:"studentId"`
	SubjectName string     `bson:"subjectName"`
	CourseWork  CourseWork `bson : "courseWork"`
}
