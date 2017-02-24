package stats

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/nathanielc/morgoth"
)

type Server struct {
	mu        sync.RWMutex
	detectors map[string]*morgoth.Detector
}

func NewServer() *Server {
	return &Server{
		detectors: make(map[string]*morgoth.Detector),
	}
}

func (s *Server) Serve() {
	http.HandleFunc("/stats", s.handleStats)
	http.ListenAndServe(":6767", nil)
}

func (s *Server) SetDetector(name string, d *morgoth.Detector) {
	s.mu.Lock()
	s.detectors[name] = d
	s.mu.Unlock()
}

func (s *Server) handleStats(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	stats := make(map[string]morgoth.DetectorStats, len(s.detectors))
	for n, d := range s.detectors {
		stats[n] = d.Stats()
	}
	s.mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
}
