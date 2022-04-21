// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0
// Source: cerbos/response/v1/response.proto

package responsev1

import (
	hash "hash"
)

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ResourcesQueryPlanResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ResourcesQueryPlanResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ResourcesQueryPlanResponse_Expression) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ResourcesQueryPlanResponse_Expression_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ResourcesQueryPlanResponse_Expression_Operand) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ResourcesQueryPlanResponse_Expression_Operand_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ResourcesQueryPlanResponse_Filter) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ResourcesQueryPlanResponse_Filter_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ResourcesQueryPlanResponse_Meta) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ResourcesQueryPlanResponse_Meta_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceSetResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourceSetResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceSetResponse_ActionEffectMap) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourceSetResponse_ActionEffectMap_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceSetResponse_Meta) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourceSetResponse_Meta_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceSetResponse_Meta_EffectMeta) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourceSetResponse_Meta_EffectMeta_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceSetResponse_Meta_ActionMeta) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourceSetResponse_Meta_ActionMeta_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceBatchResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourceBatchResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceBatchResponse_ActionEffectMap) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourceBatchResponse_ActionEffectMap_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourcesResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourcesResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourcesResponse_ResultEntry) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourcesResponse_ResultEntry_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourcesResponse_ResultEntry_Resource) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourcesResponse_ResultEntry_Resource_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourcesResponse_ResultEntry_Meta) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourcesResponse_ResultEntry_Meta_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourcesResponse_ResultEntry_Meta_EffectMeta) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_CheckResourcesResponse_ResultEntry_Meta_EffectMeta_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundFailure) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundFailure_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundFailure_Error) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundFailure_Error_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundValidateResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundValidateResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundTestResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundTestResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundTestResponse_TestResults) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundTestResponse_TestResults_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundEvaluateResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundEvaluateResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundEvaluateResponse_EvalResult) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundEvaluateResponse_EvalResult_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundEvaluateResponse_EvalResultList) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundEvaluateResponse_EvalResultList_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundProxyResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_PlaygroundProxyResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *AddOrUpdatePolicyResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_AddOrUpdatePolicyResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ListAuditLogEntriesResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ListAuditLogEntriesResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ServerInfoResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ServerInfoResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ListPoliciesResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ListPoliciesResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *GetPolicyResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_GetPolicyResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ListSchemasResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_ListSchemasResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *GetSchemaResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_response_v1_GetSchemaResponse_hashpb_sum(m, hasher, ignore)
	}
}
