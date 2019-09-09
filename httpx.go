package httpx

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func Get(url string, dest interface{}) (err error) {
	rawBytes, err := GetRaw(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawBytes, dest)
	if err != nil {
		return errors.New("Unmarshal " + string(rawBytes) + " errMsg: " + err.Error())
	}

	return nil
}

func GetRaw(url string) ([]byte, error) {
	getResp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer getResp.Body.Close()
	if getResp.StatusCode != 200 {
		return nil, errors.New("StatusCode is " + strconv.Itoa(getResp.StatusCode))
	}
	getRespBody, err := ioutil.ReadAll(getResp.Body)
	if err != nil {
		return nil, errors.New("read resp body errMsg: " + err.Error())
	}

	return getRespBody, nil
}

func PostForm(url string, data url.Values, dest interface{}) error {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("status code is " + strconv.Itoa(resp.StatusCode))
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, dest)
	if err != nil {
		return err
	}

	return nil
}
