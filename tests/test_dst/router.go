
/**
 *@Title 静态文件Byte数据
 *@Desc 该文件由工具自动生成，请勿修改
 *@Author Bearki
 *@Time 2022/04/28 18:54
 */

package test_dst

import (
	"github.com/bearki/go-endata/tests/test_dst/static/css"
	"github.com/bearki/go-endata/tests/test_dst/static/js"
	"github.com/bearki/go-endata/tests/test_dst/static/media"
	"net/http"
	"github.com/gin-gonic/gin"

)

// 请从外部传入一个路由分组即可
func InitRouter(r *gin.RouterGroup) {

	r.GET("/asset-manifest.json", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", Get_asset_manifest_json())
	})

	r.GET("/bb", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/octet-stream", Get_bb())
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/x-icon", Get_favicon_ico())
	})

	r.GET("/index.html", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", Get_index_html())
	})

	r.GET("/logo.png", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/png", Get_logo_png())
	})

	r.GET("/manifest.json", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", Get_manifest_json())
	})

	r.GET("/robots.txt", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", Get_robots_txt())
	})

	r.GET("/static/css/2.css", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/css; charset=utf-8", css.Get_2_css())
	})

	r.GET("/static/css/2.css.map", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", css.Get_2_css_map())
	})

	r.GET("/static/css/main.css", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/css; charset=utf-8", css.Get_main_css())
	})

	r.GET("/static/css/main.css.map", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", css.Get_main_css_map())
	})

	r.GET("/static/js/2.js", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", js.Get_2_js())
	})

	r.GET("/static/js/2.js.LICENSE.txt", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", js.Get_2_js_LICENSE_txt())
	})

	r.GET("/static/js/2.js.map", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", js.Get_2_js_map())
	})

	r.GET("/static/js/3.js", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", js.Get_3_js())
	})

	r.GET("/static/js/3.js.map", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", js.Get_3_js_map())
	})

	r.GET("/static/js/main.js", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", js.Get_main_js())
	})

	r.GET("/static/js/main.js.map", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", js.Get_main_js_map())
	})

	r.GET("/static/js/runtime-main.js", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", js.Get_runtime_main_js())
	})

	r.GET("/static/js/runtime-main.js.map", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", js.Get_runtime_main_js_map())
	})

	r.GET("/static/media/balance.png", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/png", media.Get_balance_png())
	})

	r.GET("/static/media/bg.png", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/png", media.Get_bg_png())
	})

	r.GET("/static/media/empty.png", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/png", media.Get_empty_png())
	})

	r.GET("/static/media/step1.png", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/png", media.Get_step1_png())
	})

	r.GET("/static/media/step2.png", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/png", media.Get_step2_png())
	})

}
