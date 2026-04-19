package apiOverview

enforcer, err := casbin.NewEnforcer("./example/model.conf", "./example/policy.csv")
if err != nil {
fmt.Printf("Error, details: %s\n", err)
}
allSubjects := enforcer.GetAllSubjects()
fmt.Println(allSubjects)