package tools

import (
	"context"
	"fmt"
	"log"
	"nanpangyou/invoice-tool/structs"
	"os"
	"strings"

	"github.com/google/go-tika/tika"
	"gopkg.in/yaml.v3"
)

func readPathFromConfig() string {
	// 获取配置中的pdf源文件路径
	configFile, readConfigErr := os.ReadFile("./config/config.yaml")
	if readConfigErr != nil {
		log.Fatal("err to read file", readConfigErr)
	}
	configStruct := structs.NewConfig()
	yamlErr := yaml.Unmarshal(configFile, &configStruct)
	if yamlErr != nil {
		log.Fatal("err to yaml", yamlErr)
	}
	return configStruct.Basic.Pdf_File_Path
}
func PdfTikaReader() {
	// 循环读取pdf目录中的所有所有文件，解析并输出至excel
	// pdf路径(相对路径)
	path := readPathFromConfig()
	// 当前目录
	currentDirPath, _ := os.Getwd()
	client := tika.NewClient(nil, "http://localhost:9998")

	de, err2 := os.ReadDir(currentDirPath + path)
	if err2 != nil {
		log.Fatal("read dir err ", err2)
	}
	var demoText []string
	for _, v := range de {
		f, err := os.Open(currentDirPath + path + "/" + v.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("v.Name(): %v\n", v.Name())
		fmt.Printf("f.Name(): %v\n", f.Name())
		defer f.Close()
		s, _ := client.Parse(context.Background(), f)
		demoText = append(demoText, s)
		// demoText.append (demoText[]byte(s))
	}
	GenenrateSheet(strings.Join(demoText, ""))

}
