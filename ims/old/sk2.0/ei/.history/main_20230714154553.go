package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tealeg/xlsx"
	"go.etcd.io/bbolt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	// File paths
	excelFile = "/home/user/sk/etcd-inventory/etcd.xlsx"
	csvFile   = "/home/user/sk/etcd-inventory/myetcd.csv"
	etcddb    = "/home/user/sk/etcd-inventory/default.etcd/member/snap/db"
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
			//fmt.Println(value)
			serverData[header] = value
		}

		// Set key-value pairs in etcd for each data field
		for header, value := range serverData {
			etcdKey := fmt.Sprintf("/servers/%s/%s/%s", serverType, serverIP, header)
			etcdValue := value
			fmt.Println(etcdKey)
			fmt.Println(etcdValue)
			_, err := etcdClient.Put(context.Background(), etcdKey, etcdValue)
			if err != nil {
				log.Printf("Failed to upload key-value to etcd: %v", err)
			}
		}

		// Set key-value pair for server data
		etcdKeyData := fmt.Sprintf("/servers/%s/%s/data", serverType, serverIP)
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

	for i := revisions; i >= revisions-4; i-- {
		revisionslist = append(revisionslist, int64(i))
	}

	values, err := getRevisionValues(etcdClient, etcdKeyData, revisionslist)
	w.Header().Set("Content-Type", "text/plain")

	for _, value := range values {
		fmt.Fprintf(w, "Value: %s\n", value)
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

func history() {
	db, err := bbolt.Open(etcddb, 0666, nil)
	if err != nil {
		log.Fatalf("Error opening the database: %v", err)
	}
	defer db.Close()

	err = db.View(func(tx *bbolt.Tx) error {

		bucketName := "key"
		bucket := tx.Bucket([]byte(bucketName))

		fmt.Printf("Current Bucket: %s\n\nAll existing buckets:\n", bucketName)

		err = tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			fmt.Println(string(name))
			return nil
		})

		fmt.Printf("\nKey-Value Data From The Bucket Is As Folows\n")

		err = bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("Key: %s, Value: %s\n", k, v)
			return nil
		})

		return nil
	})

	if err != nil {
		log.Fatalf("Error reading the database: %v", err)
	}
}

func main() {
	// Convert Excel to CSV
	convertExcelToCSV(excelFile, csvFile)
	log.Println("Excel file converted to CSV successfully.")

	// Parse command-line flags
	flag.Parse()

	// Upload CSV data to etcd
	uploadToEtcd()

	// Start API server
	log.Println("Starting API server...")
	http.HandleFunc("/servers/", getServerData)
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

}
