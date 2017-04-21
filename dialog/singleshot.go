package dialog

import (
	"fmt"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

// SingleShot is a single shot
type SingleShot struct {
	*Entry
}

func newSingleShot(in *php_serialize.PhpObject) (*SingleShot, error) {
	fmt.Println(in)
	return nil, nil
}
