# JSON to HTML service

## Usage

At startup the following arguments are expected:
- `-p` port on which server will listen (default is 8080)
- `-t` a path to an html-template which will be parsed on startup

A template should work with the following structure:

```go
type Threat struct {
	ThreatName    string    `json:"threatName"`
	Category      string    `json:"category"`
	Size          uint      `json:"size"`
	DetectionDate string    `json:"detectionDate"`
	Variants      []struct {
        Name      string    `json:"name"`
        DateAdded string    `json:"dateAdded"`
    }
}
```

## Endpoints

### GET /

Return a static page at `static/index.html` folder.

### POST /render

Received JSON is parsed into object of type `Threat` and temnplate is executed over this object. In case of missing  some key, it's value is `0` for numeric type, empty string for a string type and empty-slice in case if slice-type. In case of empty-request or malformed one the status code 400 is set.


## Architecture

The whole server is implemented as a simple go-module. All the code is kept in `src/` directory.

### render.go
This file defines `ThreadDTO` type for representing a threat with its variants along with `escape` method which is used for escaping special characters from the input. Also, it defines a `WriteRenderedHtml` function which is used as a handler for processing a `POST /render` request.

### root.go
Defines a `GetRootHandler` function which is used as a handler for `GET /` request and only returns content of the `static/index.html` file.

### json2html-server.go

Here module initialization is happening.

The `init` function sets up an argument parser, initializes a serve-multiplexer with request-handlers.

The `main` function starts the server on the provided port with a parsed template.


## How to build and run

The simpliest way to build is to run `go build -C ./src -o ../build/` from the root repository.

To run the server, run executable from the repo-root. For the correct work it requires to have `static` directory with its content located in the working directory form which server starts.
