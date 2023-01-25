/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// QuoteApiService QuoteApi service
type QuoteApiService service

type ApiGetQuoteRequest struct {
	ctx _context.Context
	ApiService *QuoteApiService
	symbols *string
	formatted *bool
	region *string
	lang *string
	includePrePost *bool
	fields *string
	corsDomain *string
}

func (r ApiGetQuoteRequest) Symbols(symbols string) ApiGetQuoteRequest {
	r.symbols = &symbols
	return r
}
func (r ApiGetQuoteRequest) Formatted(formatted bool) ApiGetQuoteRequest {
	r.formatted = &formatted
	return r
}
func (r ApiGetQuoteRequest) Region(region string) ApiGetQuoteRequest {
	r.region = &region
	return r
}
func (r ApiGetQuoteRequest) Lang(lang string) ApiGetQuoteRequest {
	r.lang = &lang
	return r
}
func (r ApiGetQuoteRequest) IncludePrePost(includePrePost bool) ApiGetQuoteRequest {
	r.includePrePost = &includePrePost
	return r
}
func (r ApiGetQuoteRequest) Fields(fields string) ApiGetQuoteRequest {
	r.fields = &fields
	return r
}
func (r ApiGetQuoteRequest) CorsDomain(corsDomain string) ApiGetQuoteRequest {
	r.corsDomain = &corsDomain
	return r
}

func (r ApiGetQuoteRequest) Execute() (QuoteResponse, *_nethttp.Response, error) {
	return r.ApiService.GetQuoteExecute(r)
}

/*
 * GetQuote Returns quotes for the specified symbols
 * Returns quotes for the specified symbols
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetQuoteRequest
 */
func (a *QuoteApiService) GetQuote(ctx _context.Context) ApiGetQuoteRequest {
	return ApiGetQuoteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return QuoteResponse
 */
func (a *QuoteApiService) GetQuoteExecute(r ApiGetQuoteRequest) (QuoteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QuoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuoteApiService.GetQuote")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v7/finance/quote"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.symbols == nil {
		return localVarReturnValue, nil, reportError("symbols is required and must be specified")
	}

	if r.formatted != nil {
		localVarQueryParams.Add("formatted", parameterToString(*r.formatted, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.lang != nil {
		localVarQueryParams.Add("lang", parameterToString(*r.lang, ""))
	}
	if r.includePrePost != nil {
		localVarQueryParams.Add("includePrePost", parameterToString(*r.includePrePost, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.corsDomain != nil {
		localVarQueryParams.Add("corsDomain", parameterToString(*r.corsDomain, ""))
	}
	localVarQueryParams.Add("symbols", parameterToString(*r.symbols, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}