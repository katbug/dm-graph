package file

import (
	"encoding/csv"
	"io"
	"os"

	db "github.com/katbug/dm-graph/db"
	"github.com/katbug/dm-graph/graph/model"
)

type File struct {
	Path string
}

func New(path string) db.Database {
	return &File{Path: path}
}

func (f *File) WritePerson(p *model.Person) {
	/*w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}*/
}

func (f *File) ReadPeople() ([]*model.Person, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}

	people := []*model.Person{}

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		person := model.Person{Name: record[0], Description: record[1]}

		people = append(people, &person)
	}
	return people, nil
}
