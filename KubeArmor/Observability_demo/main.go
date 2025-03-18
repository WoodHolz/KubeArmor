package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// 设置日志
	log.Println("Starting KubeArmor Observability Demo...")

	// 创建策略收集器
	collector := NewPoliciesCollector(10 * time.Second)
	collector.Start()
	defer collector.Stop()

	// 初始化指标
	collector.updateMetrics()

	// 启动 HTTP 服务器提供 Prometheus 指标
	http.Handle("/metrics", promhttp.Handler())
	
	// 添加状态页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>KubeArmor Observability Demo</title></head>
			<body>
				<h1>KubeArmor Observability Demo</h1>
				<p>Metrics available at <a href="/metrics">/metrics</a></p>
			</body>
		</html>`))
	})
	
	// 启动服务器
	log.Println("HTTP server listening on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
