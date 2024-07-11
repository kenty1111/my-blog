package article

import "testing"

func TestGetAll(t *testing.T) {
	got, err := GetAll()
	if err != nil {
		t.Fatal("記事の取得に失敗しました", err)
	}
	if len(got) != 4 {
		t.Fatal("記事数が4ではありません", len(got))
	}
	testEq(t, "自己紹介", got[0].Title)
	testEq(t, "日記", got[1].Title)
	testEq(t, "仕事", got[2].Title)
	testEq(t, "ブログ", got[3].Title)
}

func testEq(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Fatalf("want: %v\ngot: %v", want, got)
	}
}
