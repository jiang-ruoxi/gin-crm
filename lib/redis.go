package lib

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

//RedisConnect redis连接操作
func RedisConnect() *redis.Client {
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "111", //no password set
		DB:       0,     //use default DBls
	})

	_, err := client.Ping(cxt).Result()
	if err != nil {
		fmt.Printf("redis connect failed, err::%#v\n", err)
		panic(err)
	}
	return client
}

//RedisSet 设置缓存，指定过期时间
func RedisSet(key string, value interface{}, expiration time.Duration) error {
	if err := RedisConnect().Set(context.Background(), key, value, expiration).Err(); err != nil {
		fmt.Printf("redis cmd set failed, err::%#v\n", err)
		return err
	}
	return nil
}

//RedisGet 获取指定的缓存
func RedisGet(key string) (string, error) {
	result, err := RedisConnect().Get(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("redis cmd get failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisDel 删除指定的key
func RedisDel(keys string) (int64, error) {
	result, err := RedisConnect().Del(context.Background(), keys).Result()
	if err != nil {
		fmt.Printf("redis cmd del failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisExpire 设置指定key的过期时间
func RedisExpire(key string, expiration time.Duration) error {
	if err := RedisConnect().Expire(context.Background(), key, expiration).Err(); err != nil {
		fmt.Printf("redis cmd expire failed, err::%#v\n", err)
		return err
	}
	return nil
}

//RedisSetNX 实现分布式加锁
func RedisSetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	result, err := RedisConnect().SetNX(context.Background(), key, value, expiration).Result()
	if err != nil {
		fmt.Printf("redis cmd setnx failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisIncr 自增序列
func RedisIncr(key string) (int64, error) {
	result, err := RedisConnect().Incr(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("redis cmd incr failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisSetBit 位图操作
func RedisSetBit(key string, offset int64, value int) error {
	if err := RedisConnect().SetBit(context.Background(), key, offset, value).Err(); err != nil {
		fmt.Printf("redis cmd setbit failed, err::%#v\n", err)
		return err
	}
	return nil
}

//RedisGetBit 位图读取
func RedisGetBit(key string, offset int64) (int64, error) {
	result, err := RedisConnect().GetBit(context.Background(), key, offset).Result()
	if err != nil {
		fmt.Printf("redis cmd getbit failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisLPush 从左侧向list压入数据
func RedisLPush(key string, value ...interface{}) error {
	if err := RedisConnect().LPush(context.Background(), key, value).Err(); err != nil {
		fmt.Printf("redis cmd lpush failed, err::%#v\n", err)
		return err
	}
	return nil
}

//RedisRPush 从右侧向list压入数据
func RedisRPush(key string, value ...interface{}) error {
	if err := RedisConnect().RPush(context.Background(), key, value).Err(); err != nil {
		fmt.Printf("redis cmd rpush failed, err::%#v\n", err)
		return err
	}
	return nil
}

//RedisLRange 从
func RedisLRange(key string, start int64, stop int64) ([]string, error) {
	result, err := RedisConnect().LRange(context.Background(), key, start, stop).Result()
	if err != nil {
		fmt.Printf("redis cmd lrange failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisLPop 从左侧pop数据
func RedisLPop(key string) (string, error) {
	result, err := RedisConnect().LPop(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("redis cmd lpop failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisRPop 从右侧pop数据
func RedisRPop(key string) (string, error) {
	result, err := RedisConnect().RPop(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("redis cmd rpop failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisLLen 获取列表中的数据长度
func RedisLLen(key string) (int64, error) {
	result, err := RedisConnect().LLen(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("redis cmd llen failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisSAdd 将一个或多个member元素加入到集合key当中，已经存在于集合的member元素将被忽略
func RedisSAdd(key string, members ...interface{}) error {
	if err := RedisConnect().SAdd(context.Background(), key, members).Err(); err != nil {
		fmt.Printf("redis cmd sadd failed, err::%#v\n", err)
		return err
	}
	return nil
}

//RedisSIsMember 判断 member 元素是否集合 key 的成员
func RedisSIsMember(key string, member interface{}) (bool, error) {
	result, err := RedisConnect().SIsMember(context.Background(), key, member).Result()
	if err != nil {
		fmt.Printf("redis cmd sismember failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisSPop 移除并返回集合中的一个随机元素
func RedisSPop(key string) (string, error) {
	result, err := RedisConnect().SPop(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("redis cmd spop failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisSRandMember 返回集合中的一个随机元素
func RedisSRandMember(key string) (string, error) {
	result, err := RedisConnect().SRandMember(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("redis cmd srandmember failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisZAdd 将一个member元素及其score值加入到有序集key当中
func RedisZAdd(key string, score float64, members interface{}) (int64, error) {
	items := []*redis.Z{
		&redis.Z{Score: score, Member: members},
	}
	result, err := RedisConnect().ZAdd(context.Background(), key, items...).Result()
	if err != nil {
		fmt.Printf("redis cmd zadd failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisZScore 返回有序集key中，成员member的score值
func RedisZScore(key string, member string) (float64, error) {
	result, err := RedisConnect().ZScore(context.Background(), key, member).Result()
	if err != nil {
		fmt.Printf("redis cmd zscore failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisZIncrBy 为有序集 key 的成员 member 的 score 值加上增量 increment
func RedisZIncrBy(key string, increment float64, member string) (float64, error) {
	result, err := RedisConnect().ZIncrBy(context.Background(), key, increment, member).Result()
	if err != nil {
		fmt.Printf("redis cmd zincrby failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisZCount 返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量
func RedisZCount(key, min, max string) (int64, error) {
	result, err := RedisConnect().ZCount(context.Background(), key, min, max).Result()
	if err != nil {
		fmt.Printf("redis cmd zcount failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisZRange 返回有序集 key 中，指定区间内的成员,成员的位置按 score 值递增(从小到大)来排序。具有相同 score 值的成员按字典序(lexicographical order )来排列
func RedisZRange(key string, start int64, stop int64) ([]string, error) {
	result, err := RedisConnect().ZRange(context.Background(), key, start, stop).Result()
	if err != nil {
		fmt.Printf("redis cmd zrange failed, err::%#v\n", err)
		return result, err
	}
	return result, nil
}

//RedisZRangeByScoreWithScores 返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列。
func RedisZRangeByScoreWithScores(key, min, max string) (interface{}, error) {
	opt := &redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	res, err := RedisConnect().ZRangeByScoreWithScores(context.Background(), key, opt).Result()
	if err != nil {
		fmt.Printf("redis cmd zrangebyscorewithscores failed, err::%#v\n", err)
		return nil, err
	}
	return res, nil
}