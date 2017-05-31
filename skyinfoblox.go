//Provides an interface to the Infoblox API
package skyinfoblox

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

type InfobloxClient struct {
	Host         string
	Password     string
	Username     string
	HttpClient   *http.Client
	Timeout      time.Duration
	UseCookies   bool
	View         string
	Network_view string
	Wapi_version string
}

// NewInfobloxClient creates a new client to interface with the Infoblox API.
// Parameters are
//  host: the hostame of the infoblox device to connect to ,eg https://192.168.0.1
//  username: the username to authenticate with infoblox in plain text.
//  password: password to be used with authentication in plain text.
//  view: infoblox view
//  wapi_version: API version, default to 2.3
//  network_view: network view in infoblox
//  sslVerify: verify valid SSL certificate
//  useCookies: not implemented
//  timeout: timeout in seconds for each connection
func NewInfobloxClient(host, username, password, view, wapi_version, network_view string, sslVerify, useCookies bool, timeout int) *InfobloxClient {
	//Default values
	if network_view == "" {
		network_view = "default"
	}
	if view == "" {
		view = "default"
	}
	if wapi_version == "" {
		wapi_version = "2.3"
	}
	if timeout < 1 {
		timeout = 20
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: sslVerify},
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second, Transport: tr}

	ibxclient := &InfobloxClient{
		Host:         host,
		Password:     password,
		Username:     username,
		HttpClient:   client,
		View:         view,
		Network_view: network_view,
		Wapi_version: wapi_version,
		Timeout:      time.Duration(timeout),
	}
	return ibxclient
}

func funcName() string {
	pc, _, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s():%d ", runtime.FuncForPC(pc).Name(), line)
}

func read_httpresponse(response *http.Response) (string, error) {
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func read_jsonresponse(resp *http.Response) (ret []interface{}, err error) {

	// invalid or incomplete json may panic, catch the panic and return the error generated
	defer func() {
		info := recover()
		if info != nil {
			j := fmt.Sprintf("%s : %s", funcName(), info)
			err = errors.New(j)
		}
	}()

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	m := data.([]interface{})
	return m, nil
}

//Get_zone_auth returns the specified authoriative zone
//It take the zone name as the parameter.
//The function returns a slice of interfaces and an error
func (ibx InfobloxClient) Get_zone_auth(fqdn string) ([]interface{}, error) {

	url := fmt.Sprintf("%s/wapi/v%s/zone_auth?fqdn=%s&view=%s", ibx.Host, ibx.Wapi_version, fqdn, ibx.View)
	return ibx.Get_record(fqdn, url)
}

//Get_text_record returns the text record for the specified host.
// Parameter is a single fqdn to retrieve.
// If they parameter is found it returns a slice of interfaces.
// If it is no found an empty slice is returned.
// Otherwise an error will be returned.
func (ibx InfobloxClient) Get_text_record(fqdn string) ([]interface{}, error) {

	url := fmt.Sprintf("%s/wapi/v%s/record:txt?name=%s&view=%s", ibx.Host, ibx.Wapi_version, fqdn, ibx.View)
	return ibx.Get_record(fqdn, url)
}

//Get_network returns the specified network
func (ibx InfobloxClient) Get_network(network string) ([]interface{}, error) {

	url := fmt.Sprintf("%s/wapi/v%s/network?network=%s&network_view=%s", ibx.Host, ibx.Wapi_version, network, ibx.Network_view)
	return ibx.Get_record(network, url)
}

//Get_srv_record returns the specified SRV record
//The name parameter is required, target can be an empty string and others set to negative to be ignored, eg -1
func (ibx InfobloxClient) Get_srv_record(name, target string, weight, priority, port int) ([]interface{}, error) {

	data := "&_return_fields=ttl,target,weight,priority,port,name"
	if target != "" {
		data = fmt.Sprintf("%s&target=%s", data, target)
	}
	if weight >= 0 {
		data = fmt.Sprintf("%s&weight=%d", data, weight)
	}
	if priority >= 0 {
		data = fmt.Sprintf("%s&priority=%d", data, priority)
	}
	if port > 0 {
		data = fmt.Sprintf("%s&port=%d", data, port)
	}
	url := fmt.Sprintf("%s/wapi/v%s/record:srv?name=%s&view=%s%s", ibx.Host, ibx.Wapi_version, name, ibx.View, data)

	return ibx.Get_record(name, url)
}

//Get_host_record returns the record for the specified host.
func (ibx InfobloxClient) Get_host_record(fqdn string) ([]interface{}, error) {

	url := fmt.Sprintf("%s/wapi/v%s/record:host?name=%s&view=%s&_return_fields=ttl,ipv4addrs", ibx.Host, ibx.Wapi_version, fqdn, ibx.View)
	// url := fmt.Sprintf("%s/wapi/v%s/record:host?name=%s&view=%s", ibx.Host, ibx.Wapi_version, fqdn, ibx.View)
	return ibx.Get_record(fqdn, url)
}

//Get_cname_record returns the specified cname record
func (ibx InfobloxClient) Get_cname_record(fqdn string) ([]interface{}, error) {

	url := fmt.Sprintf("%s/wapi/v%s/record:cname?name=%s&view=%s", ibx.Host, ibx.Wapi_version, fqdn, ibx.View)
	return ibx.Get_record(fqdn, url)
}

func (ibx InfobloxClient) Get_record(fqdn, url string) ([]interface{}, error) {

	//url := fmt.Sprintf("%s/wapi/v%s/record:host?name=%s&view=%s", ibx.Host, ibx.Wapi_version, fqdn, ibx.View)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(ibx.Username, ibx.Password)
	resp, herr := ibx.HttpClient.Do(req)
	if herr != nil {
		return nil, herr
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		errresp, staterr := read_httpresponse(resp)
		if staterr != nil {
			return nil, staterr
		} else {
			return nil, errors.New(errresp)
		}
	}

	jresp, jerr := read_jsonresponse(resp)
	if jerr != nil {
		return nil, jerr
	} else {
		return jresp, nil
	}
}

//Update_host_record updates the IP address for a given host name
func (ibx InfobloxClient) Update_host_record(fqdn, address string) error {

	ret, err := ibx.Get_host_record(fqdn)
	if err != nil {
		return err
	}
	if len(ret) < 1 {
		err_txt := fmt.Sprintf("Record not found: %s", fqdn)
		return errors.New(err_txt)
	}
	data := ret[0].(map[string]interface{})
	payload := fmt.Sprintf("{\"ipv4addrs\":[{ \"ipv4addr\":\"%s\" }]}", address)
	postdata := []byte(payload)
	url := fmt.Sprintf("%s/wapi/v%s/%s", ibx.Host, ibx.Wapi_version, data["_ref"])

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(postdata))
	if err != nil {
		return err
	}

	req.SetBasicAuth(ibx.Username, ibx.Password)
	resp, err := ibx.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		errresp, staterr := read_httpresponse(resp)
		if staterr != nil {
			return staterr
		} else {
			return errors.New(errresp)
		}
	}
	return nil

}

//Delete_network deleted the specified network eg 10.77.58.0/24
func (ibx InfobloxClient) Delete_network(network string) error {

	ret, err := ibx.Get_network(network)
	if err != nil {
		return err
	}
	err = ibx.delete_record(network, ret)
	if err != nil {
		return err
	}
	return nil
}

//Delete_text_record deletes the specified text record by name
func (ibx InfobloxClient) Delete_text_record(fqdn string) error {

	ret, err := ibx.Get_text_record(fqdn)
	if err != nil {
		return err
	}
	err = ibx.delete_record(fqdn, ret)
	if err != nil {
		return err
	}
	return nil
}

//Delete_canme_record deleted the specified cname
func (ibx InfobloxClient) Delete_cname_record(fqdn string) error {

	ret, err := ibx.Get_cname_record(fqdn)
	if err != nil {
		return err
	}
	err = ibx.delete_record(fqdn, ret)
	if err != nil {
		return err
	}
	return nil
}

func (ibx InfobloxClient) Delete_srv_record(name, target string, weight, priority, port int) error {

	ret, err := ibx.Get_srv_record(name, target, weight, priority, port)
	if err != nil {
		return err
	}
	err = ibx.delete_record(name, ret)
	if err != nil {
		return err
	}
	return nil
}

//Delete_host_record deletes the specified host
func (ibx InfobloxClient) Delete_host_record(fqdn string) error {

	ret, err := ibx.Get_host_record(fqdn)
	if err != nil {
		return err
	}
	err = ibx.delete_record(fqdn, ret)
	if err != nil {
		return err
	}
	return nil
}

//Delete_zone_auth deletes the entry for an authorative zone
//returns an error on failure and nil on success
func (ibx InfobloxClient) Delete_zone_auth(fqdn string) error {

	ret, err := ibx.Get_zone_auth(fqdn)
	if err != nil {
		return err
	}
	err = ibx.delete_record(fqdn, ret)
	if err != nil {
		return err
	}
	return nil
}

func (ibx InfobloxClient) delete_record(fqdn string, ret []interface{}) error {

	if len(ret) < 1 {
		err_txt := fmt.Sprintf("Record not found: %s", fqdn)
		return errors.New(err_txt)
	}
	/*
		if len(ret) > 1 {
			err_txt := fmt.Sprintf("Multiple reocrds matched, must match only one: %s", ret)
			return errors.New(err_txt)
		}
	*/
	for _, d := range ret {
		//data := ret[0].(map[string]interface{})
		data := d.(map[string]interface{})

		url := fmt.Sprintf("%s/wapi/v%s/%s", ibx.Host, ibx.Wapi_version, data["_ref"])
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return err
		}

		req.SetBasicAuth(ibx.Username, ibx.Password)
		resp, err := ibx.HttpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 299 || resp.StatusCode < 200 {
			errresp, staterr := read_httpresponse(resp)
			if staterr != nil {
				return staterr
			} else {
				return errors.New(errresp)
			}
		}
	}
	return nil

}

//Create_network_record creates a host record, that is A record and PTR record
func (ibx InfobloxClient) Create_network(network string) (string, error) {

	url := fmt.Sprintf("%s/wapi/v%s/network", ibx.Host, ibx.Wapi_version)
	payload := fmt.Sprintf("{\"network\": \"%s\",\"network_view\": \"%s\"}", network, ibx.Network_view)

	return ibx.create_record(payload, url)
}

func (ibx InfobloxClient) Create_srv_record(name, target string, weight, priority, port, ttl int) (string, error) {

	url := fmt.Sprintf("%s/wapi/v%s/record:srv", ibx.Host, ibx.Wapi_version)

	ttl_data := ""
	if ttl >= 0 {
		ttl_data = fmt.Sprintf(",\"ttl\": %d, \"use_ttl\": %t", ttl, true)
	}

	payload := fmt.Sprintf("{\"name\": \"%s\", \"target\": \"%s\", \"view\": \"%s\", \"weight\": %d, \"priority\": %d, \"port\": %d  %s}", name, target, ibx.View, weight, priority, port, ttl_data)

	return ibx.create_record(payload, url)
}

//Create_host_record creates a host record, that is A record and PTR record.
//It returns the response from Infoblox or an error.
//  paramaters are:
//  fqdn to be created
//  address - ipv4 address to be used
//  ttl - ttl value to be used for the record in seconds. Use a negative value to leave as default
func (ibx InfobloxClient) Create_host_record(fqdn string, address string, ttl float64) (string, error) {

	url := fmt.Sprintf("%s/wapi/v%s/record:host?_return_fields=ipv4addrs", ibx.Host, ibx.Wapi_version)

	ttl_data := ""
	if ttl >= 0 {
		ttl_data = fmt.Sprintf(",\"ttl\": %.0f, \"use_ttl\": %t", ttl, true)
	}

	payload := fmt.Sprintf("{\"ipv4addrs\": [{\"configure_for_dhcp\": false,\"ipv4addr\": \"%s\"}],\"name\": \"%s\",\"view\": \"%s\" %s}", address, fqdn, ibx.View, ttl_data)

	return ibx.create_record(payload, url)
}

//Create_text_record creates a TXT record.
//parameters are the fqdn for he record and the plain text to be assoicated with it
//it returns the resonse from Infoblox or an error message if any
func (ibx InfobloxClient) Create_text_record(fqdn, text string) (string, error) {

	url := fmt.Sprintf("%s/wapi/v%s/record:txt", ibx.Host, ibx.Wapi_version)
	payload := fmt.Sprintf("{\"text\": \"%s\",\"name\": \"%s\",\"view\": \"%s\"}", text, fqdn, ibx.View)

	return ibx.create_record(payload, url)
}

func (ibx InfobloxClient) create_record(payload, url string) (string, error) {

	var postdata = []byte(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postdata))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(ibx.Username, ibx.Password)
	resp, err := ibx.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	cmdresp, staterr := read_httpresponse(resp)
	if staterr != nil {
		return "", staterr
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return "", errors.New(cmdresp)
	}
	return cmdresp, nil
}

//Create_zone_auth creates an authorative forward zone
func (ibx InfobloxClient) Create_zone_auth(name string) (string, error) {

	url := fmt.Sprintf("%s/wapi/v%s/zone_auth", ibx.Host, ibx.Wapi_version)
	payload := fmt.Sprintf("{\"create_ptr_for_hosts\": true, \"fqdn\": \"%s\", \"view\": \"%s\"}", name, ibx.View)

	return ibx.create_record(payload, url)
}

// Create_cname_record creates a CNAME record.
// The parameters are the cnonical and the name (alias) to be created.
// It returns the response from infoblox as a string and an error.
func (ibx InfobloxClient) Create_cname_record(canonical string, name string) (string, error) {

	url := fmt.Sprintf("%s/wapi/v%s/record:cname", ibx.Host, ibx.Wapi_version)
	payload := fmt.Sprintf("{\"canonical\": \"%s\",\"name\": \"%s\",\"view\": \"%s\"}", canonical, name, ibx.View)

	return ibx.create_record(payload, url)
}

func (ibx InfobloxClient) Get_all_networks(count int) ([]interface{}, error) {

	url := fmt.Sprintf("%s/wapi/v2.3/network?_max_results=%s", ibx.Host, strconv.Itoa(count))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(ibx.Username, ibx.Password)
	resp, herr := ibx.HttpClient.Do(req)
	if herr != nil {
		return nil, herr
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		errresp, staterr := read_httpresponse(resp)
		if staterr != nil {
			return nil, staterr
		} else {
			return nil, errors.New(errresp)
		}
	}

	var f interface{}
	if jerr := json.NewDecoder(resp.Body).Decode(&f); jerr != nil {
		fmt.Println(jerr)
		return nil, jerr
	}
	m := f.([]interface{})
	return m, nil
}
