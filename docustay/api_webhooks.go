/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// WebhooksAPIService WebhooksAPI service
type WebhooksAPIService service

type ApiWebhookDocumentsAiAllowanceLowRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsAiAllowanceLowRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsAiAllowanceLowRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsAiAllowanceLowRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsAiAllowanceLowExecute(r)
}

/*
WebhookDocumentsAiAllowanceLow documents.ai.allowance_low

The workspace's Scribe allowance is nearly used. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsAiAllowanceLowRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsAiAllowanceLow(ctx context.Context) ApiWebhookDocumentsAiAllowanceLowRequest {
	return ApiWebhookDocumentsAiAllowanceLowRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsAiAllowanceLowExecute(r ApiWebhookDocumentsAiAllowanceLowRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsAiAllowanceLow")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.ai.allowance_low"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsApprovalDecidedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsApprovalDecidedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsApprovalDecidedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsApprovalDecidedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsApprovalDecidedExecute(r)
}

/*
WebhookDocumentsApprovalDecided documents.approval.decided

A person approved or declined a send that a draft-only key asked for. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsApprovalDecidedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsApprovalDecided(ctx context.Context) ApiWebhookDocumentsApprovalDecidedRequest {
	return ApiWebhookDocumentsApprovalDecidedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsApprovalDecidedExecute(r ApiWebhookDocumentsApprovalDecidedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsApprovalDecided")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.approval.decided"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsApprovalRequestedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsApprovalRequestedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsApprovalRequestedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsApprovalRequestedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsApprovalRequestedExecute(r)
}

/*
WebhookDocumentsApprovalRequested documents.approval.requested

A draft-only key asked to send, and a person must approve. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsApprovalRequestedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsApprovalRequested(ctx context.Context) ApiWebhookDocumentsApprovalRequestedRequest {
	return ApiWebhookDocumentsApprovalRequestedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsApprovalRequestedExecute(r ApiWebhookDocumentsApprovalRequestedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsApprovalRequested")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.approval.requested"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentBouncedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentBouncedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentBouncedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentBouncedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentBouncedExecute(r)
}

/*
WebhookDocumentsDocumentBounced documents.document.bounced

An email to a signer could not be delivered. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentBouncedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentBounced(ctx context.Context) ApiWebhookDocumentsDocumentBouncedRequest {
	return ApiWebhookDocumentsDocumentBouncedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentBouncedExecute(r ApiWebhookDocumentsDocumentBouncedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentBounced")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.bounced"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentChangesRequestedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentChangesRequestedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentChangesRequestedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentChangesRequestedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentChangesRequestedExecute(r)
}

/*
WebhookDocumentsDocumentChangesRequested documents.document.changes_requested

A person asked for changes instead of signing. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentChangesRequestedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentChangesRequested(ctx context.Context) ApiWebhookDocumentsDocumentChangesRequestedRequest {
	return ApiWebhookDocumentsDocumentChangesRequestedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentChangesRequestedExecute(r ApiWebhookDocumentsDocumentChangesRequestedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentChangesRequested")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.changes_requested"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentCreatedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentCreatedRequest *WebhookDocumentsDocumentCreatedRequest
}

func (r ApiWebhookDocumentsDocumentCreatedRequest) WebhookDocumentsDocumentCreatedRequest(webhookDocumentsDocumentCreatedRequest WebhookDocumentsDocumentCreatedRequest) ApiWebhookDocumentsDocumentCreatedRequest {
	r.webhookDocumentsDocumentCreatedRequest = &webhookDocumentsDocumentCreatedRequest
	return r
}

func (r ApiWebhookDocumentsDocumentCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentCreatedExecute(r)
}

/*
WebhookDocumentsDocumentCreated documents.document.created

A document was made. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentCreatedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentCreated(ctx context.Context) ApiWebhookDocumentsDocumentCreatedRequest {
	return ApiWebhookDocumentsDocumentCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentCreatedExecute(r ApiWebhookDocumentsDocumentCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.created"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentCreatedRequest == nil {
		return nil, reportError("webhookDocumentsDocumentCreatedRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentCreatedRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentDeclinedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentDeclinedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentDeclinedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentDeclinedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentDeclinedExecute(r)
}

/*
WebhookDocumentsDocumentDeclined documents.document.declined

A person declined to sign. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentDeclinedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentDeclined(ctx context.Context) ApiWebhookDocumentsDocumentDeclinedRequest {
	return ApiWebhookDocumentsDocumentDeclinedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentDeclinedExecute(r ApiWebhookDocumentsDocumentDeclinedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentDeclined")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.declined"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentDelegatedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentDelegatedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentDelegatedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentDelegatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentDelegatedExecute(r)
}

/*
WebhookDocumentsDocumentDelegated documents.document.delegated

A person passed their part to someone else. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentDelegatedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentDelegated(ctx context.Context) ApiWebhookDocumentsDocumentDelegatedRequest {
	return ApiWebhookDocumentsDocumentDelegatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentDelegatedExecute(r ApiWebhookDocumentsDocumentDelegatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentDelegated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.delegated"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentDownloadedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentDownloadedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentDownloadedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentDownloadedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentDownloadedExecute(r)
}

/*
WebhookDocumentsDocumentDownloaded documents.document.downloaded

Someone downloaded the document or its sealed copy. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentDownloadedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentDownloaded(ctx context.Context) ApiWebhookDocumentsDocumentDownloadedRequest {
	return ApiWebhookDocumentsDocumentDownloadedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentDownloadedExecute(r ApiWebhookDocumentsDocumentDownloadedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentDownloaded")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.downloaded"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentExecutedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentExecutedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentExecutedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentExecutedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentExecutedExecute(r)
}

/*
WebhookDocumentsDocumentExecuted documents.document.executed

Everyone signed and the sealed copy is ready. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentExecutedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentExecuted(ctx context.Context) ApiWebhookDocumentsDocumentExecutedRequest {
	return ApiWebhookDocumentsDocumentExecutedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentExecutedExecute(r ApiWebhookDocumentsDocumentExecutedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentExecuted")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.executed"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentExpiredRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentExpiredRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentExpiredRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentExpiredRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentExpiredExecute(r)
}

/*
WebhookDocumentsDocumentExpired documents.document.expired

The document reached its deadline unsigned. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentExpiredRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentExpired(ctx context.Context) ApiWebhookDocumentsDocumentExpiredRequest {
	return ApiWebhookDocumentsDocumentExpiredRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentExpiredExecute(r ApiWebhookDocumentsDocumentExpiredRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentExpired")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.expired"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentOpenedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentOpenedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentOpenedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentOpenedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentOpenedExecute(r)
}

/*
WebhookDocumentsDocumentOpened documents.document.opened

A person opened the document. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentOpenedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentOpened(ctx context.Context) ApiWebhookDocumentsDocumentOpenedRequest {
	return ApiWebhookDocumentsDocumentOpenedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentOpenedExecute(r ApiWebhookDocumentsDocumentOpenedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentOpened")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.opened"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentReminderDueRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentReminderDueRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentReminderDueRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentReminderDueRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentReminderDueExecute(r)
}

/*
WebhookDocumentsDocumentReminderDue documents.document.reminder_due

A scheduled reminder came due. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentReminderDueRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentReminderDue(ctx context.Context) ApiWebhookDocumentsDocumentReminderDueRequest {
	return ApiWebhookDocumentsDocumentReminderDueRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentReminderDueExecute(r ApiWebhookDocumentsDocumentReminderDueRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentReminderDue")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.reminder_due"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentSentRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentSentRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentSentRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentSentRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentSentExecute(r)
}

/*
WebhookDocumentsDocumentSent documents.document.sent

A document was sent to its people. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentSentRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentSent(ctx context.Context) ApiWebhookDocumentsDocumentSentRequest {
	return ApiWebhookDocumentsDocumentSentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentSentExecute(r ApiWebhookDocumentsDocumentSentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentSent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.sent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentSignedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentSignedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentSignedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentSignedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentSignedExecute(r)
}

/*
WebhookDocumentsDocumentSigned documents.document.signed

A person signed their part. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentSignedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentSigned(ctx context.Context) ApiWebhookDocumentsDocumentSignedRequest {
	return ApiWebhookDocumentsDocumentSignedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentSignedExecute(r ApiWebhookDocumentsDocumentSignedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentSigned")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.signed"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsDocumentVoidedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsDocumentVoidedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsDocumentVoidedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsDocumentVoidedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsDocumentVoidedExecute(r)
}

/*
WebhookDocumentsDocumentVoided documents.document.voided

The sender cancelled the document. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsDocumentVoidedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsDocumentVoided(ctx context.Context) ApiWebhookDocumentsDocumentVoidedRequest {
	return ApiWebhookDocumentsDocumentVoidedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsDocumentVoidedExecute(r ApiWebhookDocumentsDocumentVoidedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsDocumentVoided")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.document.voided"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsEmailBouncedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsEmailBouncedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsEmailBouncedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsEmailBouncedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsEmailBouncedExecute(r)
}

/*
WebhookDocumentsEmailBounced documents.email.bounced

An email bounced. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsEmailBouncedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsEmailBounced(ctx context.Context) ApiWebhookDocumentsEmailBouncedRequest {
	return ApiWebhookDocumentsEmailBouncedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsEmailBouncedExecute(r ApiWebhookDocumentsEmailBouncedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsEmailBounced")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.email.bounced"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsEmailComplainedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsEmailComplainedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsEmailComplainedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsEmailComplainedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsEmailComplainedExecute(r)
}

/*
WebhookDocumentsEmailComplained documents.email.complained

A recipient marked an email as spam. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsEmailComplainedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsEmailComplained(ctx context.Context) ApiWebhookDocumentsEmailComplainedRequest {
	return ApiWebhookDocumentsEmailComplainedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsEmailComplainedExecute(r ApiWebhookDocumentsEmailComplainedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsEmailComplained")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.email.complained"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsEmailDeliveredRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsEmailDeliveredRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsEmailDeliveredRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsEmailDeliveredRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsEmailDeliveredExecute(r)
}

/*
WebhookDocumentsEmailDelivered documents.email.delivered

The mail provider delivered an email. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsEmailDeliveredRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsEmailDelivered(ctx context.Context) ApiWebhookDocumentsEmailDeliveredRequest {
	return ApiWebhookDocumentsEmailDeliveredRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsEmailDeliveredExecute(r ApiWebhookDocumentsEmailDeliveredRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsEmailDelivered")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.email.delivered"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsEmailPausedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsEmailPausedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsEmailPausedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsEmailPausedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsEmailPausedExecute(r)
}

/*
WebhookDocumentsEmailPaused documents.email.paused

Sending to an address was paused after repeated problems. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsEmailPausedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsEmailPaused(ctx context.Context) ApiWebhookDocumentsEmailPausedRequest {
	return ApiWebhookDocumentsEmailPausedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsEmailPausedExecute(r ApiWebhookDocumentsEmailPausedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsEmailPaused")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.email.paused"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsPaymentFailedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsPaymentFailedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsPaymentFailedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsPaymentFailedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsPaymentFailedExecute(r)
}

/*
WebhookDocumentsPaymentFailed documents.payment.failed

A payment, or a plan invoice, failed. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsPaymentFailedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsPaymentFailed(ctx context.Context) ApiWebhookDocumentsPaymentFailedRequest {
	return ApiWebhookDocumentsPaymentFailedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsPaymentFailedExecute(r ApiWebhookDocumentsPaymentFailedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsPaymentFailed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.payment.failed"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsPaymentHoldLapsedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsPaymentHoldLapsedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsPaymentHoldLapsedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsPaymentHoldLapsedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsPaymentHoldLapsedExecute(r)
}

/*
WebhookDocumentsPaymentHoldLapsed documents.payment.hold_lapsed

A payment hold ran out before the signer paid. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsPaymentHoldLapsedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsPaymentHoldLapsed(ctx context.Context) ApiWebhookDocumentsPaymentHoldLapsedRequest {
	return ApiWebhookDocumentsPaymentHoldLapsedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsPaymentHoldLapsedExecute(r ApiWebhookDocumentsPaymentHoldLapsedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsPaymentHoldLapsed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.payment.hold_lapsed"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsPaymentReceivedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsPaymentReceivedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsPaymentReceivedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsPaymentReceivedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsPaymentReceivedExecute(r)
}

/*
WebhookDocumentsPaymentReceived documents.payment.received

A signer paid. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsPaymentReceivedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsPaymentReceived(ctx context.Context) ApiWebhookDocumentsPaymentReceivedRequest {
	return ApiWebhookDocumentsPaymentReceivedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsPaymentReceivedExecute(r ApiWebhookDocumentsPaymentReceivedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsPaymentReceived")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.payment.received"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsTeamCreatedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsTeamCreatedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsTeamCreatedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsTeamCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsTeamCreatedExecute(r)
}

/*
WebhookDocumentsTeamCreated documents.team.created

A team was made. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsTeamCreatedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsTeamCreated(ctx context.Context) ApiWebhookDocumentsTeamCreatedRequest {
	return ApiWebhookDocumentsTeamCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsTeamCreatedExecute(r ApiWebhookDocumentsTeamCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsTeamCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.team.created"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsTeamDeletedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsTeamDeletedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsTeamDeletedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsTeamDeletedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsTeamDeletedExecute(r)
}

/*
WebhookDocumentsTeamDeleted documents.team.deleted

A team was deleted. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsTeamDeletedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsTeamDeleted(ctx context.Context) ApiWebhookDocumentsTeamDeletedRequest {
	return ApiWebhookDocumentsTeamDeletedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsTeamDeletedExecute(r ApiWebhookDocumentsTeamDeletedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsTeamDeleted")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.team.deleted"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsTeamUpdatedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsTeamUpdatedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsTeamUpdatedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsTeamUpdatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsTeamUpdatedExecute(r)
}

/*
WebhookDocumentsTeamUpdated documents.team.updated

A team was changed. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsTeamUpdatedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsTeamUpdated(ctx context.Context) ApiWebhookDocumentsTeamUpdatedRequest {
	return ApiWebhookDocumentsTeamUpdatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsTeamUpdatedExecute(r ApiWebhookDocumentsTeamUpdatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsTeamUpdated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.team.updated"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsTemplateCreatedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsTemplateCreatedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsTemplateCreatedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsTemplateCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsTemplateCreatedExecute(r)
}

/*
WebhookDocumentsTemplateCreated documents.template.created

A template was made. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsTemplateCreatedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsTemplateCreated(ctx context.Context) ApiWebhookDocumentsTemplateCreatedRequest {
	return ApiWebhookDocumentsTemplateCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsTemplateCreatedExecute(r ApiWebhookDocumentsTemplateCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsTemplateCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.template.created"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsTemplateFieldsDetectedRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsTemplateFieldsDetectedRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsTemplateFieldsDetectedRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsTemplateFieldsDetectedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsTemplateFieldsDetectedExecute(r)
}

/*
WebhookDocumentsTemplateFieldsDetected documents.template.fields_detected

Fields were found in an uploaded file. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsTemplateFieldsDetectedRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsTemplateFieldsDetected(ctx context.Context) ApiWebhookDocumentsTemplateFieldsDetectedRequest {
	return ApiWebhookDocumentsTemplateFieldsDetectedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsTemplateFieldsDetectedExecute(r ApiWebhookDocumentsTemplateFieldsDetectedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsTemplateFieldsDetected")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.template.fields_detected"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhookDocumentsUsageThresholdRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookDocumentsDocumentSentRequest *WebhookDocumentsDocumentSentRequest
}

func (r ApiWebhookDocumentsUsageThresholdRequest) WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest WebhookDocumentsDocumentSentRequest) ApiWebhookDocumentsUsageThresholdRequest {
	r.webhookDocumentsDocumentSentRequest = &webhookDocumentsDocumentSentRequest
	return r
}

func (r ApiWebhookDocumentsUsageThresholdRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhookDocumentsUsageThresholdExecute(r)
}

/*
WebhookDocumentsUsageThreshold documents.usage.threshold

The workspace crossed a usage threshold. Signed with the Standard Webhooks headers webhook-id, webhook-timestamp and webhook-signature.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDocumentsUsageThresholdRequest
*/
func (a *WebhooksAPIService) WebhookDocumentsUsageThreshold(ctx context.Context) ApiWebhookDocumentsUsageThresholdRequest {
	return ApiWebhookDocumentsUsageThresholdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhookDocumentsUsageThresholdExecute(r ApiWebhookDocumentsUsageThresholdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhookDocumentsUsageThreshold")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents.usage.threshold"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookDocumentsDocumentSentRequest == nil {
		return nil, reportError("webhookDocumentsDocumentSentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookDocumentsDocumentSentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
