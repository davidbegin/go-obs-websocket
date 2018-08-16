package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// SetCurrentProfileRequest : Set the currently active profile.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentprofile
type SetCurrentProfileRequest struct {
	// Name of the desired profile.
	// Required: Yes.
	ProfileName string `json:"profile-name"`
	_request    `json:",squash"`
}

// NewSetCurrentProfileRequest returns a new SetCurrentProfileRequest.
func NewSetCurrentProfileRequest(profileName string) SetCurrentProfileRequest {
	return SetCurrentProfileRequest{
		profileName,
		_request{
			MessageID:   getMessageID(),
			RequestType: "SetCurrentProfile",
		},
	}

}

// ID returns the request's message ID.
func (r SetCurrentProfileRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetCurrentProfileRequest) Type() string { return r.RequestType }

// Send sends the request and returns a channel to which the response will be sent.
func (r SetCurrentProfileRequest) Send(c Client) (chan SetCurrentProfileResponse, error) {
	generic, err := c.sendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan SetCurrentProfileResponse)
	go func() { future <- (<-generic).(SetCurrentProfileResponse) }()
	return future, nil
}

// SetCurrentProfileResponse : Response for SetCurrentProfileRequest.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentprofile
type SetCurrentProfileResponse _response

// ID returns the response's message ID.
func (r SetCurrentProfileResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetCurrentProfileResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetCurrentProfileResponse) Err() string { return r.Error }

// GetCurrentProfileRequest : Get the name of the current profile.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentprofile
type GetCurrentProfileRequest _request

// NewGetCurrentProfileRequest returns a new GetCurrentProfileRequest.
func NewGetCurrentProfileRequest() GetCurrentProfileRequest {
	return GetCurrentProfileRequest{MessageID: getMessageID(), RequestType: "GetCurrentProfile"}
}

// ID returns the request's message ID.
func (r GetCurrentProfileRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r GetCurrentProfileRequest) Type() string { return r.RequestType }

// Send sends the request and returns a channel to which the response will be sent.
func (r GetCurrentProfileRequest) Send(c Client) (chan GetCurrentProfileResponse, error) {
	generic, err := c.sendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan GetCurrentProfileResponse)
	go func() { future <- (<-generic).(GetCurrentProfileResponse) }()
	return future, nil
}

// GetCurrentProfileResponse : Response for GetCurrentProfileRequest.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentprofile
type GetCurrentProfileResponse struct {
	// Name of the currently active profile.
	// Required: Yes.
	ProfileName string `json:"profile-name"`
	_response   `json:",squash"`
}

// ID returns the response's message ID.
func (r GetCurrentProfileResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r GetCurrentProfileResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r GetCurrentProfileResponse) Err() string { return r.Error }

// ListProfilesRequest : Get a list of available profiles.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#listprofiles
type ListProfilesRequest _request

// NewListProfilesRequest returns a new ListProfilesRequest.
func NewListProfilesRequest() ListProfilesRequest {
	return ListProfilesRequest{MessageID: getMessageID(), RequestType: "ListProfiles"}
}

// ID returns the request's message ID.
func (r ListProfilesRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r ListProfilesRequest) Type() string { return r.RequestType }

// Send sends the request and returns a channel to which the response will be sent.
func (r ListProfilesRequest) Send(c Client) (chan ListProfilesResponse, error) {
	generic, err := c.sendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan ListProfilesResponse)
	go func() { future <- (<-generic).(ListProfilesResponse) }()
	return future, nil
}

// ListProfilesResponse : Response for ListProfilesRequest.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#listprofiles
type ListProfilesResponse struct {
	// List of available profiles.
	// Required: Yes.
	Profiles  interface{} `json:"profiles"`
	_response `json:",squash"`
}

// ID returns the response's message ID.
func (r ListProfilesResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r ListProfilesResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r ListProfilesResponse) Err() string { return r.Error }
