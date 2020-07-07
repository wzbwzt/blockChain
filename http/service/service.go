package service

import (
	"blockChain/block"
	"blockChain/chain"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	hostname string
	port int
	firstNode  *chain.Node
)


func init(){
	flag.StringVar(&hostname,"hostname","127.0.0.1","service IP or hostname")
	flag.IntVar(&port,"port",8001,"service listen port")
	firstBlock := block.CreateFirstBlock("创世区块")
	firstNode = chain.CreateFirstNode(&firstBlock)
}

func GetListHandle(writer http.ResponseWriter, request *http.Request) {
	list := chain.ShowNodeList(firstNode)
	writer.Header().Set("Content-Type","application/json")
	indent, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found"))
		log.Println("json.MarshalIndent",err)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(indent)
}

func PostAddHandle(writer http.ResponseWriter, request *http.Request) {
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Panic("ioutil.readall:",err)
	}
	defer request.Body.Close()
	var req map[string]interface{}
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		log.Panic("json.Unmarshal failed;err:",err)
	}
	data:=req["data"].(string)
	pre:=firstNode
	for pre.NextNode!=nil{
		pre=pre.NextNode
	}
	newBlock := block.CreatNewBlock(*pre.Data, data)
	newnode := chain.Add(&newBlock, pre)
	list := chain.ShowNodeList(firstNode)
	rep:=make(map[string]interface{})
	if newnode == nil {
		rep["status"]="add failed!!!"
		log.Println("Add new node failed")
	}
	log.Println("Add new node success")
	rep["status"]="add success!!!"
	rep["result"]=list

	writer.Header().Set("content-Type","application/json")
	indent, err:= json.MarshalIndent(rep, "", "\t")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found!!!"))
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(indent)
}
func Run(){
	flag.Parse()
	address := fmt.Sprintf("%s:%v", hostname, port)
	log.Println("REST service on",address)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/blockChain/list",GetListHandle).Methods("GET")
	router.HandleFunc("/blockChain/add",PostAddHandle).Methods("POST")
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Println("service listen err:",err)
	}
}



