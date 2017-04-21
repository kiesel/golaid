package dialog

import (
	"fmt"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

// Chapter represents a chapter
type Chapter struct {
}

func newChapter(in *php_serialize.PhpArray) (*Chapter, error) {
	fmt.Println(in)
	return nil, nil
}
