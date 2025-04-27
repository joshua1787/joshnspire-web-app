package main

import (
	"encoding/json"
	"net/http"
)

// Program struct for the /programs endpoint
type Program struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	// Serve static files like home.html, about.html, etc.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/programs", programsHandler)

	// Start server
	http.ListenAndServe(":8080", nil)
}

// Handler for the Home Page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

// Handler for the About Page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

// Handler for the Contact Page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

// Handler for the Programs API
func programsHandler(w http.ResponseWriter, r *http.Request) {
	programs := []Program{
		{Title: "DevOps Mastery", Description: "Become a certified DevOps Engineer with hands-on experience."},
		{Title: "Cloud Native with AWS", Description: "Deploy and scale applications in a real-world AWS environment."},
		{Title: "Docker & Kubernetes Pro", Description: "Master containerization and orchestration technologies."},
		{Title: "Site Reliability Engineering", Description: "Learn to manage reliable and scalable systems."},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(programs)
}
