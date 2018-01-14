package idcard

import (
	"io/ioutil"
	"strings"
	"fmt"
	"github.com/murlokswarm/errors"
	"strconv"
)

type IdCardInfo struct {
	Area string
	Sex string
	Birth string
	IdNo string
}

func ParseAreaFromFile(file string)  error{
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content),"\n")
	for _,line := range lines {
		if len(line) == 0 {
			continue
		}
		fields := strings.Split(line," ")
		if len(fields) != 2 {
			fmt.Printf("line :%s filed:%d\n", line, len(fields))
			continue
		}
		areaCode := fields[0]
		areaName := fields[1]
		fmt.Printf("areaMap[\"%s\"] = \"%s\"\n", areaCode, areaName)
	}
	return nil
}


func Parse(idno string) (*IdCardInfo, error) {
	if len(idno) != 18{
		return nil, errors.New("身份证号长度错误")
	}

	idinfo:= new(IdCardInfo)
	areaCode := idno[0:6]
	birth := idno[6:14]
	sex := idno[16:17]

	areaMaplock.Lock()
	if area,ok := areaMap[areaCode];ok {
		idinfo.Area = area
	} else {
		if area,ok := areaMap[areaCode[0:2] + "0000"];ok {
			idinfo.Area = area
		}
	}

	idinfo.Birth = birth[0:4] + "-" + birth[4:6] + "-" + birth[6:8]
	sexnum, _ := strconv.Atoi(sex)
	if sexnum % 2 ==0 {
		idinfo.Sex = "女"
	} else {
		idinfo.Sex = "男"
	}

	idinfo.IdNo = idno
	areaMaplock.Unlock()

	return idinfo,nil

}