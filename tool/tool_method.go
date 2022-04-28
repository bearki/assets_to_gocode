// 工具实现

package tool

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

// RunTool 运行工具
// @params params InputParams 传入参数
// @return error
func RunTool(params InputParams) error {
	// 优化输入输出路径
	srcPath := filepath.Clean(params.SrcPath)
	outPath := filepath.Clean(params.OutDirPath)

	if len(outPath) == 0 {
		outPath = "out"
	}
	// 判断输出路径是否是绝对路径
	if filepath.IsAbs(outPath) {
		return fmt.Errorf("does not support absolute path as output destination path")
	}

	// 覆盖默认数据
	globalParams.SrcPath = srcPath
	globalParams.OutDirPath = outPath
	globalParams.Author = params.Author
	globalParams.OutGinRouter = params.OutGinRouter
	globalParams.PackPrefix = params.PackPrefix

	// 配置后缀与资源的对应关系
	err := setContentType()
	if err != nil {
		return err
	}

	// 执行文件生成
	err = RangeDir(globalParams.SrcPath)
	if err != nil {
		return err
	}

	// 判断是否需要生成路由文件
	if globalParams.OutGinRouter {
		// 路由文件的包名（取输出文件夹路径的最后一个文件夹为包名）
		routerPackageName := globalParams.OutDirPath[strings.LastIndex(globalParams.OutDirPath, "/")+1:]
		// 拼接路由文件需要用到的包
		var importPackage string
		// 执行循环拼接
		for key, val := range packageCache {
			// 过滤掉与路由同级的包
			if key != routerPackageName {
				// 判断是否需要添加包前缀路径
				if len(globalParams.PackPrefix) > 0 && val {
					importPackage += fmt.Sprintf("	\"%s/%s\"\n", globalParams.PackPrefix, key)
				} else {
					importPackage += fmt.Sprintf("	\"%s\"\n", key)
				}
			}
		}
		// 创建路由文件
		data := []byte(fmt.Sprintf(
			routerFileFormat,                      // 路由文件格式
			globalParams.Author,                   // 作者
			time.Now().Format("2006/01/02 15:04"), // 创建时间
			routerPackageName,                     // 包名
			importPackage,                         // 引入包
			routerContentCache,                    // 路由内容
		))
		// 替换与路由统计的函数调用
		data = bytes.ReplaceAll(data, []byte(routerPackageName+"."), nil)
		// 创建文件
		err = ioutil.WriteFile(
			filepath.Join(globalParams.OutDirPath, "router.go"),
			data,
			0666,
		)
		// 返回错误信息
		return err
	}

	tempWd, _ := os.Getwd()
	log.Printf("编码输出目标: %s", filepath.Join(tempWd, "out"))

	// 全部成功
	return nil
}

// RangeDir 遍历文件夹
// @return error
func RangeDir(dirPath string) error {
	// 打开文件夹
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	// 遍历文件夹
	for _, fi := range fileInfos {
		// 显示下进度
		// fmt.Printf("正在处理: %d/%d\n", i+1, len(fileInfos))
		// 拼接完整路径
		filePath := filepath.Join(dirPath, fi.Name())
		// 判断是否是文件夹
		if fi.IsDir() {
			// 递归这个目录
			RangeDir(filePath)
		} else {
			// 编码处理这个文件
			err := createGoFile(filePath)
			if err != nil {
				// 生成文件失败了
				return err
			}
		}
	}

	// 生成资源Go文件成功
	return nil
}

// createGoFile 创建GO文件
// @params filePath string 文件完整路径
// @return error
func createGoFile(filePath string) error {

	// 拆分文件夹与文件名
	fileDirPath, fileBaseName := filepath.Split(filePath)

	// 将源文件夹替换成目标文件夹
	newDirPath := strings.Replace(
		fileDirPath,
		globalParams.SrcPath,
		globalParams.OutDirPath,
		1,
	)

	// 根据文件名计算函数名
	tempFunName := strings.ReplaceAll(fileBaseName, ".", "_") // 替换特殊字符
	tempFunName = strings.ReplaceAll(tempFunName, "-", "_")   // 替换特殊字符
	funcName := "Get_" + tempFunName                          // 拼接函数名

	// 根据文件夹路径计算包路径和当前文件包名
	// 剔除多余符号
	tempPackPath := filepath.Clean(newDirPath)
	// 将文件路径转换为"/"
	packagePath := strings.ReplaceAll(tempPackPath, "\\", "/")
	// 最后一个包名就是当前文件包名
	packageName := packagePath[strings.LastIndex(packagePath, "/")+1:]

	// 读取文件内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	// 文件Byte转换为字符串数组
	var tempData []string
	for i, v := range data {
		if i > 0 && i%20 == 0 {
			tempData = append(tempData, "\n")
			tempData = append(tempData, fmt.Sprintf("        0x%02x", v))
		} else {
			tempData = append(tempData, fmt.Sprintf("0x%02x", v))
		}
	}

	// 格式化内容
	data = []byte(fmt.Sprintf(
		staticFileFormat,                      // 静态文件输出格式
		globalParams.Author,                   // 作者
		time.Now().Format("2006/01/02 15:04"), // 创建时间
		packageName,                           // 包名
		funcName,                              // 函数名
		strings.ReplaceAll(strings.Join(tempData, ", "), "\n, ", "\n"), // 静态文件流
	))

	// 创建Go文件的文件夹
	err = os.MkdirAll(newDirPath, 0755)
	if err != nil {
		return err
	}
	// 将文件路径用作GO的文件名，只需要末尾添加.go即可
	goFilePath := filepath.Join(newDirPath, fileBaseName) + ".go"
	// 生成文件
	err = ioutil.WriteFile(goFilePath, data, 0666)
	if err != nil {
		return err
	}

	// 判断是否需要生成路由文件
	if globalParams.OutGinRouter {
		// 当前包路径缓存到全局
		packageCache[packagePath] = true

		// 匹配媒体资源类型
		contentType := getContentType(filepath.Ext(fileBaseName))

		//  路由地址（包路径去掉目标文件夹路径再加上原始文件名就是路由地址）
		raplaceStr := strings.ReplaceAll(globalParams.OutDirPath, "\\", "/")
		routerPath := strings.Replace(
			packagePath+"/"+fileBaseName,
			raplaceStr, // 将目标路径文件夹分隔符转换为"/",不然匹配不到无法替换
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

	return nil
}
