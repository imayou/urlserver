package application

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
	"urlserver/infrastructure"
)

var urlPage = "http://www.zidu.top/go/%s"
var incrementKey = "url:id:increment:"

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "123456", // no password set
	DB:       2,        // use default DB
})

// 短链接路径转长链接
func ShortToLong(short string) string {
	// 64进制转10进制
	//var url string
	//dec := infrastructure.B64ToDec(short)
	urlKey := fmt.Sprintf("url:short:%s", short)
	val, err := rdb.Get(ctx, urlKey).Result()
	if err != nil {
		log.Printf("查询失败: %s", err.Error())
		return ""
	}
	fmt.Println("key", val)
	return val
}

// 长链接转短链接
func LongToShort(long string, t int64) string {
	lKey := fmt.Sprintf("url:long:%s", long)
	val, _ := rdb.Get(ctx, lKey).Result()
	if val != "" {
		return val
	}
	incr := rdb.Incr(ctx, incrementKey)
	// 十进制转64进制
	strInt64 := strconv.FormatInt(incr.Val(), 10)
	id16, _ := strconv.Atoi(strInt64)
	b64 := infrastructure.DecToB64(id16)
	lUrl := fmt.Sprintf(urlPage, b64)
	rdb.Set(ctx, fmt.Sprintf("url:short:%s", b64), long, time.Duration(t))
	rdb.Set(ctx, fmt.Sprintf("url:long:%s", long), lUrl, time.Duration(t))
	// 组合短链接域名返回
	return lUrl
}
