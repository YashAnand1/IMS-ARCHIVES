package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addressServer string
var OutPutFile *os.File

var (
	itldims = &cobra.Command{
		Use:   "itldims",
		Short: "Interact with the etcd API",
		Long:  "A command-line tool to interact with the etcd API and check connection",
		Run: func(cmd *cobra.Command, args []string) { // Extracted function
			//response, err := http.Get("http://localhost:8181/servers/")
			response, err := http.Get(addressServer + "/IMS")
			if err != nil {

				fmt.Printf("\n\n ------: Failed to connect to the etcd API at http://localhost:8181/servers/:-------\n\n\n\n\n")
			}
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				fmt.Println("Connected to API. Interaction with etcd can be done.")
			}
		},
	}
)

func readconf() (addressServer string, outputfile *os.File) {
	addressServer = ""
	file, err := os.Open("./itldmis.conf")
	if err != nil {
		log.Fatalf("Failed to open itldmis.conf: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	addressServer = scanner.Text()
	file.Close()

	/////////////////////////
	outputfile, err = os.Create("./output")
	if err != nil {
		log.Fatalf("Failed to open output: %v", err)
	}
	defer OutPutFile.Close()

	return
}

func init() {

	itldims.AddCommand(GET)
	itldims.AddCommand(delete)
	itldims.AddCommand(CREATE)
	itldims.AddCommand(update)
	/////////////////////////////////////
	addressServer, OutPutFile = readconf()

	var cmd1 string
	for i := 0; i < len(os.Args); i++ {
		cmd1 = cmd1 + " " + os.Args[i]
	}
	fmt.Printf("Command given:%s\n", cmd1)
	cmd1 = fmt.Sprintf("Command given:%s\n------------------\n", cmd1)
	OutPutFile.WriteString(cmd1)

	for i := 0; i < len(os.Args); i++ {
		os.Args[i] = strings.ToUpper(os.Args[i])
	}
}

func Execute() {

	if err := itldims.Execute(); err != nil {
		log.Fatal(err)
	}

}
