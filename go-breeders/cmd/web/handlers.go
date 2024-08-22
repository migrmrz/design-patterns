package main

import (
	"fmt"
	"go-breeders/pets"
	"net/http"

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
		SetAge(3).SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = tools.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = tools.WriteJSON(w, http.StatusOK, p)
}
