package consts

import "github.com/gogf/gf/v2/os/gcache"

var (
	Cache *gcache.Cache
)

func init() {
	Cache = gcache.New()
	Cache.SetAdapter(gcache.NewAdapterMemory())
}
