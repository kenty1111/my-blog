package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

func TestHandler_Index(t *testing.T) {
	tm := template.Must(template.ParseFiles("../assets/index.html"))
	h := New(tm)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	h.Index(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatal("ステータスコードが200ではありません。", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "<h1>私のブログ</h1>") {
		t.Fatal("タイトルがありません")
	}

	titles := []string{"自己紹介", "日記", "仕事", "ブログ"}
	for _, title := range titles {
		if !strings.Contains(rec.Body.String(), fmt.Sprintf("<h2>%s</h2>", title)) {
			t.Fatal("記事のタイトルがありません", titles)
		}
	}
}
