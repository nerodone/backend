package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) RespondWithError(w http.ResponseWriter, status int, message string, logs ...string) {
	_ = s.ResponsWithJson(w, status, map[string]string{"error": message})
	if len(logs) > 0 {
		args := make([]interface{}, len(logs)+1)
		for i, v := range logs[1:] {
			args[i+1] = v
		}
		s.Log.Debug("responded with error: "+message, args)
	}
}

func (s Server) ResponsWithJson(w http.ResponseWriter, stausCode int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stausCode)
	_, err = w.Write(response)
	if err != nil {
		return err
	}
	return nil
}
