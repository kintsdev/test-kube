package main

import (
	"net/http"
	"os"
	"text/template"
)

var (
	podName   string = os.Getenv("POD_NAME")
	namespace string = os.Getenv("POD_NAMESPACE")
	nodeName  string = os.Getenv("NODE_NAME")
	hostIP    string = os.Getenv("HOST_IP")
	podIP     string = os.Getenv("POD_IP")
)

const htmlPageTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Test Kube</title>
</head>

<body>
	<h1>Welcome to K8s</h1>
	<p>Pod Name: {{.PodName}}</p>
	<p>Namespace: {{.Namespace}}</p>
	<p>Node Name: {{.NodeName}}</p>
	<p>Host IP: {{.HostIP}}</p>
	<p>Pod IP: {{.PodIP}}</p>

	<form action="/memory-stress" method="post">
		<input type="submit" value="Memory Stress">
	</form>

	<form action="/cpu-stress" method="post">
		<input type="submit" value="CPU Stress">
	</form>

	<form action="/crash" method="post">
		<input type="submit" value="Crash Pod">
	</form>
	
</body>
</html>
`

func main() {
	// http server start 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// render html page
		data := struct {
			PodName   string
			Namespace string
			NodeName  string
			HostIP    string
			PodIP     string
		}{
			PodName:   podName,
			Namespace: namespace,
			NodeName:  nodeName,
			HostIP:    hostIP,
			PodIP:     podIP,
		}
		tmpl, err := template.New("webpage").Parse(htmlPageTemplate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/memory-stress", func(w http.ResponseWriter, r *http.Request) {
		memoryStressStart()
		w.Write([]byte("Memory stress started\n"))
	})

	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	})

	http.HandleFunc("/cpu-stress", func(w http.ResponseWriter, r *http.Request) {
		for {
		}
	})

	http.ListenAndServe(":8080", nil)
}

func memoryStressStart() {
	// 1 GB = 1024 * 1024 * 1024 byte
	const size = 1024 * 1024 * 1024

	// 1 GB data
	data := make([]byte, size)

	// fill data
	for i := 0; i < size; i++ {
		data[i] = 1
	}

}
