package storage

import "context"

// ExecCloser 存取数据，并关闭
type ExecCloser interface {
	Put(ctx context.Context, data []byte) (string, error)
	Get(ctx context.Context, cid string) ([]byte, error)
	Close() error
}
