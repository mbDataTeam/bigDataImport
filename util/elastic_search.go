package util

import (
	elastigo "github.com/mattbaird/elastigo/lib"
	"flag"
	"fmt"
	"strings"
	"encoding/json"
	"bigDataImport/setting"
)
var (
	index string = setting.ES_Index //"meta_data_import"
	host *string = flag.String("host", setting.ES_Host, "Elasticsearch Host")
	port string = setting.ES_Port
)

func initConnect() elastigo.Conn {
	connect := elastigo.NewConn()
	connect.Domain = *host
	connect.Port = port
	connect.Username = setting.ES_UserName
	connect.Password = setting.ES_Password
	connect.ClusterDomains = []string {setting.ES_Cluster_Domains}
	return *connect;
}

func QueryTableMeta(metaId string) *TableSchema  {
	con := initConnect()
	searchJson := strings.Replace(`{
	    "query" : {
	        "match": {
      			"meta_id": "$$"
      			}
	    	}
	}`,"$$", metaId, -1)
	out, err := con.Search(index, "", nil, searchJson)
	var ts TableSchema
	if(err != nil){
		fmt.Println("%s", err)
	}
	if len(out.Hits.Hits) == 1{
	     byteData,_:= out.Hits.Hits[0].Source.MarshalJSON()
	     //result:= string(byteData)
	     //fmt.Print(result)
	     json.Unmarshal(byteData,&ts)
	}
	return &ts
}


