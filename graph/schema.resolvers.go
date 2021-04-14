package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/katbug/dm-graph/db/file"
	"github.com/katbug/dm-graph/graph/generated"
	"github.com/katbug/dm-graph/graph/model"
)

func (r *mutationResolver) CreatePerson(ctx context.Context, input model.NewPerson) (*model.Person, error) {
	var person model.Person
	person.Name = input.Name
	person.Description = input.Description
	return &person, nil
}

func (r *queryResolver) People(ctx context.Context) ([]*model.Person, error) {
	/*var people []*model.Person
	dummyPerson := model.Person{
		Name:        "Blah Name",
		Description: "testing description",
	}
	people = append(people, &dummyPerson) */
	fileDB := file.New("./data/people.csv")
	people, err := fileDB.ReadPeople()
	if err != nil {
		fmt.Println("Recieved error: ", err)
		return nil, err
	}
	return people, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
