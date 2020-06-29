package services

import (
	"cobadulu/models"
	"cobadulu/repository"
	"encoding/json"
	"log"
	"net/http"
)

func (ps *SchoolService) GetStudent() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ID := r.FormValue("studentID")
			dataUser := repository.GetAllStudent(ps.db)
			if ID != "" {
				for _, each := range dataUser {
					if each.StudentID == ID {
						response := models.Res{}
						response.Msg = "getStudentbyId"
						response.Data = each
						byteOfUser, err := json.Marshal(response)

						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusNotFound)
						w.Write([]byte(byteOfUser))
					}
				}
			} else {
				var resp = models.Res{Msg: "getStudent", Data: repository.GetAllStudent(ps.db)}
				studentData, err := json.Marshal(resp)
				if err != nil {
					log.Fatal(err)
				}
				w.Header().Set("content-type", "application/json")
				w.Write(studentData)
				log.Println("Endpoint hit: getStudent")

			}

		case "POST":
			var std models.Student
			json.NewDecoder(r.Body).Decode(&std)
			_, err := repository.AddNewStudent(ps.db, &std)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: postStudent")
		case "DELETE":
			studentID := r.FormValue("studentID")
			student := &models.Student{
				StudentID: studentID,
			}
			_, err := repository.DeleteStudent(ps.db, student)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: deleteStudent")
		case "PUT":
			studentID := r.FormValue("studentID")
			studentfName := r.FormValue("firstname")
			studentlName := r.FormValue("lastname")
			studentAddress := r.FormValue("address")

			student := &models.Student{
				StudentID:      studentID,
				StudentfName:   studentfName,
				StudentlName:   studentlName,
				StudentAddress: studentAddress,
			}
			_, err := repository.UpdateStudent(ps.db, student)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: UpdateStudent")
		default:
			http.Error(w, "Not Found", 404)
		}

	}
}
