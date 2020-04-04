package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type val struct {
	v string
	m sync.RWMutex
}

type app struct {
	table map[string]*val
}

func (a *app) set(k, v string) {
	for i := 0; i < 10000000; i++ {
		continue
	}

	_, exists := a.table[k]

	if !exists {
		a.table[k] = &val{}
	}

	a.table[k].m.Lock()
	a.table[k].v = v
	a.table[k].m.Unlock()
}

func (a *app) get(k string) string {
	for i := 0; i < 10000000; i++ {
		continue
	}

	val, exists := a.table[k]

	if !exists {
		return ""
	}

	val.m.RLock()
	defer val.m.RUnlock()
	return val.v
}

func logRequest(r *http.Request, t time.Time) {
	log.Printf("%s %s %v\n", r.Method, r.URL.EscapedPath(), time.Since(t))
}

func (a *app) cache(w http.ResponseWriter, r *http.Request) {
	defer logRequest(r, time.Now())

	if r.Method == http.MethodGet {
		key := r.URL.EscapedPath()[1:]

		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no req. key"))
			return
		}

		w.WriteHeader(200)

		val := a.get(key)
		w.Write([]byte(fmt.Sprintf("%s=%s", key, val)))
	} else if r.Method == http.MethodPost {
		q := r.URL.Query()

		key := r.URL.EscapedPath()[1:]
		val := q.Get("val")

		if key == "" || val == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong request"))
			return
		}

		w.WriteHeader(200)

		go a.set(key, val)

		w.Write([]byte(fmt.Sprintf("%s=%s", key, val)))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("wrong method"))
	}
}

func main() {
	// app := &app{make(map[string]*val)}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hi"))
	})

	http.ListenAndServe(":8000", mux)
}
