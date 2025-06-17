package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
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

// Sayfaları sun
func pageServe(w http.ResponseWriter, r *http.Request) {
	baseDir := "./server"

	var reqPath string = r.URL.Path
	if reqPath == "/" {
		http.ServeFile(w, r, "./homepage/index.html")
	}

	cleanPath := path.Clean(reqPath)
	fmt.Println("clean path: " + cleanPath)

	fullPath := baseDir + cleanPath + "/index.html"
	fmt.Println(fullPath)

	if !fileExists(fullPath) {
		http.Error(w, "Aradığınız sayfa bulunamadı.", 404)
		return
	}

	http.ServeFile(w, r, fullPath)
}

// Dosya varlığı kontrolü
func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}
