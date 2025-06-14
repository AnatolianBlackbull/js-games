package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Sitenin çalışacağı port loooooooooooooo
	var port int = 8080

	mux := http.NewServeMux()

	// Oyun isteği geldiğinde istenen oyunu, gerekli dosyalarla birlikte ver
	gameFS := http.FileServer(http.Dir("./games"))
	mux.Handle("/games/", http.StripPrefix("/games/", gameFS))

	// Siteyi sun
	mux.HandleFunc("/", mainPageServe)

	// Sunucuyu çalıştır
	fmt.Printf("Sunucu %d portunda çalışıyor.\n", port)
	portString := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portString, mux)

}

func mainPageServe(w http.ResponseWriter, r *http.Request) {

	http.ServeFile("/homepage/index.html")

}

// todo: Kullanıcı giriş çıkışı
func login() {
	fmt.Println("31")
}
