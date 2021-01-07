package osutils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExecutableName() string {
	temp := os.Args[0]
	i := strings.LastIndex(temp, "\\")
	j := strings.LastIndex(temp, "/")
	if i == -1 && j == -1 {
		//fmt.Println("GetExecutableName()", temp)
		return temp
	}
	if i > j && i < (len(temp)-1) {
		//fmt.Println("GetExecutableName()2 len temp ", len(temp)-i-1)
		//fmt.Println("GetExecutableName()2", temp[i+1:len(temp)-1])
		return temp[i+1:]
	}
	if j < (len(temp) - 1) {
		return temp[j+1:]
	}
	return ""
}

func GetExecutableDir() (error, string) {
	ex, err := os.Executable()
	if err != nil {
		return err, ""
	}
	exPath := filepath.Dir(ex)
	return nil, exPath
}
