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
	_ba00, _bb00 := RQST_Sdfy ([]byte (rqstt))
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
	/*--1--*/
	_ca00 :=  RSPN_Estb ()
	_ca00.Vrsn = "2.0"
	_ca00.Code = "255"
	_ca00.InsrHdrr ("name0001", "vlll0001")
	_ca00.InsrHdrr ("name0001", "vlll000x")
	_ca00.InsrHdrr ("name0002", "vlll0002")
	_ca00.InsrHdrr ("name0003", "vlll0003")
	_ca00.Core = []byte ("Hello world!!!!")
	_cb00 := _ca00.Lqfy ()
	fmt.Println ("")
	fmt.Println ("<<<< mssg:", string (_cb00))
	/*--1--*/
	_da00, _db00 := RQST_Sdfy (_cb00)
	if _da00 != nil {
		fmt.Println (_da00.Error ())
		return
	}
	fmt.Println (">>>> mthd:", _db00.Mthd)
	fmt.Println (">>>> path:", _db00.Path)
	fmt.Println (">>>> vrsn:", _db00.Vrsn)
	fmt.Println (">>>> hdrr:", _db00.Hdrr)
	fmt.Println (">>>> vlll:", _db00.HdrrVlll)
	fmt.Println (">>>> core:", string (_db00.Core))
}
