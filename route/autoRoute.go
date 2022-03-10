package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

var methodss = make(map[string]map[string]reflect.Value)
var methods = make(map[string]map[string]map[string]reflect.Value)

func AutoRoute(engine *gin.Engine, relativePath string, s interface{}) {
	if strings.HasSuffix(relativePath, "/") {
		relativePath = relativePath[1 : len(relativePath)-1]
	}
	if !strings.HasPrefix(relativePath, "/") {
		return
	}
	relativePath = relativePath + "/:action"
	engine.GET(relativePath, AutoHand(s))
	engine.POST(relativePath, AutoHand(s))
	engine.DELETE(relativePath, AutoHand(s))
	engine.PUT(relativePath, AutoHand(s))
	engine.OPTIONS(relativePath, AutoHand(s))
}
func AutoHand(s interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.Title(strings.ToLower(c.Request.Method))
		action := c.Param("action")
		realAction := strings.ToLower(action)
		rt := reflect.TypeOf(s)
		pkgName := rt.String()
		if _, ok := methods[method]; !ok {
			methods[method] = make(map[string]map[string]reflect.Value)
		}
		v, ok := methods[method][pkgName][realAction]
		if !ok && len(methods[pkgName]) == 0 {
			rv := reflect.ValueOf(s)
			methods[method][pkgName] = make(map[string]reflect.Value)
			for i := 0; i < rv.NumMethod(); i++ {
				fn := rv.Method(i)
				ft := rt.Method(i)
				//utils.Println(ft.Name,strings.Title(strings.ToLower(c.Request.Method)),action)
				if strings.HasPrefix(ft.Name, method) {
					ftname := strings.ToLower(strings.Replace(ft.Name, method, "", 1))
					methods[method][pkgName][ftname] = fn
				}
				//methods[pkgName][strings.ToLower(ft.Name)] = fn
			}
			v, ok = methods[method][pkgName][realAction]
			if !ok && (len(methods[method]) > 0 || len(methods[method][pkgName]) > 0) {
				http.NotFound(c.Writer, c.Request)
				c.Abort()
				return
			}
		} else if !ok && (len(methods[method]) > 0 || len(methods[method][pkgName]) > 0) {
			http.NotFound(c.Writer, c.Request)
			c.Abort()
			return
		}

		arguments := make([]reflect.Value, 1)
		arguments[0] = reflect.ValueOf(c) // *gin.Context
		v.Call(arguments)
	}
}

//func routeregexp(opt string, beforpath string) (bool, string) {
//	match := regexp.MustCompile(`^(` + opt + `+)(.*)$`).FindStringSubmatch(beforpath)
//	if len(match) > 0 {
//		for i := range match {
//			fmt.Println(i, match[i])
//		}
//		return true, strings.ToLower(match[2])
//	}
//	return false, ""
//}

func routeregexp(strings string) (res bool, match []string) {
	match = regexp.MustCompile(`^(Bearer\s)(.+)$`).FindStringSubmatch(strings)
	if len(match) > 0 {

		return true, match
	}
	return false, match
}