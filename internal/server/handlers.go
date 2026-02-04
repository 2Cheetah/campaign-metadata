package server

import (
	"encoding/json"
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
