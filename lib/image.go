package lib

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/kolesa-team/go-webp/webp"
	imgupload "github.com/olahol/go-imageupload"
)

// 1M
var fileLimitSize int = 1000000

func ConvertToWebp(file *imgupload.Image) (*bytes.Buffer, error) {
	if file.Size > fileLimitSize {
		return nil, fmt.Errorf("ファイルサイズが大きすぎます。1MB以下にしてください。")
	}

	var image image.Image
	readerFile := bytes.NewReader(file.Data)
	switch file.ContentType {
	case "image/jpeg":
		// JPEG画像をデコード
		img, err := jpeg.Decode(readerFile)
		if err != nil {
			return nil, fmt.Errorf("JPEG画像をデコードに失敗しました")
		}
		image = img

	case "image/png":
		// PNG画像をデコード
		img, err := png.Decode(readerFile)
		if err != nil {
			return nil, fmt.Errorf("PNG画像をデコードに失敗しました")
		}
		image = img
	case "image/webp":
		// WEBP画像をデコード
		img, err := webp.Decode(readerFile, nil)
		if err != nil {
			return nil, fmt.Errorf("WEBP画像をデコードに失敗しました")
		}
		image = img
	default:
		return nil, fmt.Errorf("対応していないファイルです")
	}

	// WebPに変換
	webpBuffer := new(bytes.Buffer)
	err := webp.Encode(webpBuffer, image, nil)
	if err != nil {
		return nil, fmt.Errorf("WEBP画像にエンコードに失敗しました")
	}

	return webpBuffer, nil
}

func UploadImage(img *bytes.Buffer, filename, place string) (string, error) {
	pass := "./image/" + place + "/" + filename

	file, err := os.Create(pass)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, img)
	if err != nil {
		return "", err
	}
	return pass, nil
}
