package rdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gomodule/redigo/redis"
	"ksd-social-api/commons/global"
	"ksd-social-api/commons/initilization"
)

/*
*
redis  SET
*/
func RdbSet(key, v string) (bool, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	b, err := redis.Bool(connect.Do("SET", global.RedisKey+key, v))
	if err != nil {
		logs.Error("set error", err.Error())
		return false, err
	}
	return b, nil
}

/*
*
redis  GET
*/
func RdbGet(key string) (string, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	val, err := redis.String(connect.Do("GET", global.RedisKey+key))
	if err != nil {
		logs.Error("get error", err.Error())
		return "", err
	}

	return val, nil
}

/*
*
redis EXPIRE
*/
func RdbSetKeyExp(key string, ex int) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	_, err := connect.Do("EXPIRE", global.RedisKey+key, ex)
	if err != nil {
		logs.Error("set error", err.Error())
		return err
	}
	return nil
}

func RdbSetExp(key, v string, ex int) (bool, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	result, err := connect.Do("SET", global.RedisKey+key, v, "EX", ex)
	if err != nil {
		logs.Error("set error", err.Error())
		return false, err
	}
	return result == "OK", nil
}

/**
 * 给key设置过期时间
 * @author feige
 * @date 2023-11-06
 * @version 1.0
 * @desc
 */
func RdbExKey(key string, second int) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	_, err := connect.Do("expire", global.RedisKey+key, second)
	if err != nil {
		logs.Error("expire error", err.Error())
		return err
	}
	return nil
}

/*
*
redis EXISTS
*/
func RdbCheck(key string) bool {
	connect := initilization.PoolConnect()
	defer connect.Close()
	b, err := redis.Bool(connect.Do("EXISTS", global.RedisKey+key))
	if err != nil {
		return false
	}
	return b
}

/*
*
redis DEL
*/
func RdbDel(key string) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	_, err := connect.Do("DEL", global.RedisKey+key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

/*
*
redis SETNX
*/
func RdbSetJson(key string, data interface{}) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	value, _ := json.Marshal(data)
	n, _ := connect.Do("SETNX", global.RedisKey+key, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}

/*
*
redis SETEXNX
*/
func RdbSetJsonEXNX(key string, data interface{}, ex int) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	value, _ := json.Marshal(data)
	n, _ := connect.Do("SET", global.RedisKey+key, value, "EX", ex, "NX") //OK
	if n == nil {
		return errors.New("set failed")
	}
	return nil
}

/*
*
redis SETEXNX
*/
func RdbSetEXNX(key string, data any, ex int) bool {
	connect := initilization.PoolConnect()
	defer connect.Close()
	n, _ := connect.Do("SET", global.RedisKey+key, data, "EX", ex, "NX") //OK
	if n == nil {
		return false
	}
	return true
}

/*
*
redis SETEX
*/
func RdbSetJsonEX(key string, data interface{}, seconds int) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	value, _ := json.Marshal(data)
	n, _ := connect.Do("SET", global.RedisKey+key, value, "EX", seconds)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}

/*
*
redis SETEX
*/
func RdbSetEX(key string, data interface{}, seconds int) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	n, _ := connect.Do("SET", global.RedisKey+key, data, "EX", seconds, "NX")
	if n != "OK" {
		return errors.New("set failed")
	}
	return nil
}

/*
*
redis GET
return map
*/
func RdbGetJson(key string) (map[string]string, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	var jsonData map[string]string
	bv, err := redis.Bytes(connect.Do("GET", global.RedisKey+key))
	if err != nil {
		logs.Error("get json error", err.Error())
		return nil, err
	}
	errJson := json.Unmarshal(bv, &jsonData)
	if errJson != nil {
		logs.Error("json nil", err.Error())
		return nil, err
	}
	return jsonData, nil
}

/*
*
redis GET
return map
*/
func RdbGetJsonAny[T any](key string) (T, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	var jsonData T
	bv, err := redis.Bytes(connect.Do("GET", global.RedisKey+key))
	if err != nil {
		logs.Error("get json error", err.Error())
		return jsonData, err
	}
	errJson := json.Unmarshal(bv, &jsonData)
	if errJson != nil {
		logs.Error("json nil", err.Error())
		return jsonData, err
	}
	return jsonData, nil
}

/*
*
redis hSet 注意 设置什么类型 取的时候需要获取对应类型
*/
func RdbHSet(key string, field string, data interface{}) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	_, err := connect.Do("HSET", global.RedisKey+key, field, data)
	if err != nil {
		logs.Error("hSet error", err.Error())
		return err
	}
	return nil
}

/*
*
redis hGet 注意 设置什么类型 取的时候需要获取对应类型
*/
func RdbHGet(key, field string) (interface{}, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	data, err := redis.String(connect.Do("HGET", global.RedisKey+key, field))
	if err != nil {
		logs.Error("hGet error", err.Error())
		return nil, err
	}
	return data, nil
}

/*
*
redis hGetAll
return map
*/
func RdbHGetAll(key string) (map[string]string, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	data, err2 := redis.StringMap(connect.Do("HGETALL", global.RedisKey+key))
	_, err := data, err2
	if err != nil {
		logs.Error("hGetAll error", err.Error())
		return nil, err
	}
	return data, nil
}

/*
*
redis INCR 将 key 中储存的数字值增一
*/
func RdbIncr(key string) (int, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	i, err := redis.Int(connect.Do("INCR", global.RedisKey+key))
	if err != nil {
		logs.Error("INCR error", err.Error())
		return 0, err
	}
	return i, nil

}

/*
*
redis INCRBY 将 key 所储存的值加上增量 n
*/
func RdbIncrBy(key string, n int) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	_, err := connect.Do("INCRBY", global.RedisKey+key, n)
	if err != nil {
		logs.Error("INCRBY error", err.Error())
		return err
	}
	return nil
}

/*
*
redis DECR 将 key 中储存的数字值减一。
*/
func RdbDecr(key string) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	_, err := connect.Do("DECR", global.RedisKey+key)
	if err != nil {
		logs.Error("DECR error", err.Error())
		return err
	}
	return nil
}

/*
*
redis DECRBY 将 key 所储存的值减去减量 n
*/
func RdbDecrBy(key string, n int) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	_, err := connect.Do("DECRBY", global.RedisKey+key, n)
	if err != nil {
		logs.Error("DECRBY error", err.Error())
		return err
	}
	return nil
}

/*
*
redis SADD 将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略。
*/
func RdbSAdd(key, v string) error {
	connect := initilization.PoolConnect()
	defer connect.Close()
	_, err := connect.Do("SADD", global.RedisKey+key, v)
	if err != nil {
		logs.Error("SADD error", err.Error())
		return err
	}
	return nil
}

/**
 * @author feige
 * @date 2023-10-16
 * @version 1.0
 * @desc
 */
func RdbSREM(key, v string) (bool, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	return redis.Bool(connect.Do("SREM", global.RedisKey+key, v))
}

/*
*
redis SMEMBERS 返回集合 key 中的所有成员。
return map
*/
func RdbSMembers(key string) (interface{}, error) {
	connect := initilization.PoolConnect()
	defer connect.Close()
	data, err := redis.Strings(connect.Do("SMEMBERS", global.RedisKey+key))
	if err != nil {
		logs.Error("json nil", err)
		return nil, err
	}
	return data, nil
}

/*
- @author feige
- @date 2023-10-16
- @version 1.0
- @desc  redis SISMEMBER 判断 member 元素是否集合 key 的成员。
- return bool
*/
func RdbSISMembers(key, v string) bool {
	connect := initilization.PoolConnect()
	defer connect.Close()
	b, err := redis.Bool(connect.Do("SISMEMBER", global.RedisKey+key, v))
	if err != nil {
		logs.Error("SISMEMBER error", err.Error())
		return false
	}
	return b
}

/*
- @author feige
- @date 2023-10-16
- @version 1.0
- @desc  redis SISMEMBER 判断 member 元素是否集合 key 的成员。
- return bool
*/
func RdbSCARD(key string) int {
	connect := initilization.PoolConnect()
	defer connect.Close()
	b, err := redis.Int(connect.Do("SCARD", global.RedisKey+key))
	if err != nil {
		logs.Error("SISMEMBER error", err.Error())
		return 0
	}
	return b
}
