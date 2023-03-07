package infoblox

import (
	"crypto/tls"
	"errors"
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
	res, err := c.httpClient.R().
		SetResult(&[]Zone{}).
		Get("/zone_auth")
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
	var records []HostRecord

	res, err := c.httpClient.R().
		SetResult(&PagedHostRecords{}).
		SetError(&InfobloxError{}).
		Get("/record:host?zone=" + zone + "&_return_as_object=1&_paging=1&_max_results=1000")
	if err != nil {
		return nil, errors.New(res.Error().(*InfobloxError).Text)
	}
	PagedRecords := *res.Result().(*PagedHostRecords)
	records = append(records, PagedRecords.Result...)

	for PagedRecords.NextPageId != "" {
		fmt.Println("Next page: " + PagedRecords.NextPageId)
		res, err := c.httpClient.R().
			SetResult(&PagedHostRecords{}).
			SetError(&InfobloxError{}).
			Get("/record:host?zone=" + zone + "&_return_as_object=1&_paging=1&_max_results=1000&_page_id=" + PagedRecords.NextPageId)
		if err != nil {
			return nil, errors.New(res.Error().(*InfobloxError).Text)
		}
		PagedRecords = *res.Result().(*PagedHostRecords)
		records = append(records, PagedRecords.Result...)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Name < records[j].Name
	})

	// return records, nil
	return records, nil
}

func (c *Client) AddHostRecord(fqdn string, ip string, view string) (any, error) {
	res, err := c.httpClient.R().
		SetBody(map[string]any{
			"name": fqdn,
			"ipv4addrs": []map[string]any{
				{
					"ipv4addr": ip,
				},
			},
			"view": view,
		}).
		Post("/record:host")
	if err != nil {
		return "", err
	}
	return res, nil
}

func (c *Client) GetHostRecord(fqdn string, view string) ([]HostRecord, error) {
	res, err := c.httpClient.R().
		SetResult(&[]HostRecord{}).
		Get("/record:host?name=" + fqdn + "&view=" + view)
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
	Ref  string `json:"_ref,omitempty"`
	Fqdn string `json:"fqdn,omitempty"`
	View string `json:"view,omitempty"`
}

type HostRecord struct {
	Ref       string `json:"_ref,omitempty"`
	Name      string `json:"name,omitempty"`
	View      string `json:"view,omitempty"`
	Ipv4Addrs []struct {
		Ref      string `json:"_ref,omitempty"`
		Host     string `json:"host,omitempty"`
		Ipv4Addr string `json:"ipv4addr,omitempty"`
	} `json:"ipv4addrs,omitempty"`
}

type PagedHostRecords struct {
	Result     []HostRecord `json:"result,omitempty"`
	NextPageId string       `json:"next_page_id,omitempty"`
}

type InfobloxError struct {
	Error string `json:"Error,omitempty"`
	Code  string `json:"code,omitempty"`
	Text  string `json:"text,omitempty"`
}
