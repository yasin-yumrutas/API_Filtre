package main

import (
	"encoding/json" // JSON işleme kütüphanesi
	"log"           // basit loglama kütüphanesi
	"net/http"      // HTTP işleme kütüphanesi
	"strconv"       // string-to-int dönüşümü yapmak için kütüphane
)

// Product adlı bir struct (yapı) oluşturuluyor. Bu, ürünlerin özelliklerini içerir.
type Product struct {
	ID    int    `json:"id"` // json tag'leri, JSON kodlama/çözme işlemi sırasında kullanılır
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Bazı örnek ürünlerin bir dilim (slice) olarak tanımlanması.
var products = []Product{
	{ID: 1, Name: "Product 1", Price: 10},
	{ID: 2, Name: "Product 2", Price: 20},
	{ID: 3, Name: "Product 3", Price: 30},
	{ID: 4, Name: "Product 4", Price: 40},
}

// Programın ana işlevi
func main() {
	// "/api/products" yolunu dinleyen bir HTTP sunucusu oluşturuluyor.
	// İstek geldiğinde, "getProducts" işlevi çağrılacak.
	http.HandleFunc("/api/products", getProducts)

	// Sunucu, :8000 adresindeki istekleri dinlemeye başlıyor.
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// "getProducts" işlevi, HTTP isteğini ve yanıtı alır.
func getProducts(w http.ResponseWriter, r *http.Request) {
	// URL'den "min_price" ve "max_price" parametrelerini alır.
	// Atoi fonksiyonu, stringi integer'a dönüştürür.
	minPrice, _ := strconv.Atoi(r.URL.Query().Get("min_price"))
	maxPrice, _ := strconv.Atoi(r.URL.Query().Get("max_price"))

	// Ürünlerde gezinir ve fiyat aralığına uygun olanları bir diziye ekler.
	var filteredProducts []Product
	for _, p := range products {
		if p.Price >= minPrice && p.Price <= maxPrice {
			filteredProducts = append(filteredProducts, p)
		}
	}

	// Bulunan ürünlerin JSON olarak kodlanması ve yanıt olarak gönderilmesi.
	json.NewEncoder(w).Encode(filteredProducts)
}
