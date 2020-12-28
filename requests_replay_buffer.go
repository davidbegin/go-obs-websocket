package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// GetReplayBufferStatusRequest : Get the status of the OBS replay buffer.
//
// Since obs-websocket version: Unreleased.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getreplaybufferstatus
type GetReplayBufferStatusRequest struct {
	_request `json:",squash"`
	response chan GetReplayBufferStatusResponse
}

// NewGetReplayBufferStatusRequest returns a new GetReplayBufferStatusRequest.
func NewGetReplayBufferStatusRequest() GetReplayBufferStatusRequest {
	return GetReplayBufferStatusRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "GetReplayBufferStatus",
			err:   make(chan error, 1),
		},
		make(chan GetReplayBufferStatusResponse, 1),
	}
}

// Send sends the request.
func (r *GetReplayBufferStatusRequest) Send(c Client) error {
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
		var resp GetReplayBufferStatusResponse
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
func (r GetReplayBufferStatusRequest) Receive() (Response, error) {
	if !r.sent {
		return GetReplayBufferStatusResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetReplayBufferStatusResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetReplayBufferStatusResponse{}, err
		case <-time.After(receiveTimeout):
			return GetReplayBufferStatusResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetReplayBufferStatusRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return GetReplayBufferStatusResponse{}, err
	}
	return r.Receive()
}

// GetReplayBufferStatusResponse : Response for GetReplayBufferStatusRequest.
//
// Since obs-websocket version: Unreleased.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getreplaybufferstatus
type GetReplayBufferStatusResponse struct {
	// Current recording status.
	// Required: Yes.
	IsReplayBufferActive bool `json:"isReplayBufferActive"`
	_response            `json:",squash"`
}

// StartStopReplayBufferRequest : Toggle the Replay Buffer on/off (depending on the current state of the replay buffer).
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startstopreplaybuffer
type StartStopReplayBufferRequest struct {
	_request `json:",squash"`
	response chan StartStopReplayBufferResponse
}

// NewStartStopReplayBufferRequest returns a new StartStopReplayBufferRequest.
func NewStartStopReplayBufferRequest() StartStopReplayBufferRequest {
	return StartStopReplayBufferRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "StartStopReplayBuffer",
			err:   make(chan error, 1),
		},
		make(chan StartStopReplayBufferResponse, 1),
	}
}

// Send sends the request.
func (r *StartStopReplayBufferRequest) Send(c Client) error {
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
		var resp StartStopReplayBufferResponse
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
func (r StartStopReplayBufferRequest) Receive() (Response, error) {
	if !r.sent {
		return StartStopReplayBufferResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartStopReplayBufferResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartStopReplayBufferResponse{}, err
		case <-time.After(receiveTimeout):
			return StartStopReplayBufferResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StartStopReplayBufferRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return StartStopReplayBufferResponse{}, err
	}
	return r.Receive()
}

// StartStopReplayBufferResponse : Response for StartStopReplayBufferRequest.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startstopreplaybuffer
type StartStopReplayBufferResponse struct {
	_response `json:",squash"`
}

// StartReplayBufferRequest : Start recording into the Replay Buffer.
// Will return an `error` if the Replay Buffer is already active or if the
// "Save Replay Buffer" hotkey is not set in OBS' settings.
// Setting this hotkey is mandatory, even when triggering saves only
// through obs-websocket.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startreplaybuffer
type StartReplayBufferRequest struct {
	_request `json:",squash"`
	response chan StartReplayBufferResponse
}

// NewStartReplayBufferRequest returns a new StartReplayBufferRequest.
func NewStartReplayBufferRequest() StartReplayBufferRequest {
	return StartReplayBufferRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "StartReplayBuffer",
			err:   make(chan error, 1),
		},
		make(chan StartReplayBufferResponse, 1),
	}
}

// Send sends the request.
func (r *StartReplayBufferRequest) Send(c Client) error {
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
		var resp StartReplayBufferResponse
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
func (r StartReplayBufferRequest) Receive() (Response, error) {
	if !r.sent {
		return StartReplayBufferResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartReplayBufferResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartReplayBufferResponse{}, err
		case <-time.After(receiveTimeout):
			return StartReplayBufferResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StartReplayBufferRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return StartReplayBufferResponse{}, err
	}
	return r.Receive()
}

// StartReplayBufferResponse : Response for StartReplayBufferRequest.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startreplaybuffer
type StartReplayBufferResponse struct {
	_response `json:",squash"`
}

// StopReplayBufferRequest : Stop recording into the Replay Buffer.
// Will return an `error` if the Replay Buffer is not active.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stopreplaybuffer
type StopReplayBufferRequest struct {
	_request `json:",squash"`
	response chan StopReplayBufferResponse
}

// NewStopReplayBufferRequest returns a new StopReplayBufferRequest.
func NewStopReplayBufferRequest() StopReplayBufferRequest {
	return StopReplayBufferRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "StopReplayBuffer",
			err:   make(chan error, 1),
		},
		make(chan StopReplayBufferResponse, 1),
	}
}

// Send sends the request.
func (r *StopReplayBufferRequest) Send(c Client) error {
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
		var resp StopReplayBufferResponse
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
func (r StopReplayBufferRequest) Receive() (Response, error) {
	if !r.sent {
		return StopReplayBufferResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopReplayBufferResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopReplayBufferResponse{}, err
		case <-time.After(receiveTimeout):
			return StopReplayBufferResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StopReplayBufferRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return StopReplayBufferResponse{}, err
	}
	return r.Receive()
}

// StopReplayBufferResponse : Response for StopReplayBufferRequest.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stopreplaybuffer
type StopReplayBufferResponse struct {
	_response `json:",squash"`
}

// SaveReplayBufferRequest : Flush and save the contents of the Replay Buffer to disk
// This is
// basically the same as triggering the "Save Replay Buffer" hotkey.
// Will return an `error` if the Replay Buffer is not active.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#savereplaybuffer
type SaveReplayBufferRequest struct {
	_request `json:",squash"`
	response chan SaveReplayBufferResponse
}

// NewSaveReplayBufferRequest returns a new SaveReplayBufferRequest.
func NewSaveReplayBufferRequest() SaveReplayBufferRequest {
	return SaveReplayBufferRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "SaveReplayBuffer",
			err:   make(chan error, 1),
		},
		make(chan SaveReplayBufferResponse, 1),
	}
}

// Send sends the request.
func (r *SaveReplayBufferRequest) Send(c Client) error {
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
		var resp SaveReplayBufferResponse
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
func (r SaveReplayBufferRequest) Receive() (Response, error) {
	if !r.sent {
		return SaveReplayBufferResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SaveReplayBufferResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SaveReplayBufferResponse{}, err
		case <-time.After(receiveTimeout):
			return SaveReplayBufferResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SaveReplayBufferRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return SaveReplayBufferResponse{}, err
	}
	return r.Receive()
}

// SaveReplayBufferResponse : Response for SaveReplayBufferRequest.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#savereplaybuffer
type SaveReplayBufferResponse struct {
	_response `json:",squash"`
}
