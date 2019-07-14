package main

import(
	"log"
	"github.com/cmoorebytes/sqlite/pkg/datacontext"
)


func main(){
	s := datacontext.GetSchedule(1)
	for _, a := range s {
		log.Printf("Date %s: Assignee: %s, %s", a.Service.ServiceDate, a.Assignee.LastName, a.Assignee.FirstName)
	}
	
}