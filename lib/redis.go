package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var redisClient *redis.Client

func GetRedisDB() *redis.Client {
	return redisClient
}

func init() {
	addr := "localhost"
	port := "6379"
	password := ""
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password,
		DB:       0,
		PoolSize: 100,
	})
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := redisClient.Ping(cxt).Result()
	if err != nil {
		fmt.Printf("redis ping err::%#v\n", err)
		panic(err)
	}
}

//RedisSet 设置缓存，指定过期时间
func RedisSet(key string, value interface{}, expiration time.Duration) error {
	if err := redisClient.Set(context.Background(), key, value, expiration).Err(); err != nil {
		fmt.Printf("redis cmd set failed, err::%#v\n", err)
		return err
	} else {
		return nil
	}
}

//RedisGet 获取指定的缓存
func RedisGet(key string) (string, error) {
	if result, err := redisClient.Get(context.Background(), key).Result(); err != nil {
		fmt.Printf("redis cmd get failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisDel 删除指定的key
func RedisDel(keys string) (int64, error) {
	if result, err := redisClient.Del(context.Background(), keys).Result(); err != nil {
		fmt.Printf("redis cmd del failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisExpire 设置指定key的过期时间
func RedisExpire(key string, expiration time.Duration) error {
	if err := redisClient.Expire(context.Background(), key, expiration).Err(); err != nil {
		fmt.Printf("redis cmd expire failed, err::%#v\n", err)
		return err
	} else {
		return nil
	}
}

//RedisSetNX 实现分布式加锁
func RedisSetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if result, err := redisClient.SetNX(context.Background(), key, value, expiration).Result(); err != nil {
		fmt.Printf("redis cmd setnx failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisIncr 自增序列
func RedisIncr(key string) (int64, error) {
	if result, err := redisClient.Incr(context.Background(), key).Result(); err != nil {
		fmt.Printf("redis cmd incr failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisSetBit 位图操作
func RedisSetBit(key string, offset int64, value int) error {
	if err := redisClient.SetBit(context.Background(), key, offset, value).Err(); err != nil {
		fmt.Printf("redis cmd setbit failed, err::%#v\n", err)
		return err
	} else {
		return nil
	}
}

//RedisGetBit 位图读取
func RedisGetBit(key string, offset int64) (int64, error) {
	if result, err := redisClient.GetBit(context.Background(), key, offset).Result(); err != nil {
		fmt.Printf("redis cmd getbit failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisLPush 从左侧向list压入数据
func RedisLPush(key string, value ...interface{}) (int64, error) {
	if result, err := redisClient.LPush(context.Background(), key, value).Result(); err != nil {
		fmt.Printf("redis cmd lpush failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisRPush 从右侧向list压入数据
func RedisRPush(key string, value ...interface{}) (int64, error) {
	if result, err := redisClient.RPush(context.Background(), key, value).Result(); err != nil {
		fmt.Printf("redis cmd rpush failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisLRange 从
func RedisLRange(key string, start int64, stop int64) ([]string, error) {
	if result, err := redisClient.LRange(context.Background(), key, start, stop).Result(); err != nil {
		fmt.Printf("redis cmd lrange failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisLPop 从左侧pop数据
func RedisLPop(key string) (string, error) {
	if result, err := redisClient.LPop(context.Background(), key).Result(); err != nil {
		fmt.Printf("redis cmd lpop failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisRPop 从右侧pop数据
func RedisRPop(key string) (string, error) {
	if result, err := redisClient.RPop(context.Background(), key).Result(); err != nil {
		fmt.Printf("redis cmd rpop failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisLLen 获取列表中的数据长度
func RedisLLen(key string) (int64, error) {
	if result, err := redisClient.LLen(context.Background(), key).Result(); err != nil {
		fmt.Printf("redis cmd llen failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisSAdd 将一个或多个member元素加入到集合key当中，已经存在于集合的member元素将被忽略
func RedisSAdd(key string, members ...interface{}) error {
	if err := redisClient.SAdd(context.Background(), key, members).Err(); err != nil {
		fmt.Printf("redis cmd sadd failed, err::%#v\n", err)
		return err
	} else {
		return nil
	}
}

//RedisSIsMember 判断 member 元素是否集合 key 的成员
func RedisSIsMember(key string, member interface{}) (bool, error) {
	if result, err := redisClient.SIsMember(context.Background(), key, member).Result(); err != nil {
		fmt.Printf("redis cmd sismember failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisSPop 移除并返回集合中的一个随机元素
func RedisSPop(key string) (string, error) {
	if result, err := redisClient.SPop(context.Background(), key).Result(); err != nil {
		fmt.Printf("redis cmd spop failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisSRandMember 返回集合中的一个随机元素
func RedisSRandMember(key string) (string, error) {
	if result, err := redisClient.SRandMember(context.Background(), key).Result(); err != nil {
		fmt.Printf("redis cmd srandmember failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisZAdd 将一个member元素及其score值加入到有序集key当中
func RedisZAdd(key string, score float64, member interface{}) (int64, error) {
	items := []*redis.Z{
		&redis.Z{Score: score, Member: member},
	}
	if result, err := redisClient.ZAdd(context.Background(), key, items...).Result(); err != nil {
		fmt.Printf("redis cmd zadd failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisZRem 用于移除有序集中的一个或多个成员，不存在的成员将被忽略
func RedisZRem(key string, member interface{}) (int64, error) {
	if result, err := redisClient.ZRem(context.Background(), key, member).Result(); err != nil {
		fmt.Printf("redis cmd zrem failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisZScore 返回有序集key中，成员member的score值
func RedisZScore(key string, member string) (float64, error) {
	if result, err := redisClient.ZScore(context.Background(), key, member).Result(); err != nil {
		fmt.Printf("redis cmd zscore failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisZIncrBy 为有序集 key 的成员 member 的 score 值加上增量 increment
func RedisZIncrBy(key string, increment float64, member string) (float64, error) {
	if result, err := redisClient.ZIncrBy(context.Background(), key, increment, member).Result(); err != nil {
		fmt.Printf("redis cmd zincrby failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisZCount 返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量
func RedisZCount(key, min, max string) (int64, error) {
	if result, err := redisClient.ZCount(context.Background(), key, min, max).Result(); err != nil {
		fmt.Printf("redis cmd zcount failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisZRange 返回有序集 key 中，指定区间内的成员,成员的位置按 score 值递增(从小到大)来排序。具有相同 score 值的成员按字典序(lexicographical order )来排列
func RedisZRange(key string, start int64, stop int64) ([]string, error) {
	if result, err := redisClient.ZRange(context.Background(), key, start, stop).Result(); err != nil {
		fmt.Printf("redis cmd zrange failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisZRangeByScore 返回有序集合中指定分数区间的成员列表。有序集成员按分数值递增(从小到大)次序排列
func RedisZRangeByScore(key, min, max string) ([]string, error) {
	opt := &redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	if result, err := redisClient.ZRangeByScore(context.Background(), key, opt).Result(); err != nil {
		fmt.Printf("redis cmd zrangebyscorewithscores failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//RedisZRangeByScoreWithScores 返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列。
func RedisZRangeByScoreWithScores(key, min, max string) (interface{}, error) {
	opt := &redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	if result, err := redisClient.ZRangeByScoreWithScores(context.Background(), key, opt).Result(); err != nil {
		fmt.Printf("redis cmd zrangebyscorewithscores failed, err::%#v\n", err)
		return result, err
	} else {
		return result, nil
	}
}

//GetQueue 组合队列名称
func GetQueue(queue string) string {
	return "yh_queues_" + queue
}

//PushDataToQueue 将数据添加到队列中
func PushDataToQueue(queue string, data interface{}, delay int64) {
	redisQueue := GetQueue(queue)
	if delay > 0 {
		redisQueue = GetQueue(queue) + "_delay"
	}
	if body, err := encodeData(redisQueue, data); err != nil {
		fmt.Printf("encodeData failed, err::%s\n", err)
	} else {
		if delay > 0 {
			res, err := RedisZAdd(redisQueue, float64(GetCurrentTimeUnix()+delay), body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
		} else {
			RedisRPush(queue, body)
		}
		fmt.Printf("data::%s put into queue succcess", body)
	}
}

//DelayTask 处理延时任务
func DelayTask(queue string) {
	currentTime := GetCurrentTimeUnix()
	nxCacheKey := GetQueue(queue) + "_nx"
	isBool, err := RedisSetNX(nxCacheKey, GetCurrentTimeUnix(), 0)
	if err != nil {
		fmt.Printf("RedisSetNX 操作失败，err::%s\n", err)
		return
	}
	if isBool {
		queueDelay := GetQueue(queue) + "_delay"
		jobs, _ := RedisZRangeByScore(queueDelay, "-inf", string(currentTime))
		if len(jobs) > 0 {
			for _, job := range jobs {
				RedisRPush(GetQueue(queue), job)
				RedisZRem(queueDelay, job)
			}
		}
		RedisDel(nxCacheKey)
	} else {
		//防止死锁
		if startTime, err := RedisGet(nxCacheKey); err != nil {
			fmt.Println("获取失败")
		} else {
			sTime, _ := strconv.Atoi(startTime)
			if GetCurrentTimeUnix() > int64(sTime+20) {
				RedisDel(nxCacheKey)
			}
		}
	}
}

//GetCurrentTimeUnix 获取当前时间戳
func GetCurrentTimeUnix() int64 {
	return time.Now().Unix()
}

//encodeData 加入队列中的数据进行编码转换
func encodeData(queue string, data interface{}) (string, error) {
	jsonDataParams := make(map[string]interface{})
	jsonDataParams["queue"] = queue
	jsonDataParams["time"] = time.Now().Unix()
	jsonDataParams["data"] = data
	if paramsByte, err := json.Marshal(jsonDataParams); err != nil {
		fmt.Printf("jsonDataParams json.Marshal failed, err::%s\n", err)
		return "", err
	} else {
		return string(paramsByte), nil
	}
}

//DefaultJsonData 默认的解码对应结构体
type DefaultJsonData struct {
	Queue string      `json:"queue"`
	Time  int         `json:"time"`
	Data  interface{} `json:"data"`
}

//decodeData 编码反解码操作
func decodeData(data string, structData interface{}) (interface{}, error) {
	fmt.Println(data)
	err := json.Unmarshal([]byte(data), &structData)
	if err != nil {
		fmt.Printf("decodeData json.Umarshal failed, err::%s\n", err)
		return false, err
	} else {
		fmt.Printf("decodeData json.Umarshal result::%#v\n", structData)
		return structData, nil
	}
}
