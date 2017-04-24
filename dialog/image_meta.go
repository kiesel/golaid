package dialog

import "time"

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
