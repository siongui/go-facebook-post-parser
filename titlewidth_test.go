package parsefb

import (
	"testing"
)

func TestTitleWidth(t *testing.T) {
	if l := TitleWidth("เราให้ความสำคัญกับอารมณ์มากเหลือเกิน ... - ธรรมะ โดย พระอาจารย์ชยสาโร"); l != 69 {
		t.Error("string length error: ", l)
		return
	}
}

func TestSetRstDocTitleWidth(t *testing.T) {
	s := "Hello Word\n####\n\n"
	r, err := SetRstDocTitleWidth(s)
	if err != nil {
		t.Error(err)
		return
	}

	if r != "Hello Word\n##########\n\n" {
		t.Error(r)
		return
	}
}
