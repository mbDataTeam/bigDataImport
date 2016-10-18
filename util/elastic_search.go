package util

import (
	elastigo "github.com/mattbaird/elastigo/lib"
	"flag"
	"fmt"
	"strings"
	"encoding/json"
)
var (
	index string = "meta_data_import"
	host *string = flag.String("host", "114.55.65.41", "Elasticsearch Host")
	port string = "29200"
)

func initConnect() elastigo.Conn {
	connect := elastigo.NewConn()
	connect.Domain = *host
	connect.Port = port
	connect.Username ="awesome_bi"
	connect.Password ="bi@@#$23(SDAFDSWE!#DASFSDGSDGSDG123eghhjjhj)"
	connect.ClusterDomains = []string {"114.55.65.41"}
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
	     result:= string(byteData)
	     fmt.Print(result)
	     json.Unmarshal(byteData,&ts)
	}
	return &ts
}


