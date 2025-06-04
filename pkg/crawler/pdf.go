package crawler

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/dslipak/pdf"
)

func ProcessPDF(file *os.File) (string, error) {
	fi, err := file.Stat()
	if err != nil {
		file.Close()
		return "", err
	}

	r, err := pdf.NewReader(file, fi.Size())
	defer file.Close()

	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	buf.ReadFrom(b)
	return buf.String(), nil
}

func DownloadFile(url string) (*os.File, error) {
	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	tempFile, err := os.CreateTemp("", "downloaded-*.pdf")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(tempFile, httpResponse.Body)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, err
	}

	_, err = tempFile.Seek(0, 0)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, err
	}

	return tempFile, nil
}
