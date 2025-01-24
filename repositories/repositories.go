package repositories

import (
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/RomainMarcazzan/ApiRest/config"
	"github.com/RomainMarcazzan/ApiRest/models"
)

func InitDB() {
	// Create users table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS Users (
		id UUID PRIMARY KEY,
		name VARCHAR(100),
		email VARCHAR(100) UNIQUE NOT NULL
	);`
	_, err := config.DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Error creating Users table: %v", err)
	}

	// Create Notif table
	createNotifTable := `
	CREATE TABLE IF NOT EXISTS Notif (
    	id UUID PRIMARY KEY,
    	message TEXT,
    	notifSongId VARCHAR(100),
    	createdAt TIMESTAMPTZ DEFAULT NOW(),
    	isView BOOLEAN,
    	receiverId UUID REFERENCES Users(id),
    	avatar VARCHAR(255)
);`

	_, err = config.DB.Exec(createNotifTable)

	if err != nil {
		log.Fatalf("Error creating Notif table: %v", err)
	}

	//Create ProPosition table
	createProPositionTable := `
	CREATE TABLE IF NOT EXISTS ProPosition (
    	id UUID PRIMARY KEY,
    	proId UUID REFERENCES Users(id),
    	latitude FLOAT,
    	longitude FLOAT,
    	timestamp TIMESTAMPTZ DEFAULT NOW()
);`

	_, err = config.DB.Exec(createProPositionTable)

	if err != nil {
		log.Fatalf("Error creating ProPosition table: %v", err)
	}

	log.Println("Database initialized successfully!")
}

// Users

func GetAllUsers() ([]models.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user models.User) error {
	_, err := config.DB.Exec("INSERT INTO Users (id, name, email) VALUES ($1, $2, $3)", user.ID, user.Name, user.Email)
	return err
}

func UpdateUser(user models.User) error {
	_, err := config.DB.Exec("UPDATE Users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
	return err
}

func DeleteUser(id uuid.UUID) error {
	_, err := config.DB.Exec("DELETE FROM Users WHERE id = $1", id)
	return err
}

// Notifs

func GetAllNotifs() ([]models.Notif, error) {
	rows, err := config.DB.Query("SELECT id, message, notifSongId, createdAt, isView, receiverId, avatar FROM Notif")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifs []models.Notif
	for rows.Next() {
		var notif models.Notif
		var createdAt *time.Time // Handle nullable time
		var avatar *string       // Handle nullable string

		if err := rows.Scan(&notif.ID, &notif.Message, &notif.NotifSongId, &createdAt, &notif.IsView, &notif.ReceiverId, &avatar); err != nil {
			return nil, err
		}
		notif.CreatedAt = createdAt
		notif.Avatar = avatar
		notifs = append(notifs, notif)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notifs, nil
}

func CreateNotif(notif models.Notif) error {
	_, err := config.DB.Exec("INSERT INTO Notif (id, message, notifSongId, createdAt, isView, receiverId, avatar) VALUES ($1, $2, $3, $4, $5, $6, $7)", notif.ID, notif.Message, notif.NotifSongId, notif.CreatedAt, notif.IsView, notif.ReceiverId, notif.Avatar)
	return err
}

func UpdateNotif(notif models.Notif) error {
	_, err := config.DB.Exec("UPDATE Notif SET message = $1, notifSongId = $2, createdAt = $3, isView = $4, receiverId = $5, avatar = $6 WHERE id = $7", notif.Message, notif.NotifSongId, notif.CreatedAt, notif.IsView, notif.ReceiverId, notif.Avatar, notif.ID)
	return err
}

func DeleteNotif(id uuid.UUID) error {
	_, err := config.DB.Exec("DELETE FROM Notif WHERE id = $1", id)
	return err
}

//PropPosition

func UpsertProPosition(proPosition models.ProPosition) error {
	query := `
	INSERT INTO ProPosition (id, proId, latitude, longitude, timestamp)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (id) 
	DO UPDATE SET
		proId = EXCLUDED.proId,
		latitude = EXCLUDED.latitude,
		longitude = EXCLUDED.longitude,
		timestamp = EXCLUDED.timestamp
	RETURNING id;
	`
	var returnedID uuid.UUID
	err := config.DB.QueryRow(query, proPosition.ID, proPosition.ProId, proPosition.Latitude, proPosition.Longitude, proPosition.Timestamp).Scan(&returnedID)
	if err != nil {
		return err
	}

	return nil
}
