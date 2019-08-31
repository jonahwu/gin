package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Foo struct {
	Value1 int
	Value2 string
	Value3 string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/array", Array)
	router.GET("/arraystruct", ArrayStruct)
	router.GET("/map", Map)
	router.GET("/mapkeys", MapKeys)
	router.GET("/mapselectkeys", MapKeys)
	router.GET("/tables", Tables)
	router.GET("/tables2", Tables2)
	router.GET("/tablesaction", TablesAction)
	router.GET("/layout", Layout)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
	router.Run(":8080")

}

func MapSelectKeys(context *gin.Context) {
	values := make(map[string]string)
	values["language"] = "Go"
	values["version"] = "1.7.4"

	context.HTML(http.StatusOK, "mapSelectKeys.html", gin.H{"myMap": values})
}

//http://172.16.155.190:8080/mapkeys
func MapKeys(context *gin.Context) {
	values := make(map[string]string)
	values["language"] = "Go"
	values["version"] = "1.7.4"

	context.HTML(http.StatusOK, "mapkeys.html", gin.H{"myMap": values})
}

//http://172.16.155.190:8080/array
func Array(context *gin.Context) {
	var values []int
	for i := 0; i < 5; i++ {
		values = append(values, i)
	}

	context.HTML(http.StatusOK, "array.tmpl", gin.H{"values": values})
	//context.String(http.StatusOK, "hello, world")
}

//important struct
//http://172.16.155.190:8080/arraystruct
func ArrayStruct(context *gin.Context) {
	var values []Foo
	for i := 0; i < 5; i++ {
		values = append(values, Foo{Value1: i, Value2: "value " + strconv.Itoa(i)})
	}

	context.HTML(http.StatusOK, "arraystruct.html", gin.H{"values": values})
}

//http://172.16.155.190:8080/tables
func Tables(context *gin.Context) {
	var values []Foo
	for i := 0; i < 5; i++ {
		values = append(values, Foo{Value1: i, Value2: "value " + strconv.Itoa(i)})
	}

	context.HTML(http.StatusOK, "tables.html", gin.H{"values": values})
}

//kind of beauty than tables
//http://172.16.155.190:8080/tables2
func Tables2(context *gin.Context) {
	var values []Foo
	for i := 0; i < 5; i++ {
		values = append(values, Foo{Value1: i, Value2: "value " + strconv.Itoa(i)})
	}

	context.HTML(http.StatusOK, "tables2.html", gin.H{"values": values})
}
func TablesAction(context *gin.Context) {
	var values []Foo
	for i := 0; i < 5; i++ {
		values = append(values, Foo{Value1: i, Value2: "value " + strconv.Itoa(i)})
	}

	context.HTML(http.StatusOK, "tables-action.html", gin.H{"values": values})
}

//http://172.16.155.190:8080/layout
func Layout(context *gin.Context) {
	var values []Foo
	for i := 0; i < 5; i++ {
		values = append(values, Foo{Value1: i, Value2: "value " + strconv.Itoa(i)})
	}

	context.HTML(http.StatusOK, "layout.html", gin.H{"title": "布局页面"})
}

//http://172.16.155.190:8080/map
func Map(context *gin.Context) {
	values := make(map[string]string)
	values["language"] = "Go"
	values["version"] = "1.7.4"

	context.HTML(http.StatusOK, "map.html", gin.H{"myMap": values})
}
