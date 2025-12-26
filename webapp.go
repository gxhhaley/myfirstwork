package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 定义处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 只处理GET请求
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 获取查询参数
		query := r.URL.Query().Get("input")

		// 设置响应头
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		// 根据输入返回不同的响应
		switch query {
		case "1":
			fmt.Fprintln(w, "hello")
		case "2":
			fmt.Fprintln(w, "world")
		case "":
			// 如果没有输入参数，显示使用说明
			fmt.Fprintln(w, "Welcome! Please add ?input=1 or ?input=2 to the URL")
			fmt.Fprintln(w, "Examples:")
			fmt.Fprintln(w, "  http://localhost:5555/?input=1  returns 'hello'")
			fmt.Fprintln(w, "  http://localhost:5555/?input=2  returns 'world'")
		default:
			fmt.Fprintf(w, "Unknown input: %s (only 1 or 2 are supported)", query)
		}
	})

	// 启动服务器
	port := "0.0.0.0:5555"
	fmt.Printf("Starting server on http://localhost%s\n", port)
	fmt.Println("Usage:")
	fmt.Println("  Open browser and visit: http://localhost:5555/?input=1")
	fmt.Println("  or visit: http://localhost:5555/?input=2")

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
