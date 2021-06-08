package main

import (
	"database/sql"
	"fmt"
	"os"

	geojson "github.com/paulmach/go.geojson"
)

type adrParc struct {
	id        string
	insee_com string
	id_adr    string
	idu       string
	type_lien string
	nb_adr    int
	nb_parc   int
	geometry  *geojson.Geometry
}

type adrBati struct {
	id         string
	insee_com  string
	id_adr     string
	id_bat     string
	type_lien  string
	nb_adr     int
	nb_bati    int
	origin_bat string
	type_bat   string
	surf_bat   float32
	haut_bat   int
	z_min_bat  float32
	z_max_bat  float32
	geometry   *geojson.Geometry
}

type batiParc struct {
	id         string
	insee_com  string
	id_bat     string
	origin_bat string
	surf_bat   float32
	haut_bat   int
	z_min_bat  float32
	z_max_bat  float32
	type_bat   string
	idu        string
	geometry   *geojson.Geometry
}

func getAdrParc(db *sql.DB, key string, value string) (*geojson.FeatureCollection, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT id, insee_com,id_adr,idu,type_lien,nb_adr,nb_parc,ST_AsGeoJSON(geom) FROM %s.adr_parc WHERE %s=$1", os.Getenv("APP_DB_SCHEMA"), key), value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	fc := geojson.NewFeatureCollection()

	for rows.Next() {
		var a adrParc
		if err := rows.Scan(&a.id, &a.insee_com, &a.id_adr, &a.idu, &a.type_lien, &a.nb_adr, &a.nb_parc, &a.geometry); err != nil {
			return nil, err
		}
		f := geojson.NewFeature(a.geometry)
		f.SetProperty("id", a.id)
		f.SetProperty("insee_com", a.insee_com)
		f.SetProperty("id_adr", a.id_adr)
		f.SetProperty("idu", a.idu)
		f.SetProperty("type_lien", a.type_lien)
		f.SetProperty("nb_adr", a.nb_adr)
		f.SetProperty("nb_parc", a.nb_parc)
		fc.AddFeature(f)
	}

	return fc, nil
}

func getAdrBati(db *sql.DB, key string, value string) (*geojson.FeatureCollection, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT id,insee_com,id_adr,id_bat,type_lien,nb_adr,nb_bati,origin_bat,type_bat,surf_bat,haut_bat,z_min_bat,z_max_bat,ST_AsGeoJSON(geom) FROM %s.adr_bati WHERE %s=$1", os.Getenv("APP_DB_SCHEMA"), key), value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	fc := geojson.NewFeatureCollection()

	for rows.Next() {
		var a adrBati
		if err := rows.Scan(&a.id, &a.insee_com, &a.id_adr, &a.id_bat, &a.type_lien, &a.nb_adr, &a.nb_bati, &a.origin_bat, &a.type_bat, &a.surf_bat, &a.haut_bat, &a.z_min_bat, &a.z_max_bat, &a.geometry); err != nil {
			return nil, err
		}
		f := geojson.NewFeature(a.geometry)
		f.SetProperty("id", a.id)
		f.SetProperty("insee_com", a.insee_com)
		f.SetProperty("id_adr", a.id_adr)
		f.SetProperty("id_bat", a.id_bat)
		f.SetProperty("type_lien", a.type_lien)
		f.SetProperty("nb_adr", a.nb_adr)
		f.SetProperty("nb_bati", a.nb_bati)
		f.SetProperty("origin_bat", a.origin_bat)
		f.SetProperty("type_bat", a.type_bat)
		f.SetProperty("surf_bat", a.surf_bat)
		f.SetProperty("haut_bat", a.haut_bat)
		f.SetProperty("z_min_bat", a.z_min_bat)
		f.SetProperty("z_max_bat", a.z_max_bat)
		fc.AddFeature(f)
	}

	return fc, nil
}

func getBatiParc(db *sql.DB, key string, value string) (*geojson.FeatureCollection, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT id,insee_com,id_bat,origin_bat,surf_bat,haut_bat,z_min_bat,z_max_bat,type_bat,idu,ST_AsGeoJSON(geom) FROM %s.bati_parc WHERE %s=$1", os.Getenv("APP_DB_SCHEMA"), key), value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	fc := geojson.NewFeatureCollection()

	for rows.Next() {
		var a batiParc
		if err := rows.Scan(&a.id, &a.insee_com, &a.id_bat, &a.origin_bat, &a.surf_bat, &a.haut_bat, &a.z_min_bat, &a.z_max_bat, &a.type_bat, &a.idu, &a.geometry); err != nil {
			return nil, err
		}
		f := geojson.NewFeature(a.geometry)
		f.SetProperty("id", a.id)
		f.SetProperty("insee_com", a.insee_com)
		f.SetProperty("id_bat", a.id_bat)
		f.SetProperty("origin_bat", a.origin_bat)
		f.SetProperty("surf_bat", a.surf_bat)
		f.SetProperty("haut_bat", a.haut_bat)
		f.SetProperty("z_min_bat", a.z_min_bat)
		f.SetProperty("z_max_bat", a.z_max_bat)
		f.SetProperty("type_bat", a.type_bat)
		f.SetProperty("idu", a.idu)
		fc.AddFeature(f)
	}

	return fc, nil
}
