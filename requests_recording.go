package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// StartStopRecordingRequest : Toggle recording on or off.
// Since obs-websocket version: 0.3.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstoprecording
type StartStopRecordingRequest _request

// NewStartStopRecordingRequest returns a new StartStopRecordingRequest.
func NewStartStopRecordingRequest() StartStopRecordingRequest {
	return StartStopRecordingRequest{MessageID: getMessageID(), RequestType: "StartStopRecording"}
}

// ID returns the request's message ID.
func (r StartStopRecordingRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r StartStopRecordingRequest) Type() string { return r.RequestType }

// Send sends the request and returns a channel to which the response will be sent.
func (r StartStopRecordingRequest) Send(c Client) (chan StartStopRecordingResponse, error) {
	generic, err := c.sendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan StartStopRecordingResponse)
	go func() { future <- (<-generic).(StartStopRecordingResponse) }()
	return future, nil
}

// StartStopRecordingResponse : Response for StartStopRecordingRequest.
// Since obs-websocket version: 0.3.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstoprecording
type StartStopRecordingResponse _response

// ID returns the response's message ID.
func (r StartStopRecordingResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r StartStopRecordingResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r StartStopRecordingResponse) Err() string { return r.Error }

// StartRecordingRequest : Start recording.
// Will return an `error` if recording is already active.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startrecording
type StartRecordingRequest _request

// NewStartRecordingRequest returns a new StartRecordingRequest.
func NewStartRecordingRequest() StartRecordingRequest {
	return StartRecordingRequest{MessageID: getMessageID(), RequestType: "StartRecording"}
}

// ID returns the request's message ID.
func (r StartRecordingRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r StartRecordingRequest) Type() string { return r.RequestType }

// Send sends the request and returns a channel to which the response will be sent.
func (r StartRecordingRequest) Send(c Client) (chan StartRecordingResponse, error) {
	generic, err := c.sendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan StartRecordingResponse)
	go func() { future <- (<-generic).(StartRecordingResponse) }()
	return future, nil
}

// StartRecordingResponse : Response for StartRecordingRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startrecording
type StartRecordingResponse _response

// ID returns the response's message ID.
func (r StartRecordingResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r StartRecordingResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r StartRecordingResponse) Err() string { return r.Error }

// StopRecordingRequest : Stop recording.
// Will return an `error` if recording is not active.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#stoprecording
type StopRecordingRequest _request

// NewStopRecordingRequest returns a new StopRecordingRequest.
func NewStopRecordingRequest() StopRecordingRequest {
	return StopRecordingRequest{MessageID: getMessageID(), RequestType: "StopRecording"}
}

// ID returns the request's message ID.
func (r StopRecordingRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r StopRecordingRequest) Type() string { return r.RequestType }

// Send sends the request and returns a channel to which the response will be sent.
func (r StopRecordingRequest) Send(c Client) (chan StopRecordingResponse, error) {
	generic, err := c.sendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan StopRecordingResponse)
	go func() { future <- (<-generic).(StopRecordingResponse) }()
	return future, nil
}

// StopRecordingResponse : Response for StopRecordingRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#stoprecording
type StopRecordingResponse _response

// ID returns the response's message ID.
func (r StopRecordingResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r StopRecordingResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r StopRecordingResponse) Err() string { return r.Error }

// SetRecordingFolderRequest : Change the current recording folder.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setrecordingfolder
type SetRecordingFolderRequest struct {
	// Path of the recording folder.
	// Required: Yes.
	RecFolder string `json:"rec-folder"`
	_request  `json:",squash"`
}

// NewSetRecordingFolderRequest returns a new SetRecordingFolderRequest.
func NewSetRecordingFolderRequest(recFolder string) SetRecordingFolderRequest {
	return SetRecordingFolderRequest{
		recFolder,
		_request{
			MessageID:   getMessageID(),
			RequestType: "SetRecordingFolder",
		},
	}

}

// ID returns the request's message ID.
func (r SetRecordingFolderRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetRecordingFolderRequest) Type() string { return r.RequestType }

// Send sends the request and returns a channel to which the response will be sent.
func (r SetRecordingFolderRequest) Send(c Client) (chan SetRecordingFolderResponse, error) {
	generic, err := c.sendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan SetRecordingFolderResponse)
	go func() { future <- (<-generic).(SetRecordingFolderResponse) }()
	return future, nil
}

// SetRecordingFolderResponse : Response for SetRecordingFolderRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setrecordingfolder
type SetRecordingFolderResponse _response

// ID returns the response's message ID.
func (r SetRecordingFolderResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetRecordingFolderResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetRecordingFolderResponse) Err() string { return r.Error }

// GetRecordingFolderRequest : Get the path of  the current recording folder.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getrecordingfolder
type GetRecordingFolderRequest _request

// NewGetRecordingFolderRequest returns a new GetRecordingFolderRequest.
func NewGetRecordingFolderRequest() GetRecordingFolderRequest {
	return GetRecordingFolderRequest{MessageID: getMessageID(), RequestType: "GetRecordingFolder"}
}

// ID returns the request's message ID.
func (r GetRecordingFolderRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r GetRecordingFolderRequest) Type() string { return r.RequestType }

// Send sends the request and returns a channel to which the response will be sent.
func (r GetRecordingFolderRequest) Send(c Client) (chan GetRecordingFolderResponse, error) {
	generic, err := c.sendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan GetRecordingFolderResponse)
	go func() { future <- (<-generic).(GetRecordingFolderResponse) }()
	return future, nil
}

// GetRecordingFolderResponse : Response for GetRecordingFolderRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getrecordingfolder
type GetRecordingFolderResponse struct {
	// Path of the recording folder.
	// Required: Yes.
	RecFolder string `json:"rec-folder"`
	_response `json:",squash"`
}

// ID returns the response's message ID.
func (r GetRecordingFolderResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r GetRecordingFolderResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r GetRecordingFolderResponse) Err() string { return r.Error }
