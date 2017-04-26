package dialog

// EntryCollection is a collection of entries
type EntryCollection struct {
	*Entry
	Entries []IEntry `php:"entries"`
}
