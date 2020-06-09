package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/roryj/das-keyboard/colour"
	"github.com/roryj/das-keyboard/keyboard"
)

const localHost = "http://localhost"
const apiVersion = "1.0"

type keyboardClient struct {
	port int
}

func NewKeyboardClient(port int) keyboardClient {
	return keyboardClient{port: port}
}

func (c *keyboardClient) CreateSignal(zone keyboard.Zone, effect keyboard.KeyEffect, colour colour.ToHexer) (CreateSignalResponse, error) {

	req := CreateSignalRequest{
		Colour:  colour.ToHex(),
		Effect:  effect,
		Message: "sending a message",
		Name:    "",
		Pid:     "DK5QPID",
		ZoneId:  zone.GetZoneName(),
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return CreateSignalResponse{}, err
	}

	result, err := http.Post(c.generateUrl("signals"), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return CreateSignalResponse{}, err
	}

	r, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return CreateSignalResponse{}, err
	}

	var response CreateSignalResponse
	err = json.Unmarshal(r, &response)
	if err != nil {
		return CreateSignalResponse{}, err
	}

	return response, nil
}

func (c *keyboardClient) DeleteSignal(id int) error {
	u := c.generateUrl("signals", url.PathEscape(strconv.Itoa(id)))
	fmt.Printf("url: %s\n", u)

	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("body: %s", string(body))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received an invalid status code. Expected 200, found %d", resp.StatusCode)
	}

	return nil
}

func (c *keyboardClient) generateUrl(requestType string, pathArgs ...string) string {
	url := fmt.Sprintf("%s:%d/api/%s/%s", localHost, c.port, apiVersion, requestType)

	if len(pathArgs) != 0 {
		url = fmt.Sprintf("%s/%s", url, strings.Join(pathArgs, "/"))
	}

	return url
}

type CreateSignalRequest struct {
	Name    string             `json:"name"`
	Message string             `json:"message"`
	ZoneId  string             `json:"zoneId"`
	Colour  string             `json:"color"` // should have some colours defined
	Effect  keyboard.KeyEffect `json:"effect"`
	Pid     string             `json:"pid"` // always DK5QPID?
}

// https://www.daskeyboard.io/api-resources/signal/resource-description/
type CreateSignalResponse struct {
	CreateSignalRequest
	Id         int    `json:"id"` // signals created via localhost have negative ids
	IsArchived bool   `json:"isArchived"`
	IsRead     bool   `json:"isRead"`
	IsMuted    bool   `json:"isMuted"`
	UserId     int    `json:"userId"`
	ClientName string `json:"clientName"`
	CreatedAt  uint64 `json:"createdAt"`
	UpdatedAt  uint64 `json:"updatedAt"`
}
