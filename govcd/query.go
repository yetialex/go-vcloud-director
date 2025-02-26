/*
 * Copyright 2019 VMware, Inc.  All rights reserved.  Licensed under the Apache v2 License.
 */

package govcd

import (
	"fmt"
	"net/http"

	"github.com/yetialex/go-vcloud-director/v2/types/v56"
)

type Results struct {
	Results *types.QueryResultRecordsType
	client  *Client
}

func NewResults(cli *Client) *Results {
	return &Results{
		Results: new(types.QueryResultRecordsType),
		client:  cli,
	}
}

func (vcdCli *VCDClient) Query(params map[string]string) (Results, error) {

	req := vcdCli.Client.NewRequest(params, http.MethodGet, vcdCli.QueryHREF, nil)
	req.Header.Add("Accept", "vnd.vmware.vcloud.org+xml;version="+vcdCli.Client.APIVersion)

	return getResult(&vcdCli.Client, req)
}

func (vdc *Vdc) Query(params map[string]string) (Results, error) {
	queryUrl := vdc.client.VCDHREF
	queryUrl.Path += "/query"
	req := vdc.client.NewRequest(params, http.MethodGet, queryUrl, nil)
	req.Header.Add("Accept", "vnd.vmware.vcloud.org+xml;version="+vdc.client.APIVersion)

	return getResult(vdc.client, req)
}

// QueryWithNotEncodedParams uses Query API to search for requested data
func (client *Client) QueryWithNotEncodedParams(params map[string]string, notEncodedParams map[string]string) (Results, error) {
	return client.QueryWithNotEncodedParamsWithApiVersion(params, notEncodedParams, client.APIVersion)
}

// QueryWithNotEncodedParams uses Query API to search for requested data
func (client *Client) QueryWithNotEncodedParamsWithApiVersion(params map[string]string, notEncodedParams map[string]string, apiVersion string) (Results, error) {
	queryUrl := client.VCDHREF
	queryUrl.Path += "/query"

	req := client.NewRequestWitNotEncodedParamsWithApiVersion(params, notEncodedParams, http.MethodGet, queryUrl, nil, apiVersion)
	req.Header.Add("Accept", "vnd.vmware.vcloud.org+xml;version="+apiVersion)

	return getResult(client, req)
}

func (vcdCli *VCDClient) QueryWithNotEncodedParams(params map[string]string, notEncodedParams map[string]string) (Results, error) {
	return vcdCli.Client.QueryWithNotEncodedParams(params, notEncodedParams)
}

func (vdc *Vdc) QueryWithNotEncodedParams(params map[string]string, notEncodedParams map[string]string) (Results, error) {
	return vdc.client.QueryWithNotEncodedParams(params, notEncodedParams)
}

func (vcdCli *VCDClient) QueryWithNotEncodedParamsWithApiVersion(params map[string]string, notEncodedParams map[string]string, apiVersion string) (Results, error) {
	return vcdCli.Client.QueryWithNotEncodedParamsWithApiVersion(params, notEncodedParams, apiVersion)
}

func (vdc *Vdc) QueryWithNotEncodedParamsWithApiVersion(params map[string]string, notEncodedParams map[string]string, apiVersion string) (Results, error) {
	return vdc.client.QueryWithNotEncodedParamsWithApiVersion(params, notEncodedParams, apiVersion)
}

func getResult(client *Client, request *http.Request) (Results, error) {
	resp, err := checkResp(client.Http.Do(request))
	if err != nil {
		return Results{}, fmt.Errorf("error retrieving query: %s", err)
	}

	results := NewResults(client)

	if err = decodeBody(types.BodyTypeXML, resp, results.Results); err != nil {
		return Results{}, fmt.Errorf("error decoding query results: %s", err)
	}

	return *results, nil
}

func (vdc *Vdc) QueryWithCustomHeaders(params, additionalHeaders map[string]string) (Results, error) {
	queryUrl := vdc.client.VCDHREF
	queryUrl.Path += "/query"
	req := vdc.client.NewRequest(params, http.MethodGet, queryUrl, nil)
	req.Header.Add("Accept", "vnd.vmware.vcloud.org+xml;version="+vdc.client.APIVersion)
	for k, v := range additionalHeaders {
		req.Header.Add(k, v)
	}
	return getResult(vdc.client, req)
}
