package handler

import (
	"apiservertest/defs"
	"apiservertest/utils"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func XuidHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//首先获取clientid，对应示例里的abc：http://192.168.250.1:8080/xuid?clientid=abc
	clientID := r.URL.Query().Get("clientid")
	var xuid string
	var err error
	mapResult, ok := defs.GlobalConfig.ResultMap.Load(clientID)
	if ok {
		xuid = mapResult.(string)
	} else {
		if len(clientID) > 0 {
			xuid, err = utils.GenRandomXUIDByClientID(10)
			if err != nil {
				w.WriteHeader(500)
				str := fmt.Sprintf("%s", err)
				_, _ = io.WriteString(w, str)
			}
			defs.GlobalConfig.ResultMap.Store(clientID,xuid)
	} else {
			w.WriteHeader(500)
			str := fmt.Sprintf("%s", "传入的clientid为空")
			_, _ = io.WriteString(w, str)
		}
	}
	_, _ = io.WriteString(w, xuid)
}

func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		xuid, _ := utils.GenRandomXUIDByClientID(defs.GlobalConfig.XuidLength)
		_, _ = io.WriteString(w, "<h1>欢迎查询xuid<br>查询语法示例：http://" +
			r.Host +
			"/xuid?clientid=abc<br>"+
			"结果示例：" + xuid +
			"</h1>")
}