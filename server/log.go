package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) RespondWithError(w http.ResponseWriter, status int, message string, logs ...string) {
	s.ResponsWithJson(w, status, map[string]string{"error": message})
	if len(logs) > 0 {
		args := make([]interface{}, len(logs))
		for i, v := range logs {
			args[i] = v
		}
		s.Log.Error("responded with error: "+message, args...)
	}
}

func (s Server) ResponsWithJson(w http.ResponseWriter, stausCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		s.Log.Error("error marshaling while responding with json", "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stausCode)
	_, err = w.Write(response)
	if err != nil {
		s.Log.Error("error responding with json", "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
