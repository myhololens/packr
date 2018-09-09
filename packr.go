package packr

import (
	"sync"

	"github.com/gobuffalo/packr/file/resolver"
)

var boxes = map[string]*Box{}
var gil = &sync.RWMutex{}

func findBox(name string) *Box {
	gil.RLock()
	defer gil.RUnlock()
	return boxes[resolver.Key(name)]
}

func placeBox(b *Box) *Box {
	gil.Lock()
	defer gil.Unlock()
	boxes[resolver.Key(b.Name)] = b
	return b
}