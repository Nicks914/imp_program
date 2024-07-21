package main

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"
)

// AjaxResponce represents the structure of an AJAX response.
type AjaxResponce struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// View is a global variable holding the template instance.
var View *template.Template

func main() {
	// Example usage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Message string `json:"message"`
		}{
			Message: "Hello, World!",
		}
		RenderTemplate(w, r, "index.html", data)
	})
	http.ListenAndServe(":8080", nil)
}

// RenderTemplate renders a template based on the HTTP request context.
func RenderTemplate(w http.ResponseWriter, r *http.Request, template string, data interface{}) {
	tmplData := make(map[string]interface{})

	if IsCurlApiRequest(r) || template == "" {
		jsonResponse, err := json.Marshal(data)
		if err != nil {
			Logger(err, "critical")
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Write(jsonResponse)
	} else {
		ajaxReq, ok := data.(AjaxResponce)
		if ok {
			AddFlash(ajaxReq.Status, ajaxReq.Message, w, r)
		}
		tmplData["data"] = data
		tmplData["flash"] = viewFlash(w, r)   // Assuming viewFlash is defined somewhere
		tmplData["session"] = fetchSession(r) // Assuming fetchSession is defined somewhere
		tmplData["appUrl"] = os.Getenv("APPURL")
		tmplData["supportEmail"] = os.Getenv("SUPPORTEMAIL")
		View.ExecuteTemplate(w, template, tmplData)
	}
}

// IsCurlApiRequest checks if the request is from a curl command or an API request.
func IsCurlApiRequest(r *http.Request) bool {
	// Implementation specific to your needs
	return false
}

// Logger logs errors with specified severity.
func Logger(err error, severity string) {
	// Implementation specific to your needs
}

// AddFlash adds flash messages to the response.
func AddFlash(status, message string, w http.ResponseWriter, r *http.Request) {
	// Implementation specific to your needs
}

// viewFlash retrieves flash messages for the view.
func viewFlash(w http.ResponseWriter, r *http.Request) interface{} {
	// Implementation specific to your needs
	return nil
}

// fetchSession retrieves session data.
func fetchSession(r *http.Request) interface{} {
	// Implementation specific to your needs
	return nil
}
