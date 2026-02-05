package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func (s *Server) PingHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Pong int `json:"pong"`
	}{
		Pong: 1,
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("couldn't encode pong response to json, error: %v", err)
	}
}

func (s *Server) GetCampaignTagsHanlder(w http.ResponseWriter, r *http.Request) {
	campaignId := r.PathValue("id")
	if campaignId == "" {
		http.Error(w, "campaign id can't be empty", http.StatusBadRequest)
		return
	}
	tags, err := s.Service.GetCampaignTags(r.Context(), campaignId)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	resp := struct {
		CampaignId string   `json:"campaignId"`
		Tags       []string `json:"tags"`
	}{
		CampaignId: campaignId,
		Tags:       tags,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("couldn't encode response to json, error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"web/templates/pages/index.html",
		"web/templates/partials/items.html",
	)
	if err != nil {
		log.Printf("couldn't parse index.html template, error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	type Item struct {
		Name string
		ID   int
	}

	data := struct {
		Items []Item
	}{
		Items: []Item{
			{Name: "Item 1", ID: 1},
			{Name: "Item 2", ID: 2},
		},
	}
	if err := tmpl.ExecuteTemplate(w, "index", data); err != nil {
		log.Printf("couldn't execute template, error: %v", err)
		http.Error(w, "couldn't render the page", http.StatusInternalServerError)
		return
	}
}
