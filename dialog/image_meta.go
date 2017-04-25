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
	WhiteBalance    int       `php:"whitebalance"`
	FocalLength     int       `php:"focalLength"`
	Make            string    `php:"make"`
	ExposureProgram int       `php:"exposureProgram"`
}

// IptcData represents iptc data
type IptcData struct {
	// 	Title                         string    `php:"-"`
	// 	Urgency                       string    `php:"-"`
	// 	Category                      string    `php:"-"`
	// 	Keywords                      []string  `php:"-"`
	// 	DateCreated                   time.Time `php:"-"`
	// 	Author                        string    `php:"-"`
	// 	AuthorPosition                string    `php:"-"`
	// 	City                          string    `php:"-"`
	// 	State                         string    `php:"-"`
	// 	Country                       string    `php:"-"`
	// 	Headline                      string    `php:"-"`
	// 	Credit                        string    `php:"-"`
	// 	Source                        string    `php:"-"`
	// 	CopyrightNotice               string    `php:"-"`
	// 	Caption                       string    `php:"-"`
	// 	Writer                        string    `php:"-"`
	// 	SpecialInstruction            string    `php:"-"`
	// 	SupplementalCategories        string    `php:"-"`
	// 	OriginalTransmissionReference string    `php:"-"`
}
