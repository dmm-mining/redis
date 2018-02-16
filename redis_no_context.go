// +build !go1.7

package redis

import (
	"github.com/dmm-mining/redis/internal/pool"
)

type baseClient struct {
	connPool pool.Pooler
	opt      *Options

	process           func(Cmder) error
	processPipeline   func([]Cmder) error
	processTxPipeline func([]Cmder) error

	onClose func() error // hook called when client is closed
}
