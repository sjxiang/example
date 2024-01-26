package auth_test

import (
	"testing"
)

var (
	// model.conf
	ACLModelFile = `
# Request definition
[request_definition]
r = sub, obj, act

# Policy definition
[policy_definition]
p = sub, obj, act

# Policy effect
[policy_effect]
e = some(where (p.eft == allow))

# Matchers
[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`
	// policy.csv
	ACLPolicyFile = `
p, root, *, produce
p, root, *, consume	
`
)

func TestAuthorize(t *testing.T) {
	// authorizer := auth.New(ACLModelFile, ACLPolicyFile)
}