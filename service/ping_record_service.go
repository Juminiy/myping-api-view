package service

import (
	"strings"

	"github.com/Juminiy/myping-api-view/db"

	"github.com/Juminiy/myping/model"
)

const (
	BASE_M = 1000000
)

func PingRecordKV(keys bool) []string {
	resultList := make([]string, 0)
	iter := db.RedisClient.Scan(
		db.GlobalCentext,
		db.REDIS_PING_RECORD_CURSOR,
		db.PING_RECORD_DEFAULT_PREFIX+"*",
		db.REDIS_PING_RECORD_COUNT,
	).Iterator()
	for iter.Next(db.GlobalCentext) {
		if keys {
			resultList = append(resultList, iter.Val())
		} else {
			val := db.RedisClient.Get(db.GlobalCentext, iter.Val()).Val()
			resultList = append(resultList, val)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return resultList
}

func PingTimestampArr() []string {
	arr := PingRecordKV(true) 
	timestampArr := make([]string,0) 
//1:time.Date(2022, time.July, 3, 15, 38, 56, 43044311, time.Local)
	for _,v := range arr {
		tmpStr := v[13:44]
		tmpArrs := strings.Split( strings.TrimSpace(tmpStr),",")
		timestampArr = append(timestampArr, tmpArrs[3]+":"+tmpArrs[4]+":"+tmpArrs[5])
	}
	return timestampArr 
}

/* 
{
"packet_recv": 16,
"packet_sent": 16,
"packet_loss": 0,
"desk_url": "cctv.com",
"ip_addr": {
	"IP": "39.107.241.57",
	"Zone": ""
},
"min_rtt": 42959923,
"max_rtt": 52186652,
"avg_rtt": 44663630,
"std_dev_rtt": 2540333,
"total_time": 15094
}
*/

func StatisticsProtoKV(recordKey string, statis model.StatisticsRecord) interface{} {
	var rtResult interface {}  
	switch recordKey {
	case "PacketRecv" :{
		rtResult = statis.PacketRecv
	}
	case "PacketSent" : {
		rtResult = statis.PacketSent
	}
	case "PacketLoss" : {
		rtResult = statis.PacketLoss
	}
	case "DestURL" : {
		rtResult = statis.DestURL
	}
	case "MinRtt" : {
		rtResult = statis.MinRtt/BASE_M
	}
	case "MaxRtt" : {
		rtResult = statis.MaxRtt/BASE_M
	}
	case "AvgRtt" : {
		rtResult = statis.AvgRtt/BASE_M
	}
	case "StdDevRtt" : {
		rtResult = statis.StdDevRtt/BASE_M
	}
	default : {
		rtResult = nil
	} 
	}
	return rtResult 
}
func PingRecordArr(recordKey string) []interface{} {
	arr := PingRecordKV(false) 
	resultArr := make([]interface{},0)
	var eachStatics model.StatisticsRecord
	for _,v := range arr {
		model.GlobalJson.UnmarshalFromString(v,&eachStatics)
		// fmt.Println(eachStatics)
		resultArr = append(resultArr, StatisticsProtoKV(recordKey,eachStatics))
	}

	return resultArr
}	