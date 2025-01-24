package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/RomainMarcazzan/ApiRest/models"
	"github.com/RomainMarcazzan/ApiRest/services"
	"github.com/google/uuid"
)

// Public handler functions exposed to the main package
func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllUsers(w, r)
	case "POST":
		createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/v1/users/")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "PUT":
		updateUser(w, r, id)
	case "DELETE":
		deleteUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Internal handler functions, not exported
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, users, http.StatusOK)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}

	user.ID = uuid.New()

	if err := services.AddUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, user, http.StatusCreated)
}

func updateUser(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id // Ensure the ID is set

	if err := services.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, user, http.StatusOK)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := services.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func HandleNotifs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllNotifs(w, r)
	case "POST":
		createNotif(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleNotifByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/v1/notif/")
	log.Println(idStr)
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "PUT":
		updateNotif(w, r, id)
	case "DELETE":
		deleteNotif(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllNotifs(w http.ResponseWriter, r *http.Request) {
	notifs, err := services.GetNotifs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, notifs, http.StatusOK)
}

func createNotif(w http.ResponseWriter, r *http.Request) {
	var notif models.Notif
	if err := json.NewDecoder(r.Body).Decode(&notif); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if notif.Message == "" || notif.ReceiverId == uuid.Nil {
		http.Error(w, "Message and ReceiverId are required", http.StatusBadRequest)
		return
	}

	if notif.CreatedAt == nil {
		now := time.Now()
		notif.CreatedAt = &now
	}

	notif.ID = uuid.New()

	if err := services.AddNotif(notif); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, notif, http.StatusCreated)
}

func updateNotif(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var notif models.Notif
	if err := json.NewDecoder(r.Body).Decode(&notif); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	notif.ID = id // Ensure the ID is set

	if err := services.UpdateNotif(notif); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, notif, http.StatusOK)
}

func deleteNotif(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := services.DeleteNotif(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func HandleProPosition(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		upsertProPosition(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func upsertProPosition(w http.ResponseWriter, r *http.Request) {
	var proPosition models.ProPosition
	if err := json.NewDecoder(r.Body).Decode(&proPosition); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if proPosition.Timestamp == nil {
		now := time.Now()
		proPosition.Timestamp = &now
	}

	proPosition.ID = uuid.New()
	if err := services.UpsertProPosition(proPosition); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, proPosition, http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
