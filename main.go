package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	println("hello world")

	// Kuo Jet: 8990F9
	// Elon Jet: A835AF, A2AE0A, A64304

	// Random jet: AA8A8A
	// https://globe.adsb.one/
	// curl https://api.adsb.one/v2/hex/AA8A8A

	jet := "AA8A8A"

	resp, _ := http.Get(fmt.Sprintf("https://api.adsb.one/v2/hex/%s", jet))

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Printf(string(body))
}
