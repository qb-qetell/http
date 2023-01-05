package http

import "fmt"

func RSPN_Estb () (*Rspn) {
	return &Rspn {
		Hdrx: make ([][]string, 0),
		Core: make ([]byte    , 0),
	}
}

type Rspn struct {
	Vrsn     string
	Code     string
	Note     string
	Hdrx     [][]string
	Core     []byte
}
	// NAME: Insert Header
	func (sbjc *Rspn) InsrHdrr (name, vlxx   string) {
		sbjc.Hdrx = append (
			sbjc.Hdrx  ,
			[]string   {name, vlxx},
		)
	}
	func (sbjc *Rspn) Lqfy () []byte {
		if sbjc.Vrsn == "" {
			sbjc.Vrsn = "1.1"
		}
		if sbjc.Code == "" {
			sbjc.Code = "200"
		}
		if sbjc.Note == "" {
			trnsTbll := map[string]string {
				"100": "Continue",
				"101": "Switching Protocols",
				"102": "Processing",
				"103": "Early Hints",
				
				"200": "OK",
				"201": "Created",
				"202": "Accepted",
				"203": "Non Authoritative Info",
				"204": "No Content",
				"205": "Reset Content",
				"206": "Partial Content",
				"207": "Multi Status",
				"208": "Already Reported",
				"226": "IM Used",
				
				"300": "Multiple Choices",
				"301": "Moved Permanently",
				"302": "Found",
				"303": "See Other",
				"304": "Not Modified",
				"305": "Use Proxy",
				"307": "Temporary Redirect",
				"308": "Permanent Redirect",
				
				"400": "Bad Request",
				"401": "Unauthorized",
				"402": "Payment Required",
				"403": "Forbidden",
				"404": "Not Found",
				"405": "Method Not Allowed",
				"406": "Not Acceptable",
				"407": "Proxy Auth Required",
				"408": "Request Timeout",
				"409": "Conflict",
				"410": "Gone",
				"411": "Length Required",
				"412": "Precondition Failed",
				"413": "Request Entity Too Large",
				"414": "Request URI Too Long",
				"415": "Unsupported Media Type",
				"416": "Requested Range Not Satisfiable",
				"417": "Expectation Failed",
				"418": "Teapot",
				"421": "Misdirected Request",
				"422": "Unprocessable Entity",
				"423": "Locked",
				"424": "Failed Dependency",
				"425": "Too Early",
				"426": "Upgrade Required",
				"428": "Precondition Required",
				"429": "Too Many Requests",
				"431": "Request Header Fields Too Large",
				"451": "Unavailable For Legal Reasons",
				
				"500": "Internal Server Error",
				"501": "Not Implemented",
				"502": "Bad Gateway",
				"503": "Service Unavailable",
				"504": "Gateway Timeout",
				"505": "HTTP Version Not Supported",
				"506": "Variant Also Negotiates",
				"507": "Insufficient Storage",
				"508": "Loop Detected",
				"510": "Not Extended",
				"511": "Network Authentication Required",
			}
			sbjc.Note = trnsTbll [sbjc.Code]
			if sbjc.Note == "" {sbjc.Note = sbjc.Code}
		}
		mssg := fmt.Sprintf ("HTTP/%s %s %s" , sbjc.Vrsn, sbjc.Code, sbjc.Note)
		for _, _bb00 := range sbjc.Hdrx {
			mssg = fmt.Sprintf ("%s\r\n%s: %s", mssg, _bb00 [0], _bb00 [1])
		}
		mss2 := []byte (mssg)
		if len (sbjc.Core) != 0 {
			mss2 = append (mss2, []byte ("\r\n\r\n")...)
			mss2 = append (mss2, sbjc.Core...)
		}
		return mss2
	}
