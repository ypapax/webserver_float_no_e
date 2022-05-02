package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FloatZeros float64

type F struct {
	Float FloatZeros
}

func (f FloatZeros) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.20f", f)), nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var result = F{Float: 0.0000000000001}
		b, _ := json.Marshal(result)
		log.Printf("result: %+v\n", string(b))
		//fmt.Fprintf(w, b)
		w.Write(b)

	})
	addr := ":8081"
	log.Printf("listening to %+v\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Printf("error: %+v\n", err)
	}
}