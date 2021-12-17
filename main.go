package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 需要编码的文件夹
var srcDirPath = "demo_assets"

// 编码后的存放文件夹
var dstDirPath = "demo_gocode"

// 作者
var author = "Bearki"

// 包前缀
var packagePrefix = "github.com/bearki/assets_to_gocode"

// 全局缓存用到的包名，值表示该包是否要拼接前缀
var packageCache = map[string]bool{
	"net/http":                 false,
	"github.com/gin-gonic/gin": false,
}

// 全局缓存路由内容，最后一起写入文件
var routerContentCache string

// 第一位小数点之后的内容，将会根据这个MAP设置Content-Type
var tempExtMap = map[string][]string{
	"text/html; charset=utf-8": {
		".html",
	},
	"application/json; charset=utf-8": {
		".json",
		".map",
	},
	"application/javascript; charset=utf-8": {
		".js",
	},
	"text/css; charset=utf-8": {
		".css",
	},
	"text/plain; charset=utf-8": {
		".txt",
	},
	"image/x-icon": {
		".ico",
	},
	"image/png": {
		".png",
	},
	"image/jpeg": {
		".jpg",
		".jpeg",
	},
	"image/gif": {
		".gif",
	},
}

// 静态文件生成格式
var staticFileFormat = `
/**
 *@Title 静态文件Byte数据
 *@Desc 该文件由工具自动生成，请勿修改
 *@Author %s
 *@Time %s
 */

package %s

func %s() []byte {
	return []byte{
		%s,
	}
}`

// 路由文件格式
var routerFileFormat = `
/**
 *@Title 静态文件Byte数据
 *@Desc 该文件由工具自动生成，请勿修改
 *@Author %s
 *@Time %s
 */

package %s

import (
%s
)

// 请从外部传入一个路由分组即可
func InitRouter(r *gin.RouterGroup) {
	%s
}
`

// 路由内容格式
var routerContentFormat = `
	r.GET("%s", func(c *gin.Context) {
		c.Data(http.StatusOK, "%s", %s.%s())
	})
`

// 入口
func main() {
	rangeDir(srcDirPath)
}

// 遍历文件夹
func rangeDir(dirPath string) {
	// 打开文件夹
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	// 遍历文件夹
	for i, fi := range fileInfos {
		// 显示下进度
		fmt.Printf("正在处理: %d/%d\n", i+1, len(fileInfos))
		// 凭借完整路径
		filePath := filepath.Join(dirPath, fi.Name())
		// 判断是否是文件夹
		if fi.IsDir() {
			// 递归这个目录
			rangeDir(filePath)
		} else {
			// 编码处理这个文件
			createGoFile(filePath)
		}
	}

	// 路由文件的包名
	routerPackageName := dstDirPath[strings.LastIndex(dstDirPath, "/")+1:]
	// 拼接路由文件需要用到的包
	var importPackage string
	for key, val := range packageCache {
		// 过滤掉与路由同级的包
		if key != routerPackageName {
			// 判断是否需要添加包前缀路径
			if len(packagePrefix) > 0 && val {
				importPackage += fmt.Sprintf("	\"%s/%s\"\n", packagePrefix, key)
			} else {
				importPackage += fmt.Sprintf("	\"%s\"\n", key)
			}
		}
	}
	// 创建路由文件
	data := []byte(fmt.Sprintf(
		routerFileFormat,                      // 路由文件格式
		author,                                // 作者
		time.Now().Format("2006/01/02 15:04"), // 创建时间
		routerPackageName,                     // 包名
		importPackage,                         // 引入包
		routerContentCache,                    // 路由内容
	))
	// 替换与路由统计的函数调用
	data = bytes.ReplaceAll(data, []byte(routerPackageName+"."), nil)
	// 创建文件
	err = ioutil.WriteFile(
		filepath.Join(dstDirPath, "router.go"),
		data,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
}

// 创建GO文件
// @params oldFilePath string 文件完整路径
func createGoFile(oldFilePath string) {
	// 将源文件夹替换成目标文件夹
	newFilePath := strings.Replace(oldFilePath, srcDirPath, dstDirPath, 1)

	// 提取文件名
	fileName := filepath.Base(newFilePath)

	// 将文件名用作GO的文件名，只需要末尾添加.go即可
	goFilePath := newFilePath + ".go"

	// 根据文件名计算函数名                  // 基本文件名
	tempFunName := strings.ReplaceAll(fileName, ".", "_")   // 替换特殊字符
	tempFunName = strings.ReplaceAll(tempFunName, "-", "_") // 替换特殊字符
	funcName := "Get_" + tempFunName                        // 拼接函数名

	// 文件路径就是包路径
	packagePath := strings.ReplaceAll(filepath.Dir(newFilePath), "\\", "/")
	// 当前包路径缓存到全局
	packageCache[packagePath] = true
	// 最后一个文件夹就是当前文件包名
	packageName := packagePath[strings.LastIndex(packagePath, "/")+1:]

	// 读取文件内容
	data, err := ioutil.ReadFile(oldFilePath)
	if err != nil {
		log.Fatal("读取文件失败:", err)
	}
	// 文件Byte转换为字符串数组
	var tempData []string
	for _, v := range data {
		tempData = append(tempData, fmt.Sprint(v))
	}

	// 格式化内容
	data = []byte(fmt.Sprintf(
		staticFileFormat,                      // 静态文件输出格式
		author,                                // 作者
		time.Now().Format("2006/01/02 15:04"), // 创建时间
		packageName,                           // 包名
		funcName,                              // 函数名
		strings.Join(tempData, ", "),          // 静态文件流
	))

	// 生成文件
	err = os.MkdirAll(filepath.Dir(goFilePath), 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(goFilePath, data, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 媒体类型，默认为HTML
	contentType := "text/html"
	// 截取第一个小数点之后的内容，根据该内容来判断文件Content-Type
	ext := filepath.Ext(fileName)
	for key, val := range tempExtMap {
		for _, v := range val {
			if v == ext {
				contentType = key
			}
		}
	}

	//  路由地址（原文件路径去掉目标文件夹，反斜杠转正斜杠）
	routerPath := strings.Replace(
		strings.ReplaceAll(newFilePath, "\\", "/"),
		dstDirPath,
		"",
		1,
	)
	// 追加路由
	routerContentCache += fmt.Sprintf(
		routerContentFormat, // 单条路由格式
		routerPath,          // 路由地址
		contentType,         // 资源类型
		packageName,         // 包名
		funcName,            // 函数名
	)
}
