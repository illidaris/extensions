package cache

// LUA_ST_INC 带判断的自增脚本，-1则为超限，否则返回当前值。
const LUA_ST_INC = `
local temp 
local num2 = redis.call('setnx',KEYS[1],0)
if (tonumber(num2))>0
then
redis.call('EXPIRE',KEYS[1],ARGV[3])
end
local num = redis.call('get',KEYS[1]) 
if tonumber(ARGV[1])>=(tonumber(num)+ARGV[2]) 
then
temp = redis.call('incrby',KEYS[1],ARGV[2]) 
return tostring(temp)
end 
return '-1'`

// `local temp
// redis.call('setnx',KEYS[1],0)
// local num = redis.call('get',KEYS[1])
// if tonumber(ARGV[1])>=(tonumber(num)+ARGV[2])
// then
// temp = redis.call('incrby',KEYS[1],ARGV[2])
// return tostring(res),tostring(temp)
// end
// return tostring(res),'-1'`
