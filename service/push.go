package service

import (
	"daily/dao"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type PushContext struct {
	Id      uint
	Content string
	Lable   string
}

func PushService() (err error) {
	var pushContext []PushContext
	err = dao.DB.Table("we_chats").Select([]string{"id", "content", "label"}).Where("deleted_at is null").Order("id ASC").Scan(&pushContext).Error
	if err != nil {
		logging.Info(err)
	}
	var msg string
	for _, v := range pushContext {
		if len(v.Lable) > 0 {
			msg += "【" + v.Lable + "】"
		}
		msg += v.Content + "@" + strconv.Itoa(int(v.Id)) + "%0D%0A"
	}
	URL := fmt.Sprintf("http://121.4.81.22:8080/wecomchan?sendkey=%v&msg=%v&msg_type=text", TOKEN, msg[:len(msg)-6])
	URL = strings.Replace(URL, " ", "%20", -1)
	logging.Info(URL)
	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", URL, nil)
	//增加header选项
	//request.Header.Add("Content-Type", "application/json")
	res, _ := client.Do(request)
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logging.Info(err)
	}
	logging.Info(string(b))
	return
}
