package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"go-breeders/models"
)

// CatBreed is simply our target interface, which defines all the methods that
// any type shich impelments this interface must have.
type CatBreeds interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
}

// RemoteService is the Adapter type. It embeds a Data interface
// (which is critical to the pattern)
type RemoteService struct {
	Remote CatBreeds
}

// GetAllBreeds is the function on RemoteService whch lets us
// call any adapter which implements the Data type.
func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

// JSONBackend is the JSON adaptee, which needs to satisfy the CatBreeds interface
// by having a GetAllCatBreeds method.
type JSONBackend struct{}

func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breeds []*models.CatBreed

	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type catBreeds struct {
		XMLName struct{}           `xml:"cat-breeds"`
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}

	var breeds catBreeds

	err = xml.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nbreeds: %v\n", breeds)

	return breeds.Breeds, nil
}
