# Understanding 'Main.go'

#### 1. package main

This is used to tell the compiler where the



It imports various packages that are required for the program.

1. It defines a type `ServerData` as a map of strings.
2. It declares and initializes variables for file paths (`excelFile`, `csvFile`) and the host address of etcd (`etcdHost`).
3. It defines a function `convertExcelToCSV` that converts an Excel file to a CSV file.
   * It opens the Excel file specified by `excelFile`.
   * It creates a new CSV file specified by `csvFile`.
   * It writes the data from the Excel file to the CSV file.
4. It defines a function `uploadToEtcd` that uploads data from the CSV file to etcd.
   * It connects to the etcd server using the specified `etcdHost`.
   * It opens the CSV file specified by `csvFile`.
   * It reads the CSV file and parses the records.
   * It iterates over the records and uploads the data to etcd.
   * It sets key-value pairs in etcd for each data field.
   * It sets a key-value pair for the server data in etcd.
5. It defines a function `getServerData` that handles HTTP requests to retrieve server data from etcd.
   * It extracts the server type and IP from the URL path.
   * It connects to the etcd server using the specified `etcdHost`.
   * It constructs the etcd key for the server data.
   * It retrieves the server data from etcd.
   * It checks if the key exists and extracts the server data value.
   * It extracts the specific field value from the server data.
   * It writes the field value as the response.
6. It defines the `main` function.
   * It converts the Excel file to CSV using the `convertExcelToCSV` function.
   * It parses command-line flags.
   * It uploads the CSV data to etcd using the `uploadToEtcd` function.
   * It starts an HTTP server to handle API requests and calls the `getServerData` function to retrieve server data from etcd.
   * It listens for incoming requests on port 8080 and logs any errors that occur.

In summary, the code converts an Excel file to a CSV file, uploads the CSV data to an etcd server, and provides an API to retrieve server data from etcd based on server type, IP, and specific field.

In this code, the `main` package is being imported along with several other packages necessary for the program. These packages provide various functionalities such as handling HTTP requests, working with CSV and Excel files, and interacting with the etcd distributed key-value store.

The code then declares and initializes several variables:

* `excelFile` represents the path to the Excel file ("/home/user/my.db/etcd-inventory/etcd.xlsx").
* `csvFile` represents the path to the CSV file ("/home/user/my.db/etcd-inventory/myetcd.csv").
* `etcdHost` represents the host address of the etcd server ("localhost:2379").

These variables specify the file paths for the Excel and CSV files and the address of the etcd server that the program will interact with.

The purpose of this code is to perform operations such as converting an Excel file to a CSV file, uploading data from the CSV file to an etcd server, and providing an API to retrieve server data from etcd.

The remaining code defines the functions and logic to carry out these tasks. It includes functions like `convertExcelToCSV`, `uploadToEtcd`, `getServerData`, and the main function.

Overall, this code sets up the necessary components and configuration to handle Excel and CSV files, communicate with an etcd server, and provide an API for retrieving server data.

In this code, the `ServerData` type is defined as a map with string keys and string values. It represents a dictionary-like structure where each key corresponds to a specific field and its associated value.

The `convertExcelToCSV` function takes two input parameters: `excelFile` and `csvFile`, which are the paths to an Excel file and a CSV file, respectively.

The function performs the following steps:

1. It opens the Excel file specified by `excelFile` using the `xlsx.OpenFile` function from the `tealeg/xlsx` package.
2. It creates the CSV file specified by `csvFile` using the `os.Create` function.
3. It defers the closing of the CSV file to ensure it is closed at the end of the function.
4. It creates a new CSV writer using the `csv.NewWriter` function, which allows writing data to the CSV file.
5. It defers the flushing of the CSV writer to ensure all data is written to the file before it is closed.
6. It iterates over each sheet in the Excel file and then iterates over each row in each sheet.
7. For each row, it creates an empty slice called `rowData` to store the cell values as strings.
8. It iterates over each cell in the row and retrieves the cell's value as a string using the `cell.String()` method.
9. It appends the cell value to the `rowData` slice.
10. After processing all the cells in the row, it checks if the row is empty by iterating over the `rowData` slice and checking if any field is not empty.
11. If the row is not empty, it writes the `rowData` slice as a row in the CSV file using the `writer.Write` method.
12. The process repeats for each row in each sheet, resulting in all non-empty rows being written to the CSV file.

In summary, the `convertExcelToCSV` function opens an Excel file, creates a corresponding CSV file, and writes the non-empty rows from the Excel file to the CSV file.

The `uploadToEtcd` function performs the following tasks:

1. It connects to an etcd server by creating a new client using the `clientv3.New` function from the `go.etcd.io/etcd/client/v3` package. It uses the `etcdHost` variable to specify the etcd server's endpoint.
2. If there is an error connecting to etcd, it logs the error and exits the program.
3. It defers the closing of the etcd client to ensure it is closed at the end of the function.
4. It opens the CSV file specified by the `csvFile` variable using the `os.Open` function.
5. If there is an error opening the CSV file, it logs the error and exits the program.
6. It defers the closing of the CSV file to ensure it is closed at the end of the function.
7. It creates a new CSV reader using the `csv.NewReader` function to read the contents of the CSV file.
8. It reads all the records from the CSV file using the `reader.ReadAll` method and stores them in the `records` variable.
9. If there is an error reading the CSV file, it logs the error and exits the program.
10. It retrieves the headers (the first record) from the CSV file and assigns them to the `headers` variable.
11. It iterates over each record in `records[1:]` (excluding the header row) to process each server's data.
12. For each record, it extracts the server IP and server type from the record.
13. It creates an empty `ServerData` map called `serverData`.
14. It iterates over the headers starting from index 2 (skipping the server IP and server type) and assigns the corresponding value from the record to the `serverData` map using the header as the key.
15. It iterates over each key-value pair in the `serverData` map.
16. For each key-value pair, it constructs an etcd key based on the server type, server IP, and header.
17. It retrieves the corresponding value from the `serverData` map.
18. It uses the etcd client to put the key-value pair into etcd using the `etcdClient.Put` method.
19. If there is an error putting the key-value pair, it logs the error.
20. It constructs an etcd key for storing the server data as a JSON string.
21. It marshals the `serverData` map into JSON using the `json.Marshal` function.
22. If there is an error marshaling the data, it logs the error and continues to the next record.
23. It uses the etcd client to put the server data key-value pair into etcd.
24. If there is an error putting the server data, it logs the error.
25. Finally, it logs a success message indicating that the server details have been added to etcd successfully.

In summary, the `uploadToEtcd` function connects to an etcd server, reads data from a CSV file, and uploads that data to etcd by creating key-value pairs for each field and storing the server data as a JSON string.

The `getServerData` function handles an HTTP request and provides server data based on the server type, IP, and a specific field.

Here's a breakdown of what the function does:

1. It extracts the server type and IP from the URL path by splitting the path using the forward slash ("/") as a delimiter. The server type is assigned to the `serverType` variable, and the server IP is assigned to the `serverIP` variable.
2. It connects to the etcd server by creating a new etcd client using the `clientv3.New` function from the `go.etcd.io/etcd/client/v3` package. It uses the `etcdHost` variable to specify the etcd server's endpoint.
3. If there is an error connecting to etcd, it logs the error, returns an HTTP 500 Internal Server Error status, and sends an error response to the client.
4. It defers the closing of the etcd client to ensure it is closed at the end of the function.
5. It constructs the etcd key for retrieving the server data by formatting a string using the server type and IP.
6. It retrieves the server data from etcd using the etcd client's `Get` method, passing the etcd key as an argument.
7. If there is an error getting the server data from etcd, it logs the error, returns an HTTP 500 Internal Server Error status, and sends an error response to the client.
8. It checks if the key exists in the etcd response. If the key does not exist, it returns an HTTP 404 Not Found status and sends a "Server data not found" error response to the client.
9. It extracts the server data value from the etcd response. The value is unmarshaled into a `ServerData` variable using the `json.Unmarshal` function.
10. If there is an error unmarshaling the server data, it logs the error, returns an HTTP 500 Internal Server Error status, and sends an error response to the client.
11. It extracts the specific field value from the server data based on the URL path. The field is extracted from the `parts` slice at index 4.
12. It checks if the field exists in the server data. If the field does not exist, it returns an HTTP 404 Not Found status and sends a "Field not found" error response to the client.
13. It sets the response header's content type to "text/plain".
14. It writes the field value as the response body to the `http.ResponseWriter` using the `Write` method.

In summary, the `getServerData` function extracts server type, IP, and a specific field from the URL path, connects to an etcd server, retrieves the server data based on the provided information, and sends the corresponding field value as the HTTP response.

The `main` function is the entry point of the program. It executes several tasks to convert an Excel file to CSV, upload CSV data to an etcd server, and start an API server to handle HTTP requests.

Here's a breakdown of what the `main` function does:

1. It calls the `convertExcelToCSV` function to convert the Excel file to CSV. It passes the `excelFile` and `csvFile` variables as arguments. After the conversion is completed, it logs a message indicating that the Excel file has been successfully converted to CSV.
2. It parses the command-line flags using the `flag.Parse` function. This allows the program to read and handle command-line arguments if any are provided.
3. It calls the `uploadToEtcd` function to upload the CSV data to the etcd server. This function reads the CSV file, processes the data, and stores it in the etcd server. If any errors occur during the upload process, they will be logged.
4. It logs a message indicating that the API server is starting.
5. It registers a request handler function, `getServerData`, to handle requests with the URL path prefix "/servers/". This function retrieves server data from the etcd server based on the server type, IP, and specific field.
6. It starts the API server using the `http.ListenAndServe` function, specifying that it should listen on port 8080. If there is an error starting the server, it logs the error and exits the program with a fatal error message.

In summary, the `main` function orchestrates the process of converting an Excel file to CSV, uploading the CSV data to an etcd server, and starting an API server to handle requests for server data.
