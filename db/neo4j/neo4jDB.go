package neo4jDB

import (
	"github.com/katbug/dm-graph/graph/model"
	"github.com/neo4j/neo4j-go-driver"
)

type Neo4jDB struct {
	Driver neo4j.Driver
}

func New(username, password string) Database {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}

	return &Neo4jDB{Driver: driver}
}

func (n *Neo4jDB) WritePerson(p *model.Person) {

}

func (n *Neo4jDB) ReadPeople() ([]*model.Person, error) {
	return []*model.Person{}
}
