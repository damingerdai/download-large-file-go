package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := gin.New()
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Handle("GET", "ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	app.Handle("GET", "/", homePage)
	app.Handle("POST", "/query", downloadFile)

	app.Run(":8080")
}

func homePage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.Writer.WriteString(`
<html>
<body>
open window and to download:
<a href="javascript:download()">download</a>
<script>
function download(){
    var handle = window.open("about:blank", "my_download_window");
	document.forms[0].target = "my_download_window";
	document.forms[0].json.value="ahfu test";
	document.forms[0].submit();
}
</script>
<form action="/query" method="POST" enctype="multipart/form-data">
<input type="hidden" name="json" value=""/>
</form>
</body>
</html>
	`)
}

func downloadFile(ctx *gin.Context) {
	_, has := ctx.GetPostForm("json")
	if !has {
		ctx.Data(400, "text/plain", []byte("not found json form data"))
		return
	}

	ctx.Writer.WriteHeader(200)
	ctx.Header("Content-Type", "text/plain;charset=utf-8")
	ctx.Header("Transfer-Encoding", "chunked") // 告诉浏览器，分段的流式输出
	ctx.Header("Content-Encoding", "gzip")
	now := time.Now()
	fileName := now.Format("20060102_150405.csv")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment;filename=\"%s\"", fileName))
	str := "damingerdai"
	for i := 0; i < 100; i++ {
		ctx.Writer.WriteString("\"")
		ctx.Writer.WriteString(str)
		ctx.Writer.WriteString("\"\t")
		ctx.Writer.WriteString("\"")
		ctx.Writer.WriteString(time.Now().Format("2006-01-02 15:04:05"))
		ctx.Writer.WriteString("\"\n")
		ctx.Writer.Flush() // 产生一定的数据后， flush到浏览器端
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}
