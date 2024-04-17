package main

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
	"strings"
)

// Data Transfer Object type for threat which is parsed from a provided JSON
type ThreatDTO struct {
	ThreatName    string `json:"threatName"`
	Category      string `json:"category"`
	Size          uint   `json:"size"`
	DetectionDate string `json:"detectionDate"`
	Variants      []ThreatVariantDTO
}
type ThreatVariantDTO struct {
	Name      string `json:"name"`
	DateAdded string `json:"dateAdded"`
}

const jsonInputFormName = "json_input"

// Escape all the strings in ThreatDTO
func (t *ThreatDTO) escape() (result ThreatDTO) {
	result.ThreatName = html.EscapeString(t.ThreatName)
	result.Category = html.EscapeString(t.Category)
	result.DetectionDate = html.EscapeString(t.DetectionDate)
	result.Variants = make([]ThreatVariantDTO, len(t.Variants))
	result.Size = t.Size

	for i, variant := range t.Variants {
		result.Variants[i].Name = html.EscapeString(variant.Name)
		result.Variants[i].DateAdded = html.EscapeString(variant.DateAdded)
	}
	return
}

// Handler function for processing a provided JSON and executing html-template on the provided data
func WriteRenderedHtml(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	request.ParseForm()

	// in case of not submitted form return status code 400
	if !request.Form.Has(jsonInputFormName) {
		log.Println("Error: form was not submitted")
		w.WriteHeader(400)
		return
	}
	bodyString := request.Form.Get(jsonInputFormName)

	var threat ThreatDTO
	// Parsing provided JSON into DTO
	decoder := json.NewDecoder(strings.NewReader(bodyString))
	decodeErr := decoder.Decode(&threat)
	if decodeErr != nil {
		log.Println("JSON parsing error: ", decodeErr)
		log.Println("Problematic JSON ", bodyString)
		w.WriteHeader(400)
		return
	}

	// escaping non-valid characters to avoid injecting html into json strings
	threat = threat.escape()
	log.Println("Threat parsed: ", threat)

	// applying the template on data
	errExecute := StartupServerModel.template.Execute(w, threat)
	if errExecute != nil {
		log.Print(errExecute)
		w.WriteHeader(400)
		return
	}

}
