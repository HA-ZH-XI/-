package global

import (
	"container/list"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/beego/beego/v2/core/config"
	"github.com/patrickmn/go-cache"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"sync"
)

/*
*
这些对象为什么要定义成全局的，说明这对象要在启动的时候就必须初始化，
然后才能够在controllers/servcice等进行使用。
*/
var (
	List       list.List
	Env        string
	Log        *zap.Logger
	SugarLog   *zap.SugaredLogger
	Lock       sync.RWMutex
	Yaml       map[string]interface{}
	Config     config.Configer
	LocalCache *cache.Cache
	BlackCache local_cache.Cache
	OssClient  *oss.Client
	RedisKey   string
)
