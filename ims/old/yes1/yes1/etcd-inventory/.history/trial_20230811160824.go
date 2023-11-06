package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	itldims = &cobra.Command{
		Use:   "itldims",
		Short: "For checking connectivity with ETCD API",
		Long:  "For checking connectivity - lets user know if connected or not",
		Run:   checkConnection,
	}
)

func checkConnection(cmd *cobra.Command, args []string) {
	{ // Extracted function
		response, err := http.Get("http://localhost:8181/servers/")
		if err != nil {
			log.Fatalf("Failed to connect to the etcd API.")
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			fmt.Println("Connected to API. Interaction with etcd can be done.")
		}
	}
}

func main() {
	if err := itldims.Execute(); err != nil {
		log.Fatal(err)
	}
}
