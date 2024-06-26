package model

import (
	"encoding/json"

	api "github.com/flightctl/flightctl/api/v1alpha1"
	"github.com/flightctl/flightctl/internal/util"
)

var (
	EnrollmentRequestAPI      = "v1alpha1"
	EnrollmentRequestKind     = "EnrollmentRequest"
	EnrollmentRequestListKind = "EnrollmentRequestList"
)

type EnrollmentRequest struct {
	Resource

	// The desired state of the enrollment request, stored as opaque JSON object.
	Spec *JSONField[api.EnrollmentRequestSpec]

	// The last reported state of the enrollment request, stored as opaque JSON object.
	Status *JSONField[api.EnrollmentRequestStatus]
}

type EnrollmentRequestList []EnrollmentRequest

func (e EnrollmentRequest) String() string {
	val, _ := json.Marshal(e)
	return string(val)
}

func NewEnrollmentRequestFromApiResource(resource *api.EnrollmentRequest) *EnrollmentRequest {
	if resource == nil || resource.Metadata.Name == nil {
		return &EnrollmentRequest{}
	}

	status := api.EnrollmentRequestStatus{Conditions: []api.Condition{}}
	if resource.Status != nil {
		status = *resource.Status
	}
	return &EnrollmentRequest{
		Resource: Resource{
			Name:   *resource.Metadata.Name,
			Labels: util.LabelMapToArray(resource.Metadata.Labels),
		},
		Spec:   MakeJSONField(resource.Spec),
		Status: MakeJSONField(status),
	}
}

func (e *EnrollmentRequest) ToApiResource() api.EnrollmentRequest {
	if e == nil {
		return api.EnrollmentRequest{}
	}

	status := api.EnrollmentRequestStatus{Conditions: []api.Condition{}}
	if e.Status != nil {
		status = e.Status.Data
	}

	metadataLabels := util.LabelArrayToMap(e.Resource.Labels)

	return api.EnrollmentRequest{
		ApiVersion: EnrollmentRequestAPI,
		Kind:       EnrollmentRequestKind,
		Metadata: api.ObjectMeta{
			Name:              util.StrToPtr(e.Name),
			CreationTimestamp: util.TimeToPtr(e.CreatedAt.UTC()),
			Labels:            &metadataLabels,
			ResourceVersion:   GetResourceVersion(e.UpdatedAt),
		},
		Spec:   e.Spec.Data,
		Status: &status,
	}
}

func (el EnrollmentRequestList) ToApiResource(cont *string, numRemaining *int64) api.EnrollmentRequestList {
	if el == nil {
		return api.EnrollmentRequestList{
			ApiVersion: EnrollmentRequestAPI,
			Kind:       EnrollmentRequestListKind,
			Items:      []api.EnrollmentRequest{},
		}
	}

	enrollmentRequestList := make([]api.EnrollmentRequest, len(el))
	for i, enrollmentRequest := range el {
		enrollmentRequestList[i] = enrollmentRequest.ToApiResource()
	}
	ret := api.EnrollmentRequestList{
		ApiVersion: EnrollmentRequestAPI,
		Kind:       EnrollmentRequestListKind,
		Items:      enrollmentRequestList,
		Metadata:   api.ListMeta{},
	}
	if cont != nil {
		ret.Metadata.Continue = cont
		ret.Metadata.RemainingItemCount = numRemaining
	}
	return ret
}
