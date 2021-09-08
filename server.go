package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

var server http.Server

var CounterMap = PasswordMap{m: make(map[int]string)}

var ResponseStat = ResponseStats{Average: 0.00, Total: 0}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type ctxKey struct{}

var routes = []route{
	newRoute("GET", "/hash/([0-9]+)", GetHashHttp),
	newRoute("POST", "/hash", PostHashHttp),
	newRoute("GET", "/stats", getStats),
	newRoute("GET", "/shutdown", shutdown),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func Serve(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf("Serving %s %s", r.Method, r.URL.Path)
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				fmt.Println(allow)
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			if strings.HasSuffix(r.URL.Path, "/hash") {
				end := time.Now()
				defer ResponseStat.countAvarage(start, end)
			}
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

//method to get field value from request
func GetField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

func getStats(w http.ResponseWriter, r *http.Request) {
	response, _ := ResponseStat.toJson()
	w.Write(response)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

func StartHttpServer() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Serve)
	server = http.Server{
		Addr:    os.Getenv("host"),
		Handler: mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutingdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting....")
}
