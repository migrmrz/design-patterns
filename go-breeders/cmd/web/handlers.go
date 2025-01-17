package main

import (
	"fmt"
	"go-breeders/pets"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	_ = tools.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	_ = tools.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadGateway)
	}

	_ = tools.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadGateway)
	}

	_ = tools.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	dogBreeds, err := app.App.Models.DogBreed.GetAll()
	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadRequest)

		return
	}

	_ = tools.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	// create a dog using the builder pattern
	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("A mixed breed of unknown origin. Probably has some German Shepherd heritage.").
		SetColor("Black and White").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = tools.WriteJSON(w, http.StatusOK, p)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	// create a cat using the builder pattern
	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("ragdoll").
		SetWeight(15).
		SetDescription("Ragdoll").
		SetColor("White with black and brown").
		SetGeographicOrigin("United States").
		SetAge(2).
		SetAgeEstimated(false).
		Build()

	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = tools.WriteJSON(w, http.StatusOK, p)
}

// GetAllCatBreeds get all cat breeds from some source (using an adapter) and returns it as JSON.
func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	// Since we are using the adapter pattern, this handler does not care where it get the data from
	// and it will simply use whatever is stored in app.catService.
	catBreeds, err := app.App.CatService.GetAllBreeds()
	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = tools.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	// Setup toolbox
	var tools toolbox.Tools

	// Get species from URL itself
	species := chi.URLParam(r, "species")

	// Get breed from the URL
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)

	// Create a pet from abstract factory
	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadRequest)
	}

	// Write the result as JSON.
	_ = tools.WriteJSON(w, http.StatusOK, pet)
}
