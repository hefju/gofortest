package jutil

import (
	"log"
	"strings"
)

var IniConfiger map[string]map[string]string //存储Ini配置的数据

func ParseIni(filename string) {
	IniConfiger = make(map[string]map[string]string)

	ReadLine(filename, handle) //从指定文件中逐行读取字符串, 第二参数是每行字符的处理方法
}

var section string

func handle(line string) {
	length := len(line)
	switch {
	case length == 0: //没有字符不处理
	case line[0] == ';': //注释不处理

	case line[0] == '[' && line[length-1] == ']': //处理section
		section = line[1 : length-1]
		IniConfiger[section] = make(map[string]string)

	default: //处理key-value
		i := strings.Index(line, "=")
		key := strings.TrimSpace(line[0:i])
		value := strings.TrimSpace(line[i+1 : length])
		if len(section) == 0 {
			log.Fatalln("IniConfig.go, section is empty.")
		}
		IniConfiger[section][key] = value
	}
}
