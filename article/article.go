package article

type Article struct {
	Title string
}

func GetAll() ([]Article, error) {
	articles := []Article{
		{
			Title: "自己紹介",
		},
		{
			Title: "日記",
		},
		{
			Title: "仕事",
		},
		{
			Title: "ブログ",
		},
	}
	return articles, nil
}
