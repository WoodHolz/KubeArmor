package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// 已应用策略总数
	TotalPoliciesApplied = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "kubearmor_policies_applied_total",
			Help: "Total number of KubeArmor security policies currently applied",
		},
	)

	// 按类型(容器/主机)分类的策略数量
	PoliciesByType = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kubearmor_policies_by_type",
			Help: "Number of KubeArmor security policies by type",
		},
		[]string{"type"}, // type: container/host
	)

	// 按命名空间分类的策略数量
	PoliciesByNamespace = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kubearmor_policies_by_namespace",
			Help: "Number of KubeArmor security policies by namespace",
		},
		[]string{"namespace"}, // namespace name
	)
)
