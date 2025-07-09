package main

import (
	"gee/gee"
	"net/http"
)

//type Engine struct{}

//func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	switch req.URL.Path {
//	case "/":
//		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//	case "/hello":
//		for k, v := range req.Header {
//			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
//		}
//	default:
//		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
//	}
//}
//func main() {
//	engine := new(Engine)
//	log.Fatal(http.ListenAndServe(":8081", engine))
//}

//func main() {
//	http.HandleFunc("/", indexHandler)
//	http.HandleFunc("/hello", helloHandler)
//	log.Fatal(http.ListenAndServe(":8081", nil))
//}

//func indexHandler(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//}
//
//func helloHandler(w http.ResponseWriter, req *http.Request) {
//	for k, v := range req.Header {
//		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
//	}
//}

//func main() {
//	r := gee.New()
//	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "funck URL.Path = %q\n", r.URL.Path)
//	})
//	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
//		for k, v := range r.Header {
//			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
//		}
//	})
//	r.Run(":8081")
//}

/***************************day2 ***************************/
func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s , you are at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
