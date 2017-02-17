package parse

import(
	//"database/sql"

	//pq "github.com/lib/pq"
	//"github.com/golang/protobuf/ptypes/timestamp"
	//"github.com/jmoiron/sqlx/reflectx"
)

struct SMSGroupMstr {
	SmsGroupSid    int64
	LoginSid       int64
	ClientSid      int64
	CountrySid     int64
	Name           string
	SmsMessageSid  int64
	CallerIds      []string
	StartTime      *google_protobuf.Timestamp
	StopTime       *google_protobuf.Timestamp
	SendsPerMinute int32
	Status         int32
	TotalCost      float64
}

/*  goal of this project: take SMSGroupMstr struct,  => row in the database => smsgroupmstr struct */


