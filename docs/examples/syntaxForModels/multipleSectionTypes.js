// 1. set an `EnforceContext`
const enforceContext = new EnforceContext('r2', 'p2', 'e2', 'm2');
class EnforceContext {
    constructor(rType, pType, eType, mType) {
        this.pType = pType;
        this.eType = eType;
        this.mType = mType;
        this.rType = rType;
    }
}

// 2. `.enforce(previouslyDefinedEnforceContext)`
// Pass in a suffix as a parameter to NewEnforceContext, such as 2 or 3, and it will create r2, p2, etc.
const enforceContext = new NewEnforceContext('2');

// You can also specify a certain type individually
enforceContext.eType = "e"

// Don't pass in EnforceContext; the default is r, p, e, m
e.Enforce("alice", "data2", "read")        // true

// Pass in EnforceContext
e.Enforce(enforceContext, {Age: 70}, "/data1", "read")        //false
e.Enforce(enforceContext, {Age: 30}, "/data1", "read")        //true
