package models

import "time"

type PetType struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex;not null"`
}

type Specialty struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex;not null"`
}

type Vet struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Specialties []Specialty `json:"specialties" gorm:"many2many:vet_specialties;"`
}

type Owner struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
	Pets      []Pet  `json:"pets"`
}

type Pet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	PetTypeID uint      `json:"petTypeId"`
	PetType   PetType   `json:"petType"`
	OwnerID   uint      `json:"ownerId"`
	Visits    []Visit   `json:"visits"`
}

type Visit struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	PetID       uint      `json:"petId"`
}
