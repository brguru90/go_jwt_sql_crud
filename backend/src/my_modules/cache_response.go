package my_modules

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"learn_go/src/database"
	"math"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	log "github.com/sirupsen/logrus"
)

const one_sec = 1000000000

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GetCachedResponse(view_func func(*gin.Context), table_name string, cache_ttl_secs time.Duration, custom_cache_prefix func(*gin.Context) string) func(c *gin.Context) {

	return func(c *gin.Context) {

		REDIS_CACHE := cache.New(&cache.Options{
			Redis:      database.REDIS_DB_CONNECTION,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		})
		_uri := c.Request.RequestURI
		route_path := c.FullPath()
		_raw_dt, _ := c.GetRawData()
		_params, _ := json.Marshal(c.Params)
		_prefix := ""
		if custom_cache_prefix != nil {
			_prefix = "___" + custom_cache_prefix(c)
		}

		h := sha1.New()
		h.Write([]byte(_uri + string(_params) + string(_raw_dt)))
		cache_key := table_name + _prefix + "___" + route_path + "___" + string(h.Sum(nil))

		var responseCache ResponseCacheStruct
		var cache_mis_err error
		{
			// getting data from cache
			cache_mis_err = REDIS_CACHE.Get(c.Request.Context(), cache_key, &responseCache)
			if cache_mis_err == nil {
				_now := time.Now()

				log.Debugln("cache hit --> " + route_path)

				c.Writer.Header().Set("Content-Type", responseCache.ContentType)
				c.Writer.Header().Set("From-cache", "true")
				c.Writer.Header().Set("From-cache-TTL-left-Secs", fmt.Sprintf("%v", (cache_ttl_secs-_now.Sub(responseCache.LastModified)).Seconds()))
				c.String(responseCache.HTTPStatusCode, string(responseCache.ResponseData))

				{
					// renewing cache expiry if 25% of Time To Live(TTL) value is elapsed
					cache_ttl_sec_3quarter := time.Duration(one_sec * int(math.Floor(cache_ttl_secs.Seconds()*0.75)))
					if _now.Sub(responseCache.LastModified) >= cache_ttl_sec_3quarter {
						log.Debugln("cache Renewing --> " + route_path)
						responseCache.LastModified = _now
						err := REDIS_CACHE.Set(&cache.Item{
							Ctx:   c.Request.Context(),
							Key:   cache_key,
							Value: responseCache,
							TTL:   cache_ttl_secs,
						})
						if err != nil {
							log.WithFields(log.Fields{
								"err":     err,
								"_uri":    _uri,
								"_params": _params,
								"_raw_dt": _raw_dt,
							}).Error("while caching response")
						}
					}
				}
				return
			}
			if strings.Contains(strings.ToLower(fmt.Sprintf("%s", cache_mis_err)), "key is missing") {
				log.Debugln("cache miss --> " + route_path)
			} else {
				log.WithFields(log.Fields{
					"err":     cache_mis_err,
					"_uri":    _uri,
					"_params": _params,
					"_raw_dt": _raw_dt,
				}).Error("while getting response from cache")
			}
		}

		w := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w

		view_func(c)

		{
			// creating a new cache from response
			responseCache = ResponseCacheStruct{
				ResponseData:   w.body.String(),
				ContentType:    c.Writer.Header().Get("Content-Type"),
				HTTPStatusCode: c.Writer.Status(),
				LastModified:   time.Now(),
			}
			err := REDIS_CACHE.Set(&cache.Item{
				Ctx:   c.Request.Context(),
				Key:   cache_key,
				Value: responseCache,
				TTL:   cache_ttl_secs,
			})
			if err != nil {
				log.WithFields(log.Fields{
					"err":     err,
					"_uri":    _uri,
					"_params": _params,
					"_raw_dt": _raw_dt,
				}).Error("while caching response")
			}

			REDIS_CACHE.Set(&cache.Item{
				Ctx:   c.Request.Context(),
				Key:   "users_table",
				Value: "cached",
				TTL:   cache_ttl_secs,
			})
		}
	}

}
