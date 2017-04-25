package dialog

// Image represents Dialog Image
type Image struct {
	Name     string   `php:"name"`
	Width    int64    `php:"width"`
	Height   int64    `php:"height"`
	ExifData ExifData `php:"exifData"`
	// IptcData IptcData `php:"iptcData"`
}
