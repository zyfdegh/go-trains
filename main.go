// LogTest project main.go
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
)

const ParamID = "_id"

func main() {

	//	seeDifference()
	//	documentLocationDemo()
	//	multipleHeader()
	//	loopPointerArray()
	getSite()
}

func multipleHeader() {
	client := &http.Client{}
	url := "http://baidu.com"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", "1123")
	req.Header.Set("X-Auth-Token", "1123")
	req.Header.Set("X-Auth-Token", "1123")
	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()
}

func documentLocationDemo() {
	url := "http://192.168.10.198:10002/v1/cluster/569c9038e138239fa125d453"
	req := &restful.Request{}
	req.Request, _ = http.NewRequest("GET", url, nil)

	id := "569c9038e138239fa125d453"

	ret := documentLocation(req, id)
	fmt.Println(ret)

}

//
// Return document location URL
//
func documentLocation(req *restful.Request, id string) (location string) {
	// Get current location url
	location = strings.TrimRight(req.Request.URL.RequestURI(), "/")

	// Remove id from current location url if any
	if reqId := req.PathParameter(ParamID); reqId != "" {
		logrus.Infof("reqId found")
		idlen := len(reqId)
		// If id is in current location remove it
		if noid := len(location) - idlen; noid > 0 {
			if id := location[noid : noid+idlen]; id == reqId {
				location = location[:noid]
			}
		}
		location = strings.TrimRight(location, "/")
	}

	// Add id of the document
	return location + "/" + id
}

func getSite() {
	var url = "http://www.linkernetworks.com"
	req, _ := http.NewRequest("GET", url, nil)
	client := http.DefaultClient
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	logrus.Debugf("GET %s,status code:%s", url, resp.StatusCode)
	fmt.Println(resp.Status)

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Xiaoy", Age: 23}
	fmt.Println(p)
}

func seeDifference() {
	req := &restful.Request{}
	req.Request, _ = http.NewRequest("GET", "http://192.168.10.198:10002/v1/cluster/569c9038e138239fa125d453/hosts", nil)
	fmt.Println(req.Request.URL.String())
	fmt.Println(req.Request.URL.RequestURI())
}

func printInOneLine() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("\rnum:%d", i)
	}

	for i := 0; i < 10; i++ {
		//update this line
		fmt.Printf("\r Count,%d s", i)
		time.Sleep(1 * time.Second)
	}
}

func switchFallThrough() {
	switch 2 {
	//no automatic fall through
	case 1:
	case 2:
	case 3:
		fmt.Println("123")
	case 4:
	case 5:
		fmt.Println("45")
	default:
		fmt.Println("default")
	}

	switch 2 {
	//no automatic fall through
	case 1, 2, 3:
		fmt.Println("123")
	case 4:
	case 5:
		fmt.Println("45")
	default:
		fmt.Println("default")
	}
}

func replaceString() {
	fmt.Println("<Some value>")
	var body = []byte(`some string with <VAL> to <VAL> replace`)

	var bs = []byte("VALUE")
	var body2 = bytes.Replace(body, []byte(`<VAL>`), bs, -1)
	fmt.Println(string(body2))
}

func timeFormat() {
	//time test
	var t_str string = time.Now().Format(time.RFC3339)
	var t time.Time
	t, err := time.Parse(time.RFC3339, t_str)
	if err != nil {
		logrus.Errorln("error parse time.%v", err)
	}
	fmt.Println(t_str)
	fmt.Println(t.Day())

	//	t2 := time.Parse(time.RFC3339, time.Now())
	//	fmt.Println(time.Now().Day())
	//	fmt.Println(time.Now().Format()

	fmt.Println(time.Now().Format(time.RFC3339))
	t2, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	//	fmt.Println()
}
