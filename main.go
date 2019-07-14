package main

import(
	"log"
	"github.com/cmoorebytes/sqlite-example/pkg/datacontext"
)


func main(){
	p := datacontext.GetPerson(1)
	log.Printf("id %d: Assignee: %s, %s", p.Id, p.LastName, p.FirstName)
	
	s := datacontext.GetSchedule(1)
	for _, a := range s {
		log.Printf("Date %s: Assignee: %s, %s", a.Service.ServiceDate, a.Assignee.LastName, a.Assignee.FirstName)
	}
}