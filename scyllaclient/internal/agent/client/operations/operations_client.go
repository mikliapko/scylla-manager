// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/url"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CoreBwlimit sets the bandwidth limit

This sets the bandwidth limit to that passed in
*/
func (a *Client) CoreBwlimit(params *CoreBwlimitParams) (*CoreBwlimitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoreBwlimitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoreBwlimit",
		Method:             "POST",
		PathPattern:        "/rclone/core/bwlimit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoreBwlimitReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*CoreBwlimitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CoreBwlimit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CoreGroupList groups names

Returns list of group names
*/
func (a *Client) CoreGroupList(params *CoreGroupListParams) (*CoreGroupListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoreGroupListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoreGroupList",
		Method:             "POST",
		PathPattern:        "/rclone/core/group-list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoreGroupListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*CoreGroupListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CoreGroupList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CoreStats stats about transfers

Returns stats about current transfers
*/
func (a *Client) CoreStats(params *CoreStatsParams) (*CoreStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoreStatsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoreStats",
		Method:             "POST",
		PathPattern:        "/rclone/core/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoreStatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*CoreStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CoreStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CoreStatsReset resets all or specific stats group

Resets stats
*/
func (a *Client) CoreStatsReset(params *CoreStatsResetParams) (*CoreStatsResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoreStatsResetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoreStatsReset",
		Method:             "POST",
		PathPattern:        "/rclone/core/stats-reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoreStatsResetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*CoreStatsResetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CoreStatsReset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CoreTransferred completeds transfers

Returns stats about completed transfers
*/
func (a *Client) CoreTransferred(params *CoreTransferredParams) (*CoreTransferredOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoreTransferredParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoreTransferred",
		Method:             "POST",
		PathPattern:        "/rclone/core/transferred",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoreTransferredReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*CoreTransferredOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CoreTransferred: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
JobStatus jobs status

Reads the status of the job ID
*/
func (a *Client) JobStatus(params *JobStatusParams) (*JobStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JobStatus",
		Method:             "POST",
		PathPattern:        "/rclone/job/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &JobStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*JobStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for JobStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
JobStop stops async job

Stops job with provided ID
*/
func (a *Client) JobStop(params *JobStopParams) (*JobStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JobStop",
		Method:             "POST",
		PathPattern:        "/rclone/job/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &JobStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*JobStopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for JobStop: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
NodeInfo gets information about scylla node

Get information about Scylla node
*/
func (a *Client) NodeInfo(params *NodeInfoParams) (*NodeInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NodeInfo",
		Method:             "GET",
		PathPattern:        "/node_info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*NodeInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for NodeInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OperationsAbout abouts remote

Get usage information from the remote
*/
func (a *Client) OperationsAbout(params *OperationsAboutParams) (*OperationsAboutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsAboutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsAbout",
		Method:             "POST",
		PathPattern:        "/rclone/operations/about",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsAboutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*OperationsAboutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OperationsAbout: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OperationsCat cats remote

Concatenate any files and send them in response
*/
func (a *Client) OperationsCat(params *OperationsCatParams) (*OperationsCatOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsCatParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsCat",
		Method:             "POST",
		PathPattern:        "/rclone/operations/cat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsCatReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*OperationsCatOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OperationsCat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OperationsCopyfile copies a file

Copy a file from source remote to destination remote
*/
func (a *Client) OperationsCopyfile(params *OperationsCopyfileParams) (*OperationsCopyfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsCopyfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsCopyfile",
		Method:             "POST",
		PathPattern:        "/rclone/operations/copyfile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsCopyfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*OperationsCopyfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OperationsCopyfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OperationsDeletefile deletes file

Remove the single file pointed to
*/
func (a *Client) OperationsDeletefile(params *OperationsDeletefileParams) (*OperationsDeletefileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsDeletefileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsDeletefile",
		Method:             "POST",
		PathPattern:        "/rclone/operations/deletefile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsDeletefileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*OperationsDeletefileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OperationsDeletefile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OperationsList lists remote

List the given remote and path
*/
func (a *Client) OperationsList(params *OperationsListParams) (*OperationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsList",
		Method:             "POST",
		PathPattern:        "/rclone/operations/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*OperationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OperationsList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OperationsMovefile moves a file

Move a file from source remote to destination remote
*/
func (a *Client) OperationsMovefile(params *OperationsMovefileParams) (*OperationsMovefileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsMovefileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsMovefile",
		Method:             "POST",
		PathPattern:        "/rclone/operations/movefile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsMovefileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*OperationsMovefileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OperationsMovefile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OperationsPurge purges container

Remove a directory or container and all of its contents
*/
func (a *Client) OperationsPurge(params *OperationsPurgeParams) (*OperationsPurgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsPurgeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsPurge",
		Method:             "POST",
		PathPattern:        "/rclone/operations/purge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsPurgeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*OperationsPurgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OperationsPurge: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SyncCopy copies directory

Copy a directory from source remote to destination remote
*/
func (a *Client) SyncCopy(params *SyncCopyParams) (*SyncCopyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncCopyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SyncCopy",
		Method:             "POST",
		PathPattern:        "/rclone/sync/copy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SyncCopyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			err = e.Err
		}
		return nil, err
	}
	success, ok := result.(*SyncCopyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SyncCopy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
