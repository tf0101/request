package request

import(
	"time"
    "net/http"
)

type Request_paramater struct{
	Header *Header_list
}

type Header_list struct{
	Header_json *map[string]string
}

func Request(url string, rp *Request_paramater)(*http.Response,error){

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,err
	}

	//set header
	if rp.Header != nil{
		for key,value := range *rp.Header.Header_json{
			req.Header.Set(key,value)
		}
	}

	//timeout 90 second
	timeout := time.Duration(90 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

    resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}

	return resp,err
}

func Request_select(url string, rp *Request_paramater, http_method string)(*http.Response,error){

	req, err := http.NewRequest(http_method, url, nil)
	if err != nil {
		return nil,err
	}

	//set header
	if rp.Header != nil{
		for key,value := range *rp.Header.Header_json{
			req.Header.Set(key,value)
		}
	}

	//timeout 90 second
	timeout := time.Duration(90 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

    resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}

	return resp,err
}