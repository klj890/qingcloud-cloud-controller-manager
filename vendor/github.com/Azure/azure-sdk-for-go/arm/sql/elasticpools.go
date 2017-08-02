package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ElasticPoolsClient is the the Azure SQL Database management API provides a
// RESTful set of web services that interact with Azure SQL Database services
// to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type ElasticPoolsClient struct {
	ManagementClient
}

// NewElasticPoolsClient creates an instance of the ElasticPoolsClient client.
func NewElasticPoolsClient(subscriptionID string) ElasticPoolsClient {
	return NewElasticPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewElasticPoolsClientWithBaseURI creates an instance of the
// ElasticPoolsClient client.
func NewElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolsClient {
	return ElasticPoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new elastic pool or updates an existing elastic
// pool. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. elasticPoolName is the
// name of the elastic pool to be operated on (updated or created). parameters
// is the required parameters for creating or updating an elastic pool.
func (client ElasticPoolsClient) CreateOrUpdate(resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool, cancel <-chan struct{}) (<-chan ElasticPool, <-chan error) {
	resultChan := make(chan ElasticPool, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result ElasticPool
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, elasticPoolName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ElasticPoolsClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) CreateOrUpdateResponder(resp *http.Response) (result ElasticPool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the elastic pool.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. elasticPoolName is the
// name of the elastic pool to be deleted.
func (client ElasticPoolsClient) Delete(resourceGroupName string, serverName string, elasticPoolName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ElasticPoolsClient) DeletePreparer(resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an elastic pool.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. elasticPoolName is the
// name of the elastic pool to be retrieved.
func (client ElasticPoolsClient) Get(resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPool, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ElasticPoolsClient) GetPreparer(resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) GetResponder(resp *http.Response) (result ElasticPool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDatabase gets a database inside of an elastic pool.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. elasticPoolName is the
// name of the elastic pool to be retrieved. databaseName is the name of the
// database to be retrieved.
func (client ElasticPoolsClient) GetDatabase(resourceGroupName string, serverName string, elasticPoolName string, databaseName string) (result Database, err error) {
	req, err := client.GetDatabasePreparer(resourceGroupName, serverName, elasticPoolName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "GetDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "GetDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.GetDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "GetDatabase", resp, "Failure responding to request")
	}

	return
}

// GetDatabasePreparer prepares the GetDatabase request.
func (client ElasticPoolsClient) GetDatabasePreparer(resourceGroupName string, serverName string, elasticPoolName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/databases/{databaseName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetDatabaseSender sends the GetDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) GetDatabaseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetDatabaseResponder handles the response to the GetDatabase request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) GetDatabaseResponder(resp *http.Response) (result Database, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListActivity returns elastic pool activities.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. elasticPoolName is the
// name of the elastic pool for which to get the current activity.
func (client ElasticPoolsClient) ListActivity(resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPoolActivityListResult, err error) {
	req, err := client.ListActivityPreparer(resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListActivity", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListActivitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListActivity", resp, "Failure sending request")
		return
	}

	result, err = client.ListActivityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListActivity", resp, "Failure responding to request")
	}

	return
}

// ListActivityPreparer prepares the ListActivity request.
func (client ElasticPoolsClient) ListActivityPreparer(resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/elasticPoolActivity", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListActivitySender sends the ListActivity request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) ListActivitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListActivityResponder handles the response to the ListActivity request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListActivityResponder(resp *http.Response) (result ElasticPoolActivityListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer returns a list of elastic pools in a server.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server.
func (client ElasticPoolsClient) ListByServer(resourceGroupName string, serverName string) (result ElasticPoolListResult, err error) {
	req, err := client.ListByServerPreparer(resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client ElasticPoolsClient) ListByServerPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListByServerResponder(resp *http.Response) (result ElasticPoolListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListDatabaseActivity returns activity on databases inside of an elastic
// pool.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. elasticPoolName is the
// name of the elastic pool.
func (client ElasticPoolsClient) ListDatabaseActivity(resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPoolDatabaseActivityListResult, err error) {
	req, err := client.ListDatabaseActivityPreparer(resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListDatabaseActivity", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListDatabaseActivitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListDatabaseActivity", resp, "Failure sending request")
		return
	}

	result, err = client.ListDatabaseActivityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListDatabaseActivity", resp, "Failure responding to request")
	}

	return
}

// ListDatabaseActivityPreparer prepares the ListDatabaseActivity request.
func (client ElasticPoolsClient) ListDatabaseActivityPreparer(resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/elasticPoolDatabaseActivity", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListDatabaseActivitySender sends the ListDatabaseActivity request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) ListDatabaseActivitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListDatabaseActivityResponder handles the response to the ListDatabaseActivity request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListDatabaseActivityResponder(resp *http.Response) (result ElasticPoolDatabaseActivityListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListDatabases returns a list of databases in an elastic pool.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. elasticPoolName is the
// name of the elastic pool to be retrieved.
func (client ElasticPoolsClient) ListDatabases(resourceGroupName string, serverName string, elasticPoolName string) (result DatabaseListResult, err error) {
	req, err := client.ListDatabasesPreparer(resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListDatabases", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListDatabasesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListDatabases", resp, "Failure sending request")
		return
	}

	result, err = client.ListDatabasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListDatabases", resp, "Failure responding to request")
	}

	return
}

// ListDatabasesPreparer prepares the ListDatabases request.
func (client ElasticPoolsClient) ListDatabasesPreparer(resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/databases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListDatabasesSender sends the ListDatabases request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) ListDatabasesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListDatabasesResponder handles the response to the ListDatabases request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListDatabasesResponder(resp *http.Response) (result DatabaseListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
