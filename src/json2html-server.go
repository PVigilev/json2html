package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

// A simple structure for keeping setting for server startup
var StartupServerModel struct {
	template *template.Template // parsed html-template
	port     uint               // parsed porton which server will listen
	serveMux *http.ServeMux     // a pointer to a serve multiplexer with handlers for all the supported endpoints
}

const templateParam = "t"
const templateDescription = "A path to json-to-html template"
const portParam = "p"
const portDescription = "A port number on which server will listen"
const defaultTemplateName = "threat.html.tmpl"
const defaultTemplatePath = "templates/threat.html.tmpl"
const defaultPortNumber = 8080
const rootPageAddress = "static/index.html"

// Initialize module by setting program arguments parser
// and initialize ServeMux with hadlers for endpoints
func init() {
	var templateFilename string
	flag.StringVar(&templateFilename, templateParam, defaultTemplatePath, templateDescription)
	flag.UintVar(&StartupServerModel.port, portParam, defaultPortNumber, portDescription)
	template, err := template.New(defaultTemplateName).ParseFiles(templateFilename)
	if err != nil {
		log.Fatal("Template parsing at ", templateFilename, " failed with error: ", err)
	} else {
		log.Println("Template at ", templateFilename, " parsed")
	}
	StartupServerModel.template = template

	StartupServerModel.serveMux = http.NewServeMux()
	StartupServerModel.serveMux.HandleFunc("POST /render", WriteRenderedHtml)
	StartupServerModel.serveMux.HandleFunc("GET /", GetRootHandler)
}

func main() {
	flag.Parse()
	s := &http.Server{
		Addr:           ":" + fmt.Sprint(StartupServerModel.port),
		Handler:        StartupServerModel.serveMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server will start on port", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Server starting failed with error: ", err)
	}
}
