#Rextro
[Rextro](https://github.com/varver/rextro) is a simple api consumer in golang (go) , make api calls easily using this library.
Easily read json response values like : **response.success** if response looks like : 
**{"response" : {"success" : true, "error" : "none" , "code" : "200"}}**

See examples below for better understanding and usage.


#Installation
```
go get -u -v github.com/varver/rextro
```

#How to Use ?

For example lets say you want to use this api : 
https://www.mashape.com/imagevision/nudity-recognition-nudity-filter-for-images

**The curl request looks like this**
```
curl -X POST --include 'https://nuditysearch.p.mashape.com/nuditySearch/image' \
  -H 'X-Mashape-Key: ad40Z3nH9NmshlvepLaKhd0oDfTnp1GrSPNjsn7C8MsJAum8J6' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Accept: application/json' \
  -d 'objecturl=http://i.imgur.com/4hGni1I.jpg' \
  -d 'setting=3'
```

**Json Response looks like this**
```
{
  "version": "config_v8_2014-01-31",
  "message": "SUCCESS",
  "transactionid": "1234",
  "objecturl": "http://i.imgur.com/4hGni1I.jpg",
  "setting": 2,
  "score": 94,
  "classification": "NUDITY"
}
```

**This is how the code will look like to make above call**
```
package main

import (
	"fmt"
	"github.com/varver/rextro"
)

func Mashape(image string, api_key string) {
	req := rextro.NewTequest("https://nuditysearch.p.mashape.com/nuditySearch/image")

	//headers to be sent
	req.Headers["X-Mashape-Key"] = api_key
	req.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	req.Headers["Accept"] = "application/json"

	// parameters to be sent
	req.Body["setting"] = "2"
	req.Body["objecturl"] = image

	// get response as json converted // you can also make 'GET' request instead of 'POST'
	container, err := req.FetchJson("POST")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(container.String())

	/*
		// use this to get raw response as byte array
		// byteArray , err := req.Fetch("POST")

		// OR

		// use this to get response as a string
		// respString , err := req.FetchString("POST")

		// The Request is internally done using https://godoc.org/net/http#NewRequest
	*/

	//////////////////////////////////////////////////////////////////
	/// Lets fetch json values from Response looks like this ////////
	/////////////////////////////////////////////////////////////////
	/*
		{
		  "version": "config_v8_2014-01-31",
		  "message": "SUCCESS",
		  "transactionid": "1234",
		  "objecturl": "http://i.imgur.com/4hGni1I.jpg",
		  "setting": 2,
		  "score": 94,
		  "classification": "NUDITY"
		}
	*/

	message := container.Path("message").Data().(string)
	if message == "SUCCESS" {
		fmt.Println(container.Path("classification").Data().(string))
	} else {
		fmt.Print("Something went wrong")
	}

}

func main() {
	key := "YOUR_API_KEY"
	image_url := "http://i.imgur.com/4hGni1I.jpg"
	Mashape(image_url, key)
}

```

#Note
The package internally uses : https://github.com/Jeffail/gabs to parse json , you can refer its godoc here : https://godoc.org/github.com/Jeffail/gabs to do more stuff on FetchJson output , as it returns the <a href="https://godoc.org/github.com/Jeffail/gabs#Container">Container</a> of gabs package. 


