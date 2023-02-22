package infoblox

import (
	"crypto/tls"
	"fmt"
	"sort"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	httpClient *resty.Client
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(host string, username string, password string) Client {
	baseURL := fmt.Sprintf("https://%s/wapi/v2.10", host)
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBaseURL(baseURL)
	client.SetBasicAuth(username, password)
	return Client{
		httpClient: client,
	}
}

func (c *Client) GetZones() ([]Zone, error) {
	res, err := c.httpClient.R().SetResult(&[]Zone{}).Get("/zone_auth")
	if err != nil {
		return nil, err
	}
	zones := *res.Result().(*[]Zone)
	sort.Slice(zones, func(i, j int) bool {
		return zones[i].Fqdn < zones[j].Fqdn
	})
	return zones, nil
}

func (c *Client) GetZoneRecords(zone string) ([]HostRecord, error) {
	res, err := c.httpClient.R().SetResult(&[]HostRecord{}).Get("/record:host?zone=" + zone)
	if err != nil {
		return nil, err
	}
	records := *res.Result().(*[]HostRecord)
	sort.Slice(records, func(i, j int) bool {
		return records[i].Name < records[j].Name
	})

	return records, nil
}

func (c *Client) AddHostRecord(fqdn string, ip string, view string) (any, error) {
	res, err := c.httpClient.R().SetBody(map[string]any{
		"name": fqdn,
		"ipv4addrs": []map[string]any{
			{
				"ipv4addr": ip,
			},
		},
		"view": view,
	}).Post("/record:host")
	if err != nil {
		return "", err
	}
	return res, nil
}

func (c *Client) GetHostRecord(fqdn string, view string) ([]HostRecord, error) {
	res, err := c.httpClient.R().SetResult(&[]HostRecord{}).Get("/record:host?name=" + fqdn + "&view=" + view)
	if err != nil {
		return []HostRecord{}, err
	}
	return *res.Result().(*[]HostRecord), nil
}

func (c *Client) RemoveHostRecord(fqdn string, view string) (any, error) {

	// get host record ref from GetHostRecord
	record, err := c.GetHostRecord(fqdn, view)
	if err != nil {
		return "", err
	}

	if len(record) == 0 {
		return "", fmt.Errorf("no record found for %s", fqdn)
	}

	// delete host record
	res, err := c.httpClient.R().Delete("/" + record[0].Ref)
	if err != nil {
		return "", err
	}
	return res, nil
}

type Zone struct {
	Ref  string `json:"_ref"`
	Fqdn string `json:"fqdn"`
	View string `json:"view"`
}

type HostRecord struct {
	Ref       string `json:"_ref"`
	Name      string `json:"name"`
	View      string `json:"view"`
	Ipv4Addrs []struct {
		Ref      string `json:"_ref"`
		Host     string `json:"host"`
		Ipv4Addr string `json:"ipv4addr"`
	} `json:"ipv4addrs"`
}
