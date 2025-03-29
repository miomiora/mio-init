-- 参数说明：
-- KEYS[1]: refresh token 的 key (格式: refresh:<token>)
-- KEYS[2]: access token 黑名单的 key (格式: blacklist:<token>)
-- ARGV[1]: access token 剩余的 TTL (秒)

-- 1. 删除 refresh token
local refreshDeleted = redis.call("DEL", KEYS[1])

-- 2. 将 access token 加入黑名单（仅当 refresh token 存在时才操作）
if refreshDeleted == 1 then
    redis.call("SETEX", KEYS[2], ARGV[1], "1")
    return 1  -- 成功
else
    return 0  -- refresh token 不存在或已过期
end