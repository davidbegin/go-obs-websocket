package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// SetCurrentSceneCollectionRequest : Change the active scene collection.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setcurrentscenecollection
type SetCurrentSceneCollectionRequest struct {
	// Name of the desired scene collection.
	// Required: Yes.
	ScName   string `json:"sc-name"`
	_request `json:",squash"`
	response chan SetCurrentSceneCollectionResponse
}

// NewSetCurrentSceneCollectionRequest returns a new SetCurrentSceneCollectionRequest.
func NewSetCurrentSceneCollectionRequest(scName string) SetCurrentSceneCollectionRequest {
	return SetCurrentSceneCollectionRequest{
		scName,
		_request{
			ID_:   GetMessageID(),
			Type_: "SetCurrentSceneCollection",
			err:   make(chan error, 1),
		},
		make(chan SetCurrentSceneCollectionResponse, 1),
	}
}

// Send sends the request.
func (r *SetCurrentSceneCollectionRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp SetCurrentSceneCollectionResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r SetCurrentSceneCollectionRequest) Receive() (Response, error) {
	if !r.sent {
		return SetCurrentSceneCollectionResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetCurrentSceneCollectionResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetCurrentSceneCollectionResponse{}, err
		case <-time.After(receiveTimeout):
			return SetCurrentSceneCollectionResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetCurrentSceneCollectionRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return SetCurrentSceneCollectionResponse{}, err
	}
	return r.Receive()
}

// SetCurrentSceneCollectionResponse : Response for SetCurrentSceneCollectionRequest.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setcurrentscenecollection
type SetCurrentSceneCollectionResponse struct {
	_response `json:",squash"`
}

// GetCurrentSceneCollectionRequest : Get the name of the current scene collection.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getcurrentscenecollection
type GetCurrentSceneCollectionRequest struct {
	_request `json:",squash"`
	response chan GetCurrentSceneCollectionResponse
}

// NewGetCurrentSceneCollectionRequest returns a new GetCurrentSceneCollectionRequest.
func NewGetCurrentSceneCollectionRequest() GetCurrentSceneCollectionRequest {
	return GetCurrentSceneCollectionRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "GetCurrentSceneCollection",
			err:   make(chan error, 1),
		},
		make(chan GetCurrentSceneCollectionResponse, 1),
	}
}

// Send sends the request.
func (r *GetCurrentSceneCollectionRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetCurrentSceneCollectionResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetCurrentSceneCollectionRequest) Receive() (Response, error) {
	if !r.sent {
		return GetCurrentSceneCollectionResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetCurrentSceneCollectionResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetCurrentSceneCollectionResponse{}, err
		case <-time.After(receiveTimeout):
			return GetCurrentSceneCollectionResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetCurrentSceneCollectionRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return GetCurrentSceneCollectionResponse{}, err
	}
	return r.Receive()
}

// GetCurrentSceneCollectionResponse : Response for GetCurrentSceneCollectionRequest.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getcurrentscenecollection
type GetCurrentSceneCollectionResponse struct {
	// Name of the currently active scene collection.
	// Required: Yes.
	ScName    string `json:"sc-name"`
	_response `json:",squash"`
}

// ListSceneCollectionsRequest : List available scene collections.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#listscenecollections
type ListSceneCollectionsRequest struct {
	_request `json:",squash"`
	response chan ListSceneCollectionsResponse
}

// NewListSceneCollectionsRequest returns a new ListSceneCollectionsRequest.
func NewListSceneCollectionsRequest() ListSceneCollectionsRequest {
	return ListSceneCollectionsRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "ListSceneCollections",
			err:   make(chan error, 1),
		},
		make(chan ListSceneCollectionsResponse, 1),
	}
}

// Send sends the request.
func (r *ListSceneCollectionsRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp ListSceneCollectionsResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r ListSceneCollectionsRequest) Receive() (Response, error) {
	if !r.sent {
		return ListSceneCollectionsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ListSceneCollectionsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ListSceneCollectionsResponse{}, err
		case <-time.After(receiveTimeout):
			return ListSceneCollectionsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r ListSceneCollectionsRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return ListSceneCollectionsResponse{}, err
	}
	return r.Receive()
}

// ListSceneCollectionsResponse : Response for ListSceneCollectionsRequest.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#listscenecollections
type ListSceneCollectionsResponse struct {
	// Scene collections list.
	// Required: Yes.
	SceneCollections []string `json:"scene-collections"`
	// Scene collection name.
	// Required: Yes.
	SceneCollectionsScName string `json:"scene-collections.*.sc-name"`
	_response              `json:",squash"`
}
