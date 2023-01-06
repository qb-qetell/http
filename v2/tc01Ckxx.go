package  http

import  "fmt"
import  "time"

func CKXX_Estb (name string) (*Ckxx) {
	return &Ckxx {
		Name : name ,
		Scrt : false,
		HOXX : false,
		SSXX : false,
	}
}

type Ckxx struct {
	Dmnx string
	Path string
	Scrt   bool
	Name string
	Vlxx string
	ETXX time.Duration
	HOXX   bool
	SSXX   bool
}
	func (s *Ckxx) SetxDmnx (d        string) (error) { s.Dmnx =    d; return nil }
	func (s *Ckxx) SetxPath (d        string) (error) { s.Path =    d; return nil }
	func (s *Ckxx) EnfrScrt (               ) (error) { s.Scrt = true; return nil }
	func (s *Ckxx) SetxVlxx (d        string) (error) { s.Vlxx =    d; return nil }
	func (s *Ckxx) SetxETXX (d time.Duration) (error) { s.ETXX =    d; return nil }
	func (s *Ckxx) EnfrHOXX (               ) (error) { s.HOXX = true; return nil }
	func (s *Ckxx) EnfrSSXX (               ) (error) { s.SSXX = true; return nil }
	func (s *Ckxx) Lqfy ( ) (error ,  string) {
		_ba00     :=   ""
		/*--1--*/
		if s.Dmnx !=   "" {
			_ba00 = fmt.Sprintf ("Domain=%s"  , s.Dmnx)
		}
		/*--1--*/
		if s.Path !=   "" {
			_ba00 = fmt.Sprintf ("%s; Path=%s",  _ba00, s.Path)
		}
		/*--1--*/
		if s.Scrt == true {
			_ba00 = fmt.Sprintf ("%s; Secure" ,  _ba00)
		}
		/*--1--*/
		if  _ba00 !=   "" {_ba00 = "; " + _ba00}
		_ba00  = fmt.Sprintf ("%s=%s%s",  s.Name  , s.Vlxx,  _ba00)
		/*--1--*/
		if s.ETXX !=    0 {
			_ca00 := time.Now ().In (time.FixedZone("+0000", 0)).Add(s.ETXX).Format (
				"Mon, 02 Jan 2006 03:04:05 GMT",
			)
			_ba00  = fmt.Sprintf ("%s; Expires=%s",  _ba00, _ca00)
		}
		/*--1--*/
		if s.HOXX == true {
			_ba00  = fmt.Sprintf ("%s; HttpOnly"  ,  _ba00)
		}
		/*--1--*/
		if s.SSXX == true {
			_ba00  = fmt.Sprintf ("%s; SameSite"  ,  _ba00)
		}
		/*--1--*/
		return nil, _ba00		
	}
