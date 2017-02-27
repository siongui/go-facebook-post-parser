package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/www.masterchingche.org/photos/a.1602539563313910.1073741827.1601992876701912/1928528887381641/?type=3"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "你如果現在思考模式還是很衝動，最後死路一條！"
	post.Summary = "這個古代的人善根深厚，你懾服他，你罵他，他起慚愧心。他馬上跪下來說：請和尚開示如何讓我醒過來？我不遠千里而來，就是希望您指點我一條明路，我不知道怎麼下手。雪峰禪師講了一句話而已：出去。然後轉身就走了。"
	tmplpath, filename := GetTemplatePath(post)
	rst, err := ToreStructuredText(post, tmplpath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rst)

	err = SaveRst(filename, rst)
	if err != nil {
		t.Error(err)
	}
}
