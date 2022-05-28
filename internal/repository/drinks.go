package repository

import (
	"barbot/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type DrinksRepository interface {
	GetAll() []domain.Drink
}

type drinksRepository struct {
	filepath string
	db       []domain.Drink
}

func NewDrinksRepository(filepath string) DrinksRepository {
	r := drinksRepository{filepath: filepath}

	if err := r.ReadFile(); err != nil {
		log.Fatalln(err)
	}

	return &r
}

func (r *drinksRepository) ReadFile() error {
	if _, err := os.Stat(r.filepath); errors.Is(err, os.ErrNotExist) {
		return errors.New(fmt.Sprintf("file %s doesn't exists", r.filepath))
	}

	// Open file and create the reader
	file, err := os.Open(r.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &r.db)
	return err
}

func (r *drinksRepository) GetAll() []domain.Drink {
	return r.db
}
