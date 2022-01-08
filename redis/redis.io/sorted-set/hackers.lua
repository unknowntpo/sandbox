--! file: hackers.lua
-- Ref: Data type intro - Redis Sorted sets
-- https://redis.io/topics/data-types-intro

local hackers = {}
hackers["Alan Kay"] = 1940
hackers["Sophie Wilson"] = 1957
hackers["Richard Stallman"] = 1953
hackers["Anita Borg"] = 1949
hackers["Yukihiro Matsumoto"] = 1965
hackers["Hedy Lamarr"] = 1914
hackers["Claude Shannon"] = 1916
hackers["Linus Torvalds"] = 1964
hackers["Alan Turing"] = 1912

for name, year in pairs(hackers) do
    redis.call("ZADD", "hackers", year, name)
    print(name, year)
end
