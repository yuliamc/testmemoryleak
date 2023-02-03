package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"yuliamchandra/utils/constant"
	"yuliamchandra/utils/file"
	"yuliamchandra/utils/null"

	_ "net/http/pprof"
)

func main() {
	go startPPROF()

	http.HandleFunc("/check", PostCheckHandler)

	fmt.Println("Server will listen on :8090")
	http.ListenAndServe(":8090", nil)
}

func startPPROF() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}

type PostCheckHandlerBody struct {
	Url *string `json:"url"`
}

func PostCheckHandler(w http.ResponseWriter, req *http.Request) {
	// POST only.
	if req.Method != "POST" {
		fmt.Fprint(w, "Method Not Supported!")
		return
	}

	// Gets body.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "err read body %s", err.Error())
		return
	}

	var reqBody PostCheckHandlerBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		fmt.Fprintf(w, "err unmarshal %s", err.Error())
		return
	}

	// Handler.
	log.Println(*reqBody.Url)
	keyValues := []constant.KeyValue{
		{Key: "KTP", Value: *reqBody.Url},
	}

	maxFileSize := int64(3)
	err = CheckRemoteFileSize(&keyValues, &maxFileSize)
	if err != nil {
		fmt.Fprintf(w, "err check remote file size %s", err.Error())
	} else {
		fmt.Fprint(w, "no err")
	}
}

func CheckRemoteFileSize(keyUrlMap *[]constant.KeyValue, maxFileSize *int64) error {
	var (
		invalidSizeFileKey = make([]string, 0)
	)

	if null.IsNil(keyUrlMap) {
		return errors.New("INTERNAL_SERVER_ERROR")
	}

	for _, item := range *keyUrlMap {
		if len(strings.TrimSpace(item.Key)) == 0 || len(strings.TrimSpace(item.Value)) == 0 {
			continue
		}
		if isValid, err := file.IsRemoteFileValidSize(&item.Value, maxFileSize); err != nil {
			invalidSizeFileKey = append(invalidSizeFileKey, item.Key)
		} else if !*isValid {
			invalidSizeFileKey = append(invalidSizeFileKey, item.Key)
		}
	}
	if len(invalidSizeFileKey) > 0 {
		return errors.New("UPLOAD_TO_S3_FAILED")
	}

	return nil
}
