# syncdebug

### drop-in replacement for "sync" while debugging

Have a mutex that just wont unlock somewhere? Maybe this will help.

### limited

If you program uses any more of the "sync" package than these, it wont work.
  * sync.Mutex
  * sync.WaitGroup

### example

```go
import (
	//	"sync"
	sync "github.com/aerth/syncdebug"
)
```
