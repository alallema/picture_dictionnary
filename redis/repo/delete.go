package repo

import "github.com/go-redis/redis"

func DelHSet(id string) bool {
	keys, err := Client.HGetAll(id).Result()
	if err != redis.Nil {
		for key, _ := range keys {
			Client.HDel(id, key)
		}
		return true
	}
	return false
}

func DelMembersInSet(set string, id string) bool {
	if Client.SRem(set, id).Val() == 1 {
		return true
	}
	return false
}

// func DelZSet(set string) bool {
// 	if Client.ZRem(set, "-inf", "+inf").Val() == 1 {
// 		return true
// 	}
// 	return false
// }

func DelAll(set string) bool {
	if Client.Del(set).Val() == 1 {
		return true
	}
	return false
}
