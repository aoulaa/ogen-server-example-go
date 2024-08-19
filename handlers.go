package main

import (
	"context"
	petstore "ogen-go/petstore"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddPet implements addPet operation.
	//
	// Add a new pet to the store.
	//
	// POST /pet
	AddPet(ctx context.Context, req *petstore.Pet) (*petstore.Pet, error)
	// DeletePet implements deletePet operation.
	//
	// Deletes a pet.
	//
	// DELETE /pet/{petId}
	DeletePet(ctx context.Context, params petstore.DeletePetParams) error
	// GetPetById implements getPetById operation.
	//
	// Returns a single pet.
	//
	// GET /pet/{petId}
	GetPetById(ctx context.Context, params petstore.GetPetByIdParams) (petstore.GetPetByIdRes, error)
	// UpdatePet implements updatePet operation.
	//
	// Updates a pet in the store.
	//
	// POST /pet/{petId}
	UpdatePet(ctx context.Context, params petstore.UpdatePetParams) error
}
