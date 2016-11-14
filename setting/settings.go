package setting

import (
	"fmt"
	"os"
	"runtime"
	"github.com/creamdog/gonfig"
)
var (
	Environment string
	DataQuery string
	ES_Index string
	ES_Host  string
	ES_Port string
	ES_UserName string
	ES_Password string
	ES_Cluster_Domains string
	Limit int
	Top int
	SignUrl string
	SecondUrl string
)

// initialize config file information
func init()  {
	var fileName string
	var err error
	pwd, _ := os.Getwd()
	if runtime.GOOS == "windows" {
		fileName = fmt.Sprintf("%s\\conf\\sysConfig.json",pwd)
	}else {
		fileName = fmt.Sprintf("%s/conf/sysConfig.json",pwd)
	}
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Sprint(err)
	}
	defer f.Close();
	config, err := gonfig.FromJson(f)
	if err != nil {
		fmt.Sprint(err)
	}
	Environment, err = config.GetString("env", "")
	DataQuery, err = config.GetString("data_query", "") //prod 10.24.35.212:7777 -- dev 192.168.174.139:8085
	Limit,err = config.GetInt("limit", "10000")
	Top, err = config.GetInt("top",2000)
	SignUrl, err = config.GetString("signUrl","http://databi.ifuli.cn:39200")
	ES_Index, err = config.GetString("elastic_search/index","")
	ES_Host,err = config.GetString("elastic_search/host","")
	ES_Port,err = config.GetString("elastic_search/port","")
	ES_UserName,err = config.GetString("elastic_search/user_name","")
	ES_Password, err = config.GetString("elastic_search/password","")
	ES_Cluster_Domains, err = config.GetString("elastic_search/cluster_domains","")
}
