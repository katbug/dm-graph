package db

import "github.com/katbug/dm-graph/graph/model"

type Database interface {
	ReadPeople() ([]*model.Person, error)
	WritePerson(p *model.Person)
}
