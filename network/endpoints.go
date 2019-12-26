package network

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
}

func (s *srv) insertHandler(w http.ResponseWriter, r *http.Request) {
	val, err := strconv.Atoi(mux.Vars(r)["val"])
	if err != nil {
		http.Error(w, "Invalid value to insert.", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusAccepted)
	go func() {
		c := make(chan int)
		go insertIntoBt(s, c)
		c <- val
	}()
}

func insertIntoBt(s *srv, val <-chan int) {
	value := <-val
	s.m.Lock()
	defer s.m.Unlock()
	s.bt.Insert(value)
}

func (s *srv) containsHandler(w http.ResponseWriter, r *http.Request) {
	val, err := strconv.Atoi(mux.Vars(r)["val"])
	if err != nil {
		http.Error(w, "Invalid value check.", http.StatusBadRequest)
	}
	w.Write([]byte(strconv.FormatBool(s.bt.Contains(val))))
}

func (s *srv) getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.bt.PrintLtr()))
}
