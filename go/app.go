package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	geojson "github.com/paulmach/go.geojson"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname, hostname string) {

	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, dbname, hostname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func (a *App) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cleabs := vars["id"]
	path := r.URL.Path

	var premium *geojson.FeatureCollection
	var err error
	if strings.Contains(path, "adr-parc") {
		premium, err = getAdrParc(a.DB, "id", cleabs)
	} else if strings.Contains(path, "adr-bat") {
		premium, err = getAdrBati(a.DB, "id", cleabs)
	} else if strings.Contains(path, "bati-parc") {
		premium, err = getBatiParc(a.DB, "id", cleabs)
	}

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Aucun résultat trouvé")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, premium)
}

func (a *App) findById(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	cleabs := v.Get("id")
	path := r.URL.Path

	var premium *geojson.FeatureCollection
	var err error
	if strings.Contains(path, "findByAdrId") && strings.Contains(path, "adr-parc") {
		premium, err = getAdrParc(a.DB, "id_adr", cleabs)
	} else if strings.Contains(path, "findByAdrId") && strings.Contains(path, "adr-bati") {
		premium, err = getAdrBati(a.DB, "id_adr", cleabs)
	} else if strings.Contains(path, "findByBatId") {
		premium, err = getBatiParc(a.DB, "id_bat", cleabs)
	} else if strings.Contains(path, "findByParcId") {
		premium, err = getBatiParc(a.DB, "idu", cleabs)
	}

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Aucun résultat trouvé")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, premium)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/adr-parc/{id:ADR_PARC[0-9]+}", a.getById).Methods("GET")
	a.Router.HandleFunc("/adr-bati/{id:ADR_BATI[0-9]+}", a.getById).Methods("GET")
	a.Router.HandleFunc("/bati-parc/{id:BAT_PARC[0-9]+}", a.getById).Methods("GET")
	a.Router.HandleFunc("/adr-parc/findByAdrId", a.findById).Queries("id", "{id:ADRNIVX_[0-9]+}").Methods("GET")
	a.Router.HandleFunc("/adr-bati/findByAdrId", a.findById).Queries("id", "{id:ADRNIVX_[0-9]+}").Methods("GET")
	a.Router.HandleFunc("/bati-parc/findByBatId", a.findById).Queries("id", "{id:BATIMENT[0-9]+}").Methods("GET")
	a.Router.HandleFunc("/bati-parc/findByParcId", a.findById).Queries("id", "{id}").Methods("GET")
}
