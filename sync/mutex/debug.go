package mutex

import "sync/atomic"

type debugInfo struct {
	ownerStack atomic.Pointer[string]
}
