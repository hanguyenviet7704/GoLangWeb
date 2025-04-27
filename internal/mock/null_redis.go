package mock

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type NullRedisClient struct{}

func (n *NullRedisClient) Pipeline() redis.Pipeliner {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TxPipeline() redis.Pipeliner {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Command(ctx context.Context) *redis.CommandsInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CommandList(ctx context.Context, filter *redis.FilterBy) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CommandGetKeys(ctx context.Context, commands ...interface{}) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CommandGetKeysAndFlags(ctx context.Context, commands ...interface{}) *redis.KeyFlagsCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientGetName(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Echo(ctx context.Context, message interface{}) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Quit(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BgRewriteAOF(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BgSave(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientList(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientInfo(ctx context.Context) *redis.ClientInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientUnpause(ctx context.Context) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientID(ctx context.Context) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientUnblock(ctx context.Context, id int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClientUnblockWithError(ctx context.Context, id int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ConfigGet(ctx context.Context, parameter string) *redis.MapStringStringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ConfigResetStat(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ConfigRewrite(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) DBSize(ctx context.Context) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FlushAll(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FlushAllAsync(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FlushDB(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FlushDBAsync(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Info(ctx context.Context, section ...string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LastSave(ctx context.Context) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Save(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Shutdown(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ShutdownSave(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ShutdownNoSave(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SlowLogGet(ctx context.Context, num int64) *redis.SlowLogCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Time(ctx context.Context) *redis.TimeCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) DebugObject(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ModuleLoadex(ctx context.Context, conf *redis.ModuleLoadexConfig) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ACLDryRun(ctx context.Context, username string, command ...interface{}) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ACLLog(ctx context.Context, count int64) *redis.ACLLogCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ACLLogReset(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitPosSpan(ctx context.Context, key string, bit int8, start, end int64, span string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitField(ctx context.Context, key string, values ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BitFieldRO(ctx context.Context, key string, values ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterMyShardID(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterShards(ctx context.Context) *redis.ClusterShardsCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterLinks(ctx context.Context) *redis.ClusterLinksCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterNodes(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterResetSoft(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterResetHard(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterInfo(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterSaveConfig(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterFailover(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ReadOnly(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ReadWrite(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFunctionLoad(ctx context.Context, lib string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFunctionLoadArgs(ctx context.Context, lib string, options *redis.TFunctionLoadOptions) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFunctionDelete(ctx context.Context, libName string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFunctionList(ctx context.Context) *redis.MapStringInterfaceSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFunctionListArgs(ctx context.Context, options *redis.TFunctionListOptions) *redis.MapStringInterfaceSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFCall(ctx context.Context, libName string, funcName string, numKeys int) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFCallArgs(ctx context.Context, libName string, funcName string, numKeys int, options *redis.TFCallOptions) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFCallASYNC(ctx context.Context, libName string, funcName string, numKeys int) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TFCallASYNCArgs(ctx context.Context, libName string, funcName string, numKeys int, options *redis.TFCallOptions) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Dump(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ExpireNX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ExpireXX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ExpireGT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ExpireLT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ObjectFreq(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ObjectRefCount(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ObjectEncoding(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Persist(ctx context.Context, key string) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PTTL(ctx context.Context, key string) *redis.DurationCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) RandomKey(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SortRO(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Type(ctx context.Context, key string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) *redis.GeoSearchLocationCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoSearchStore(ctx context.Context, key, store string, q *redis.GeoSearchStoreQuery) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HLen(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HRandField(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HRandFieldWithValues(ctx context.Context, key string, count int) *redis.KeyValueSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HPersist(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HPExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) HPTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PFCount(ctx context.Context, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LLen(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LMPop(ctx context.Context, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LPop(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LPos(ctx context.Context, key string, value string, args redis.LPosArgs) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LPosCount(ctx context.Context, key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) RPop(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LMove(ctx context.Context, source, destination, srcpos, destpos string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFAdd(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFCard(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFExists(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFInfo(ctx context.Context, key string) *redis.BFInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFInfoArg(ctx context.Context, key, option string) *redis.BFInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFInfoCapacity(ctx context.Context, key string) *redis.BFInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFInfoSize(ctx context.Context, key string) *redis.BFInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFInfoFilters(ctx context.Context, key string) *redis.BFInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFInfoItems(ctx context.Context, key string) *redis.BFInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFInfoExpansion(ctx context.Context, key string) *redis.BFInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFInsert(ctx context.Context, key string, options *redis.BFInsertOptions, elements ...interface{}) *redis.BoolSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFMAdd(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFMExists(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity, expansion int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFReserveWithArgs(ctx context.Context, key string, options *redis.BFReserveOptions) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFScanDump(ctx context.Context, key string, iterator int64) *redis.ScanDumpCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFAdd(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFAddNX(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFCount(ctx context.Context, key string, element interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFDel(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFExists(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFInfo(ctx context.Context, key string) *redis.CFInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFInsert(ctx context.Context, key string, options *redis.CFInsertOptions, elements ...interface{}) *redis.BoolSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFInsertNX(ctx context.Context, key string, options *redis.CFInsertOptions, elements ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFMExists(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFReserve(ctx context.Context, key string, capacity int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFReserveWithArgs(ctx context.Context, key string, options *redis.CFReserveOptions) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFScanDump(ctx context.Context, key string, iterator int64) *redis.ScanDumpCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CMSIncrBy(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CMSInfo(ctx context.Context, key string) *redis.CMSInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CMSInitByDim(ctx context.Context, key string, width, height int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CMSInitByProb(ctx context.Context, key string, errorRate, probability float64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) CMSQuery(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKAdd(ctx context.Context, key string, elements ...interface{}) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKCount(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKIncrBy(ctx context.Context, key string, elements ...interface{}) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKInfo(ctx context.Context, key string) *redis.TopKInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKList(ctx context.Context, key string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKListWithCount(ctx context.Context, key string) *redis.MapStringIntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKQuery(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKReserve(ctx context.Context, key string, k int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TopKReserveWithOptions(ctx context.Context, key string, k int64, width, depth int64, decay float64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestAdd(ctx context.Context, key string, elements ...float64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestByRank(ctx context.Context, key string, rank ...uint64) *redis.FloatSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *redis.FloatSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestCDF(ctx context.Context, key string, elements ...float64) *redis.FloatSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestCreate(ctx context.Context, key string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestInfo(ctx context.Context, key string) *redis.TDigestInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestMax(ctx context.Context, key string) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestMin(ctx context.Context, key string) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestMerge(ctx context.Context, destKey string, options *redis.TDigestMergeOptions, sourceKeys ...string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestQuantile(ctx context.Context, key string, elements ...float64) *redis.FloatSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestRank(ctx context.Context, key string, values ...float64) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestReset(ctx context.Context, key string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestRevRank(ctx context.Context, key string, values ...float64) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile, highCutQuantile float64) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SPublish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PubSubNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PubSubShardChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) PubSubShardNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ScriptKill(ctx context.Context) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionLoad(ctx context.Context, code string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionLoadReplace(ctx context.Context, code string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionDelete(ctx context.Context, libName string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionFlush(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionKill(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionFlushAsync(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionList(ctx context.Context, q redis.FunctionListQuery) *redis.FunctionListCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionDump(ctx context.Context) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionRestore(ctx context.Context, libDump string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FunctionStats(ctx context.Context) *redis.FunctionStatsCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FCall(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FCallRo(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FCallRO(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FT_List(ctx context.Context) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTAggregate(ctx context.Context, index string, query string) *redis.MapStringInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTAggregateWithArgs(ctx context.Context, index string, query string, options *redis.FTAggregateOptions) *redis.AggregateCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTAliasAdd(ctx context.Context, index string, alias string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTAliasDel(ctx context.Context, alias string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTAliasUpdate(ctx context.Context, index string, alias string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTConfigGet(ctx context.Context, option string) *redis.MapMapStringInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTConfigSet(ctx context.Context, option string, value interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTCreate(ctx context.Context, index string, options *redis.FTCreateOptions, schema ...*redis.FieldSchema) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTCursorDel(ctx context.Context, index string, cursorId int) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTCursorRead(ctx context.Context, index string, cursorId int, count int) *redis.MapStringInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTDictAdd(ctx context.Context, dict string, term ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTDictDel(ctx context.Context, dict string, term ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTDictDump(ctx context.Context, dict string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTDropIndex(ctx context.Context, index string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTDropIndexWithArgs(ctx context.Context, index string, options *redis.FTDropIndexOptions) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTExplain(ctx context.Context, index string, query string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTExplainWithArgs(ctx context.Context, index string, query string, options *redis.FTExplainOptions) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTInfo(ctx context.Context, index string) *redis.FTInfoCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTSpellCheck(ctx context.Context, index string, query string) *redis.FTSpellCheckCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTSpellCheckWithArgs(ctx context.Context, index string, query string, options *redis.FTSpellCheckOptions) *redis.FTSpellCheckCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTSearch(ctx context.Context, index string, query string) *redis.FTSearchCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTSearchWithArgs(ctx context.Context, index string, query string, options *redis.FTSearchOptions) *redis.FTSearchCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTSynDump(ctx context.Context, index string) *redis.FTSynDumpCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTSynUpdate(ctx context.Context, index string, synGroupId interface{}, terms []interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId interface{}, options *redis.FTSynUpdateOptions, terms []interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) FTTagVals(ctx context.Context, index string, field string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SCard(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SInterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SMIsMember(ctx context.Context, key string, members ...interface{}) *redis.BoolSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SPop(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SRandMember(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZAdd(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZAddLT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZAddGT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZAddNX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZAddXX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZCard(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZInter(ctx context.Context, store *redis.ZStore) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZInterWithScores(ctx context.Context, store *redis.ZStore) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZInterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZMPop(ctx context.Context, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZMScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRank(ctx context.Context, key, member string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRankWithScore(ctx context.Context, key, member string) *redis.RankWithScoreCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRevRank(ctx context.Context, key, member string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRevRankWithScore(ctx context.Context, key, member string) *redis.RankWithScoreCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZScore(ctx context.Context, key, member string) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRandMember(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZRandMemberWithScores(ctx context.Context, key string, count int) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZUnion(ctx context.Context, store redis.ZStore) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZUnionWithScores(ctx context.Context, store redis.ZStore) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZDiffWithScores(ctx context.Context, keys ...string) *redis.ZSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Append(ctx context.Context, key, value string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) Decr(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) GetDel(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) LCS(ctx context.Context, q *redis.LCSQuery) *redis.LCSCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SetArgs(ctx context.Context, key string, value interface{}, a redis.SetArgs) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) StrLen(ctx context.Context, key string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XLen(ctx context.Context, stream string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XRevRange(ctx context.Context, stream string, start, stop string) *redis.XMessageSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XPending(ctx context.Context, stream, group string) *redis.XPendingCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XAutoClaim(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XAutoClaimJustID(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimJustIDCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XTrimMinID(ctx context.Context, key string, minID string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XInfoStreamFull(ctx context.Context, key string, count int) *redis.XInfoStreamFullCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) XInfoConsumers(ctx context.Context, key string, group string) *redis.XInfoConsumersCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSAdd(ctx context.Context, key string, timestamp interface{}, value float64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSAddWithArgs(ctx context.Context, key string, timestamp interface{}, value float64, options *redis.TSOptions) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSCreate(ctx context.Context, key string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSCreateWithArgs(ctx context.Context, key string, options *redis.TSOptions) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSAlter(ctx context.Context, key string, options *redis.TSAlterOptions) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator redis.Aggregator, bucketDuration int) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator redis.Aggregator, bucketDuration int, options *redis.TSCreateRuleOptions) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSIncrBy(ctx context.Context, Key string, timestamp float64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *redis.TSIncrDecrOptions) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSDecrBy(ctx context.Context, Key string, timestamp float64) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *redis.TSIncrDecrOptions) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSDel(ctx context.Context, Key string, fromTimestamp int, toTimestamp int) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSGet(ctx context.Context, key string) *redis.TSTimestampValueCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSGetWithArgs(ctx context.Context, key string, options *redis.TSGetOptions) *redis.TSTimestampValueCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSInfo(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSInfoWithArgs(ctx context.Context, key string, options *redis.TSInfoOptions) *redis.MapStringInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSMAdd(ctx context.Context, ktvSlices [][]interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSQueryIndex(ctx context.Context, filterExpr []string) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.TSTimestampValueSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *redis.TSRevRangeOptions) *redis.TSTimestampValueSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.TSTimestampValueSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *redis.TSRangeOptions) *redis.TSTimestampValueSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *redis.MapStringSliceInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *redis.TSMRangeOptions) *redis.MapStringSliceInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *redis.MapStringSliceInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *redis.TSMRevRangeOptions) *redis.MapStringSliceInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSMGet(ctx context.Context, filters []string) *redis.MapStringSliceInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) TSMGetWithArgs(ctx context.Context, filters []string, options *redis.TSMGetOptions) *redis.MapStringSliceInterfaceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONArrAppend(ctx context.Context, key, path string, values ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONArrIndex(ctx context.Context, key, path string, value ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONArrIndexWithArgs(ctx context.Context, key, path string, options *redis.JSONArrIndexArgs, value ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONArrInsert(ctx context.Context, key, path string, index int64, values ...interface{}) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONArrLen(ctx context.Context, key, path string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONArrPop(ctx context.Context, key, path string, index int) *redis.StringSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONArrTrim(ctx context.Context, key, path string) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONArrTrimWithArgs(ctx context.Context, key, path string, options *redis.JSONArrTrimArgs) *redis.IntSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONClear(ctx context.Context, key, path string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONDebugMemory(ctx context.Context, key, path string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONDel(ctx context.Context, key, path string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONForget(ctx context.Context, key, path string) *redis.IntCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONGet(ctx context.Context, key string, paths ...string) *redis.JSONCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONGetWithArgs(ctx context.Context, key string, options *redis.JSONGetArgs, paths ...string) *redis.JSONCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONMerge(ctx context.Context, key, path string, value string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONMSetArgs(ctx context.Context, docs []redis.JSONSetArgs) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONMSet(ctx context.Context, params ...interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONMGet(ctx context.Context, path string, keys ...string) *redis.JSONSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONNumIncrBy(ctx context.Context, key, path string, value float64) *redis.JSONCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONObjKeys(ctx context.Context, key, path string) *redis.SliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONObjLen(ctx context.Context, key, path string) *redis.IntPointerSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONSet(ctx context.Context, key, path string, value interface{}) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONSetMode(ctx context.Context, key, path string, value interface{}, mode string) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONStrAppend(ctx context.Context, key, path, value string) *redis.IntPointerSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONStrLen(ctx context.Context, key, path string) *redis.IntPointerSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONToggle(ctx context.Context, key, path string) *redis.IntPointerSliceCmd {
	//TODO implement me
	panic("implement me")
}

func (n *NullRedisClient) JSONType(ctx context.Context, key, path string) *redis.JSONSliceCmd {
	//TODO implement me
	panic("implement me")
}

var _ redis.Cmdable = (*NullRedisClient)(nil)

func (n *NullRedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	return redis.NewStatusResult("PONG", nil)
}

func (n *NullRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return redis.NewStringResult("", redis.Nil) // Mocked to return Nil
}

func (n *NullRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redis.NewStatusResult("OK", nil)
}

func (n *NullRedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	return redis.NewIntResult(1, nil) // Mocked to return 1
}

func (n *NullRedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return redis.NewBoolResult(true, nil)
}

func (n *NullRedisClient) TTL(ctx context.Context, key string) *redis.DurationCmd {
	return redis.NewDurationResult(time.Minute, nil)
}

func (n *NullRedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return redis.NewIntResult(1, nil) // Mocked to return 1
}

func (n *NullRedisClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	return redis.NewIntResult(1, nil) // Mocked to return 1
}
