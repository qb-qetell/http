package http

import "fmt"
import "testing"

func Test_ (t *testing.T) {
	rqstt := `POST / HTTP/1.1
Content-Type: application/json
User-Agent: PostmanRuntime/7.29.2
Accept: */*
Postman-Token: cd8fdf02-cf13-455b-836d-3a9fa168830f
Host: 127.0.0.1:10001
Accept-Encoding: gzip, deflate, br
Content-Length: 19
Content-Length: 24

Hey man. Come back!`
	_ba00, _bb00 := Parse ([]byte (rqstt))
	if _ba00 != nil {
		fmt.Println (_ba00.Error ())
		return
	}
	fmt.Println (">>>> mthd:", _bb00.Mthd)
	fmt.Println (">>>> path:", _bb00.Path)
	fmt.Println (">>>> vrsn:", _bb00.Vrsn)
	fmt.Println (">>>> hdrr:", _bb00.Hdrr)
	fmt.Println (">>>> vlll:", _bb00.HdrrVlll)
	fmt.Println (">>>> core:", string (_bb00.Core))
}
