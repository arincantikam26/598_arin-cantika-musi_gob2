package controllers

import (
	"html/template"
	"net/http"

	"path"
	"reload/helpers"
)

func DataController(w http.ResponseWriter, r *http.Request) {
	var FilePath = path.Join("frontend", "index.html")
	var temp, err = template.ParseFiles(FilePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	getData := helpers.GetData()
	numWater := getData.Status.Water
	numWind := getData.Status.Wind

	waterUnit := helpers.WaterUnit(numWater)
	windUnit := helpers.WindUnit(numWind)
	waterStatus := helpers.WaterStatus(numWater)
	windStatus := helpers.WindStatus(numWind)

	var data = map[string]interface{}{
		"waterUnit":   waterUnit,
		"windUnit":    windUnit,
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
