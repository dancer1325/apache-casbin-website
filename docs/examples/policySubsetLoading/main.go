package PolicySubsetLoading

import (
	"github.com/casbin/casbin/v3"
	fileadapter "github.com/casbin/casbin/v3/persist/file-adapter"
)

enforcer, _ := casbin.NewEnforcer()

adapter := fileadapter.NewFilteredAdapter("examples/rbac_with_domains_policy.csv")
enforcer.InitWithAdapter("examples/rbac_with_domains_model.conf", adapter)

// ONLY load policies -- for -- "domain1"
filter := &fileadapter.Filter{
P: []string{"", "domain1"},
G: []string{"", "", "domain1"},
}
enforcer.LoadFilteredPolicy(filter)
