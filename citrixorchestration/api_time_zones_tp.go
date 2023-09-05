/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// TimeZonesTPApiService TimeZonesTPApi service
type TimeZonesTPApiService service

type ApiTimeZonesTPGetTimeZonesRequest struct {
	ctx context.Context
	ApiService *TimeZonesTPApiService
	customerid string
	siteid string
	authorization *string
	citrixTransactionId *string
	xActionName *string
}

// Citrix Cloud authorization header: Bearer {token}
func (r ApiTimeZonesTPGetTimeZonesRequest) Authorization(authorization string) ApiTimeZonesTPGetTimeZonesRequest {
	r.authorization = &authorization
	return r
}

// Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned.
func (r ApiTimeZonesTPGetTimeZonesRequest) CitrixTransactionId(citrixTransactionId string) ApiTimeZonesTPGetTimeZonesRequest {
	r.citrixTransactionId = &citrixTransactionId
	return r
}

// Orchestration Action Name
func (r ApiTimeZonesTPGetTimeZonesRequest) XActionName(xActionName string) ApiTimeZonesTPGetTimeZonesRequest {
	r.xActionName = &xActionName
	return r
}

func (r ApiTimeZonesTPGetTimeZonesRequest) Execute() (*TimeZoneResponseModelCollection, *http.Response, error) {
	return r.ApiService.TimeZonesTPGetTimeZonesExecute(r)
}

/*
TimeZonesTPGetTimeZones Get a list of time zones supported by the site.

Get a (non-exhaustive) list of time zones supported by the site,
with time zone names represented in the caller's locale.


Note that all IANA (https://www.iana.org/time-zones) time zones are
supported as inputs to time zone APIs.  However sometimes callers
need a suitable list of time zones to display as options within a
UI. This API provides that capability. The options returned
represent a reasonable palette of options for a caller to choose
from when selecting a time zone.


The localized time zone names come from the Unicode CLDR
(http://cldr.unicode.org/) project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerid
 @param siteid
 @return ApiTimeZonesTPGetTimeZonesRequest
*/
func (a *TimeZonesTPApiService) TimeZonesTPGetTimeZones(ctx context.Context, customerid string, siteid string) ApiTimeZonesTPGetTimeZonesRequest {
	return ApiTimeZonesTPGetTimeZonesRequest{
		ApiService: a,
		ctx: ctx,
		customerid: customerid,
		siteid: siteid,
	}
}

// Execute executes the request
//  @return TimeZoneResponseModelCollection
func (a *TimeZonesTPApiService) TimeZonesTPGetTimeZonesExecute(r ApiTimeZonesTPGetTimeZonesRequest) (*TimeZoneResponseModelCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimeZoneResponseModelCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeZonesTPApiService.TimeZonesTPGetTimeZones")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/techpreview/{customerid}/{siteid}/TimeZones/All"
	localVarPath = strings.Replace(localVarPath, "{"+"customerid"+"}", url.PathEscape(parameterValueToString(r.customerid, "customerid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"siteid"+"}", url.PathEscape(parameterValueToString(r.siteid, "siteid")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	if r.citrixTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-TransactionId", r.citrixTransactionId, "")
	}
	if r.xActionName != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-ActionName", r.xActionName, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorData
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
			var v ErrorData
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
			var v ErrorData
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorData
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 406 {
			var v ErrorData
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
			var v ErrorData
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorData
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
