package apiOverview

// reason	== matching policy
ok, reason, err := enforcer.EnforceEx("amber", "data1", "read")
fmt.Println(ok, reason) // true [admin data1 read]

// Amber has the `admin` role -> the policy `p, admin, data1, read` allows the request

