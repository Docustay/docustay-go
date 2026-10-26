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


// EmailAPIService EmailAPI service
type EmailAPIService service

type ApiListEmailsRequest struct {
	ctx context.Context
	ApiService *EmailAPIService
	documentId *string
	to *string
	status *string
	kind *string
	limit *int32
	cursor *string
}

func (r ApiListEmailsRequest) DocumentId(documentId string) ApiListEmailsRequest {
	r.documentId = &documentId
	return r
}

func (r ApiListEmailsRequest) To(to string) ApiListEmailsRequest {
	r.to = &to
	return r
}

// Only emails in this state; any other value is refused (R13-D205).
func (r ApiListEmailsRequest) Status(status string) ApiListEmailsRequest {
	r.status = &status
	return r
}

// Only emails of this kind; any other value is refused (R13-D346).
func (r ApiListEmailsRequest) Kind(kind string) ApiListEmailsRequest {
	r.kind = &kind
	return r
}

func (r ApiListEmailsRequest) Limit(limit int32) ApiListEmailsRequest {
	r.limit = &limit
	return r
}

func (r ApiListEmailsRequest) Cursor(cursor string) ApiListEmailsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiListEmailsRequest) Execute() (*ListEmails200Response, *http.Response, error) {
	return r.ApiService.ListEmailsExecute(r)
}

/*
ListEmails The log of email about your documents

Every email sent about a document, to a person named on it: kind, subject, status (sent, delivered, bounced, complained, failed, paused). Filter by `documentId`, `to`, `status`, `kind`. No bodies, no links.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListEmailsRequest
*/
func (a *EmailAPIService) ListEmails(ctx context.Context) ApiListEmailsRequest {
	return ApiListEmailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListEmails200Response
func (a *EmailAPIService) ListEmailsExecute(r ApiListEmailsRequest) (*ListEmails200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListEmails200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailAPIService.ListEmails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/emails"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.documentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "documentId", r.documentId, "form", "")
	}
	if r.to != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.kind != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "kind", r.kind, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
