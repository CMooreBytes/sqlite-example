package datacontext

type Service struct{
	Id int
	ServiceDate string
}

type Person struct {
	Id int
	FirstName string
	LastName string
}

type Assignment struct {
	Assignee Person
	Service Service
}
