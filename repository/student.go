package repository

import (
	"cobadulu/models"
	"database/sql"
	"log"
)

func GetAllStudent(db *sql.DB) []*models.Student {
	var student []*models.Student
	rows, err := db.Query(`select * from student`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var std = new(models.Student)
		err := rows.Scan(&std.StudentID, &std.StudentfName, &std.StudentlName, &std.StudentAddress)
		if err != nil {
			log.Fatal(err)
		}
		student = append(student, std)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return student
}

func AddNewStudent(db *sql.DB, student *models.Student) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("insert into student values	(uuid(),?,?,?)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(student.StudentfName, student.StudentlName, student.StudentAddress); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
func DeleteStudent(db *sql.DB, student *models.Student) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from student where studentID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(student.StudentID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func UpdateStudent(db *sql.DB, student *models.Student) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("update student set studentfName = ?, studentlName = ?, studentAddress = ? where studentID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(student.StudentfName, student.StudentlName, student.StudentAddress, student.StudentID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
