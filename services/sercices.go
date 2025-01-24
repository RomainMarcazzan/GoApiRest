package services

import (
	"github.com/RomainMarcazzan/ApiRest/models"
	"github.com/RomainMarcazzan/ApiRest/repositories"
	"github.com/google/uuid"
)

func GetUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func AddUser(user models.User) error {
	return repositories.CreateUser(user)
}

func UpdateUser(user models.User) error {
	return repositories.UpdateUser(user)
}

func DeleteUser(id uuid.UUID) error {
	return repositories.DeleteUser(id)
}

func GetNotifs() ([]models.Notif, error) {
	return repositories.GetAllNotifs()
}

func AddNotif(notif models.Notif) error {
	return repositories.CreateNotif(notif)
}

func UpdateNotif(notif models.Notif) error {
	return repositories.UpdateNotif(notif)
}

func DeleteNotif(id uuid.UUID) error {
	return repositories.DeleteNotif(id)
}

func UpsertProPosition(proPosition models.ProPosition) error {
	return repositories.UpsertProPosition(proPosition)
}
