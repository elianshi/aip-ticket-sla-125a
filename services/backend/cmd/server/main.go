package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type TicketMVP struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Severity  string    `json:"severity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type MVPStore struct {
	mu      sync.RWMutex
	tickets []TicketMVP
}

func newMVPMux(st *MVPStore) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "sprint": "1", "project": "云原生故障工单与SLA敏捷监控台"})
	})
	mux.HandleFunc("GET /api/tickets", func(w http.ResponseWriter, _ *http.Request) {
		st.mu.RLock()
		defer st.mu.RUnlock()
		_ = json.NewEncoder(w).Encode(map[string]any{"items": st.tickets})
	})
	mux.HandleFunc("POST /api/tickets", func(w http.ResponseWriter, r *http.Request) {
		var in TicketMVP
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Title == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		st.mu.Lock()
		in.ID = fmt.Sprintf("INC-%d", 100+len(st.tickets)+1)
		in.Status = "open"
		in.CreatedAt = time.Now().UTC()
		st.tickets = append(st.tickets, in)
		st.mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(in)
	})
	return mux
}

func main() {
	port := flag.Int("port", 18800, "listen port")
	flag.Parse()
	st := &MVPStore{tickets: []TicketMVP{{ID: "INC-101", Title: "Sprint 1 MVP 初始告警样本", Severity: "P1", Status: "open", CreatedAt: time.Now().UTC()}}}
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", *port), newMVPMux(st)))
}
