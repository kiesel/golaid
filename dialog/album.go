package dialog

// Album represents Dialog Album
type Album struct {
	*Entry     `recurse:"true"`
	Highlights []Image   `php:"highlights"`
	Chapters   []Chapter `php:"chapters"`
}
