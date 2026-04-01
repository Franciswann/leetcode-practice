package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// captureOutput 捕獲函數輸出到 stdout 的內容
func captureOutput(fn func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

// TestDFS 測試深度優先搜尋的主要功能
func TestDFS(t *testing.T) {
	output := captureOutput(func() {
		printDFS("pics")
	})

	// DFS 應該按照深度優先順序印出檔案：先進入子目錄，再處理檔案
	// 預期順序：a.png, space.png, odyssey.png
	lines := strings.TrimSpace(output)
	expected := "a.png\nspace.png\nodyssey.png"

	if lines != expected {
		t.Errorf("DFS 輸出順序錯誤。\n實際輸出:\n%s\n預期輸出:\n%s", lines, expected)
	}
}

// TestBFS 測試廣度優先搜尋的主要功能
func TestBFS(t *testing.T) {
	output := captureOutput(func() {
		printBFS("pics")
	})

	// BFS 應該按照層級順序印出檔案：同層級的先處理完
	// 預期順序：odyssey.png, a.png, space.png
	lines := strings.TrimSpace(output)
	expected := "odyssey.png\na.png\nspace.png"

	if lines != expected {
		t.Errorf("BFS 輸出順序錯誤。\n實際輸出:\n%s\n預期輸出:\n%s", lines, expected)
	}
}
