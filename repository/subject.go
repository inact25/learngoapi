package repository

import (
	"cobadulu/models"
	"database/sql"
	"log"
)

func GetAllSubject(db *sql.DB) []*models.Subject {
	rows, err := db.Query(`select * from subject`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	subject := make([]*models.Subject, 0)
	for rows.Next() {
		var s = new(models.Subject)
		err := rows.Scan(&s.SubjectID, &s.SubjectDesc)
		if err != nil {
			log.Fatal(err)
		}
		subject = append(subject, s)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return subject
}
func AddNewSubject(db *sql.DB, subject *models.Subject) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("insert into subject values	(uuid(),?)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(subject.SubjectDesc); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func DeleteSubject(db *sql.DB, subject *models.Subject) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from subject where subjectID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(subject.SubjectID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func UpdateSubject(db *sql.DB, subject *models.Subject) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("update subject set subjectDesc = ? where subjectID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(subject.SubjectDesc, subject.SubjectID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
