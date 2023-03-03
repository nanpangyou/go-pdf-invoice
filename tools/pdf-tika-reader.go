package tools

import (
	"context"
	"fmt"
	"log"
	"nanpangyou/invoice-tool/structs"
	"os"
	"regexp"
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
		if strings.HasSuffix(v.Name(), ".pdf") {
			f, err := os.Open(currentDirPath + path + "/" + v.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("v.Name(): %v\n", v.Name())
			fmt.Printf("f.Name(): %v\n", f.Name())
			defer f.Close()
			s, _ := client.Parse(context.Background(), f)
			// fmt.Printf("s: %v\n", s)
			// fmt.Printf("demoText: %v\n", demoText)
			c := strings.ReplaceAll(s, " ", "")
			// fmt.Printf("c: %v\n", c)
			c = strings.ReplaceAll(c, "\n", "")
			c = strings.ReplaceAll(c, "\t", "")
			r := regexp.MustCompile("<body>(.*?)</body>")
			s1 := r.FindAllString(c, -1)
			r2 := regexp.MustCompile("<p>(.*?)</p>")
			s3 := r2.FindAllString(strings.Join(s1, ""), -1)

			// fmt.Printf("s3: %v\n", s3)
			// fmt.Printf("s2: %v\n", s3)
			// fmt.Printf("len(s3): %v\n", len(s3))
			// fmt.Printf("s2[len(s2)-6:]: %v\n", s3[len(s3)-6:])
			demoText = append(demoText, strings.Join(s3[len(s3)-10:], ""))
		}
	}
	// fmt.Printf("demoText[:]: %v\n", )
	GenenrateSheet(strings.Join(demoText, ""))

}
