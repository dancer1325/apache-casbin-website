
// 1. set an `EnforceContext`
EnforceContext enforceContext = new EnforceContext("2");
public class EnforceContext {
    private String pType;
    private String eType;
    private String mType;
    private String rType;
    public EnforceContext(String suffix) {
      this.pType = "p" + suffix;
      this.eType = "e" + suffix;
      this.mType = "m" + suffix;
      this.rType = "r" + suffix;
    }
}

// 2. `.enforce(previouslyDefinedEnforceContext)`
// Pass in a suffix as a parameter to NewEnforceContext, such as 2 or 3, and it will create r2, p2, etc.
EnforceContext enforceContext = new EnforceContext("2");
// You can also specify a certain type individually
enforceContext.seteType("e");
// Don't pass in EnforceContext; the default is r, p, e, m
e.enforce("alice", "data2", "read");  // true
// Pass in EnforceContext
// TestEvalRule is located in https://github.com/casbin/jcasbin/blob/master/src/test/java/org/casbin/jcasbin/main/AbacAPIUnitTest.java#L56
e.enforce(enforceContext, new AbacAPIUnitTest.TestEvalRule("alice", 70), "/data1", "read"); // false
e.enforce(enforceContext, new AbacAPIUnitTest.TestEvalRule("alice", 30), "/data1", "read"); // true
