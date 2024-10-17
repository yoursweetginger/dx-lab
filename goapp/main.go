package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	decort "repository.basistech.ru/BASIS/decort-golang-sdk"
	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/k8s"
)

func main() {
	client, k8sid, err := initDecort()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()

	r.StaticFile("/", "./static/index.html")

	r.GET("/api/k8s", func(c *gin.Context) {

		data, err := client.CloudAPI().K8S().Get(c, k8s.GetRequest{K8SID: k8sid})
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, data)
	})
	r.Run()
}

func initDecort() (*decort.DecortClient, uint64, error) {
	cfg := config.Config{
		AppID:     os.Getenv("DECORT_APP_ID"),
		AppSecret: os.Getenv("DECORT_APP_SECRET"),
		SSOURL:    os.Getenv("DECORT_SSO_URL"),
		DecortURL: os.Getenv("DECORT_URL"),
		Retries:   1,
	}
	cfg.SetTimeout(60 * time.Second)

	client := decort.New(cfg)

	k8sid, err := strconv.ParseUint(os.Getenv("DECORT_K8SID"), 10, 64)

	if err != nil {
		log.Fatal("Invalid k8sid")
		return nil, 0, err
	}
	return client, k8sid, err
}
