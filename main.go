package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// コマンドラインフラグを定義
	insecure := flag.Bool("k", false, "allow insecure server connections when using SSL")
	flag.BoolVar(insecure, "insecure", false, "allow insecure server connections when using SSL")
	headOnly := flag.Bool("I", false, "show response headers only")
	flag.BoolVar(headOnly, "head", false, "show response headers only")
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

	// HTTPトランスポートを作成（プロキシ環境変数サポートのためDefaultTransportをベースにする）
	transport := http.DefaultTransport.(*http.Transport).Clone()

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

	// 標準出力がターミナルかどうかを確認（パイプ/リダイレクト検出）
	stat, err := os.Stdout.Stat()
	isPiped := err != nil || (stat.Mode()&os.ModeCharDevice) == 0

	// 出力ロジック
	showHeaders := !isPiped || *headOnly
	showBody := !*headOnly

	// ヘッダーを表示
	if showHeaders {
		fmt.Println(resp.Status)
		fmt.Println()
		for key, values := range resp.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", key, value)
			}
		}
	}

	// ヘッダーとボディの間に空行を追加
	if showHeaders && showBody {
		fmt.Println()
	}

	// ボディを表示
	if showBody {
		io.Copy(os.Stdout, resp.Body)
	}
}
