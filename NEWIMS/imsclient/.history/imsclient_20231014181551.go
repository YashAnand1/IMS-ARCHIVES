package main

//"golang.org/x/crypto/ssh/terminal"
//https://developer.ibm.com/articles/a-tour-of-the-kubernetes-source-code/
//ansible
// Varad Gupta
// 10:29
// https://developer.ibm.com/articles/a-tour-of-the-kubernetes-source-code/
// https://jvns.ca/blog/2017/06/04/learning-about-kubernetes/
// Varad Gupta
// 10:45
// https://www.reddit.com/r/kubernetes/comments/12y24c7/lets_read_the_kubernetes_source_code_video/?rdt=32771

// https://www.reddit.com/r/kubernetes/comments/12y24c7/lets_read_the_kubernetes_source_code_video/?rdt=32771
// https://developer.ibm.com/articles/a-tour-of-the-kubernetes-source-code/
//Harsh Choudhary
//9456014301

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/tealeg/xlsx"
	"golang.org/x/term"
)

type InputData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var selectedClient, loggedInUser string
var loginStatus int = 0

// //////////////////////
var addressServer string
var OutPutFile *os.File
var linkcells = map[int][]string{}
var cmdName = map[int][]string{}
var tableName = map[int][]string{}
var tableHeader = map[int][]string{}

// //////////////////////
func main() {
	// //////// temp ////

	// loggedInUser = "admin"
	// loginStatus = 2

	////////////////////
	readconf()
	for {
		cmd := getCommand()
		if strings.ToUpper(cmd[0]) == "EXIT" || strings.ToUpper(cmd[0]) == "CLOSE" || strings.ToUpper(cmd[0]) == "QUIT" || strings.ToUpper(cmd[0]) == "Q" {
			break
		}
		cmdProcess(cmd)
	}

	//////////////////////

}

// /////////////////
func cmdProcess(cmd []string) {
	var f int = 0

	////////////////////
	if strings.ToUpper(cmd[0]) == "CREATE" && strings.ToUpper(cmd[1]) == "--CLIENT" {
		if loginStatus < 2 {
			fmt.Printf("You are not Logged in as ADMIN\n")
			return
		}
		creatClient(cmd)
		f = 1
		fmt.Printf("\n\n")
	}
	///////////////////////
	if strings.ToUpper(cmd[0]) == "CREATE" && strings.ToUpper(cmd[1]) == "--KEY" {
		if loginStatus < 2 {
			fmt.Printf("You are not Logged in as ADMIN\n")
			return
		}
		creatKey(cmd)
		f = 1
		fmt.Printf("\n\n")
	}
	// //////////////////////////
	if strings.ToUpper(cmd[0]) == "DELETE" && strings.ToUpper(cmd[1]) == "--KEY" {
		if loginStatus < 2 {
			fmt.Printf("You are not Logged in as ADMIN\n")
			return
		}
		deleteKey(cmd)
		f = 1
		fmt.Printf("\n\n")
	}
	/////////////////////////
	if strings.ToUpper(cmd[0]) == "UPDATE" && strings.ToUpper(cmd[1]) == "--KEY" {
		if loginStatus < 2 {
			fmt.Printf("You are not Logged in as ADMIN\n")
			return
		}
		updateKey(cmd)
		f = 1
		fmt.Printf("\n\n")
	}

	/////////////////////
	// if strings.ToUpper(cmd[0]) == "CREATE" {
	// 	if loginStatus < 2 {
	// 		fmt.Printf("You are not Logged in as ADMIN\n")
	// 		return
	// 	}
	// 	create(cmd)
	// 	f = 1
	// 	fmt.Printf("\n\n")
	// }

	/////////////////////

	if strings.ToUpper(cmd[0]) == "LIST" {
		if loginStatus == 0 {
			fmt.Printf("You are not Logged in\n")
			return
		}
		listColumn(cmd)
		f = 1
		fmt.Printf("\n\n")
	}
	////////////////////////////
	if strings.ToUpper(cmd[0]) == "GET" {
		if loginStatus == 0 {
			fmt.Printf("You are not Logged in\n")
			return
		}
		getCmdResult(cmd)
		f = 1
		fmt.Printf("\n\n")
	}

	/////////////////////
	if strings.ToUpper(cmd[0]) == "SELECT" && strings.ToUpper(cmd[1]) == "CLIENT" {
		if loginStatus == 0 {
			fmt.Printf("You are not Logged in\n")
			return
		}
		selectClient(cmd)
		f = 1
		fmt.Printf("\n\n")
	}

	/////////////////////
	if strings.ToUpper(cmd[0]) == "LOGIN" {
		login(cmd)
		f = 1
		fmt.Printf("\n\n")
	}

	//////////////////changePasswd /////////
	if strings.ToUpper(cmd[0]) == "PASSWD" {
		if loginStatus == 0 {
			fmt.Printf("You are not Logged in\n")
			return
		}
		changePasswd(cmd)
		f = 1
		fmt.Printf("\n\n")
	}
	// ///////////// Upload xls file ////
	if strings.ToUpper(cmd[0]) == "UPLOAD" {
		if loginStatus < 2 {
			fmt.Printf("You are not Logged in as ADMIN\n")
			return
		}
		uploadFile(cmd)
		f = 1
		fmt.Printf("\n\n")
	}

	//////////////////////
	if f == 0 {
		fmt.Printf("Not a valid command! (Type \"exit\" and enter for close)\n\n")
	}
}

//////////////

// //////////////////////////////////
func getCommand() []string {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a command:")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	line = strings.TrimSpace(line)
	cmd := strings.Split(line, " ")
	//////////////////////
	return cmd
}

// /////////////////
func creatClient(cmd []string) {
	t := strings.ToUpper(cmd[1])
	key, value := FindKeyValue(t[2:], strings.ToUpper(cmd[2]))
	if value == strings.ToUpper(cmd[2]) {
		fmt.Printf("Client: %s already exists, No need to create", value)
	}

	if key == "" {
		fmt.Printf("\n Going to create Client: %s \n", value)
		var key = "/" + strings.ToUpper(cmd[2]) + "/1/" + "CLIENT"
		k, v := CreateKeyValue(key, strings.ToUpper(cmd[2]))
		fmt.Printf("Key: %s  | Value: %s\n", k, v)
		fmt.Printf("Client %s Created", v)
	}
}

// /////////////////////
func creatKey(cmd []string) {
	//t := strings.ToUpper(cmd[1])

	key, _ := FindKeyValue(strings.ToUpper(cmd[2]), cmd[3])
	if key == strings.ToUpper(cmd[2]) {
		fmt.Printf("Key: %s already exists, No need to create", strings.ToUpper(cmd[2]))
	}

	if key == "" {
		//fmt.Printf("\n Going to create Client: %s \n", value)
		//var key = "/" + strings.ToUpper(cmd[2]) + "/1/" + "CLIENT"
		key = strings.ToUpper(cmd[2])
		k, v := CreateKeyValue(key, cmd[3])
		fmt.Printf("Key: %s  | Value: %s  created\n", k, v)

	}
}

////////////////////////////

func deleteKey(cmd []string) {
	//t := strings.ToUpper(cmd[1])

	key, _ := FindKeyValue(strings.ToUpper(cmd[2]), cmd[3])
	if key == strings.ToUpper(cmd[2]) {
		//fmt.Printf("Key: %s already exists, No need to create", strings.ToUpper(cmd[2]))

		k, v := DeleteKeyValue(key, cmd[3])
		fmt.Printf("Key: %s  | Value: %s  deleted\n", k, v)

	}

	if key == "" {
		fmt.Printf("Key: %s does not exists!", strings.ToUpper(cmd[2]))

	}
}

// /////////////
func updateKey(cmd []string) {
	// /////////
	var v string
	_, value := FindKeyValue(strings.ToUpper(cmd[2]), "*")
	if value != "" {
		x := map[string]string{}
		json.Unmarshal([]byte(value), &x)
		v = x[strings.ToUpper(cmd[2])]

	}
	if v == "" {
		fmt.Printf("Key: %s does not exists!", strings.ToUpper(cmd[2]))
	}

	if v != "" {
		k, v := CreateKeyValue(strings.ToUpper(cmd[2]), cmd[3])
		fmt.Printf("Key: %s  | Value: %s  updated\n", k, v)

	}

}

// /////////////////
func create(cmd []string) {

	t := strings.ToUpper(cmd[1])

	key, value := FindKeyValue(t, strings.ToUpper(cmd[2]))

	/////////////////////////////////
	if value == strings.ToUpper(cmd[2]) {
		fmt.Printf("Client: %s already exists, No need to create", value)
	}

	if key == "" {
		fmt.Printf("\n Going to create Client: %s \n", value)
		var key = "/" + strings.ToUpper(cmd[2]) + "/1/" + "CLIENT"
		k, v := CreateKeyValue(key, strings.ToUpper(cmd[2]))
		fmt.Printf("Key: %s  | Value: %s\n", k, v)
		fmt.Printf("Client %s Created", v)
	}
}

//////////////////////

func FindKeyValue(k string, v string) (kk string, vv string) {

	// ////////////////////////////////
	var url string = "http://localhost:8181/api/v1/find/keyvalue"

	var d1 InputData

	d1.Key = k
	d1.Value = v
	jsonStr, _ := json.Marshal(d1)
	responseBody := bytes.NewBuffer(jsonStr)
	////////////////////////////////
	client1 := &http.Client{}
	req, err := http.NewRequest("GET", url, responseBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client1.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body) //Reads response body using the io.ReadAll
	if err != nil {
		fmt.Printf("%s", err)
	}
	x := map[string]string{}
	json.Unmarshal([]byte(data), &x)
	kk = x["key"]
	vv = x["value"]
	return
}

///////////////

func CreateKeyValue(k string, v string) (kk string, vv string) {

	//fmt.Printf("IN create Key Function\n")
	//fmt.Printf("Key: %s| Value: %s \n", k, v)

	// ////////////////////////////////
	var url string = "http://localhost:8181/api/v1/create/keyvalue"
	var d1 InputData
	d1.Key = k
	d1.Value = v
	jsonStr, _ := json.Marshal(d1)
	responseBody := bytes.NewBuffer(jsonStr)
	////////////////////////////////
	client1 := &http.Client{}
	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client1.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body) //Reads response body using the io.ReadAll
	if err != nil {
		fmt.Printf("%s", err)
	}
	x := map[string]string{}
	json.Unmarshal([]byte(data), &x)
	kk = x["key"]
	vv = x["value"]
	return
}

// /////////////////////////
func DeleteKeyValue(k string, v string) (kk string, vv string) {

	//fmt.Printf("IN create Key Function\n")
	//fmt.Printf("Key: %s| Value: %s \n", k, v)

	// ////////////////////////////////
	var url string = "http://localhost:8181/api/v1/delete/keyvalue"
	var d1 InputData
	d1.Key = k
	d1.Value = v
	jsonStr, _ := json.Marshal(d1)
	responseBody := bytes.NewBuffer(jsonStr)
	////////////////////////////////
	client1 := &http.Client{}
	req, err := http.NewRequest("DELETE", url, responseBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client1.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body) //Reads response body using the io.ReadAll
	if err != nil {
		fmt.Printf("%s", err)
	}
	x := map[string]string{}
	json.Unmarshal([]byte(data), &x)
	kk = x["key"]
	vv = x["value"]
	return
}

////////////////

func listClients(cmd []string) {

	//_, value := FindKeyValue("CLIENT", "*")
	_, value := FindKeyValue("PROJECT", "*")

	if value != "" {
		x := map[string]string{}
		json.Unmarshal([]byte(value), &x)

		var header []string
		header = append(header, "KEY")
		header = append(header, "VALUE")
		printTable(header, x)

		//fmt.Printf("Client: %s already exists, No need to create", value)
	}
	//key = key
}

func listColumn(cmd []string) {

	t := strings.ToUpper(cmd[1])
	_, value := FindKeyValue(t, "*")

	if value != "" {
		x := map[string]string{}
		json.Unmarshal([]byte(value), &x)

		var header []string
		header = append(header, "KEY")
		header = append(header, t)
		printTable(header, x)

		//fmt.Printf("Client: %s already exists, No need to create", value)
	}
	//key = key
}

//////////////////////

func getCmdResult(cmd []string) {
	////////////////////////////////////////////
	counter := -1
	for i := 0; i < len(cmdName); i++ {
		if strings.ToUpper(cmd[1]) == strings.ToUpper(cmdName[i][0]) {
			counter = i
			//getCmdResult(counter, data)
			fmt.Printf("Command is :%s|%d\n", strings.ToUpper(cmdName[counter][0]), counter)
		}
	}
	////////////////////////////////
	// t := strings.ToUpper(cmd[1])
	// _, value := FindKeyValue(t, "*")

	// if value != "" {
	// 	x := map[string]string{}
	// 	json.Unmarshal([]byte(value), &x)

	// 	var header []string
	// 	header = append(header, "KEY")
	// 	header = append(header, t)
	// 	printTable(header, x)

	// 	//fmt.Printf("Client: %s already exists, No need to create", value)
	// }
	// //key = key
}

// ////////////////////////
func selectClient(cmd []string) {

	// //////////////////////////
	file, err := os.Open("./.ims.inf")
	if err != nil {
		log.Fatalf("Failed to open itldmis.conf: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	selectedClient = scanner.Text()
	// scanner.Scan()
	// time = scanner.Text()
	file.Close()
	// ////////////////////////

	if selectedClient == "" {

	}
}

// /////////////////
func printTable(header []string, x map[string]string) {

	rowHeader := make(table.Row, len(header))
	for i := 0; i < len(header); i++ {
		rowHeader[i] = strings.ToUpper(header[i])
	}
	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	uu := t.Style()
	uu.Options.SeparateRows = true
	t.AppendHeader(table.Row(rowHeader))

	for k, v := range x {
		row := make(table.Row, len(header))
		row[0] = k
		row[1] = v
		t.AppendRow(table.Row(row))
	}

	fmt.Println(t.Render())

}

////////////

func login(cmd []string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("UserName:")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	usr := strings.ToUpper(scanner.Text())
	fmt.Printf("Password:")
	//////////////////////
	//var fd int
	b, _ := term.ReadPassword(0)
	pass := string(b)
	//fmt.Printf("UserName: %s |Password: %s\n", usr, pass)

	// /////////////////////////////////
	_, value := FindKeyValue("USR", "*")
	// if value != "" {
	// 	fmt.Printf("NO USR EXIST\n")
	// }

	x := map[string]string{}
	json.Unmarshal([]byte(value), &x)

	var password string = ""
	for i := 1; i < len(x)+1; i++ {
		k := "/USR/" + fmt.Sprint(i) + "/USR"
		v := "/USR/" + fmt.Sprint(i) + "/PASS"
		if x[k] == usr {
			password = x[v]
			loggedInUser = usr
			break
		}
	}

	if password != pass {
		fmt.Printf("\nWrong Password! Login again\n")
	}

	if password == pass {
		fmt.Printf("\nLogged in successfuly !\n")
		loginStatus = 1
		if usr == "ADMIN" {
			loginStatus = 2
		}
	}

	// /////////////////////////////
	// var header []string
	// header = append(header, "KEY")
	// header = append(header, "VALUE")
	// printTable(header, x)

	/////////////////////////////

}

// ////////////
func changePasswd(cmd []string) {

	fmt.Printf("Enter New Password  :")
	b, _ := term.ReadPassword(0)
	pass1 := string(b)
	fmt.Printf("\nReEnter New Password:")
	b, _ = term.ReadPassword(0)
	pass2 := string(b)

	if pass1 != pass2 {
		println("\nBoth Password are not Same--- Relogin\n")
		return
	}

	if loggedInUser == "ADMIN" {
		return
	}

	////////////////////
	_, value := FindKeyValue("USR", "*")
	x := map[string]string{}
	json.Unmarshal([]byte(value), &x)

	//var password string = ""
	for i := 1; i < len(x)+1; i++ {
		k := "/USR/" + fmt.Sprint(i) + "/USR"
		v := "/USR/" + fmt.Sprint(i) + "/PASS"
		if x[k] == loggedInUser {
			CreateKeyValue(v, pass1)
			println("\nPassword changed")
			break
		}
	}
}

// ////////////////
func uploadFile(cmd []string) {
	file := cmd[1]
	convertExcelToCSV(file, "xls.csv")
	uploadToEtcd("xls.csv")
}

// /////////////////////////////////////
// ///////////////////
func convertExcelToCSV(excelFile string, csvFile string) {
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

// ////////////////////////
func uploadToEtcd(csvFile string) {
	// Connect to etcd
	log.Println("Entered into function")
	// // etcdClient, err := clientv3.New(clientv3.Config{
	// // 	Endpoints: []string{etcdHost},
	// // })
	// if err != nil {
	// 	log.Fatalf("Failed to connect to etcd: %v", err)
	// }
	// defer etcdClient.Close()

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

	type ServerData map[string]string

	for _, record := range records[1:] {
		serverIP := record[0]
		serverType := record[1]
		serverData := make(ServerData)

		// Create server data dictionary
		//for i := 0; i < len(headers); i++ {
		for i := 2; i < len(headers); i++ {
			if headers[i] == "" {
				continue
			}
			header := headers[i]
			value := record[i]
			serverData[header] = value
		}

		// Set key-value pairs in etcd for each data field
		// Set key-value pairs in etcd for each data field
		for header, value := range serverData {
			if header == "" {
				continue
			}
			etcdKey := strings.ToUpper(fmt.Sprintf("/%s/%s/%s", serverIP, serverType, header))
			etcdValue := value
			// fmt.Println(etcdKey)
			// fmt.Println(etcdValue)
			k, v := CreateKeyValue(etcdKey, etcdValue)
			fmt.Printf("Key: %s  | Value: %s\n", k, v)

		}

		///////////////////////
		// // Set key-value pair for server data
		// etcdKeyData := strings.ToUpper(fmt.Sprintf("/%s/%s/data", serverIP, serverType))
		// etcdValueData, err := json.Marshal(serverData)
		// fmt.Printf("from data upload\n\n\n")
		// fmt.Printf("key:%s|data:%s\n", etcdKeyData, etcdValueData)
		// if err != nil {
		// 	log.Printf("Failed to marshal server data: %v", err)
		// 	continue
		// }
		// _, err = etcdClient.Put(context.Background(), etcdKeyData, string(etcdValueData))
		// if err != nil {
		// 	log.Printf("Failed to upload server data to etcd: %v", err)
		// }

		//////////////////////

	}

	log.Println("Server details added to etcd successfully.")
}

// /////////////////////
func readconf() {
	addressServer = ""
	file, err := os.Open("./ims.conf")
	if err != nil {
		log.Fatalf("Failed to open itldmis.conf: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	addressServer = scanner.Text()
	scanner.Scan()

	numbers_linkcells, _ := strconv.Atoi(scanner.Text())
	for j := 0; j < numbers_linkcells; j++ {
		scanner.Scan()
		linkcells[j] = append(linkcells[j], scanner.Text())

	}

	////////////////////////
	scanner.Scan()
	numbers_cmd, _ := strconv.Atoi(scanner.Text())
	for j := 0; j < numbers_cmd; j++ {
		scanner.Scan()
		cmdName[j] = append(cmdName[j], scanner.Text())
		scanner.Scan()
		tableName[j] = append(tableName[j], scanner.Text())
		scanner.Scan()
		tableHeader[j] = append(tableHeader[j], strings.ToUpper(scanner.Text()))
	}

	file.Close()

	/////////////////////////
	OutPutFile, err = os.Create("./output.txt")
	if err != nil {
		log.Fatalf("Failed to open output: %v", err)
	}
	defer OutPutFile.Close()

	return
}

/////////////////////////////
