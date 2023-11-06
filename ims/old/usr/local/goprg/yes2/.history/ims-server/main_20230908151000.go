package main

import (
	"context"       // For telling code when and how to perfomr certain task
	"encoding/csv"  // used for reading writing csv files
	"encoding/json" // only used for marshalling during upload
	"flag"
	"fmt"
	"io"
	"log"      //logging errors
	"net/http" //for interacting with etcd
	"os"       // mainly used for opening closing files
	"strings"  //for handling string types

	"github.com/tealeg/xlsx"             //for interacting with xlsx files, opening & reading them
	clientv3 "go.etcd.io/etcd/client/v3" //helps storing & fetching with the etcd api
)

var (
	// File paths
	excelFile = "./etcd.xlsx"
	csvFile   = "./myetcd.csv"
	etcdHost  = "localhost:2379"
)

type ServerData map[string]string

func convertExcelToCSV(excelFile, csvFile string) {
	// Open the Excel file
	xlFile, err := xlsx.OpenFile(excelFile)
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}

	// Create the CSV file
	file, err := os.Create(csvFile)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	// Write data to the CSV file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Iterate over sheets and rows in the Excel file
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var rowData []string
			for _, cell := range row.Cells {
				text := cell.String()
				rowData = append(rowData, text)
			}

			// Check if the row is empty
			isEmptyRow := true
			for _, field := range rowData {
				if field != "" {
					isEmptyRow = false
					break
				}
			}

			// Skip empty rows
			if !isEmptyRow {
				writer.Write(rowData)
			}
		}
	}
}

func uploadToEtcd() {
	// Connect to etcd
	log.Println("Entered into function")
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	// Read the CSV file
	file, err := os.Open(csvFile)
	log.Println("reading file")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Iterate over the records and upload to etcd
	headers := records[0]
	for _, record := range records[1:] {
		serverIP := record[0]
		serverType := record[1]
		serverData := make(ServerData)

		// Create server data dictionary
		for i := 2; i < len(headers); i++ {
			header := headers[i]
			value := record[i]
			serverData[header] = value
		}

		// Set key-value pairs in etcd for each data field
		for header, value := range serverData {
			etcdKey := strings.ToUpper(fmt.Sprintf("/%s/%s/%s", serverType, serverIP, header))
			etcdValue := value
			fmt.Println(etcdKey)
			fmt.Println(etcdValue)
			_, err := etcdClient.Put(context.Background(), etcdKey, etcdValue)
			if err != nil {
				log.Printf("Failed to upload key-value to etcd: %v", err)
			}
		}

		// Set key-value pair for server data
		etcdKeyData := strings.ToUpper(fmt.Sprintf("/%s/%s/data", serverType, serverIP))
		etcdValueData, err := json.Marshal(serverData)
		if err != nil {
			log.Printf("Failed to marshal server data: %v", err)
			continue
		}
		_, err = etcdClient.Put(context.Background(), etcdKeyData, string(etcdValueData))
		if err != nil {
			log.Printf("Failed to upload server data to etcd: %v", err)
		}
	}

	log.Println("Server details added to etcd successfully.")
}

func getServerData(w http.ResponseWriter, r *http.Request) {
	// Extract the etcd key from the URL path
	etcdKeyData := r.URL.Path[len("/"):]
	if r.Method == "GET" {
		// Check if the key is empty or if it starts with "/"
		// if etcdKeyData == "" || strings.HasPrefix(etcdKeyData, "/") {
		// 	listAll(w, r)
		// } else {
		// 	getSpecificKey(w, r)
		// }
		/////////////   anoop /////
		// Check if the key is empty or if it starts with "/"
		if etcdKeyData == "" || strings.HasPrefix(etcdKeyData, "/") {
			listAll(w, r)
		} else {
			getSpecificKeyAnoop(w, r)
		}
	}

	////////////////////
	if r.Method == "POST" {
		postSpecificKeyAnoop(w, r)
	}
	////////////////////

	if r.Method == "DELETE" {
		deleteSpecificKeyAnoop(w, r)
	}
	/////////////
}

// listall function is the only part of this code created by Yash Anand
func listAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("response %v", r.URL.Path)
	ctx := context.TODO()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Printf("Failed to connect to etcd: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer etcdClient.Close()

	// Get all keys with values

	//response, err := etcdClient.Get(ctx, "/servers/", clientv3.WithPrefix())
	response, err := etcdClient.Get(ctx, "/", clientv3.WithPrefix())
	if err != nil {
		log.Printf("Failed to retrieve keys from etcd: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	for _, kv := range response.Kvs {
		fmt.Fprintf(w, "Key: %s\n", string(kv.Key))
		fmt.Fprintf(w, "Value: %s\n", string(kv.Value))
		fmt.Fprintf(w, "----------------------------\n")
	}
}

// ////////// Anoop  ////////////////////////////
func getSpecificKeyAnoop(w http.ResponseWriter, r *http.Request) {
	// Extract the server type and IP from the URL path
	log.Printf("response %v", r.URL.Path)
	//parts := strings.Split(r.URL.Path, "/")
	//serverType := parts[2]
	//serverIP := parts[3]

	// Connect to etcd
	ctx := context.TODO()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Printf("Failed to connect to etcd: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer etcdClient.Close()

	// Construct the etcd key for the server data
	etcdKeyData := fmt.Sprintf(r.URL.Path)

	response, err := etcdClient.Get(ctx, etcdKeyData, clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "text/plain")
	for _, kv := range response.Kvs {
		fmt.Fprintf(w, "Key: %s\n", string(kv.Key))
		fmt.Fprintf(w, "Value: %s\n", string(kv.Value))
		fmt.Fprintf(w, "----------------------------\n")
	}
}

// ///////////////////////////////////////////////
func postSpecificKeyAnoop(w http.ResponseWriter, r *http.Request) {
	// Extract the server type and IP from the URL path
	log.Printf("response %v", r.URL.Path)
	// Read response body
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response:", string(responseBody))

	type Keydata struct {
		EtcdKey   string
		EtcdValue string
	}
	var keydata1 Keydata

	err1 := json.Unmarshal(responseBody, &keydata1)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("Struct is:", keydata1)
	fmt.Printf("key:%s | value:%s\n", keydata1.EtcdKey, keydata1.EtcdValue)

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()
	_, err = etcdClient.Put(context.Background(), keydata1.EtcdKey, keydata1.EtcdValue)
	if err != nil {
		log.Printf("Failed to upload server data to etcd: %v", err)
	}

}

// /////////////
func deleteSpecificKeyAnoop(w http.ResponseWriter, r *http.Request) {
	// Extract the server type and IP from the URL path
	log.Printf("response %v", r.URL.Path)
	// Read response body
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response:", string(responseBody))

	type Keydata struct {
		EtcdKey string
	}
	var keydata1 Keydata

	err1 := json.Unmarshal(responseBody, &keydata1)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("Struct is:", keydata1)
	fmt.Printf("key:%s \n", keydata1.EtcdKey)

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	_, err = etcdClient.Delete(context.Background(), keydata1.EtcdKey)
	if err != nil {
		log.Printf("Failed to delete key frometcd: %v", err)
	}

}

// ////////////////////////////////
func getSpecificKey(w http.ResponseWriter, r *http.Request) {

	// Extract the server type and IP from the URL path
	log.Printf("response %v", r.URL.Path)
	//parts := strings.Split(r.URL.Path, "/")
	//serverType := parts[2]
	//serverIP := parts[3]

	// Connect to etcd
	ctx := context.TODO()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Printf("Failed to connect to etcd: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer etcdClient.Close()

	// Construct the etcd key for the server data
	etcdKeyData := fmt.Sprintf(r.URL.Path)
	//response, err := etcdClient.Get(ctx, etcdKeyData, clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))

	//etcdKeyData := "/servers/VM/10.249.221.21/RAM"
	//revisions := []int64{1066, 1065, 1064, 1063, 1062}
	//revisions := make(map[int64]string)

	var revisions int

	response, err := etcdClient.Get(ctx, etcdKeyData, clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))
	//response, err := etcdClient.Get(ctx, etcdKeyData, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByModRevision, clientv3.SortDescend))
	//print(response)
	log.Printf("response %v", response)
	for _, kv := range response.Kvs {
		revisions = int(kv.ModRevision)
		log.Printf("revisions %v", revisions)
	}
	log.Printf("revision response: %v", revisions)

	var revisionslist []int64

	for i := revisions; i >= revisions-5; i-- {
		revisionslist = append(revisionslist, int64(i))
	}

	values, err := getRevisionValues(etcdClient, etcdKeyData, revisionslist)
	w.Header().Set("Content-Type", "text/plain")

	for _, value := range values {
		fmt.Fprintf(w, "Value: %s\n", value)
		fmt.Fprintf(w, "---------------------------\n")
	}
}

func getRevisionValues(client *clientv3.Client, key string, revisions []int64) ([]string, error) {
	ctx := context.TODO()

	//values := make(map[int64]string)
	var values []string

	for _, rev := range revisions {
		response, err := client.Get(ctx, key, clientv3.WithRev(rev))
		if err != nil {
			return nil, err
		}

		if len(response.Kvs) > 0 {
			value := string(response.Kvs[0].Value)
			//values[rev] = value
			values = append(values, value)
			print(value)

		} else {
			print("Value not found")
		}
	}

	return values, nil
}

func main() {

	// for removing all data remove directory /var/lib/etcd/default   and restart etcd service
	// Convert Excel to CSV
	// for deleting all old records removing directory /var/lib/etcd/default and restart service etcd

	convertExcelToCSV(excelFile, csvFile)
	log.Println("Excel file converted to CSV successfully.")

	// Parse command-line flags
	flag.Parse()

	// Upload CSV data to etcd
	uploadToEtcd()

	// Start API server
	log.Println("Starting API server...")
	http.HandleFunc("/", getServerData)
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}
