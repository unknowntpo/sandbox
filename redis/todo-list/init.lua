--! file:init.lua 
--[[
INCR list:id

HMSET list:1 name "Main Todo List" slug "main" locked 0

HGETALL list:1

ZADD lists 1 list:1

ZRANGE lists 0 -1

]]--
