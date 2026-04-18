local Enforcer = require("casbin")
local e = Enforcer:new("path/to/model.conf", "path/to/policy.csv") -- The Casbin Enforcer



// check permissions
if e:enforce("alice", "data1", "read") then
   -- permit alice to read data1
else
   -- deny the request, show an error
end