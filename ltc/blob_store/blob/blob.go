package blob

import "time"

type Blob struct {
	Path    string
	Created time.Time
	Size    int64
}
