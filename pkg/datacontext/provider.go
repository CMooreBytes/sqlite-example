package datacontext

import(
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Queryable interface{
	Query(index int)
}

func GetSchedule(index int) []Assignment{
	schedule := []Assignment{}
	query := `SELECT Persons.id as person_id, Persons.first_name, Persons.last_name, Services.id as service_id, Services.service_date 
			  FROM Assignments  INNER JOIN Persons ON Assignments.person_id = Persons.id 
			  INNER JOIN Services ON Assignments.service_id = Services.id
			  WHERE person_id = ?`
	db, err := sql.Open("sqlite3", "file:main.db")
	if err != nil { log.Fatal(err) }
	defer db.Close()
	rows, err := db.Query(query, index)
	if err != nil { log.Fatal(err) }
	defer rows.Close()
	for rows.Next() {
		a := new(Assignment)
		a.Service = Service{}
		a.Assignee = Person{}
		if err := rows.Scan(&a.Assignee.Id, &a.Assignee.FirstName, &a.Assignee.LastName, &a.Service.Id, &a.Service.ServiceDate); err != nil { log.Fatal(err) }
		schedule = append(schedule, *a)
	}
	return schedule
}

func GetPerson(index int) *Person {
	p := new(Person)
	query := "SELECT * FROM Persons WHERE id = ?"
	queryItem(query, index, func(row *sql.Row){
		if err := row.Scan(&p.Id, &p.FirstName, &p.LastName); err != nil { log.Fatal(err) }
		 
	})
	return p
}

func GetService(index int) *Service{
	s := new(Service)
	query := "SELECT * FROM Services WHERE id = ?"
	queryItem(query, index, func(row *sql.Row){
		if err := row.Scan(&s.Id, &s.ServiceDate); err != nil { log.Fatal(err) }	
	})
	return s
}


func queryItem(query string, index int, callback func(*sql.Row)){
	db, err := sql.Open("sqlite3", "file:main.db")
	if err != nil { log.Fatal(err) }
	defer db.Close()
	row := db.QueryRow(query, index)
	if err != nil { log.Fatal(err) }
	callback(row)
}

func query(query string, index int, callback func(*sql.Rows)){
	db, err := sql.Open("sqlite3", "file:main.db")
	if err != nil { log.Fatal(err) }
	defer db.Close()
	rows, err := db.Query(query, index)
	if err != nil { log.Fatal(err) }
	callback(rows)
}