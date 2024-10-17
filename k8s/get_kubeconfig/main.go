package main

import (
	"context"
	"log"
	"os"
	decort "repository.basistech.ru/BASIS/decort-golang-sdk"
	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/k8s"
	"strconv"
	"time"
)

func main() {
	cfg := config.Config{
		AppID:     os.Getenv("DECORT_APP_ID"),
		AppSecret: os.Getenv("DECORT_APP_SECRET"),
		SSOURL:    os.Getenv("DECORT_SSO_URL"),
		DecortURL: os.Getenv("DECORT_URL"),
		Retries:   1,
	}

	k8sid, err := strconv.ParseUint(os.Getenv("DECORT_K8SID"), 10, 64)

	if err != nil {
		log.Fatal("Invalid k8sid")
		return
	}

	cfg.SetTimeout(60 * time.Second)

	client := decort.New(cfg)

	k8sRecord, err := client.CloudAPI().K8S().GetConfig(context.Background(), k8s.GetConfigRequest{
		K8SID: k8sid,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	print(k8sRecord, "\n")
}
