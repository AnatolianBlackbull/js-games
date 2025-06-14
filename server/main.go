package main

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	// Sitenin çalışacağı port
	var port int = 8080

	mux := http.NewServeMux()

	// Oyun isteği geldiğinde istenen oyunu, gerekli dosyalarla birlikte ver
	gameFS := http.FileServer(http.Dir("./games"))
	mux.Handle("/games/", http.StripPrefix("/games/", gameFS))

	// Siteyi sun
	mux.HandleFunc("/", pageServe)

	// Sunucuyu çalıştır
	fmt.Printf("Sunucu %d portunda çalışıyor.\n", port)
	portString := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portString, mux)

}

func pageServe(w http.ResponseWriter, r *http.Request) {
	baseDir := "./server"

	var reqPath string = r.URL.Path
	if reqPath == "/" {
		reqPath = "/homepage"
	}

	cleanPath := path.Clean(reqPath)
	fullPath := filepath.Join(baseDir, cleanPath)

	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		http.Error(w, "Sunucu içi hata", 500)
		return
	}

	absFullPath, err := filepath.Abs(fullPath)
	if err != nil {
		http.Error(w, "Aradığınız site bulunamadı.", 500)
		return
	}

	if !strings.HasPrefix(absFullPath, absBaseDir) {
		http.Error(w, "İzinsiz erişim talebi", http.StatusForbidden)
		return
	}

	pagePath := filepath.Join(absFullPath, "index.html")

	finalPath, err := filepath.Abs(pagePath)
	if err != nil {
		http.Error(w, "Aradığınız sayfa bulunamadı.", 404)
		return
	}

	http.ServeFile(w, r, finalPath)
}
