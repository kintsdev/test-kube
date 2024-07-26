package main

import (
	"net/http"
	"os"
)

var (
	podName   string = os.Getenv("POD_NAME")
	namespace string = os.Getenv("POD_NAMESPACE")
	nodeName  string = os.Getenv("NODE_NAME")
	hostIP    string = os.Getenv("HOST_IP")
	podIP     string = os.Getenv("POD_IP")
)

func main() {
	// http server start 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to K8s\n"))
		w.Write([]byte("Pod Name: " + podName + "\n"))
		w.Write([]byte("Namespace: " + namespace + "\n"))
		w.Write([]byte("Node Name: " + nodeName + "\n"))
		w.Write([]byte("Host IP: " + hostIP + "\n"))
		w.Write([]byte("Pod IP: " + podIP + "\n"))
	})

	http.ListenAndServe(":8080", nil)
}
