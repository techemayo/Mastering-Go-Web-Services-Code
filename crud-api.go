package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type testData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

func main() {
	http.HandleFunc("/get-test", func(rw http.ResponseWriter, r *http.Request) {

		url := "http://root:asif4106@127.0.0.1:5984/newdb/_all_docs"
		fmt.Println("URL:>", url)

		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	})
	http.HandleFunc("/post-test", func(rw http.ResponseWriter, r *http.Request) {

		url := "http://root:asif4106@127.0.0.1:5984/newdb"
		fmt.Println("URL:>", url)
		var data testData
		_ = json.NewDecoder(r.Body).Decode(&data)
		fmt.Println(data)
		newJson := `{"Title":"` + data.Topic + `","Data":"` + data.Data + `"}`
		fmt.Println(newJson)
		var jsonStr = []byte(newJson)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	})
	type ReqBody struct {
		Topic      string `json:"topic"`
		Data       string `json:"data"`
		ID         string `json:"id"`
		RevisionID string `json:"revisionID"`
	}
	http.HandleFunc("/put-test", func(rw http.ResponseWriter, r *http.Request) {
		var putReqBody ReqBody
		_ = json.NewDecoder(r.Body).Decode(&putReqBody)

		ID := putReqBody.ID
		updateValue := `'{"topic" : "` + putReqBody.Topic + `, "_rev" : "` + putReqBody.RevisionID + `"}'`
		url := "http://root:asif4106@127.0.0.1:5984/newdb/" + ID + " -d " + updateValue
		fmt.Println(url)
		client := http.Client{}
		req, err := http.NewRequest("PUT", url, nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	})
	http.ListenAndServe(":8000", nil)
}

//const da = `{"title":"new title"}`
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// func main() {

// 	url := "http://root:asif4106@127.0.0.1:5984/newdb"

// 	payload := strings.NewReader("{\n \"data\": \"some data\",\n \"title\": \"my new data\"\n}")

// 	req, _ := http.NewRequest("POST", url, payload)

// 	req.Header.Add("content-type", "application/json")
// 	req.Header.Add("cache-control", "no-cache")
// 	req.Header.Add("postman-token", "7e078c2a-598c-5f73-3e62-786b324becf7")

// 	res, _ := http.DefaultClient.Do(req)

// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)

// 	fmt.Println(res)
// 	fmt.Println(string(body))

// }

/*
	var data testData
	status := json.NewDecoder(r.Body).Decode(&data)
	if status != nil {
		fmt.Println(status)
	}
	//fmt.Println(data)
	// fmt.Println(data.Data)
	// fmt.Println(data.Topic)
	newData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(newData)
	req, err := http.NewRequest("POST", "http://root:asif4106@127.0.0.1:5984/newdb", bytes.NewBuffer(newData))
	if err != nil {
		fmt.Println(err)
	}
	http.DefaultClient.Do(req)
*/
// url := "http://root:asif4106@127.0.0.1:5984/newdb"
// fmt.Println("URL:>", url)
// req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
// req.Header.Set("X-Custom-Header", "myvalue")
// req.Header.Set("Content-Type", "application/json")

// client := &http.Client{}
// resp, err := client.Do(req)
// if err != nil {
// 	panic(err)
// }
// defer resp.Body.Close()

// fmt.Println("response Status:", resp.Status)
// fmt.Println("response Headers:", resp.Header)
// body, _ := ioutil.ReadAll(resp.Body)
// fmt.Println("response Body:", string(body))

//client := &http.Client{}
//newData := testData{}
// bodyData, err := ioutil.ReadAll(r.Body)
// if err != nil {
// 	fmt.Println("something happend")
// }
//fmt.Println(bodyData)

// newData.Data = string(bodyData[:len(bodyData)/2])
// newData.Topic = string(bodyData[len(bodyData)/2:])
// newData = string(bodyData)
// postData, err := json.Marshal(newData)
// postData, err := json.Marshal(bodyData)
// if err != nil {
// 	fmt.Println("Something went wrong")
// }
// fmt.Println(string(postData))
// //sendData := strings.NewReader(bodyData)
//req, err := http.NewRequest("POST", "http://root:asif4106@127.0.0.1:5984/newdb", sendData)
//req, err := http.NewRequest("POST", "http://root:asif4106@127.0.0.1:5984/newdb", bytes.NewBuffer(postData))
// req, err := http.NewRequest("POST", "http://root:asif4106@127.0.0.1:5984/newdb", bytes.NewBuffer(bodyData))
// req.Header.Add("content-type", "application/json")
// req.Header.Add("cache-control", "no-cache")

// //fmt.Fprintf(rw, string(postData))
// if err != nil {
// 	fmt.Println("Somthing went wrong in new request")
// }
// //_, err = client.Do(req)
// resp, err := http.DefaultClient.Do(req)

// if err != nil {
// 	fmt.Println(err)
// }
// defer resp.Body.Close()
// body, _ := ioutil.ReadAll(resp.Body)
// fmt.Println(resp)

// fmt.Println(body)
