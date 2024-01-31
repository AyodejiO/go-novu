package lib

import (
	"context"
	"net/http"
	"strconv"
)

type WorkflowService service

// CreateWorkflow creates a new workflow with the specified parameters.
// ctx: Context for request cancellation and deadline.
// workflow: The workflow object to be created, structured as a CreateWorkflowRequest.
// Returns: JsonResponse containing the result of the workflow creation and an error if any.
// On error, an empty JsonResponse and the error are returned.
func (e *WorkflowService) CreateWorkflow(ctx context.Context, workflow CreateWorkflowRequest) (JsonResponse, error) {
	URL := e.client.config.BackendURL.JoinPath("workflows")
	payload := WorkflowCreatePayload{Data: &workflow}

	resp, err := e.client.makeHTTPRequest(ctx, http.MethodPost, URL.String(), payload)
	if err != nil {
		return JsonResponse{}, err
	}
	return resp, nil
}

// UpdateWorkflow updates an existing workflow identified by the given identifier.
// ctx: Context for request cancellation and deadline.
// identifier: A unique string identifying the workflow to be updated.
// workflow: The updated workflow data, structured as a pointer to UpdateWorkflowRequest.
// Returns: JsonResponse containing the result of the update operation and an error if any.
// On error, the JsonResponse received up to the point of error and the error are returned.
func (e *WorkflowService) UpdateWorkflow(ctx context.Context, identifier string, workflow *UpdateWorkflowRequest) (JsonResponse, error) {
	URL := e.client.config.BackendURL.JoinPath("workflows", identifier)
	payload := WorkflowUpdatePayload{Data: workflow}

	resp, err := e.client.makeHTTPRequest(ctx, http.MethodPut, URL.String(), payload)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UpdateWorkflowStatus updates the active status of a workflow.
// ctx: Context for request cancellation and deadline.
// identifier: A unique string identifying the workflow whose status is to be updated.
// status: A boolean value representing the new active status of the workflow.
// Returns: JsonResponse containing the result of the status update and an error if any.
// On error, the JsonResponse received up to the point of error and the error are returned.
func (e *WorkflowService) UpdateWorkflowStatus(ctx context.Context, identifier string, status bool) (JsonResponse, error) {
	URL := e.client.config.BackendURL.JoinPath("workflows", identifier)
	payload := WorkflowStatusUpdatePayload{
		Data: struct {
			Active bool `json:"active"`
		}{
			Active: status,
		},
	}

	resp, err := e.client.makeHTTPRequest(ctx, http.MethodPut, URL.String(), payload)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetWorkflows retrieves a list of workflows with pagination support.
// ctx: Context for request cancellation and deadline.
// page: Integer specifying the page number in the pagination.
// limit: Integer specifying the number of items per page.
// Returns: JsonResponse containing the list of workflows and an error if any.
// On error, the JsonResponse received up to the point of error and the error are returned.
func (e *WorkflowService) GetWorkflows(ctx context.Context, page int, limit int) (JsonResponse, error) {
	URL := e.client.config.BackendURL.JoinPath("workflows")
	v := URL.Query()
	v.Set("page", strconv.Itoa(page))
	v.Set("limit", strconv.Itoa(limit))
	URL.RawQuery = v.Encode()

	resp, err := e.client.makeHTTPRequest(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetTenant retrieves details of a specific tenant (workflow) identified by the identifier.
// ctx: Context for request cancellation and deadline.
// identifier: A unique string identifying the tenant (workflow).
// Returns: JsonResponse containing the tenant's details and an error if any.
// On error, the JsonResponse received up to the point of error and the error are returned.
func (e *WorkflowService) GetTenant(ctx context.Context, identifier string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("workflows", identifier)

	resp, err := e.client.makeHTTPRequest(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DeleteWorkflow deletes a workflow identified by the given identifier.
// ctx: Context for request cancellation and deadline.
// identifier: A unique string identifying the workflow to be deleted.
// Returns: JsonResponse confirming the deletion and an error if any.
// On error, the JsonResponse received up to the point of error and the error are returned.
func (e *WorkflowService) DeleteWorkflow(ctx context.Context, identifier string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("workflows", identifier)

	resp, err := e.client.makeHTTPRequest(ctx, http.MethodDelete, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
