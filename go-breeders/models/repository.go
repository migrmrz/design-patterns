package models

import "database/sql"

// Repository is the database repository. Anything that implements
// this interface must implement all the methods included here.
type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
}

// mySQLRepository is a simple wrapper for the *sql.DB type.
// This is used to return a MySQL/MariaDB repository.
type mySQLRepository struct {
	DB *sql.DB
}

// newMySQLRepository is a convenience factory
// method to return a new newMySQLRepository.
func newMySQLRepository(conn *sql.DB) Repository {
	return &mySQLRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
