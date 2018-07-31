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
type RevId struct {
	Rev string `json:"_rev`
}

func revGetter(db string, id string) {

	getRevIdURL := `http://root:asif4106@127.0.0.1:5984/` + db + "/" + id
	revReq, err := http.NewRequest("GET", getRevIdURL, nil)
	revReq.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(revReq)
	if err != nil {
		fmt.Println("somthing went wrong in getting rev")
	}
	var revID RevId
	_ = json.Decoder(resp.Body).Decode(&revID)

	return revID.Rev
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
		ID         string `json:"_id"`
		RevisionID string `json:"_rev"`
	}

	http.HandleFunc("/put-test", func(rw http.ResponseWriter, r *http.Request) {
		var putReqBody ReqBody
		_ = json.NewDecoder(r.Body).Decode(&putReqBody)

		//ID := putReqBody.ID

		getRevIdURL := `http://root:asif4106@127.0.0.1:5984/newdb/` + putReqBody.ID

		client := http.Client{}
		revReq, err := http.NewRequest("GET", getRevIdURL, nil)
		revReq.Header.Set("Content-Type", "application/json")
		revResp, err := client.Do(revReq)
		if err != nil {
			panic(err)
		}
		// defer resp.Body.Close()

		// revBody, _ := ioutil.ReadAll(revResp.Body)
		// fmt.Println("response Body:", string(body))
		var putReqRevBody ReqBody
		_ = json.NewDecoder(revResp.Body).Decode(&putReqRevBody)

		updateValue := `'{"topic" : "` + putReqBody.Topic + `, "data" : "` + putReqBody.Data + `, "_rev" : "` + putReqRevBody.RevisionID + `"}'`
		url := "http://root:asif4106@127.0.0.1:5984/newdb/" + putReqBody.ID + " -d " + updateValue
		fmt.Println(url)

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
