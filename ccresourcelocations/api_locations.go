/*
Resource Locations APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccresourcelocations

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// LocationsDAASService LocationsDAAS service
type LocationsDAASService service

type ApiLocationsCreateRequest struct {
	ctx context.Context
	ApiService *LocationsDAASService
	authorization *string
	citrixCustomerId *string
	accept *string
	model *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
}

// The access token.
func (r ApiLocationsCreateRequest) Authorization(authorization string) ApiLocationsCreateRequest {
	r.authorization = &authorization
	return r
}

// ID of the customer.
func (r ApiLocationsCreateRequest) CitrixCustomerId(citrixCustomerId string) ApiLocationsCreateRequest {
	r.citrixCustomerId = &citrixCustomerId
	return r
}

// Only supports application/json
func (r ApiLocationsCreateRequest) Accept(accept string) ApiLocationsCreateRequest {
	r.accept = &accept
	return r
}

// The resource location model.
func (r ApiLocationsCreateRequest) Model(model CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel) ApiLocationsCreateRequest {
	r.model = &model
	return r
}

func (r ApiLocationsCreateRequest) Execute() (*CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel, *http.Response, error) {
	return r.ApiService.LocationsCreateExecute(r)
}

/*
LocationsCreate Create a resource location for a customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLocationsCreateRequest
*/
func (a *LocationsDAASService) LocationsCreate(ctx context.Context) ApiLocationsCreateRequest {
	return ApiLocationsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
func (a *LocationsDAASService) LocationsCreateExecute(r ApiLocationsCreateRequest) (*CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsDAASService.LocationsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.citrixCustomerId == nil {
		return localVarReturnValue, nil, reportError("citrixCustomerId is required and must be specified")
	}
	if r.accept == nil {
		return localVarReturnValue, nil, reportError("accept is required and must be specified")
	}
	if r.model == nil {
		return localVarReturnValue, nil, reportError("model is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-CustomerId", r.citrixCustomerId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	// body params
	localVarPostBody = r.model
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

	if a.client.cfg.ResponseMiddleware != nil {
		err = a.client.cfg.ResponseMiddleware(localVarHTTPResponse, localVarBody)
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiLocationsDeleteRequest struct {
	ctx context.Context
	ApiService *LocationsDAASService
	authorization *string
	id string
	citrixCustomerId *string
	accept *string
}

// The access token.
func (r ApiLocationsDeleteRequest) Authorization(authorization string) ApiLocationsDeleteRequest {
	r.authorization = &authorization
	return r
}

// ID of the customer.
func (r ApiLocationsDeleteRequest) CitrixCustomerId(citrixCustomerId string) ApiLocationsDeleteRequest {
	r.citrixCustomerId = &citrixCustomerId
	return r
}

// Only supports application/json
func (r ApiLocationsDeleteRequest) Accept(accept string) ApiLocationsDeleteRequest {
	r.accept = &accept
	return r
}

func (r ApiLocationsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.LocationsDeleteExecute(r)
}

/*
LocationsDelete Delete a resource location.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The resource location id.
 @return ApiLocationsDeleteRequest
*/
func (a *LocationsDAASService) LocationsDelete(ctx context.Context, id string) ApiLocationsDeleteRequest {
	return ApiLocationsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *LocationsDAASService) LocationsDeleteExecute(r ApiLocationsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsDAASService.LocationsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
	}
	if r.citrixCustomerId == nil {
		return nil, reportError("citrixCustomerId is required and must be specified")
	}
	if r.accept == nil {
		return nil, reportError("accept is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-CustomerId", r.citrixCustomerId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
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

	if a.client.cfg.ResponseMiddleware != nil {
		err = a.client.cfg.ResponseMiddleware(localVarHTTPResponse, localVarBody)
		if err != nil {
			return localVarHTTPResponse, err
		}
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

type ApiLocationsGetRequest struct {
	ctx context.Context
	ApiService *LocationsDAASService
	authorization *string
	id string
	citrixCustomerId *string
	accept *string
}

// The access token.
func (r ApiLocationsGetRequest) Authorization(authorization string) ApiLocationsGetRequest {
	r.authorization = &authorization
	return r
}

// ID of the customer.
func (r ApiLocationsGetRequest) CitrixCustomerId(citrixCustomerId string) ApiLocationsGetRequest {
	r.citrixCustomerId = &citrixCustomerId
	return r
}

// Only supports application/json
func (r ApiLocationsGetRequest) Accept(accept string) ApiLocationsGetRequest {
	r.accept = &accept
	return r
}

func (r ApiLocationsGetRequest) Execute() (*CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel, *http.Response, error) {
	return r.ApiService.LocationsGetExecute(r)
}

/*
LocationsGet Get a resource location from id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The resource location id
 @return ApiLocationsGetRequest
*/
func (a *LocationsDAASService) LocationsGet(ctx context.Context, id string) ApiLocationsGetRequest {
	return ApiLocationsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
func (a *LocationsDAASService) LocationsGetExecute(r ApiLocationsGetRequest) (*CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsDAASService.LocationsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.citrixCustomerId == nil {
		return localVarReturnValue, nil, reportError("citrixCustomerId is required and must be specified")
	}
	if r.accept == nil {
		return localVarReturnValue, nil, reportError("accept is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-CustomerId", r.citrixCustomerId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

	if a.client.cfg.ResponseMiddleware != nil {
		err = a.client.cfg.ResponseMiddleware(localVarHTTPResponse, localVarBody)
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiLocationsGetAllRequest struct {
	ctx context.Context
	ApiService *LocationsDAASService
	authorization *string
	citrixCustomerId *string
	accept *string
}

// The access token.
func (r ApiLocationsGetAllRequest) Authorization(authorization string) ApiLocationsGetAllRequest {
	r.authorization = &authorization
	return r
}

// ID of the customer.
func (r ApiLocationsGetAllRequest) CitrixCustomerId(citrixCustomerId string) ApiLocationsGetAllRequest {
	r.citrixCustomerId = &citrixCustomerId
	return r
}

// Only supports application/json
func (r ApiLocationsGetAllRequest) Accept(accept string) ApiLocationsGetAllRequest {
	r.accept = &accept
	return r
}

func (r ApiLocationsGetAllRequest) Execute() (*CitrixCloudServicesRegistryApiModelsLocationsResourceLocationsResultsModel, *http.Response, error) {
	return r.ApiService.LocationsGetAllExecute(r)
}

/*
LocationsGetAll Get all resource locations for a customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLocationsGetAllRequest
*/
func (a *LocationsDAASService) LocationsGetAll(ctx context.Context) ApiLocationsGetAllRequest {
	return ApiLocationsGetAllRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CitrixCloudServicesRegistryApiModelsLocationsResourceLocationsResultsModel
func (a *LocationsDAASService) LocationsGetAllExecute(r ApiLocationsGetAllRequest) (*CitrixCloudServicesRegistryApiModelsLocationsResourceLocationsResultsModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationsResultsModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsDAASService.LocationsGetAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.citrixCustomerId == nil {
		return localVarReturnValue, nil, reportError("citrixCustomerId is required and must be specified")
	}
	if r.accept == nil {
		return localVarReturnValue, nil, reportError("accept is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-CustomerId", r.citrixCustomerId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

	if a.client.cfg.ResponseMiddleware != nil {
		err = a.client.cfg.ResponseMiddleware(localVarHTTPResponse, localVarBody)
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiLocationsUpdateRequest struct {
	ctx context.Context
	ApiService *LocationsDAASService
	authorization *string
	citrixCustomerId *string
	accept *string
	id string
	model *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel
}

// The access token.
func (r ApiLocationsUpdateRequest) Authorization(authorization string) ApiLocationsUpdateRequest {
	r.authorization = &authorization
	return r
}

// ID of the customer.
func (r ApiLocationsUpdateRequest) CitrixCustomerId(citrixCustomerId string) ApiLocationsUpdateRequest {
	r.citrixCustomerId = &citrixCustomerId
	return r
}

// Only supports application/json
func (r ApiLocationsUpdateRequest) Accept(accept string) ApiLocationsUpdateRequest {
	r.accept = &accept
	return r
}

// The resource location model.
func (r ApiLocationsUpdateRequest) Model(model CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) ApiLocationsUpdateRequest {
	r.model = &model
	return r
}

func (r ApiLocationsUpdateRequest) Execute() (*CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel, *http.Response, error) {
	return r.ApiService.LocationsUpdateExecute(r)
}

/*
LocationsUpdate Update the customer resource location information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The resource location id.
 @return ApiLocationsUpdateRequest
*/
func (a *LocationsDAASService) LocationsUpdate(ctx context.Context, id string) ApiLocationsUpdateRequest {
	return ApiLocationsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
func (a *LocationsDAASService) LocationsUpdateExecute(r ApiLocationsUpdateRequest) (*CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsDAASService.LocationsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.citrixCustomerId == nil {
		return localVarReturnValue, nil, reportError("citrixCustomerId is required and must be specified")
	}
	if r.accept == nil {
		return localVarReturnValue, nil, reportError("accept is required and must be specified")
	}
	if r.model == nil {
		return localVarReturnValue, nil, reportError("model is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-CustomerId", r.citrixCustomerId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	// body params
	localVarPostBody = r.model
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

	if a.client.cfg.ResponseMiddleware != nil {
		err = a.client.cfg.ResponseMiddleware(localVarHTTPResponse, localVarBody)
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
