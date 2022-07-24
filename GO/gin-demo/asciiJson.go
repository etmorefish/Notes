package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Vcenter struct {
	VcenterPort     string `json:"vcenterPort"`
	VcenterIp       string `json:"vcenterIp"`
	VcenterPassword string `json:"vcenterPassword"`
	VcenterUser     string `json:"vcenterUser"`
}

type GetCmdbResponse struct {
	Data []Vcenter `json:"data"`
}

type GetCloudAccountReq struct {
	TypeName string `json:"typeName"`
}

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang":    "GO语言",
			"tag":     "<br>",
			"version": 1,
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.POST("/getCloudAccount", func(c *gin.Context) {
		var body GetCloudAccountReq
		if err := c.ShouldBindJSON(&body); err != nil {
			fmt.Println("err: this request is not body")
			return
		}
		if body.TypeName == "" {
			fmt.Println("typeName is empty")
			return
		} else {
			fmt.Printf("typename is %s", body.TypeName)
		}

		data := GetCmdbResponse{
			Data: []Vcenter{
				{VcenterPort: "111",
					VcenterIp:       "123",
					VcenterPassword: "333",
					VcenterUser:     "444"},
			},
		}
		// c.AsciiJSON(http.StatusOK, data)
		c.JSONP(http.StatusOK, data)

	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
