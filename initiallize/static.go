package initiallize

import (
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const frontendSubDir = "web/dist"

// MountFrontend 将嵌入的前端 FS 挂载到 Gin，处理静态资源与 SPA History 回退。
// 若 frontendFS 为 nil 则不挂载。
func MountFrontend(r *gin.Engine, frontendFS fs.FS) {
	if frontendFS == nil {
		return
	}
	subFS, err := fs.Sub(frontendFS, frontendSubDir)
	if err != nil {
		panic("frontend fs.Sub: " + err.Error())
	}
	r.NoRoute(serveFrontend(subFS))
}

func serveFrontend(subFS fs.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/api") {
			c.Status(http.StatusNotFound)
			return
		}
		name := strings.TrimPrefix(path, "/")
		if name == "" {
			name = "index.html"
		}
		f, err := subFS.Open(name)
		if err != nil {
			serveIndex(subFS, c)
			return
		}
		defer f.Close()
		stat, _ := f.Stat()
		if stat.IsDir() {
			index, _ := subFS.Open(name + "/index.html")
			if index != nil {
				defer index.Close()
				c.Header("Content-Type", "text/html; charset=utf-8")
				_, _ = io.Copy(c.Writer, index)
				return
			}
			c.Status(http.StatusNotFound)
			return
		}
		content, _ := io.ReadAll(f)
		c.Data(http.StatusOK, contentType(name), content)
	}
}

func serveIndex(subFS fs.FS, c *gin.Context) {
	index, err := subFS.Open("index.html")
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	defer index.Close()
	c.Header("Content-Type", "text/html; charset=utf-8")
	_, _ = io.Copy(c.Writer, index)
}

func contentType(name string) string {
	switch {
	case strings.HasSuffix(name, ".html"):
		return "text/html; charset=utf-8"
	case strings.HasSuffix(name, ".js"):
		return "application/javascript; charset=utf-8"
	case strings.HasSuffix(name, ".css"):
		return "text/css; charset=utf-8"
	case strings.HasSuffix(name, ".ico"):
		return "image/x-icon"
	case strings.HasSuffix(name, ".png"):
		return "image/png"
	case strings.HasSuffix(name, ".svg"):
		return "image/svg+xml"
	case strings.HasSuffix(name, ".woff2"):
		return "font/woff2"
	case strings.HasSuffix(name, ".woff"):
		return "font/woff"
	default:
		return "application/octet-stream"
	}
}
