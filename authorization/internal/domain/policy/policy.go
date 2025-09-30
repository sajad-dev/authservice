package policy


type AuthorizerService interface {
	AddGroupPolicyService(req *addgroupreq.AddGroupRequest) (policyres.PolicyResponse, error)
	AddPolicyService(req *addpolicyreq.AddPolicyRequest) (policyres.PolicyResponse, error)
}
