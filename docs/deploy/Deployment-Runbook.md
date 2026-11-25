# 《云原生故障工单与SLA敏捷监控台》本地与受控服务器部署手册

- **运维负责人**：Titan（运维部署工程师）
- **编译与测试命令**：`cd services/backend && go test ./... && go build -o /tmp/aip-deployed-app ./cmd/server`
- **本地启动命令**：`/tmp/aip-deployed-app -port 18800 -web-dir .`
- **健康探针**：`GET /healthz` 预期返回 `200 OK`
