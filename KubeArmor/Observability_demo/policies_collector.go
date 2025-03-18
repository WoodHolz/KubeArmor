package main

import (
	"log"
	"net/http"
	"time"
)

// 简化的策略结构体
type Policy struct {
	Metadata struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Selector struct {
			MatchLabels map[string]string `json:"matchLabels"`
		} `json:"selector"`
	} `json:"spec"`
}

// 简化的 kubectl 输出结构体
type KubectlOutput struct {
	Items []Policy `json:"items"`
}

// 策略收集器
type PoliciesCollector struct {
	updateInterval time.Duration
	stopCh         chan struct{}
}

// 创建新的策略收集器
func NewPoliciesCollector(interval time.Duration) *PoliciesCollector {
	return &PoliciesCollector{
		updateInterval: interval,
		stopCh:         make(chan struct{}),
	}
}

// 开始收集策略指标
func (pc *PoliciesCollector) Start() {
	go func() {
		ticker := time.NewTicker(pc.updateInterval)
		defer ticker.Stop()

		for {
			select {
			case <-pc.stopCh:
				return
			case <-ticker.C:
				pc.updateMetrics()
			}
		}
	}()
}

// 停止收集器
func (pc *PoliciesCollector) Stop() {
	close(pc.stopCh)
}

// 更新策略指标
func (pc *PoliciesCollector) updateMetrics() {
	// 使用临时数据来模拟 kubectl 输出
	// 实际实现应该使用 kubectl 或 KubeArmor API
	containerPolicies, err := getContainerPolicies()
	if err != nil {
		log.Printf("Error getting container policies: %v", err)
	}

	hostPolicies, err := getHostPolicies()
	if err != nil {
		log.Printf("Error getting host policies: %v", err)
	}

	// 计算总策略数
	totalPolicies := len(containerPolicies) + len(hostPolicies)
	TotalPoliciesApplied.Set(float64(totalPolicies))

	// 更新按类型分类的策略数量
	PoliciesByType.WithLabelValues("container").Set(float64(len(containerPolicies)))
	PoliciesByType.WithLabelValues("host").Set(float64(len(hostPolicies)))

	// 更新按命名空间分类的策略数量
	namespaceCount := make(map[string]int)
	for _, policy := range containerPolicies {
		namespace := policy.Metadata.Namespace
		namespaceCount[namespace]++
	}

	// 重置之前的指标
	PoliciesByNamespace.Reset()
	
	// 更新命名空间指标
	for namespace, count := range namespaceCount {
		PoliciesByNamespace.WithLabelValues(namespace).Set(float64(count))
	}

	log.Printf("Updated policy metrics: %d total policies (%d container, %d host)", 
		totalPolicies, len(containerPolicies), len(hostPolicies))
}

// 获取容器安全策略
// 在实际实现中，应该使用 kubectl 或 KubeArmor API
func getContainerPolicies() ([]Policy, error) {
	// 临时加载演示数据
	// 在实际实现中，可以使用类似以下命令:
	// kubectl get KubeArmorPolicies --all-namespaces -o json
	
	/*
	cmd := exec.Command("kubectl", "get", "KubeArmorPolicies", "--all-namespaces", "-o", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute kubectl: %w", err)
	}

	var result KubectlOutput
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kubectl output: %w", err)
	}

	return result.Items, nil
	*/

	// 演示数据
	policies := []Policy{
		createDemoPolicy("policy1", "default"),
		createDemoPolicy("policy2", "default"),
		createDemoPolicy("policy3", "kube-system"),
		createDemoPolicy("policy4", "app"),
		createDemoPolicy("policy5", "app"),
	}
	
	return policies, nil
}

// 获取主机安全策略
func getHostPolicies() ([]Policy, error) {
	/*
	cmd := exec.Command("kubectl", "get", "KubeArmorHostPolicies", "-o", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute kubectl: %w", err)
	}

	var result KubectlOutput
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kubectl output: %w", err)
	}

	return result.Items, nil
	*/

	// 演示数据
	policies := []Policy{
		createDemoPolicy("host-policy1", ""),
		createDemoPolicy("host-policy2", ""),
	}
	
	return policies, nil
}

// 创建演示策略
func createDemoPolicy(name string, namespace string) Policy {
	policy := Policy{}
	policy.Metadata.Name = name
	policy.Metadata.Namespace = namespace
	policy.Spec.Selector.MatchLabels = map[string]string{"app": "demo"}
	return policy
}

func main() {
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", nil))
}