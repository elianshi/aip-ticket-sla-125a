# Sprint 2 后端增量交付记录（云原生故障工单与SLA敏捷监控台-Go 后端工程师）

- 在 Sprint 1 MVP 基础上增量实现：
  - `POST /api/tickets/{{id}}/transition` 多态流转与操作人时间线
  - SLA 违约窗口动态计算（P0 15m / P1 1h / P2 4h / P3 24h）
  - `GET /api/services` & `POST /api/services/{{id}}/status` 微服务健康切换与故障注入
  - 扩展单元测试 `TestCreateAndTransitionTicket` 与 `TestServiceStatusToggle`
