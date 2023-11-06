//https://github.com/Jeffail/gabs

package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	clientv3 "go.etcd.io/etcd/client/v3" //helps storing & fetching with the etcd api
)

type InputData struct {
	gorm.Model
	Key   string `json:"key"`
	Value string `json:"value"`
}

var (
	// File paths
	//excelFile = "./api.xlsx"
	excelFile = "./app.xlsx"
	csvFile   = "./myetcd.csv"
	etcdHost  = "localhost:2379"
)

type ServerData map[string]string

func CreateKeyValue(c *fiber.Ctx) {
	fmt.Println("In Create Key")

	data := new(InputData)
	if err := c.BodyParser(data); err != nil {
		c.Status(503).Send(err)
		return
	}
	fmt.Printf("Key:%s\nValue:%s\n", data.Key, data.Value)
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	_, err = etcdClient.Put(context.Background(), data.Key, data.Value)
	if err != nil {
		log.Printf("Failed to upload server data to etcd: %v", err)
		c.JSON(data)
	}
	if err == nil {
		updateData(data.Key)
		c.JSON(data)
	}

}

// //////////////////////////////
func DeleteKeyValue(c *fiber.Ctx) {
	fmt.Println("In Create Key")

	data := new(InputData)
	if err := c.BodyParser(data); err != nil {
		c.Status(503).Send(err)
		return
	}
	fmt.Printf("Key:%s\nValue:%s\n", data.Key, data.Value)
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	//_, err = etcdClient.Put(context.Background(), data.Key, data.Value)
	_, err = etcdClient.Delete(context.Background(), data.Key)
	if err != nil {
		log.Printf("Failed to upload server data to etcd: %v", err)
		c.JSON(data)
	}
	if err == nil {
		updateData(data.Key)
		c.JSON(data)
	}

}

// ///////////////////////////////////////////////////
func FindKeyValue(c *fiber.Ctx) {

	data := new(InputData)
	if err := c.BodyParser(data); err != nil {
		c.Status(503).Send(err)
		return
	}

	fmt.Printf("Key:%s\n", data.Key)
	fmt.Printf("Value:%s\n", data.Value)
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	//st := data.Key
	response, err := etcdClient.Get(context.Background(), "/", clientv3.WithPrefix())
	if err != nil {
		log.Printf("Failed to retrieve keys from etcd: %v", err)
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// //////////////////
	x := map[string]string{}
	for _, kv := range response.Kvs {
		if strings.Contains(strings.ToUpper(string(kv.Key)), strings.ToUpper((data.Key))) {
			x[string(kv.Key)] = string(kv.Value)
		}
		//x[string(kv.Key)] = string(kv.Value)
	}

	var key, value string

	if data.Value != "*" {
		for kk, kv := range x {
			if strings.ToUpper(kv) == strings.ToUpper(data.Value) {
				key = kk
				value = kv
			}
		}
	}

	var kkk, vvv []string
	if data.Value == "*" {
		for kk, kv := range x {
			kkk = append(kkk, kk)
			vvv = append(vvv, kv)

		}
	}

	var out InputData
	out.Key = key
	out.Value = value
	if data.Value == "*" {
		m, _ := json.Marshal(x)
		// c.JSON(string(m))
		out.Key = data.Key
		out.Value = string(m)
	}
	c.JSON(out)
}

///////////////

// //////////////////
func updateData(key1 string) {

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Printf("Failed to connect to etcd: %v", err)
		return
	}
	defer etcdClient.Close()
	ll := strings.Split(key1, "/")
	var Key string = "/" + ll[1] + "/" + ll[2]
	response, err := etcdClient.Get(context.Background(), Key, clientv3.WithPrefix())

	var valueData string = "{"
	var counter int = 0
	for _, kv := range response.Kvs {
		kk := strings.Split(string(kv.Key), "/")
		if kk[3] == "DATA" {
			continue
		}
		if counter != 0 {
			valueData = valueData + fmt.Sprintf(",")
		}
		valueData = valueData + fmt.Sprintf("%c%s%c:", 34, kk[3], 34)
		valueData = valueData + fmt.Sprintf("%c%s%c", 34, string(kv.Value), 34)
		counter++
	}
	valueData = valueData + "}"

	fmt.Printf("\n\n\n")

	tt, err := json.Marshal(valueData)
	tt1 := string(tt)

	fmt.Println(tt1)
	// ////////////////////////
	serverData := make(ServerData)

	for _, kv := range response.Kvs {
		kk := strings.Split(string(kv.Key), "/")

		if kk[3] == "DATA" {
			continue
		}
		header := kk[3]
		value := string(kv.Value)
		serverData[header] = value

	}

	fmt.Printf("\n\n\n  From Last Line\n")
	etcdValueData, err := json.Marshal(serverData)
	etcdKeyData := Key + "/DATA"

	fmt.Printf("key:%s|data:%s\n", etcdKeyData, etcdValueData)
	_, err = etcdClient.Put(context.Background(), etcdKeyData, string(etcdValueData))
	if err != nil {
		log.Printf("Failed to upload server data to etcd: %v", err)
	}
	////////////////////////

}

////////////////////
