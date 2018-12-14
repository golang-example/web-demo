package util

import (
	"io/ioutil"
	"net/http"
	. "web-demo/log"
	. "web-demo/util/threadlocal"
	"strconv"
	"bytes"
	"encoding/json"
	"io"
)

type httpClient struct {
	url 		string
	header 		map[string]string
	data 		interface{}
}

func NewHttpClient() *httpClient {
	return new(httpClient)
}

func (hClient *httpClient) Header(header map[string]string) *httpClient {
	hClient.header = header
	return hClient
}

func (hClient *httpClient) Param(data interface{}) *httpClient {
	hClient.data = data
	return hClient
}

func (hClient *httpClient) Url(url string) *httpClient {
	hClient.url = url
	return hClient
}

func (hClient *httpClient) Get() []byte {
	var b []byte = nil
	var err1 error = nil
	var data io.Reader = nil
	if hClient.data != nil {
		b, err1 = json.Marshal(hClient.data)
		if err1 != nil {
			Log.Error(err1.Error())
			return nil
		}
		data = bytes.NewReader(b)
	}

	InnerRRLog.Info("GetUrl:" + hClient.url)
	InnerRRLog.Info("GetParam:" + string(b))

	req, err2 := http.NewRequest("GET", hClient.url, data)
	if err2 != nil {
		Log.Error(err2.Error())
		return nil
	}

	if len(hClient.header) != 0 {
		for key, value := range hClient.header {
			req.Header.Add(key, value)
		}
	}

	setResquestChainHeader(req)

	client := &http.Client{}
	resp, err3 := client.Do(req)
	if err3 != nil {
		Log.Error(err3.Error())
		return nil
	}
	defer resp.Body.Close()

	InnerRRLog.Info("status:" + resp.Status)
	InnerRRLog.Info("status code:" + strconv.Itoa(resp.StatusCode))

	if resp.StatusCode != 200 {
		Log.Error(strconv.Itoa(resp.StatusCode))
		Log.Error("%s", resp.Body)
		return nil
	}

	body, err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		Log.Error(err4.Error())
		return nil
	}

	InnerRRLog.Info("response:" + string(body))

	return body
}

func (hClient *httpClient) Post() []byte {
	var b []byte = nil
	var err1 error = nil
	var data io.Reader = nil
	if hClient.data != nil {
		b, err1 = json.Marshal(hClient.data)
		if err1 != nil {
			Log.Error(err1.Error())
			return nil
		}
		data = bytes.NewReader(b)
	}

	InnerRRLog.Info("PostUrl:" + hClient.url)
	InnerRRLog.Info("PostParam:" + string(b))

	req, err2 := http.NewRequest("POST", hClient.url, data)
	if err2 != nil {
		Log.Error(err2.Error())
		return nil
	}

	if len(hClient.header) != 0 {
		for key, value := range hClient.header {
			req.Header.Add(key, value)
		}
	}

	setResquestChainHeader(req)

	client := &http.Client{}
	resp, err3 := client.Do(req)
	if err3 != nil {
		Log.Error(err3.Error())
		return nil
	}
	defer resp.Body.Close()

	InnerRRLog.Info("status:" + resp.Status)
	InnerRRLog.Info("status code:" + strconv.Itoa(resp.StatusCode))

	if resp.StatusCode != 200 {
		Log.Error(strconv.Itoa(resp.StatusCode))
		Log.Error("%s", resp.Body)
		return nil
	}

	body, err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		Log.Error(err4.Error())
		return nil
	}

	InnerRRLog.Info("response:" + string(body))

	return body
}

func setResquestChainHeader(req *http.Request)  {
	if rid, ok := Mgr.GetValue(Rid); ok {
		req.Header.Add("X-Web-Demo-RequestId", rid.(string))
	}
}
