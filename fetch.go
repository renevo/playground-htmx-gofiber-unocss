package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch(endpoint string) []byte {
	fmt.Fprintf(os.Stderr, "GET: %q\n", endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch error %q: %v\n", endpoint, err)
		return nil
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Fetch Bad Response %q: %s\n", endpoint, resp.Status)
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read error %q: %v\n", endpoint, err)
		return nil
	}

	return data
}
