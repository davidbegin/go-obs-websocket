package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// GetSceneItemListRequest : Get a list of all scene items in a scene.
//
// Since obs-websocket version: Unreleased.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getsceneitemlist
type GetSceneItemListRequest struct {
	// Name of the scene to get the list of scene items from.
	// Defaults to the current scene if not specified.
	// Required: No.
	SceneName string `json:"sceneName"`
	_request  `json:",squash"`
	response  chan GetSceneItemListResponse
}

// NewGetSceneItemListRequest returns a new GetSceneItemListRequest.
func NewGetSceneItemListRequest(sceneName string) GetSceneItemListRequest {
	return GetSceneItemListRequest{
		sceneName,
		_request{
			ID_:   GetMessageID(),
			Type_: "GetSceneItemList",
			err:   make(chan error, 1),
		},
		make(chan GetSceneItemListResponse, 1),
	}
}

// Send sends the request.
func (r *GetSceneItemListRequest) Send(c Client) error {
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
		var resp GetSceneItemListResponse
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
func (r GetSceneItemListRequest) Receive() (Response, error) {
	if !r.sent {
		return GetSceneItemListResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneItemListResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneItemListResponse{}, err
		case <-time.After(receiveTimeout):
			return GetSceneItemListResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetSceneItemListRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return GetSceneItemListResponse{}, err
	}
	return r.Receive()
}

// GetSceneItemListResponse : Response for GetSceneItemListRequest.
//
// Since obs-websocket version: Unreleased.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getsceneitemlist
type GetSceneItemListResponse struct {
	// Name of the requested (or current) scene.
	// Required: Yes.
	SceneName string `json:"sceneName"`
	// Array of scene items.
	// Required: Yes.
	SceneItems []map[string]interface{} `json:"sceneItems"`
	// Unique item id of the source item.
	// Required: Yes.
	SceneItemsItemID int `json:"sceneItems.*.itemId"`
	// ID if the scene item's source.
	// For example `vlc_source` or `image_source`.
	// Required: Yes.
	SceneItemsSourceKind string `json:"sceneItems.*.sourceKind"`
	// Name of the scene item's source.
	// Required: Yes.
	SceneItemsSourceName string `json:"sceneItems.*.sourceName"`
	// Type of the scene item's source.
	// Either `input`, `group`, or `scene`.
	// Required: Yes.
	SceneItemsSourceType string `json:"sceneItems.*.sourceType"`
	_response            `json:",squash"`
}

// GetSceneItemPropertiesRequest : Gets the scene specific properties of the specified source item.
// Coordinates are relative to the item's parent (the scene or group it belongs to).
//
// Since obs-websocket version: 4.3.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getsceneitemproperties
type GetSceneItemPropertiesRequest struct {
	// Name of the scene the scene item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// Scene Item name (if this field is a string) or specification (if it is an object).
	// Required: Yes.
	Item interface{} `json:"item"`
	// Scene Item name (if the `item` field is an object).
	// Required: No.
	ItemName string `json:"item.name"`
	// Scene Item ID (if the `item` field is an object).
	// Required: No.
	ItemID   int `json:"item.id"`
	_request `json:",squash"`
	response chan GetSceneItemPropertiesResponse
}

// NewGetSceneItemPropertiesRequest returns a new GetSceneItemPropertiesRequest.
func NewGetSceneItemPropertiesRequest(
	sceneName string,
	item interface{},
	itemName string,
	itemID int,
) GetSceneItemPropertiesRequest {
	return GetSceneItemPropertiesRequest{
		sceneName,
		item,
		itemName,
		itemID,
		_request{
			ID_:   GetMessageID(),
			Type_: "GetSceneItemProperties",
			err:   make(chan error, 1),
		},
		make(chan GetSceneItemPropertiesResponse, 1),
	}
}

// Send sends the request.
func (r *GetSceneItemPropertiesRequest) Send(c Client) error {
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
		var resp GetSceneItemPropertiesResponse
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
func (r GetSceneItemPropertiesRequest) Receive() (Response, error) {
	if !r.sent {
		return GetSceneItemPropertiesResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneItemPropertiesResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneItemPropertiesResponse{}, err
		case <-time.After(receiveTimeout):
			return GetSceneItemPropertiesResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetSceneItemPropertiesRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return GetSceneItemPropertiesResponse{}, err
	}
	return r.Receive()
}

// GetSceneItemPropertiesResponse : Response for GetSceneItemPropertiesRequest.
//
// Since obs-websocket version: 4.3.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getsceneitemproperties
type GetSceneItemPropertiesResponse struct {
	// Scene Item name.
	// Required: Yes.
	Name string `json:"name"`
	// Scene Item ID.
	// Required: Yes.
	ItemID int `json:"itemId"`
	// The x position of the source from the left.
	// Required: Yes.
	PositionX float64 `json:"position.x"`
	// The y position of the source from the top.
	// Required: Yes.
	PositionY float64 `json:"position.y"`
	// The point on the source that the item is manipulated from.
	// The sum of 1=Left or 2=Right, and 4=Top or 8=Bottom, or omit to center on that axis.
	// Required: Yes.
	PositionAlignment int `json:"position.alignment"`
	// The clockwise rotation of the item in degrees around the point of alignment.
	// Required: Yes.
	Rotation float64 `json:"rotation"`
	// The x-scale factor of the source.
	// Required: Yes.
	ScaleX float64 `json:"scale.x"`
	// The y-scale factor of the source.
	// Required: Yes.
	ScaleY float64 `json:"scale.y"`
	// The number of pixels cropped off the top of the source before scaling.
	// Required: Yes.
	CropTop int `json:"crop.top"`
	// The number of pixels cropped off the right of the source before scaling.
	// Required: Yes.
	CropRight int `json:"crop.right"`
	// The number of pixels cropped off the bottom of the source before scaling.
	// Required: Yes.
	CropBottom int `json:"crop.bottom"`
	// The number of pixels cropped off the left of the source before scaling.
	// Required: Yes.
	CropLeft int `json:"crop.left"`
	// If the source is visible.
	// Required: Yes.
	Visible bool `json:"visible"`
	// If the source is muted.
	// Required: Yes.
	Muted bool `json:"muted"`
	// If the source's transform is locked.
	// Required: Yes.
	Locked bool `json:"locked"`
	// Type of bounding box.
	// Can be "OBS_BOUNDS_STRETCH", "OBS_BOUNDS_SCALE_INNER", "OBS_BOUNDS_SCALE_OUTER", "OBS_BOUNDS_SCALE_TO_WIDTH", "OBS_BOUNDS_SCALE_TO_HEIGHT", "OBS_BOUNDS_MAX_ONLY" or "OBS_BOUNDS_NONE".
	// Required: Yes.
	BoundsType string `json:"bounds.type"`
	// Alignment of the bounding box.
	// Required: Yes.
	BoundsAlignment int `json:"bounds.alignment"`
	// Width of the bounding box.
	// Required: Yes.
	BoundsX float64 `json:"bounds.x"`
	// Height of the bounding box.
	// Required: Yes.
	BoundsY float64 `json:"bounds.y"`
	// Base width (without scaling) of the source.
	// Required: Yes.
	SourceWidth int `json:"sourceWidth"`
	// Base source (without scaling) of the source.
	// Required: Yes.
	SourceHeight int `json:"sourceHeight"`
	// Scene item width (base source width multiplied by the horizontal scaling factor).
	// Required: Yes.
	Width float64 `json:"width"`
	// Scene item height (base source height multiplied by the vertical scaling factor).
	// Required: Yes.
	Height float64 `json:"height"`
	// Name of the item's parent (if this item belongs to a group).
	// Required: No.
	ParentGroupName string `json:"parentGroupName"`
	// List of children (if this item is a group).
	// Required: No.
	GroupChildren []*SceneItemTransform `json:"groupChildren"`
	_response     `json:",squash"`
}

// SetSceneItemPropertiesRequest : Sets the scene specific properties of a source
// Unspecified properties will remain unchanged.
// Coordinates are relative to the item's parent (the scene or group it belongs to).
//
// Since obs-websocket version: 4.3.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemproperties
type SetSceneItemPropertiesRequest struct {
	// Name of the scene the source item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// Scene Item name (if this field is a string) or specification (if it is an object).
	// Required: Yes.
	Item interface{} `json:"item"`
	// Scene Item name (if the `item` field is an object).
	// Required: No.
	ItemName string `json:"item.name"`
	// Scene Item ID (if the `item` field is an object).
	// Required: No.
	ItemID int `json:"item.id"`
	// The new x position of the source.
	// Required: No.
	PositionX float64 `json:"position.x"`
	// The new y position of the source.
	// Required: No.
	PositionY float64 `json:"position.y"`
	// The new alignment of the source.
	// Required: No.
	PositionAlignment int `json:"position.alignment"`
	// The new clockwise rotation of the item in degrees.
	// Required: No.
	Rotation float64 `json:"rotation"`
	// The new x scale of the item.
	// Required: No.
	ScaleX float64 `json:"scale.x"`
	// The new y scale of the item.
	// Required: No.
	ScaleY float64 `json:"scale.y"`
	// The new amount of pixels cropped off the top of the source before scaling.
	// Required: No.
	CropTop int `json:"crop.top"`
	// The new amount of pixels cropped off the bottom of the source before scaling.
	// Required: No.
	CropBottom int `json:"crop.bottom"`
	// The new amount of pixels cropped off the left of the source before scaling.
	// Required: No.
	CropLeft int `json:"crop.left"`
	// The new amount of pixels cropped off the right of the source before scaling.
	// Required: No.
	CropRight int `json:"crop.right"`
	// The new visibility of the source.
	// 'true' shows source, 'false' hides source.
	// Required: No.
	Visible bool `json:"visible"`
	// The new locked status of the source.
	// 'true' keeps it in its current position, 'false' allows movement.
	// Required: No.
	Locked bool `json:"locked"`
	// The new bounds type of the source.
	// Can be "OBS_BOUNDS_STRETCH", "OBS_BOUNDS_SCALE_INNER", "OBS_BOUNDS_SCALE_OUTER", "OBS_BOUNDS_SCALE_TO_WIDTH", "OBS_BOUNDS_SCALE_TO_HEIGHT", "OBS_BOUNDS_MAX_ONLY" or "OBS_BOUNDS_NONE".
	// Required: No.
	BoundsType string `json:"bounds.type"`
	// The new alignment of the bounding box.
	// (0-2, 4-6, 8-10).
	// Required: No.
	BoundsAlignment int `json:"bounds.alignment"`
	// The new width of the bounding box.
	// Required: No.
	BoundsX float64 `json:"bounds.x"`
	// The new height of the bounding box.
	// Required: No.
	BoundsY  float64 `json:"bounds.y"`
	_request `json:",squash"`
	response chan SetSceneItemPropertiesResponse
}

// NewSetSceneItemPropertiesRequest returns a new SetSceneItemPropertiesRequest.
func NewSetSceneItemPropertiesRequest(
	sceneName string,
	item interface{},
	itemName string,
	itemID int,
	positionX float64,
	positionY float64,
	positionAlignment int,
	rotation float64,
	scaleX float64,
	scaleY float64,
	cropTop int,
	cropBottom int,
	cropLeft int,
	cropRight int,
	visible bool,
	locked bool,
	boundsType string,
	boundsAlignment int,
	boundsX float64,
	boundsY float64,
) SetSceneItemPropertiesRequest {
	return SetSceneItemPropertiesRequest{
		sceneName,
		item,
		itemName,
		itemID,
		positionX,
		positionY,
		positionAlignment,
		rotation,
		scaleX,
		scaleY,
		cropTop,
		cropBottom,
		cropLeft,
		cropRight,
		visible,
		locked,
		boundsType,
		boundsAlignment,
		boundsX,
		boundsY,
		_request{
			ID_:   GetMessageID(),
			Type_: "SetSceneItemProperties",
			err:   make(chan error, 1),
		},
		make(chan SetSceneItemPropertiesResponse, 1),
	}
}

// Send sends the request.
func (r *SetSceneItemPropertiesRequest) Send(c Client) error {
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
		var resp SetSceneItemPropertiesResponse
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
func (r SetSceneItemPropertiesRequest) Receive() (Response, error) {
	if !r.sent {
		return SetSceneItemPropertiesResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemPropertiesResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemPropertiesResponse{}, err
		case <-time.After(receiveTimeout):
			return SetSceneItemPropertiesResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetSceneItemPropertiesRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return SetSceneItemPropertiesResponse{}, err
	}
	return r.Receive()
}

// SetSceneItemPropertiesResponse : Response for SetSceneItemPropertiesRequest.
//
// Since obs-websocket version: 4.3.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemproperties
type SetSceneItemPropertiesResponse struct {
	_response `json:",squash"`
}

// ResetSceneItemRequest : Reset a scene item.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#resetsceneitem
type ResetSceneItemRequest struct {
	// Name of the scene the scene item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// Scene Item name (if this field is a string) or specification (if it is an object).
	// Required: Yes.
	Item interface{} `json:"item"`
	// Scene Item name (if the `item` field is an object).
	// Required: No.
	ItemName string `json:"item.name"`
	// Scene Item ID (if the `item` field is an object).
	// Required: No.
	ItemID   int `json:"item.id"`
	_request `json:",squash"`
	response chan ResetSceneItemResponse
}

// NewResetSceneItemRequest returns a new ResetSceneItemRequest.
func NewResetSceneItemRequest(
	sceneName string,
	item interface{},
	itemName string,
	itemID int,
) ResetSceneItemRequest {
	return ResetSceneItemRequest{
		sceneName,
		item,
		itemName,
		itemID,
		_request{
			ID_:   GetMessageID(),
			Type_: "ResetSceneItem",
			err:   make(chan error, 1),
		},
		make(chan ResetSceneItemResponse, 1),
	}
}

// Send sends the request.
func (r *ResetSceneItemRequest) Send(c Client) error {
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
		var resp ResetSceneItemResponse
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
func (r ResetSceneItemRequest) Receive() (Response, error) {
	if !r.sent {
		return ResetSceneItemResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ResetSceneItemResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ResetSceneItemResponse{}, err
		case <-time.After(receiveTimeout):
			return ResetSceneItemResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r ResetSceneItemRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return ResetSceneItemResponse{}, err
	}
	return r.Receive()
}

// ResetSceneItemResponse : Response for ResetSceneItemRequest.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#resetsceneitem
type ResetSceneItemResponse struct {
	_response `json:",squash"`
}

// SetSceneItemRenderRequest : Show or hide a specified source item in a specified scene.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemrender
type SetSceneItemRenderRequest struct {
	// Name of the scene the scene item belongs to.
	// Defaults to the currently active scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// Scene Item name.
	// Required: Yes.
	Source string `json:"source"`
	// true = shown ; false = hidden.
	// Required: Yes.
	Render   bool `json:"render"`
	_request `json:",squash"`
	response chan SetSceneItemRenderResponse
}

// NewSetSceneItemRenderRequest returns a new SetSceneItemRenderRequest.
func NewSetSceneItemRenderRequest(
	sceneName string,
	source string,
	render bool,
) SetSceneItemRenderRequest {
	return SetSceneItemRenderRequest{
		sceneName,
		source,
		render,
		_request{
			ID_:   GetMessageID(),
			Type_: "SetSceneItemRender",
			err:   make(chan error, 1),
		},
		make(chan SetSceneItemRenderResponse, 1),
	}
}

// Send sends the request.
func (r *SetSceneItemRenderRequest) Send(c Client) error {
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
		var resp SetSceneItemRenderResponse
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
func (r SetSceneItemRenderRequest) Receive() (Response, error) {
	if !r.sent {
		return SetSceneItemRenderResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemRenderResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemRenderResponse{}, err
		case <-time.After(receiveTimeout):
			return SetSceneItemRenderResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetSceneItemRenderRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return SetSceneItemRenderResponse{}, err
	}
	return r.Receive()
}

// SetSceneItemRenderResponse : Response for SetSceneItemRenderRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemrender
type SetSceneItemRenderResponse struct {
	_response `json:",squash"`
}

// SetSceneItemPositionRequest : Sets the coordinates of a specified source item.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemposition
type SetSceneItemPositionRequest struct {
	// Name of the scene the scene item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// Scene Item name.
	// Required: Yes.
	Item string `json:"item"`
	// X coordinate.
	// Required: Yes.
	X float64 `json:"x"`
	// Y coordinate.
	// Required: Yes.
	Y        float64 `json:"y"`
	_request `json:",squash"`
	response chan SetSceneItemPositionResponse
}

// NewSetSceneItemPositionRequest returns a new SetSceneItemPositionRequest.
func NewSetSceneItemPositionRequest(
	sceneName string,
	item string,
	x float64,
	y float64,
) SetSceneItemPositionRequest {
	return SetSceneItemPositionRequest{
		sceneName,
		item,
		x,
		y,
		_request{
			ID_:   GetMessageID(),
			Type_: "SetSceneItemPosition",
			err:   make(chan error, 1),
		},
		make(chan SetSceneItemPositionResponse, 1),
	}
}

// Send sends the request.
func (r *SetSceneItemPositionRequest) Send(c Client) error {
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
		var resp SetSceneItemPositionResponse
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
func (r SetSceneItemPositionRequest) Receive() (Response, error) {
	if !r.sent {
		return SetSceneItemPositionResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemPositionResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemPositionResponse{}, err
		case <-time.After(receiveTimeout):
			return SetSceneItemPositionResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetSceneItemPositionRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return SetSceneItemPositionResponse{}, err
	}
	return r.Receive()
}

// SetSceneItemPositionResponse : Response for SetSceneItemPositionRequest.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemposition
type SetSceneItemPositionResponse struct {
	_response `json:",squash"`
}

// SetSceneItemTransformRequest : Set the transform of the specified source item.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemtransform
type SetSceneItemTransformRequest struct {
	// Name of the scene the scene item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// Scene Item name.
	// Required: Yes.
	Item string `json:"item"`
	// Width scale factor.
	// Required: Yes.
	XScale float64 `json:"x-scale"`
	// Height scale factor.
	// Required: Yes.
	YScale float64 `json:"y-scale"`
	// Source item rotation (in degrees).
	// Required: Yes.
	Rotation float64 `json:"rotation"`
	_request `json:",squash"`
	response chan SetSceneItemTransformResponse
}

// NewSetSceneItemTransformRequest returns a new SetSceneItemTransformRequest.
func NewSetSceneItemTransformRequest(
	sceneName string,
	item string,
	xScale float64,
	yScale float64,
	rotation float64,
) SetSceneItemTransformRequest {
	return SetSceneItemTransformRequest{
		sceneName,
		item,
		xScale,
		yScale,
		rotation,
		_request{
			ID_:   GetMessageID(),
			Type_: "SetSceneItemTransform",
			err:   make(chan error, 1),
		},
		make(chan SetSceneItemTransformResponse, 1),
	}
}

// Send sends the request.
func (r *SetSceneItemTransformRequest) Send(c Client) error {
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
		var resp SetSceneItemTransformResponse
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
func (r SetSceneItemTransformRequest) Receive() (Response, error) {
	if !r.sent {
		return SetSceneItemTransformResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemTransformResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemTransformResponse{}, err
		case <-time.After(receiveTimeout):
			return SetSceneItemTransformResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetSceneItemTransformRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return SetSceneItemTransformResponse{}, err
	}
	return r.Receive()
}

// SetSceneItemTransformResponse : Response for SetSceneItemTransformRequest.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemtransform
type SetSceneItemTransformResponse struct {
	_response `json:",squash"`
}

// SetSceneItemCropRequest : Sets the crop coordinates of the specified source item.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemcrop
type SetSceneItemCropRequest struct {
	// Name of the scene the scene item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// Scene Item name.
	// Required: Yes.
	Item string `json:"item"`
	// Pixel position of the top of the source item.
	// Required: Yes.
	Top int `json:"top"`
	// Pixel position of the bottom of the source item.
	// Required: Yes.
	Bottom int `json:"bottom"`
	// Pixel position of the left of the source item.
	// Required: Yes.
	Left int `json:"left"`
	// Pixel position of the right of the source item.
	// Required: Yes.
	Right    int `json:"right"`
	_request `json:",squash"`
	response chan SetSceneItemCropResponse
}

// NewSetSceneItemCropRequest returns a new SetSceneItemCropRequest.
func NewSetSceneItemCropRequest(
	sceneName string,
	item string,
	top int,
	bottom int,
	left int,
	right int,
) SetSceneItemCropRequest {
	return SetSceneItemCropRequest{
		sceneName,
		item,
		top,
		bottom,
		left,
		right,
		_request{
			ID_:   GetMessageID(),
			Type_: "SetSceneItemCrop",
			err:   make(chan error, 1),
		},
		make(chan SetSceneItemCropResponse, 1),
	}
}

// Send sends the request.
func (r *SetSceneItemCropRequest) Send(c Client) error {
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
		var resp SetSceneItemCropResponse
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
func (r SetSceneItemCropRequest) Receive() (Response, error) {
	if !r.sent {
		return SetSceneItemCropResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemCropResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneItemCropResponse{}, err
		case <-time.After(receiveTimeout):
			return SetSceneItemCropResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetSceneItemCropRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return SetSceneItemCropResponse{}, err
	}
	return r.Receive()
}

// SetSceneItemCropResponse : Response for SetSceneItemCropRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setsceneitemcrop
type SetSceneItemCropResponse struct {
	_response `json:",squash"`
}

// DeleteSceneItemRequest : Deletes a scene item.
//
// Since obs-websocket version: 4.5.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#deletesceneitem
type DeleteSceneItemRequest struct {
	// Name of the scene the scene item belongs to.
	// Defaults to the current scene.
	// Required: No.
	Scene string `json:"scene"`
	// Scene item to delete (required).
	// Required: Yes.
	Item map[string]interface{} `json:"item"`
	// Scene Item name (prefer `id`, including both is acceptable).
	// Required: Yes.
	ItemName string `json:"item.name"`
	// Scene Item ID.
	// Required: Yes.
	ItemID   int `json:"item.id"`
	_request `json:",squash"`
	response chan DeleteSceneItemResponse
}

// NewDeleteSceneItemRequest returns a new DeleteSceneItemRequest.
func NewDeleteSceneItemRequest(
	scene string,
	item map[string]interface{},
	itemName string,
	itemID int,
) DeleteSceneItemRequest {
	return DeleteSceneItemRequest{
		scene,
		item,
		itemName,
		itemID,
		_request{
			ID_:   GetMessageID(),
			Type_: "DeleteSceneItem",
			err:   make(chan error, 1),
		},
		make(chan DeleteSceneItemResponse, 1),
	}
}

// Send sends the request.
func (r *DeleteSceneItemRequest) Send(c Client) error {
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
		var resp DeleteSceneItemResponse
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
func (r DeleteSceneItemRequest) Receive() (Response, error) {
	if !r.sent {
		return DeleteSceneItemResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return DeleteSceneItemResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return DeleteSceneItemResponse{}, err
		case <-time.After(receiveTimeout):
			return DeleteSceneItemResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r DeleteSceneItemRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return DeleteSceneItemResponse{}, err
	}
	return r.Receive()
}

// DeleteSceneItemResponse : Response for DeleteSceneItemRequest.
//
// Since obs-websocket version: 4.5.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#deletesceneitem
type DeleteSceneItemResponse struct {
	_response `json:",squash"`
}

// AddSceneItemRequest : Creates a scene item in a scene
// In other words, this is how you add a source into a scene.
//
// Since obs-websocket version: Unreleased.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#addsceneitem
type AddSceneItemRequest struct {
	// Name of the scene to create the scene item in.
	// Required: Yes.
	SceneName string `json:"sceneName"`
	// Name of the source to be added.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// Whether to make the sceneitem visible on creation or not.
	// Default `true`.
	// Required: Yes.
	SetVisible bool `json:"setVisible"`
	_request   `json:",squash"`
	response   chan AddSceneItemResponse
}

// NewAddSceneItemRequest returns a new AddSceneItemRequest.
func NewAddSceneItemRequest(
	sceneName string,
	sourceName string,
	setVisible bool,
) AddSceneItemRequest {
	return AddSceneItemRequest{
		sceneName,
		sourceName,
		setVisible,
		_request{
			ID_:   GetMessageID(),
			Type_: "AddSceneItem",
			err:   make(chan error, 1),
		},
		make(chan AddSceneItemResponse, 1),
	}
}

// Send sends the request.
func (r *AddSceneItemRequest) Send(c Client) error {
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
		var resp AddSceneItemResponse
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
func (r AddSceneItemRequest) Receive() (Response, error) {
	if !r.sent {
		return AddSceneItemResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return AddSceneItemResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return AddSceneItemResponse{}, err
		case <-time.After(receiveTimeout):
			return AddSceneItemResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r AddSceneItemRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return AddSceneItemResponse{}, err
	}
	return r.Receive()
}

// AddSceneItemResponse : Response for AddSceneItemRequest.
//
// Since obs-websocket version: Unreleased.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#addsceneitem
type AddSceneItemResponse struct {
	// Numerical ID of the created scene item.
	// Required: Yes.
	ItemID    int `json:"itemId"`
	_response `json:",squash"`
}

// DuplicateSceneItemRequest : Duplicates a scene item.
//
// Since obs-websocket version: 4.5.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#duplicatesceneitem
type DuplicateSceneItemRequest struct {
	// Name of the scene to copy the item from.
	// Defaults to the current scene.
	// Required: No.
	FromScene string `json:"fromScene"`
	// Name of the scene to create the item in.
	// Defaults to the current scene.
	// Required: No.
	ToScene string `json:"toScene"`
	// Scene Item to duplicate from the source scene (required).
	// Required: Yes.
	Item map[string]interface{} `json:"item"`
	// Scene Item name (prefer `id`, including both is acceptable).
	// Required: Yes.
	ItemName string `json:"item.name"`
	// Scene Item ID.
	// Required: Yes.
	ItemID   int `json:"item.id"`
	_request `json:",squash"`
	response chan DuplicateSceneItemResponse
}

// NewDuplicateSceneItemRequest returns a new DuplicateSceneItemRequest.
func NewDuplicateSceneItemRequest(
	fromScene string,
	toScene string,
	item map[string]interface{},
	itemName string,
	itemID int,
) DuplicateSceneItemRequest {
	return DuplicateSceneItemRequest{
		fromScene,
		toScene,
		item,
		itemName,
		itemID,
		_request{
			ID_:   GetMessageID(),
			Type_: "DuplicateSceneItem",
			err:   make(chan error, 1),
		},
		make(chan DuplicateSceneItemResponse, 1),
	}
}

// Send sends the request.
func (r *DuplicateSceneItemRequest) Send(c Client) error {
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
		var resp DuplicateSceneItemResponse
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
func (r DuplicateSceneItemRequest) Receive() (Response, error) {
	if !r.sent {
		return DuplicateSceneItemResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return DuplicateSceneItemResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return DuplicateSceneItemResponse{}, err
		case <-time.After(receiveTimeout):
			return DuplicateSceneItemResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r DuplicateSceneItemRequest) SendReceive(c Client) (Response, error) {
	if err := r.Send(c); err != nil {
		return DuplicateSceneItemResponse{}, err
	}
	return r.Receive()
}

// DuplicateSceneItemResponse : Response for DuplicateSceneItemRequest.
//
// Since obs-websocket version: 4.5.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#duplicatesceneitem
type DuplicateSceneItemResponse struct {
	// Name of the scene where the new item was created.
	// Required: Yes.
	Scene string `json:"scene"`
	// New item info.
	// Required: Yes.
	Item map[string]interface{} `json:"item"`
	// New item ID.
	// Required: Yes.
	ItemID int `json:"item.id"`
	// New item name.
	// Required: Yes.
	ItemName  string `json:"item.name"`
	_response `json:",squash"`
}
