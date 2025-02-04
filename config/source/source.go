// Package source is the interface for sources
package source

import (
	"errors"
	"time"
)

var (
	// ErrWatcherStopped is returned when source watcher has been stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)

// Source is the source from which config is loaded
type Source interface {
	Read() (*ChangeSet, error)
	Watch() (Watcher, error)
	String() string
}

// ChangeSet represents a set of changes from a source
type ChangeSet struct {
	Data      []byte
	Checksum  string // 校验和
	Format    string // 数据的格式json等
	Source    string
	Timestamp time.Time
}

// Watcher watches a source for changes
type Watcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}
