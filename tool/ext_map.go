// 后缀与资源类型的关系映射

package tool

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 第一位小数点之后的内容，将会根据这个MAP设置Content-Type
var extMapContentType = map[string][]string{
	"text/html; charset=utf-8": {
		".jsp",
		".html",
		".htx",
		".plg",
		".xhtml",
		".htm",
		".stm",
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
	"image/gif": {
		".gif",
	},
	"application/x-001": {
		".001",
	},
	"text/h323": {
		".323",
	},
	"drawing/907": {
		".907",
	},
	"audio/x-mei-aac": {
		".acp",
	},
	"audio/aiff": {
		".aif",
		".aiff",
		".aifc",
	},
	"text/asa": {
		".asa",
	},
	"text/asp": {
		".asp",
	},
	"audio/basic": {
		".snd",
		".au",
	},
	"application/vnd.adobe.workflow": {
		".awf",
	},
	"application/x-bmp": {
		".bmp",
	},
	"application/x-c4t": {
		".c4t",
	},
	"application/x-cals": {
		".cal",
	},
	"application/x-netcdf": {
		".cdf",
	},
	"application/x-cel": {
		".cel",
	},
	"application/x-g4": {
		".g4",
		".ig4",
		".cg4",
	},
	"application/x-cit": {
		".cit",
	},
	"application/x-cmx": {
		".cmx",
	},
	"application/pkix-crl": {
		".crl",
	},
	"application/x-csi": {
		".csi",
	},
	"application/x-cut": {
		".cut",
	},
	"application/x-dbm": {
		".dbm",
	},
	"application/x-x509-ca-cert": {
		".cer",
		".crt",
		".der",
	},
	"application/x-dib": {
		".dib",
	},
	"application/msword": {
		".rtf",
		".wiz",
		".dot",
		".doc",
	},
	"application/x-drw": {
		".drw",
	},
	"Model/vnd.dwf": {
		".dwf",
	},
	"application/x-dwg": {
		".dwg",
	},
	"application/x-dxf": {
		".dxf",
	},
	"application/x-emf": {
		".emf",
	},
	"application/x-ps": {
		".ps",
		".eps",
	},
	"application/x-ebx": {
		".etd",
	},
	"image/fax": {
		".fax",
	},
	"application/fractals": {
		".fif",
	},
	"application/x-frm": {
		".frm",
	},
	"application/x-gbr": {
		".gbr",
	},
	"application/x-gp4": {
		".gp4",
	},
	"application/x-hmr": {
		".hmr",
	},
	"application/x-hpl": {
		".hpl",
	},
	"application/x-hrf": {
		".hrf",
	},
	"text/x-component": {
		".htc",
	},
	"application/x-iff": {
		".iff",
	},
	"application/x-igs": {
		".igs",
	},
	"application/x-img": {
		".img",
	},
	"application/x-internet-signup": {
		".ins",
		".isp",
	},
	"java/*": {
		".java",
		".class",
	},
	"image/jpeg": {
		".jfif",
		".jpg",
		".jpeg",
	},
	"application/x-jpg": {
		".jpg",
	},
	"application/x-laplayer-reg": {
		".lar",
	},
	"audio/x-liquid-secure": {
		".lavs",
	},
	"audio/x-la-lms": {
		".lmsff",
	},
	"application/x-ltr": {
		".ltr",
	},
	"video/x-mpeg": {
		".mps",
		".m1v",
		".mpe",
		".m2v",
	},
	"video/mpeg4": {
		".mpv2",
		".mp4",
		".m4e",
	},
	"application/x-troff-man": {
		".man",
	},
	"application/msaccess": {
		".mdb",
	},
	"application/x-shockwave-flash": {
		".mfp",
		".swf",
	},
	"message/rfc822": {
		".nws",
		".eml",
		".mht",
		".mhtml",
	},
	"audio/mid": {
		".mid",
		".rmi",
		".midi",
	},
	"application/x-mil": {
		".mil",
	},
	"audio/x-musicnet-download": {
		".mnd",
	},
	"application/x-javascript": {
		".ls",
		".js",
		".mocha",
	},
	"audio/mp1": {
		".mp1",
	},
	"video/mpeg": {
		".mp2v",
	},
	"application/vnd.ms-project": {
		".mpp",
		".mpt",
		".mpx",
		".mpw",
		".mpd",
	},
	"video/mpg": {
		".mpv",
		".mpg",
		".mpeg",
	},
	"audio/rn-mpeg": {
		".mpga",
	},
	"image/pnetvue": {
		".net",
	},
	"application/x-out": {
		".out",
	},
	"application/x-pkcs12": {
		".pfx",
		".p12",
	},
	"application/pkcs7-mime": {
		".p7m",
		".p7c",
	},
	"application/x-pkcs7-certreqresp": {
		".p7r",
	},
	"application/x-pc5": {
		".pc5",
	},
	"application/x-pcl": {
		".pcl",
	},
	"application/pdf": {
		".pdf",
	},
	"application/vnd.adobe.pdx": {
		".pdx",
	},
	"application/x-pgl": {
		".pgl",
	},
	"application/vnd.ms-pki.pko": {
		".pko",
	},
	"application/x-plt": {
		".plt",
	},
	"application/x-png": {
		".png",
	},
	"application/vnd.ms-powerpoint": {
		".ppa",
		".pps",
		".pwz",
		".ppt",
		".pot",
	},
	"application/x-ppt": {
		".ppt",
	},
	"application/pics-rules": {
		".prf",
	},
	"application/x-prt": {
		".prt",
	},
	"application/postscript": {
		".ai",
		".eps",
		".ps",
	},
	"audio/vnd.rn-realaudio": {
		".ra",
	},
	"application/x-ras": {
		".ras",
	},
	"application/x-red": {
		".red",
	},
	"application/vnd.rn-realsystem-rjs": {
		".rjs",
	},
	"application/x-rlc": {
		".rlc",
	},
	"application/vnd.rn-realmedia": {
		".rm",
	},
	"application/vnd.rn-realmedia-secure": {
		".rms",
	},
	"application/vnd.rn-realsystem-rmx": {
		".rmx",
	},
	"image/vnd.rn-realpix": {
		".rp",
	},
	"application/vnd.rn-rsml": {
		".rsml",
	},
	"video/vnd.rn-realvideo": {
		".rv",
	},
	"application/x-sat": {
		".sat",
	},
	"application/x-sdw": {
		".sdw",
	},
	"application/x-slb": {
		".slb",
	},
	"drawing/x-slk": {
		".slk",
	},
	"application/smil": {
		".smi",
		".smil",
	},
	"text/plain": {
		".txt",
		".sol",
		".sor",
	},
	"application/futuresplash": {
		".spl",
	},
	"application/streamingmedia": {
		".ssm",
	},
	"application/vnd.ms-pki.stl": {
		".stl",
	},
	"application/x-sty": {
		".sty",
	},
	"application/x-tg4": {
		".tg4",
	},
	"image/tiff": {
		".tif",
		".tiff",
	},
	"drawing/x-top": {
		".top",
	},
	"application/x-icq": {
		".uin",
	},
	"text/x-vcard": {
		".vcf",
	},
	"application/vnd.visio": {
		".vdx",
		".vst",
		".vsw",
		".vtx",
		".vsd",
		".vss",
		".vsx",
	},
	"application/x-vpeg005": {
		".vpg",
	},
	"application/x-vsd": {
		".vsd",
	},
	"audio/wav": {
		".wav",
	},
	"application/x-wb1": {
		".wb1",
	},
	"application/x-wb3": {
		".wb3",
	},
	"application/x-wk4": {
		".wk4",
	},
	"application/x-wks": {
		".wks",
	},
	"audio/x-ms-wma": {
		".wma",
	},
	"application/x-wmf": {
		".wmf",
	},
	"video/x-ms-wmv": {
		".wmv",
	},
	"application/x-ms-wmz": {
		".wmz",
	},
	"application/x-wpd": {
		".wpd",
	},
	"application/vnd.ms-wpl": {
		".wpl",
	},
	"application/x-wr1": {
		".wr1",
	},
	"application/x-wrk": {
		".wrk",
	},
	"application/x-ws": {
		".ws",
		".ws2",
	},
	"application/vnd.adobe.xdp": {
		".xdp",
	},
	"application/vnd.adobe.xfd": {
		".xfd",
	},
	"application/x-xls": {
		".xls",
	},
	"text/xml": {
		".dcd",
		".ent",
		".mtx",
		".rdf",
		".tsd",
		".wsdl",
		".xml",
		".xq",
		".xquery",
		".xsl",
		".biz",
		".cml",
		".dtd",
		".fo",
		".math",
		".mml",
		".spp",
		".svg",
		".tld",
		".vml",
		".vxml",
		".xdr",
		".xql",
		".xsd",
		".xslt",
	},
	"application/x-xwd": {
		".xwd",
	},
	"application/vnd.symbian.install": {
		".sisx",
		".sis",
	},
	"application/x-x_t": {
		".x_t",
	},
	"application/vnd.android.package-archive": {
		".apk",
	},
	"application/x-301": {
		".301",
	},
	"application/x-906": {
		".906",
	},
	"application/x-a11": {
		".a11",
	},
	"application/x-anv": {
		".anv",
	},
	"video/x-ms-asf": {
		".asx",
		".asf",
	},
	"video/avi": {
		".avi",
	},
	"application/x-bot": {
		".bot",
	},
	"application/x-c90": {
		".c90",
	},
	"application/vnd.ms-pki.seccat": {
		".cat",
	},
	"application/x-cdr": {
		".cdr",
	},
	"application/x-cgm": {
		".cgm",
	},
	"application/x-cmp": {
		".cmp",
	},
	"application/x-cot": {
		".cot",
	},
	"text/css": {
		".css",
	},
	"application/x-dbf": {
		".dbf",
	},
	"application/x-dbx": {
		".dbx",
	},
	"application/x-dcx": {
		".dcx",
	},
	"application/x-dgn": {
		".dgn",
	},
	"application/x-msdownload": {
		".exe",
		".dll",
	},
	"application/x-dwf": {
		".dwf",
	},
	"application/x-dxb": {
		".dxb",
	},
	"application/vnd.adobe.edn": {
		".edn",
	},
	"application/x-epi": {
		".epi",
	},
	"application/vnd.fdf": {
		".fdf",
	},
	"application/x-": {
		".",
	},
	"application/x-gl2": {
		".gl2",
	},
	"application/x-hgl": {
		".hgl",
	},
	"application/x-hpgl": {
		".hpg",
	},
	"application/mac-binhex40": {
		".hqx",
	},
	"application/hta": {
		".hta",
	},
	"text/webviewhtml": {
		".htt",
	},
	"application/x-icb": {
		".icb",
	},
	"application/x-ico": {
		".ico",
	},
	"application/x-iphone": {
		".iii",
	},
	"video/x-ivf": {
		".IVF",
	},
	"application/x-jpe": {
		".jpe",
	},
	"audio/x-liquid-file": {
		".la1",
	},
	"application/x-latex": {
		".latex",
	},
	"application/x-lbm": {
		".lbm",
	},
	"audio/mpegurl": {
		".m3u",
	},
	"application/x-mac": {
		".mac",
	},
	"application/x-mdb": {
		".mdb",
	},
	"application/x-mi": {
		".mi",
	},
	"audio/x-musicnet-stream": {
		".mns",
	},
	"video/x-sgi-movie": {
		".movie",
	},
	"audio/mp2": {
		".mp2",
	},
	"audio/mp3": {
		".mp3",
	},
	"video/x-mpg": {
		".mpa",
	},
	"application/x-mmxp": {
		".mxp",
	},
	"application/x-nrf": {
		".nrf",
	},
	"text/x-ms-odc": {
		".odc",
	},
	"application/pkcs10": {
		".p10",
	},
	"application/x-pkcs7-certificates": {
		".p7b",
		".spc",
	},
	"application/pkcs7-signature": {
		".p7s",
	},
	"application/x-pci": {
		".pci",
	},
	"application/x-pcx": {
		".pcx",
	},
	"application/x-pic": {
		".pic",
	},
	"application/x-perl": {
		".pl",
	},
	"audio/scpls": {
		".xpl",
		".pls",
	},
	"application/x-ppm": {
		".ppm",
	},
	"application/x-pr": {
		".pr",
	},
	"application/x-prn": {
		".prn",
	},
	"application/x-ptn": {
		".ptn",
	},
	"text/vnd.rn-realtext3d": {
		".r3t",
	},
	"application/rat-file": {
		".rat",
	},
	"application/vnd.rn-recording": {
		".rec",
	},
	"application/x-rgb": {
		".rgb",
	},
	"application/vnd.rn-realsystem-rjt": {
		".rjt",
	},
	"application/x-rle": {
		".rle",
	},
	"application/vnd.adobe.rmf": {
		".rmf",
	},
	"application/vnd.rn-realsystem-rmj": {
		".rmj",
	},
	"application/vnd.rn-rn_music_package": {
		".rmp",
	},
	"application/vnd.rn-realmedia-vbr": {
		".rmvb",
	},
	"application/vnd.rn-realplayer": {
		".rnx",
	},
	"audio/x-pn-realaudio-plugin": {
		".ram",
		".rmm",
		".rpm",
	},
	"text/vnd.rn-realtext": {
		".rt",
	},
	"application/x-rtf": {
		".rtf",
	},
	"application/x-sam": {
		".sam",
	},
	"application/sdp": {
		".sdp",
	},
	"application/x-stuffit": {
		".sit",
	},
	"application/x-sld": {
		".sld",
	},
	"application/x-smk": {
		".smk",
	},
	"application/vnd.ms-pki.certstore": {
		".sst",
	},
	"application/x-tdf": {
		".tdf",
	},
	"application/x-tga": {
		".tga",
	},
	"application/x-tif": {
		".tif",
	},
	"application/x-bittorrent": {
		".torrent",
	},
	"text/iuls": {
		".uls",
	},
	"application/x-vda": {
		".vda",
	},
	"application/x-vst": {
		".vst",
	},
	"audio/x-ms-wax": {
		".wax",
	},
	"application/x-wb2": {
		".wb2",
	},
	"image/vnd.wap.wbmp": {
		".wbmp",
	},
	"application/x-wk3": {
		".wk3",
	},
	"application/x-wkq": {
		".wkq",
	},
	"video/x-ms-wm": {
		".wm",
	},
	"application/x-ms-wmd": {
		".wmd",
	},
	"text/vnd.wap.wml": {
		".wml",
	},
	"video/x-ms-wmx": {
		".wmx",
	},
	"application/x-wp6": {
		".wp6",
	},
	"application/x-wpg": {
		".wpg",
	},
	"application/x-wq1": {
		".wq1",
	},
	"application/x-wri": {
		".wri",
	},
	"text/scriptlet": {
		".wsc",
	},
	"video/x-ms-wvx": {
		".wvx",
	},
	"application/vnd.adobe.xfdf": {
		".xfdf",
	},
	"application/vnd.ms-excel": {
		".xls",
	},
	"application/x-xlw": {
		".xlw",
	},
	"application/x-x_b": {
		".x_b",
	},
	"application/vnd.iphone": {
		".ipa",
	},
	"application/x-silverlight-app": {
		".xap",
	},
}

// CreateExtMapJson 创建MAP映射文件
// @return error 错误信息
func CreateExtMapJson() error {
	data, err := json.MarshalIndent(extMapContentType, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("extmap.json", data, 0666)
}

// setContentType 配置后缀对应的资源类型
// @return error 错误信息
func setContentType() error {
	// 判断JSON文件是否存在
	data, err := ioutil.ReadFile("extmap.json")
	if err != nil {
		// 判断是否是文件不存在
		if os.IsNotExist(err) { // 文件不存在，使用内置数据
			return nil
		}
		// 其他错误
		return err
	}

	// 文件错误，解析JSON为MAP
	var tempMap = make(map[string][]string)
	err = json.Unmarshal(data, &tempMap)
	if err != nil {
		return err
	}
	// 覆盖内置的
	extMapContentType = tempMap
	return nil
}

// getContentType 获取后缀对印的隐射关系
// @params ext string 文件后缀
// @return     string 资源类型
func getContentType(ext string) string {
	// 匹配对应关系
	for key, val := range extMapContentType {
		// 遍历资源类型对应的后缀
		for _, v := range val {
			// 该资源类型
			if v == ext {
				return key
			}
		}
	}
	// 默认为文件流
	return "application/octet-stream"
}
