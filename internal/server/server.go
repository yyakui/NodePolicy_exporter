package server

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"

    "nodepolicy_exporter/internal/collector"
    "nodepolicy_exporter/internal/types"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer() {
    http.Handle("/metrics", promhttp.Handler())

    // 调度器任务入口
    http.HandleFunc("/run_tasks", func(w http.ResponseWriter, r *http.Request) {
        body, _ := ioutil.ReadAll(r.Body)
        var tasks []types.Task
        json.Unmarshal(body, &tasks)

        for _, t := range tasks {
            collector.RunTask(t)
        }
        w.Write([]byte("ok"))
    })

    // GET 五元组探测
    http.HandleFunc("/probe", func(w http.ResponseWriter, r *http.Request) {
        var t types.Task
        t.TaskID = "api-probe"
        t.SourceIP = r.URL.Query().Get("source_ip")
        t.TargetIP = r.URL.Query().Get("target_ip")
        fmt.Sscanf(r.URL.Query().Get("target_port"), "%d", &t.TargetPort)
        t.Protocol = "TCP"
        t.TimeoutMs = 500

        // 直接执行探测
        collector.RunTask(t)

        // 返回探测结果（1 或 0）
        w.Write([]byte("ok"))
    })

    // POST 五元组探测
    http.HandleFunc("/probe_json", func(w http.ResponseWriter, r *http.Request) {
        body, _ := ioutil.ReadAll(r.Body)
        var t types.Task
        json.Unmarshal(body, &t)

        collector.RunTask(t)
        w.Write([]byte("ok"))
    })

    log.Println("nodepolicy_exporter running on :9100")
    http.ListenAndServe(":9100", nil)
}

