package main

import(
	"github.com/tcncloud/protoc-gen-persist/parse-example/service"
)
type SMSGroupMstr struct {
	SmsGroupSid    int64
	LoginSid       int64
	ClientSid      int64
	CountrySid     int64
	Name           string
	SmsMessageSid  int64
	//CallerIds      []string
	//StartTime      *google_protobuf.Timestamp
	//StopTime       *google_protobuf.Timestamp
	SendsPerMinute int32
	Status         int32
	TotalCost      float64
}

func main() {
	smsGroupMstr := &SMSGroupMstr {
		SmsGroupSid: 1,
		LoginSid: 2,
		ClientSid: 3,
		CountrySid: 4,
		Name: "name",
		SmsMessageSid: 5,
		SendsPerMinute: 1,
		Status: 2,
		TotalCost: 1.1,
	}
	service.ParseStruct(smsGroupMstr)

}
