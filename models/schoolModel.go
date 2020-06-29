package models

type Res struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type Student struct {
	StudentID      string `json:"studentID"`
	StudentfName   string `json:"firstName"`
	StudentlName   string `json:"lastName"`
	StudentAddress string `json:"address"`
}

type Subject struct {
	SubjectID   string `json:"subjectID"`
	SubjectDesc string `json:"subjectDesc"`
}

type Teacher struct {
	TeacherID      string `json:"teacherID"`
	TeacherfName   string `json:"firstName"`
	TeacherlName   string `json:"lastName"`
	TeacherAddress string `json:"address"`
}
