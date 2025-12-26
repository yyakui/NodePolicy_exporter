# NodePolicy_exporter
Prometheus-compatible agent for network connectivity probing, scheduler tasks, and real-time API detection.


# 🚀 nodepolicy_exporter

A lightweight, high‑performance network connectivity & policy probe agent for Prometheus.  
Supports scheduler‑based tasks, real‑time GET/POST probing, Prometheus metrics, and enterprise‑grade RPM/YUM deployment.

---

## ✨ Features

- 🔥 **Five‑tuple TCP connectivity probing**
- 📡 **Scheduler task execution** (`/run_tasks`)
- ⚡ **Real‑time GET/POST probe API** (`/probe`, `/probe_json`)
- 📊 **Prometheus metrics export** (`/metrics`)
- 📦 **RPM packaging & internal YUM repository support**
- 🧩 **Extensible Linux Ops features** (CPU, memory, ports, processes)
- 🧵 **High‑performance multi‑thread probe engine** (planned)
- 🛠 **Zero dependency, single binary, production‑ready**

---

## 🏗 Architecture

```
Scheduler → nodepolicy_exporter → Probe → Metrics → Prometheus → Grafana
```

### Detailed Flow

```
┌──────────────┐      ┌──────────────────────────┐
│  Scheduler   │      │  nodepolicy_exporter     │
│ /api/tasks   │ ---> │  - /run_tasks            │
└──────┬───────┘      │  - /probe /probe_json    │
       │              │  - /metrics              │
       ▼              └───────────┬──────────────┘
   JSON Tasks                      │
                                   ▼
                          Prometheus Scrape
                                   │
                                   ▼
                                Grafana
```

---

## 🚀 Quick Start

### 1. Run binary

```bash
./nodepolicy_exporter
```

### 2. Test GET probe

```bash
curl "http://127.0.0.1:9100/probe?source_ip=127.0.0.1&target_ip=127.0.0.1&target_port=22"
```

### 3. Test POST probe

```bash
curl -X POST http://127.0.0.1:9100/probe_json \
  -H "Content-Type: application/json" \
  -d '{"task_id":"t1","source_ip":"127.0.0.1","target_ip":"127.0.0.1","target_port":22,"protocol":"TCP","timeout_ms":500}'
```

### 4. Test scheduler task

```bash
curl -X POST http://127.0.0.1:9100/run_tasks \
  -H "Content-Type: application/json" \
  -d '[{"task_id":"probe-22","source_ip":"127.0.0.1","target_ip":"127.0.0.1","target_port":22,"protocol":"TCP","timeout_ms":500}]'
```

---

## 📊 Prometheus Integration

Add this to your `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'nodepolicy_exporter'
    scrape_interval: 5s
    static_configs:
      - targets: ['127.0.0.1:9100']
```

Query example:

```promql
nodepolicy_connectivity_status
```

---

## 📦 RPM / YUM Deployment

### Install RPM

```bash
yum install -y nodepolicy_exporter-1.0.0.rpm
```

### Start service

```bash
systemctl enable --now nodepolicy_exporter
```

### Check status

```bash
systemctl status nodepolicy_exporter
```

---

## 📁 Directory Structure

```
nodepolicy_exporter/
├── cmd/
├── internal/
│   ├── server/
│   ├── collector/
│   ├── probe/
│   └── types/
├── build/
│   └── rpm/
├── README.md
├── LICENSE
└── examples/
```

---

## 🛣 Roadmap

- [ ] Multi‑thread probe engine
- [ ] UDP / ICMP probe support
- [ ] Auto scheduler pull mode (Agent mode)
- [ ] CMDB‑based dynamic task generation
- [ ] Port listening auto‑discovery
- [ ] Process health monitoring
- [ ] Grafana dashboard templates
- [ ] TLS support for APIs
