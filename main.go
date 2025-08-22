package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// コマンドラインフラグを定義
	insecure := flag.Bool("k", false, "allow insecure server connections when using SSL")
	flag.BoolVar(insecure, "insecure", false, "allow insecure server connections when using SSL")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [URL]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	// URLを取得
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	url := args[0]

	// HTTPトランスポートを作成
	transport := &http.Transport{}

	// insecureオプションが指定されている場合、TLS証明書の検証をスキップ
	if *insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// HTTPクライアントを作成（リダイレクト無効）
	client := &http.Client{
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// GETリクエストを実行
	resp, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// レスポンスヘッダを表示
	fmt.Println(resp.Status)
	fmt.Println()
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
