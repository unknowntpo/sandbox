

local article = "article"
local votes = "votes"
local head = "head-line"

-- function upVote(id)

local function upVote(id)
    local key = article .. ":" .. id .. ":"  .. votes 
    redis.call("INCR", key)
end
-- function downVote(id)
-- function showResults(id)
local function showResults(id)
    local voteKey = article .. ":" .. id .. ":"  .. votes 
    return redis.call("GET", voteKey)
end
--
upVote(3)
print(showResults(3))
