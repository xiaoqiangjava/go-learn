package main

import "fmt"

/*
map 基本语法：
var name map[keytype]valtype
1. map在使用前一定要用make函数来开辟空间
2. map的key是不可以重复的
3. map中的kv是无序的
*/
func main() {
	// 方式一： 定义一个map, 容量是10
	var kv map[string]string = make(map[string]string, 10)
	fmt.Println("kv before: ", kv)
	kv["张三"] = "java"
	kv["李四"] = "python"
	fmt.Println(kv)
	// 方式二：不指定容量
	var m = make(map[int]string)
	m[1] = "A"
	m[2] = "B"
	fmt.Println(m)
	// 方式三：直接赋值
	var m1 = map[int]string{1: "a", 2: "b"}
	fmt.Println(m1)

	var metrics []Metric
	metrics = append(metrics, Metric{Stage: "asr", Count: 1, Success: 1, Cost: 200})
	metrics = append(metrics, Metric{Stage: "asr", Count: 1, Success: 1, Cost: 200})
	metrics = append(metrics, Metric{Stage: "asr", Count: 1, Success: 1, Cost: 200})
	metrics = append(metrics, Metric{Stage: "asr", Count: 1, Success: 1, Cost: 300})
	metrics = append(metrics, Metric{Stage: "asr", Count: 1, Success: 1, Cost: 200})
	metrics = append(metrics, Metric{Stage: "asr", Count: 1, Success: 1, Cost: 300})

	var metricMap = make(map[string]*Metric, 5)
	for i := 0; i < len(metrics); i++ {
		metric := &metrics[i]
		if s, ok := metricMap[metric.Stage]; ok {
			s.Count += metric.Count
			s.Success += metric.Success
			s.Cost = (s.Cost + metric.Cost) / 2
		} else {
			metricMap[metric.Stage] = metric
		}
	}

	fmt.Printf("metric: %+v", *metricMap["asr"])
}

type Metric struct {
	Stage   string
	Count   float64
	Success float64
	Cost    float64
}
