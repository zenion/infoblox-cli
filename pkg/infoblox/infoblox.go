package infoblox

import (
	"encoding/json"
	"fmt"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func NewClient(host string, username string, password string) (ibclient.IBObjectManager, error) {
	hostConfig := ibclient.HostConfig{
		Host:    host,
		Version: "2.10",
	}
	authConfig := ibclient.AuthConfig{
		Username: username,
		Password: password,
	}
	transportConfig := ibclient.NewTransportConfig("false", 20, 10)
	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}
	conn, err := ibclient.NewConnector(hostConfig, authConfig, transportConfig, requestBuilder, requestor)
	if err != nil {
		return nil, fmt.Errorf("error creating connector: %s", err)
	}
	defer conn.Logout()
	objMgr := ibclient.NewObjectManager(conn, "myclient", "")
	return objMgr, nil
}

func MarshalAndPrint(obj interface{}) error {
	json, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(json))
	return nil
}
