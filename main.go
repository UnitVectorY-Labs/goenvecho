package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime/debug"
)

// Version is the application version, injected at build time via ldflags
var Version = "dev"

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	env := make(map[string]string)
	for _, e := range os.Environ() {
		pair := splitEnvVar(e)
		env[pair[0]] = pair[1]
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-App-Version", Version)

	encoder := json.NewEncoder(w)
	// Check if PRETTY_PRINT is set to "true"
	if os.Getenv("PRETTY_PRINT") == "true" {
		encoder.SetIndent("", "  ")
	}
	encoder.Encode(env)
}

func splitEnvVar(envVar string) []string {
	for i := 0; i < len(envVar); i++ {
		if envVar[i] == '=' {
			return []string{envVar[:i], envVar[i+1:]}
		}
	}
	return []string{envVar, ""}
}

func main() {
	// Set the build version from the build info if not set by the build system
	if Version == "dev" || Version == "" {
		if bi, ok := debug.ReadBuildInfo(); ok {
			if bi.Main.Version != "" && bi.Main.Version != "(devel)" {
				Version = bi.Main.Version
			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	http.HandleFunc("GET /", handler)
	http.ListenAndServe(":"+port, nil)
}
