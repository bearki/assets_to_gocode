// 工具实现

package tool

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

	// 提取根包名
	rootPackage := filepath.Base(outPath)

	// 覆盖默认数据
	globalParams.SrcPath = srcPath
	globalParams.OutDirPath = outPath
	globalParams.Author = params.Author
	globalParams.PackPrefix = params.PackPrefix
	globalParams.ShowDetail = params.ShowDetail
	globalParams.rootPackage = rootPackage

	fmt.Println(params.PackPrefix)

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
	if len(globalParams.PackPrefix) > 0 {
		// 路由文件的包名（取输出文件夹路径的最后一个文件夹为包名）
		routerPackageName := rootPackage
		// 拼接路由文件需要用到的包
		var importPackage string
		// 执行循环拼接
		for key, _ := range packageCache {
			// 过滤掉与路由同级的包
			if key != fmt.Sprintf("%s/%s", globalParams.PackPrefix, routerPackageName) {
				importPackage += fmt.Sprintf("\t\"%s\"\n", key)
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
// @params srcFilePath string 文件完整路径
// @return             error
func createGoFile(srcFilePath string) error {
	// 拆分文件夹与文件名
	fileDirPath, fileBaseName := filepath.Split(srcFilePath)
	// 将源文件夹替换成目标文件夹
	newDirPath := strings.Replace(
		fileDirPath,
		globalParams.SrcPath,
		globalParams.OutDirPath,
		1,
	)
	// GO的文件名，只需要末尾添加.go即可
	goFilePath := filepath.Join(newDirPath, fileBaseName) + ".go"

	// 是否需要显示详情
	if globalParams.ShowDetail {
		// 显示下进度
		log.Printf("正在处理: %s\n", srcFilePath)
	}

	// 根据文件名计算函数名
	funcName := strings.ReplaceAll(fileBaseName, ".", "_") // 替换特殊字符
	funcName = strings.ReplaceAll(funcName, "-", "_")      // 替换特殊字符
	funcName = "Get_" + funcName                           // 拼接函数名

	// 根据文件夹路径计算当前包完整路径和当前包
	// 移除目标文件夹
	packagePath := strings.Replace(newDirPath, globalParams.OutDirPath, "", 1)
	// 清除多余符号
	packagePath = filepath.Clean(packagePath)
	// 将"\"替换为"/"
	packagePath = strings.ReplaceAll(packagePath, "\\", "/")
	// 当开头为"/"时需要移除
	ok, err := regexp.MatchString(`^/.*$`, packagePath)
	if err != nil {
		return err
	}
	if ok {
		// 移除
		packagePath = packagePath[1:]
	}
	// 拼接上前缀得到的便是包完整路径
	if len(globalParams.PackPrefix) > 0 {
		if len(packagePath) > 0 {
			// 包路径 = 包前缀 + 根包名 + 当前包名
			packagePath = globalParams.PackPrefix + "/" + globalParams.rootPackage + "/" + packagePath
		} else {
			// 包路径 = 包前缀 + 根包名 + 当前包名
			packagePath = globalParams.PackPrefix + "/" + globalParams.rootPackage
		}
	} else {
		if len(packagePath) > 0 {
			// 包路径 = 根包名 + 当前包名
			packagePath = globalParams.rootPackage + "/" + packagePath
		} else {
			// 包路径 = 根包名
			packagePath = globalParams.rootPackage
		}
	}

	// 最后一个路径就是当前文件包名
	packageName := filepath.Base(packagePath)

	// 读取文件内容
	data, err := ioutil.ReadFile(srcFilePath)
	if err != nil {
		return err
	}
	// 文件Byte转换为字符串数组
	var tempData []string
	for i, v := range data {
		if i > 0 && i%20 == 0 {
			tempData = append(tempData, fmt.Sprintf("\n\t\t0x%02x", v))
		} else {
			tempData = append(tempData, fmt.Sprintf("0x%02x", v))
		}
	}

	// 将要写入的数据
	var writeData string = strings.Join(tempData, ", ")
	if len(tempData) > 0 {
		writeData += ","
	}

	// 格式化内容
	data = []byte(fmt.Sprintf(
		staticFileFormat,                      // 静态文件输出格式
		globalParams.Author,                   // 作者
		time.Now().Format("2006/01/02 15:04"), // 创建时间
		packageName,                           // 包名
		funcName,                              // 函数名
		writeData,                             // 静态文件流
	))

	// 创建Go文件的文件夹
	err = os.MkdirAll(newDirPath, 0755)
	if err != nil {
		return err
	}

	// 生成文件
	err = ioutil.WriteFile(goFilePath, data, 0666)
	if err != nil {
		return err
	}

	// 判断是否需要生成路由文件
	if len(globalParams.PackPrefix) > 0 {
		// 当前包路径缓存到全局
		packageCache[packagePath] = true

		// 匹配媒体资源类型
		contentType := getContentType(filepath.Ext(fileBaseName))

		//  路由地址
		routerPath := strings.Replace(
			packagePath+"/"+fileBaseName,
			globalParams.PackPrefix+"/"+globalParams.rootPackage,
			"",
			1,
		)
		// 非根包需要前置包名
		if packageName != globalParams.rootPackage {
			funcName = fmt.Sprintf("%s.%s", packageName, funcName)
		}
		// 追加路由
		routerContentCache += fmt.Sprintf(
			routerContentFormat, // 单条路由格式
			routerPath,          // 路由地址
			contentType,         // 资源类型
			funcName,            // 函数名
		)
	}

	// 是否需要显示详情
	if globalParams.ShowDetail {
		// 显示下进度
		log.Printf("处理完成: %s\n", goFilePath)
	}
	return nil
}
