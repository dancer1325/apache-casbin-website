import org.casbin.jcasbin.main.Enforcer;

Enforcer e = new Enforcer("path/to/model.conf", "path/to/policy.csv");



// check permissions
String sub = "alice"; // the user that wants to access a resource.
String obj = "data1"; // the resource that is going to be accessed.
String act = "read"; // the operation that the user performs on the resource.

if (e.enforce(sub, obj, act) == true) {
    // permit alice to read data1
} else {
    // deny the request, show an error
}
