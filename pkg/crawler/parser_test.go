package crawler

import (
	"context"
	"testing"
)

func TestMakeReadable(t *testing.T) {
	tt := []struct {
		Name       string
		URL        string
		WantedType ContentType
	}{
		{Name: "HTML page", URL: "https://m.povar.ru/recipes/oladi_na_drojjah_na_vode-40123.html", WantedType: Text},
		{Name: "PDF file", URL: "https://arxiv.org/pdf/2407.16833", WantedType: PDF},
	}

	for _, test := range tt {
		t.Run(test.Name, func(t *testing.T) {

			raw_content, err := GetContent(context.Background(), test.URL)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(raw_content.Content)
		})
	}
}
