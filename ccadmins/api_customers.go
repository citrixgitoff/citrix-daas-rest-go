/*
Citrix.CloudServices.Administrators.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CustomersAPIService CustomersAPI service
type CustomersAPIService service

type ApiCustomerCustomersIdDeleteRequest struct {
	ctx context.Context
	ApiService *CustomersAPIService
	id string
	customer string
	citrixConsistencyToken *string
	xCwsTransactionId *string
	citrixTimeoutMs *float64
}

// CosmosDB consistency token.
func (r ApiCustomerCustomersIdDeleteRequest) CitrixConsistencyToken(citrixConsistencyToken string) ApiCustomerCustomersIdDeleteRequest {
	r.citrixConsistencyToken = &citrixConsistencyToken
	return r
}

// Used for Citrix Cloud telemetry correlation.
func (r ApiCustomerCustomersIdDeleteRequest) XCwsTransactionId(xCwsTransactionId string) ApiCustomerCustomersIdDeleteRequest {
	r.xCwsTransactionId = &xCwsTransactionId
	return r
}

// Maximum time caller is willing to wait for a response before timing out the request, in milliseconds
func (r ApiCustomerCustomersIdDeleteRequest) CitrixTimeoutMs(citrixTimeoutMs float64) ApiCustomerCustomersIdDeleteRequest {
	r.citrixTimeoutMs = &citrixTimeoutMs
	return r
}

func (r ApiCustomerCustomersIdDeleteRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.CustomerCustomersIdDeleteExecute(r)
}

/*
CustomerCustomersIdDelete Delete customer's administrators. [ServiceKey][RootOnly]

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the customer.
 @param customer
 @return ApiCustomerCustomersIdDeleteRequest
*/
func (a *CustomersAPIService) CustomerCustomersIdDelete(ctx context.Context, id string, customer string) ApiCustomerCustomersIdDeleteRequest {
	return ApiCustomerCustomersIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		customer: customer,
	}
}

// Execute executes the request
//  @return bool
func (a *CustomersAPIService) CustomerCustomersIdDeleteExecute(r ApiCustomerCustomersIdDeleteRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomersAPIService.CustomerCustomersIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{customer}/customers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customer"+"}", url.PathEscape(parameterValueToString(r.customer, "customer")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.citrixConsistencyToken != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-ConsistencyToken", r.citrixConsistencyToken, "")
	}
	if r.xCwsTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Cws-TransactionId", r.xCwsTransactionId, "")
	}
	if r.citrixTimeoutMs != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-Timeout-ms", r.citrixTimeoutMs, "")
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

type ApiCustomerCustomersPostRequest struct {
	ctx context.Context
	ApiService *CustomersAPIService
	customer string
	citrixConsistencyToken *string
	xCwsTransactionId *string
	citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel *CitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel
}

// CosmosDB consistency token.
func (r ApiCustomerCustomersPostRequest) CitrixConsistencyToken(citrixConsistencyToken string) ApiCustomerCustomersPostRequest {
	r.citrixConsistencyToken = &citrixConsistencyToken
	return r
}

// Used for Citrix Cloud telemetry correlation.
func (r ApiCustomerCustomersPostRequest) XCwsTransactionId(xCwsTransactionId string) ApiCustomerCustomersPostRequest {
	r.xCwsTransactionId = &xCwsTransactionId
	return r
}

func (r ApiCustomerCustomersPostRequest) CitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel(citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel CitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel) ApiCustomerCustomersPostRequest {
	r.citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel = &citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel
	return r
}

func (r ApiCustomerCustomersPostRequest) Execute() (*SystemNetHttpHttpResponseMessage, *http.Response, error) {
	return r.ApiService.CustomerCustomersPostExecute(r)
}

/*
CustomerCustomersPost Creates a new customer. [ServiceKey][RootOnly]

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customer
 @return ApiCustomerCustomersPostRequest
*/
func (a *CustomersAPIService) CustomerCustomersPost(ctx context.Context, customer string) ApiCustomerCustomersPostRequest {
	return ApiCustomerCustomersPostRequest{
		ApiService: a,
		ctx: ctx,
		customer: customer,
	}
}

// Execute executes the request
//  @return SystemNetHttpHttpResponseMessage
func (a *CustomersAPIService) CustomerCustomersPostExecute(r ApiCustomerCustomersPostRequest) (*SystemNetHttpHttpResponseMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SystemNetHttpHttpResponseMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomersAPIService.CustomerCustomersPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{customer}/customers"
	localVarPath = strings.Replace(localVarPath, "{"+"customer"+"}", url.PathEscape(parameterValueToString(r.customer, "customer")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.citrixConsistencyToken != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-ConsistencyToken", r.citrixConsistencyToken, "")
	}
	if r.xCwsTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Cws-TransactionId", r.xCwsTransactionId, "")
	}
	// body params
	localVarPostBody = r.citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel
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
