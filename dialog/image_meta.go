package dialog

import (
	"time"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

// ExifData represents exif data
type ExifData struct {
	ApertureFNumber string    `php:"apertureFNumber"`
	Software        string    `php:"software"`
	ExposureTime    string    `php:"exposureTime"`
	MeteringMode    int       `php:"meteringMode"`
	Flash           int       `php:"flash"`
	Orientation     int       `php:"orientation"`
	FileSize        int       `php:"fileSize"`
	DateTime        time.Time `php:"dateTime"`
	Model           string    `php:"model"`
	MimeType        string    `php:"mimeType"`
	WhiteBalance    int       `php:"whiteBalance"`
	FocalLength     int       `php:"focalLength"`
	Make            string    `php:"make"`
	ExposureProgram int       `php:"exposureProgram"`
}

// IptcData represents iptc data
type IptcData struct {
}

func newExifData(in *php_serialize.PhpObject) (ExifData, error) {
	out, err := newObject(&ExifData{}, in)
	if err != nil {
		return ExifData{}, err
	}

	return out.(ExifData), nil
}
