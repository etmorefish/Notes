package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Resp struct {
	UserInfo struct {
		ID                 string `json:"id"`
		LoginName          string `json:"loginName"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		Type               string `json:"type"`
		Status             string `json:"status"`
		Email              string `json:"email"`
		Telephone          string `json:"telephone"`
		IsDeleted          int    `json:"isDeleted"`
		CreateTime         string `json:"createTime"`
		UpdateTime         string `json:"updateTime"`
		PasswordUpdateTime string `json:"passwordUpdateTime"`
		Aggregation        struct {
			Roles []struct {
				ID          string      `json:"id"`
				Name        string      `json:"name"`
				Description string      `json:"description"`
				Type        string      `json:"type"`
				IsDefault   bool        `json:"isDefault"`
				CreateTime  string      `json:"createTime"`
				UpdateTime  string      `json:"updateTime"`
				Users       interface{} `json:"users"`
				Permissions interface{} `json:"permissions"`
				UserCount   interface{} `json:"userCount"`
			} `json:"roles"`
			UserVdcs []struct {
				ID          string        `json:"id"`
				Name        string        `json:"name"`
				Description string        `json:"description"`
				IsDeleted   int           `json:"isDeleted"`
				Creator     string        `json:"creator"`
				CreateTime  string        `json:"createTime"`
				UpdateTime  string        `json:"updateTime"`
				DcIds       []interface{} `json:"dcIds"`
			} `json:"userVdcs"`
		} `json:"aggregation"`
	} `json:"data"`
}

func main() {

	url := "http://identity:11304/getUser"
	method := "POST"

	payload := strings.NewReader(`{"id":"1b729924-5d92-49bf-9942-5bed03ad7f1b"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
	var resp Resp
	err = json.Unmarshal(body, &resp)

	// fmt.Printf("resp: %#v\n", resp)
	var vids []string
	for _, v := range resp.UserInfo.Aggregation.UserVdcs {
		vids = append(vids, v.ID)
	}

	fmt.Println(vids)

}
