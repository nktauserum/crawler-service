package crawler

import "testing"

func TestContentType(t *testing.T) {
	tt := []struct {
		Name       string
		URL        string
		WantedType ContentType
	}{
		{Name: "PDF", URL: "https://arxiv.org/pdf/2407.16833", WantedType: PDF},
		{Name: "HTML", URL: "https://m.povar.ru/recipes/oladi_na_drojjah_na_vode-40123.html", WantedType: Text},
		{Name: "Unknown", URL: "https://masterpiecer-images.s3.yandex.net/67222f887d3b11ee9eb2ceda526c50ab:upscaled", WantedType: Unknown},
	}

	for _, test := range tt {
		t.Run(test.Name, func(t *testing.T) {
			c_type, err := CheckContentType(test.URL)
			if err != nil {
				t.Fatal(err)
			}

			if c_type != test.WantedType {
				t.Errorf("got content type %q, want %q", c_type, test.WantedType)
			}
		})
	}
}
