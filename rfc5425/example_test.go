package rfc5425

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func output(out interface{}) {
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Dump(out)
}

func Example() {
	r := strings.NewReader("48 <1>1 2003-10-11T22:14:15.003Z host.local - - - -25 <3>1 - host.local - - - -38 <2>1 - host.local su - - - κόσμε")
	output(NewParser(r, WithBestEffort()).Parse())
	// Output:
	// ([]rfc5425.Result) (len=3) {
	//  (rfc5425.Result) {
	//   Message: (*rfc5424.SyslogMessage)({
	//    Priority: (*uint8)(1),
	//    facility: (*uint8)(0),
	//    severity: (*uint8)(1),
	//    Version: (uint16) 1,
	//    Timestamp: (*time.Time)(2003-10-11 22:14:15.003 +0000 UTC),
	//    Hostname: (*string)((len=10) "host.local"),
	//    Appname: (*string)(<nil>),
	//    ProcID: (*string)(<nil>),
	//    MsgID: (*string)(<nil>),
	//    StructuredData: (*map[string]map[string]string)(<nil>),
	//    Message: (*string)(<nil>)
	//   }),
	//   MessageError: (error) <nil>,
	//   Error: (error) <nil>
	//  },
	//  (rfc5425.Result) {
	//   Message: (*rfc5424.SyslogMessage)({
	//    Priority: (*uint8)(3),
	//    facility: (*uint8)(0),
	//    severity: (*uint8)(3),
	//    Version: (uint16) 1,
	//    Timestamp: (*time.Time)(<nil>),
	//    Hostname: (*string)((len=10) "host.local"),
	//    Appname: (*string)(<nil>),
	//    ProcID: (*string)(<nil>),
	//    MsgID: (*string)(<nil>),
	//    StructuredData: (*map[string]map[string]string)(<nil>),
	//    Message: (*string)(<nil>)
	//   }),
	//   MessageError: (error) <nil>,
	//   Error: (error) <nil>
	//  },
	//  (rfc5425.Result) {
	//   Message: (*rfc5424.SyslogMessage)({
	//    Priority: (*uint8)(2),
	//    facility: (*uint8)(0),
	//    severity: (*uint8)(2),
	//    Version: (uint16) 1,
	//    Timestamp: (*time.Time)(<nil>),
	//    Hostname: (*string)((len=10) "host.local"),
	//    Appname: (*string)((len=2) "su"),
	//    ProcID: (*string)(<nil>),
	//    MsgID: (*string)(<nil>),
	//    StructuredData: (*map[string]map[string]string)(<nil>),
	//    Message: (*string)((len=11) "κόσμε")
	//   }),
	//   MessageError: (error) <nil>,
	//   Error: (error) <nil>
	//  }
	// }
}

func Example_handler() {
	// Defining a custom handler to send the parsing results into a channel
	EmittingParsing := func(p *Parser) chan Result {
		c := make(chan Result)

		toChannel := func(r *Result) {
			c <- *r
		}

		go func() {
			defer close(c)
			p.ParseExecuting(toChannel)
		}()

		return c
	}

	// Use it
	r := strings.NewReader("16 <1>1 - - - - - -17 <2>12 A B C D E -16 <1>1")
	results := EmittingParsing(NewParser(r, WithBestEffort()))

	for r := range results {
		output(r)
	}
	// Output:
	// (rfc5425.Result) {
	//  Message: (*rfc5424.SyslogMessage)({
	//   Priority: (*uint8)(1),
	//   facility: (*uint8)(0),
	//   severity: (*uint8)(1),
	//   Version: (uint16) 1,
	//   Timestamp: (*time.Time)(<nil>),
	//   Hostname: (*string)(<nil>),
	//   Appname: (*string)(<nil>),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)(<nil>),
	//   StructuredData: (*map[string]map[string]string)(<nil>),
	//   Message: (*string)(<nil>)
	//  }),
	//  MessageError: (error) <nil>,
	//  Error: (error) <nil>
	// }
	// (rfc5425.Result) {
	//  Message: (*rfc5424.SyslogMessage)({
	//   Priority: (*uint8)(2),
	//   facility: (*uint8)(0),
	//   severity: (*uint8)(2),
	//   Version: (uint16) 12,
	//   Timestamp: (*time.Time)(<nil>),
	//   Hostname: (*string)(<nil>),
	//   Appname: (*string)(<nil>),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)(<nil>),
	//   StructuredData: (*map[string]map[string]string)(<nil>),
	//   Message: (*string)(<nil>)
	//  }),
	//  MessageError: (*errors.errorString)(expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 6]),
	//  Error: (error) <nil>
	// }
	// (rfc5425.Result) {
	//  Message: (*rfc5424.SyslogMessage)({
	//   Priority: (*uint8)(1),
	//   facility: (*uint8)(0),
	//   severity: (*uint8)(1),
	//   Version: (uint16) 1,
	//   Timestamp: (*time.Time)(<nil>),
	//   Hostname: (*string)(<nil>),
	//   Appname: (*string)(<nil>),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)(<nil>),
	//   StructuredData: (*map[string]map[string]string)(<nil>),
	//   Message: (*string)(<nil>)
	//  }),
	//  MessageError: (*errors.errorString)(parsing error [col 4]),
	//  Error: (*errors.errorString)(found EOF after "<1>1", expecting a SYSLOGMSG containing 16 octets)
	// }
}
