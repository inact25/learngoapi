package services

import (
	"cobadulu/models"
	"cobadulu/repository"
	"encoding/json"
	"log"
	"net/http"
)

func (ps *SchoolService) GetTeacher() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ID := r.FormValue("teacherID")
			teacher := repository.GetAllTeacher(ps.db)
			if ID != "" {
				for _, each := range teacher {
					if each.TeacherID == ID {
						response := models.Res{}
						response.Msg = "getTeacherbyId"
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
				var resp = models.Res{Msg: "getTeacher", Data: repository.GetAllTeacher(ps.db)}
				teacherData, err := json.Marshal(resp)
				if err != nil {
					log.Fatal(err)
				}
				w.Header().Set("content-type", "application/json")
				w.Write(teacherData)
				log.Println("Endpoint hit: getTeacher")

			}
		case "POST":
			var tcs models.Teacher
			json.NewDecoder(r.Body).Decode(&tcs)
			_, err := repository.AddNewTeacher(ps.db, &tcs)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: postSubject")
		case "DELETE":
			teacherID := r.FormValue("teacherID")
			teacher := &models.Teacher{
				TeacherID: teacherID,
			}
			_, err := repository.DeleteTeacher(ps.db, teacher)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: deleteTeacher")
		case "PUT":
			teacherID := r.FormValue("teacherID")
			teacherfName := r.FormValue("firstname")
			teacherlName := r.FormValue("lastname")
			teacherAddress := r.FormValue("address")

			teacher := &models.Teacher{
				TeacherID:      teacherID,
				TeacherfName:   teacherfName,
				TeacherlName:   teacherlName,
				TeacherAddress: teacherAddress,
			}
			_, err := repository.UpdateTeacher(ps.db, teacher)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: UpdateTeacher")
		default:
			http.Error(w, "Not Found", 404)
		}
	}
}
