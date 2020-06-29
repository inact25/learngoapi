package repository

import (
	"cobadulu/models"
	"database/sql"
	"log"
)

func GetAllTeacher(db *sql.DB) []*models.Teacher {
	var teacher []*models.Teacher
	rows, err := db.Query(`select * from teacher`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t = new(models.Teacher)
		err := rows.Scan(&t.TeacherID, &t.TeacherfName, &t.TeacherlName, &t.TeacherAddress)
		if err != nil {
			log.Fatal(err)
		}
		teacher = append(teacher, t)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return teacher
}

func AddNewTeacher(db *sql.DB, teacher *models.Teacher) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("insert into teacher values	(uuid(),?,?,?)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(teacher.TeacherfName, teacher.TeacherlName, teacher.TeacherAddress); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func DeleteTeacher(db *sql.DB, teacher *models.Teacher) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from teacher where teacherID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(teacher.TeacherID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func UpdateTeacher(db *sql.DB, teacher *models.Teacher) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("update teacher set teacherfName = ?, teacherlName = ?, teacherAddress = ? where teacherID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(teacher.TeacherfName, teacher.TeacherlName, teacher.TeacherAddress, teacher.TeacherID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
