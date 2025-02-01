package main



import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type Car struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Model  string `json:"model"`
	Status string `json:"status"`
}

var db *gorm.DB
var err error

func initDB() {
	db, err = gorm.Open(sqlite.Open("cars.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	
	db.AutoMigrate(&Car{})
}


func getCars(w http.ResponseWriter, r *http.Request) {
	var cars []Car
	db.Find(&cars)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}


func getCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var car Car

	if err := db.First(&car, id).Error; err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "mkch", http.StatusNotFound)
		} else {
			http.Error(w, "mfihach db", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}
func createCar(w http.ResponseWriter, r *http.Request) {
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "mfihach", http.StatusBadRequest)
		return
	}
	car.Status = "available"

	db.Create(&car)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}


func updateCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var car Car
	if err := db.First(&car, id).Error; err != nil {
		http.Error(w, "mkch", http.StatusNotFound)
		return
	}
	var updatedData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "mfihach", http.StatusBadRequest)
		return
	}
	if status, exists := updatedData["status"]; exists {
		car.Status = status
		db.Save(&car)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var car Car

	if err := db.First(&car, id).Error; err != nil {
		http.Error(w, "mkch cars", http.StatusNotFound)
		return
	}
	db.Delete(&car)
	w.WriteHeader(http.StatusNoContent)
}

func handleRequests() {
	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCars(w, r)
		case "POST":
			createCar(w, r)
		default:
			http.Error(w, "Mkch methode", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/car", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCar(w, r)
		case "PUT":
			updateCar(w, r)
		case "DELETE":
			deleteCar(w, r)
		default:
			http.Error(w, "Mkch methode", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	initDB()
	handleRequests()
}
