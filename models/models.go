package models

// Model of student
type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// Model of exam
type Exam struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Model of enrollment
type Enrollment struct {
	StudentId string `json:"fk_student_id"`
	ExamId    string `json:"fk_exam_id"`
}

// Model of question
type Question struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	ExamId   string `json:"fk_exam_id"`
}

// Model of answer
type Answer struct {
	Id         string `json:"id"`
	Answer     string `json:"answer"`
	QuestionId string `json:"fk_question_id"`
}
