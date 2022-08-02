package repository

import (
	"context"

	"github.com/ChrisCodeX/gRPC/models"
)

// Repository interface
type Repository interface {
	/* Student Service */
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	/* Exam Service */
	//Unary Methods
	GetExam(ctx context.Context, id string) (*models.Exam, error)
	SetExam(ctx context.Context, exam *models.Exam) error
	// Stream from client
	SetQuestion(ctx context.Context, question *models.Question) error
}

// Assign Repository
var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

// Abstract Implementations

/* Service Students Operations */
// Get a student
func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

// Insert a student
func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

/* Service Exams Operations */
// Get a Exam
func GetExam(ctx context.Context, id string) (*models.Exam, error) {
	return implementation.GetExam(ctx, id)
}

// Insert a Exam
func SetExam(ctx context.Context, exam *models.Exam) error {
	return implementation.SetExam(ctx, exam)
}

// Insert Questions
func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}
