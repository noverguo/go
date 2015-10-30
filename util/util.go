package util
import (
	"log"
	"os"
	"bytes"
	"mime/multipart"
	"io"
	"net/http"
)



// 检查错误
func CheckErr(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}


func CleanDir(dir string) string {
	endVal := dir[len(dir)-1]
	if endVal != '/' && endVal != '\\' {
		dir += "/"
	}
	return dir
}

// 文件拷贝
func Copy(inPath, outPath string) {
	inFile, err := os.Open(inPath)
	CheckErr(err)
	defer inFile.Close()
	outFile, err := os.Create(outPath)
	CheckErr(err)
	defer outFile.Close()
	io.Copy(outFile, inFile)
}

// 包装文件上传请求
func NewFileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request, err
}