package http

import "errors"
import "fmt"
import "regexp"
import "strings"

func RQST_Sdfy (data []byte) (error, *Rqst) {
	if len (data) == 0 {
		return errors.New ("BA00: Data provided is empty."), nil
	}
	if len (data) >  2147483647 {
		return errors.New ("BB00: Data provided is too long."), nil
	}
	/*--1--*/
	extn  := 0
	iBA00 := 0
	for iBA00 = 1;   iBA00 <= len (data); iBA00 ++ {
		extn = iBA00
		if data [iBA00 - 1] == '\n' {
			break
		}
	}
	prtsBB00 := strings.Split (string (data [:extn]), " ")
	if len (prtsBB00) != 3 {
		_ca00 := fmt.Sprintf ("BC00: Request '%s' is invalid.",
			string (data [:extn - 1]),
		)
		return errors.New (_ca00), nil
	}
	rqst := &Rqst {
		Hdrr:     make ([]string,       0),
		HdrrVlxx: make (map[string]string),
		Ckxx:     make ([]string,       0),
		CkxxVlxx: make (map[string]string),
	}
	rqst.Mthd = prtsBB00 [0]
	rqst.Path = prtsBB00 [1]
	rqst.Vrsn = prtsBB00 [2]
	rqst.Vrsn = strings.ReplaceAll (rqst.Vrsn, "HTTP/", "")
	rqst.Vrsn = strings.ReplaceAll (rqst.Vrsn, "\r", "")
	rqst.Vrsn = strings.ReplaceAll (rqst.Vrsn, "\n", "")
	if extn == len (data) {
		return nil, rqst
	}
	/*--1--*/
	more := true
	for more == true {
		lastExtn := extn
		iCA00 := 0
		for iCA00 = extn + 1; iCA00 <= len (data); iCA00 ++ {
			extn = iCA00
			if data [iCA00 - 1] == '\n' {
				break
			}
		}
		_ca50 := string (data [lastExtn:extn])
		_ca50  = strings.Trim (_ca50, " 	")
		_ca50  = strings.ReplaceAll (_ca50, "\r", "")
		_ca50  = strings.ReplaceAll (_ca50, "\n", "")
		if _ca50 == "" {
			break
		}
		prtsCB00 := strings.SplitN (_ca50, ": ", 2)
		if len (prtsCB00) < 2 {
			_da00 := fmt.Sprintf ("An header '%s' is invalid.",
				string (data [lastExtn:extn]),
			)
			return errors.New (_da00), nil
		}
		rqst.Hdrr = append (rqst.Hdrr, prtsCB00 [0])
		rqst.HdrrVlxx [prtsCB00 [0]] = strings.ReplaceAll (prtsCB00 [1], "\r", "")
		rqst.HdrrVlxx [prtsCB00 [0]] = strings.ReplaceAll (prtsCB00 [1], "\n", "")
		/*--2--*/
		if (extn + 0) == len (data) {
			more = false
			continue
		}
		if (data [(extn + 1) - 1] == '\n') {
			more = false
			extn = extn + 1
			continue
		}
		if (extn + 1) == len (data) {
			more = false
			extn = extn + 1
			continue
		}
		more = true
	}
	/*--1--*/
	if extn < len (data) {
		rqst.Core = data [  extn: ]
	}
	/*--1--*/
	_bc00 := ""
	for _ , _be00 := range rqst.Hdrr  {
		_ca00 := strings.ToLower  (_be00)
		if _ca00 == "cookie" {
			_bc00 = _be00
			break
		}
	}
	if _bc00 != "" && rqst.HdrrVlxx [_bc00] != "" {
		_ca00 := regexp.MustCompile (`; *`).Split (rqst.HdrrVlxx [_bc00], -1)
		for _ , _cb00 := range _ca00 {
			_da00 := strings.Split (_cb00, "=")
			 rqst.Ckxx = append (rqst.Ckxx, _da00 [0])
			 rqst.CkxxVlxx [_da00 [0]]  =   _da00 [1]
		}
	}
	/*--1--*/
	return nil, rqst
}
type Rqst struct {
	Mthd     string
	Path     string
	Vrsn     string
	Hdrr     []string
	HdrrVlxx map[string]string
	Core     []byte

	Ckxx     []string
	CkxxVlxx map[string]string
}
