package services

import (
	"cobadulu/models"
	"cobadulu/repository"
	"encoding/json"
	"log"
	"net/http"
)

func (ps *SchoolService) GetSubject() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ID := r.FormValue("subjectID")
			subject := repository.GetAllSubject(ps.db)
			if ID != "" {
				for _, each := range subject {
					if each.SubjectID == ID {
						response := models.Res{}
						response.Msg = "getSubjectbyId"
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
				var resp = models.Res{Msg: "getSubject", Data: repository.GetAllSubject(ps.db)}
				subjectData, err := json.Marshal(resp)
				if err != nil {
					log.Fatal(err)
				}
				w.Header().Set("content-type", "application/json")
				w.Write(subjectData)
				log.Println("Endpoint hit: getSubject")

			}
		case "POST":
			var sbj models.Subject
			json.NewDecoder(r.Body).Decode(&sbj)
			_, err := repository.AddNewSubject(ps.db, &sbj)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: postSubject")
		case "DELETE":
			subjectID := r.FormValue("subjectID")
			subject := &models.Subject{
				SubjectID: subjectID,
			}
			_, err := repository.DeleteSubject(ps.db, subject)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: deleteSubject")
		case "PUT":
			subjectID := r.FormValue("subjectID")
			subjectDesc := r.FormValue("subjectDesc")
			subject := &models.Subject{
				SubjectID:   subjectID,
				SubjectDesc: subjectDesc,
			}
			_, err := repository.UpdateSubject(ps.db, subject)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Endpoint hit: updateSubject")

		default:
			http.Error(w, "Not Found", 404)

		}
	}
}
