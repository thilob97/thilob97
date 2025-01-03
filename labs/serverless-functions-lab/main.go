package serverlessfunctionslab

import "net/http"

func loadPluginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loading plugins..."))
}

func main() {
	http.HandleFunc("/load-plugins", loadPluginHandler)
}
