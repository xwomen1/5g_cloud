package main

import (
	"encoding/json"
	"fmt"

	"main/data"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Các route khác...

	// AI Events routes
	router.HandleFunc("/api/aievents", GetAIEventsHandler).Methods("GET")
	router.HandleFunc("/api/aievents/{id}", GetAIEventHandler).Methods("GET")
	router.HandleFunc("/api/aievents", CreateAIEventHandler).Methods("POST")
	router.HandleFunc("/api/aievents/{id}", UpdateAIEventHandler).Methods("PUT")
	router.HandleFunc("/api/aievents/{id}", DeleteAIEventHandler).Methods("DELETE")

	// CORS middleware nếu cần
	router.Use(enableCors)

	http.ListenAndServe(":8080", router)
}

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Streaming Video Endpoint"))
}

func PlaybackHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Playback Video Endpoint"))
}

func AIEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AI Event Endpoint"))
}

// Thêm các handler mới
func StreamCameraHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cameraID := vars["id"]
	// Logic để lấy luồng video của camera tương ứng
	videoURL := fmt.Sprintf("/videos/live_camera_%s.mp4", cameraID)
	json.NewEncoder(w).Encode(map[string]string{"videoUrl": videoURL})
}

func GetAIEventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.AIEvents)

}
func GetAIEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, event := range data.AIEvents {
		if event.ID == id {
			json.NewEncoder(w).Encode(event)
			return
		}
	}
	http.Error(w, "Event not found", http.StatusNotFound)
}
func CreateAIEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEvent models.AIEvent
	json.NewDecoder(r.Body).Decode(&newEvent)
	newEvent.ID = len(data.AIEvents) + 1
	data.AIEvents = append(data.AIEvents, newEvent)
	json.NewEncoder(w).Encode(newEvent)
}

func UpdateAIEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var updatedEvent models.AIEvent
	json.NewDecoder(r.Body).Decode(&updatedEvent)

	for index, event := range data.AIEvents {
		if event.ID == id {
			updatedEvent.ID = id
			data.AIEvents[index] = updatedEvent
			json.NewEncoder(w).Encode(updatedEvent)
			return
		}
	}
	http.Error(w, "Event not found", http.StatusNotFound)
}

func DeleteAIEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, event := range data.AIEvents {
		if event.ID == id {
			data.AIEvents = append(data.AIEvents[:index], data.AIEvents[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Event not found", http.StatusNotFound)
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
