// 工具的全局定义

package tool

// InputParams 参数结构
type InputParams struct {
	rootPackage string // 根包名
	SrcPath     string // 需要编码的文件或文件夹路径
	OutDirPath  string // 编码后输出的文件夹路径
	Author      string // 文件作者（默认为空）
	PackPrefix  string // Gin路由文件引入时的包名前缀
	ShowDetail  bool   // 是否显示详细信息
}

// globalParams 全局参数
var globalParams = InputParams{
	rootPackage: "out", // 根包名
	SrcPath:     "",    // 需要编码的文件或文件夹路径
	OutDirPath:  "",    // 编码后输出的文件夹路径
	Author:      "",    // 文件作者（默认为空）
	PackPrefix:  "",    // Gin路由文件引入时的包名前缀
	ShowDetail:  false, // 是否显示详细信息
}

// 全局缓存用到的包名，值表示该包是否要拼接前缀
var packageCache = map[string]bool{
	"net/http":                 false,
	"github.com/gin-gonic/gin": false,
}

// 全局缓存路由内容，最后一起写入文件
var routerContentCache string

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
		%s
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
		c.Data(http.StatusOK, "%s", %s())
	})
`
