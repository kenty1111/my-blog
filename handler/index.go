package handler

import (
	"log"
	"my-blog/article"
	"net/http"
)

func (h *Handler) Index(w http.ResponseWriter, req *http.Request) {
	articles, err := article.GetAll()
	if err != nil {
		log.Println("記事の一覧の取得に失敗しました", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := struct {
		Title    string
		Articles []article.Article
	}{
		Title:    "私のブログ",
		Articles: articles,
	}
	h.templateIndex.Execute(w, params)
}
