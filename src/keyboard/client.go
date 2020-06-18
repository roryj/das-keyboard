package keyboard

import (
	log "github.com/sirupsen/logrus"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/roryj/das-keyboard/src/colour"
	"go.uber.org/ratelimit"
)

const DefaultPort = 27301
const localHost = "http://localhost"
const apiVersion = "1.0"
const pid = "DK5QPID"

type Client interface {
	CreateSignal(zone Zone, effect KeyEffect, colour colour.Hex) (SignalResponse, error)
	DeleteSignal(id int) error
	DeleteSignalAtZone(zone Zone) error
	GetSignal(zone Zone) (SignalResponse, error)
	ClearAllSignals()
}

type keyboardClient struct {
	port    int
	limiter ratelimit.Limiter
}

func NewKeyboardClient(port int) Client {
	return &keyboardClient{port: port, limiter: ratelimit.New(20)}
}

func (c *keyboardClient) CreateSignal(zone Zone, effect KeyEffect, colour colour.Hex) (SignalResponse, error) {
	c.limiter.Take()

	req := CreateSignalRequest{
		Colour:  colour.Hex(),
		Effect:  effect,
		Message: "",
		Name:    "",
		Pid:     pid,
		ZoneId:  zone.GetZoneName(),
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return SignalResponse{}, err
	}

	result, err := http.Post(c.generateUrl("signals"), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return SignalResponse{}, err
	}

	r, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return SignalResponse{}, err
	}

	var response SignalResponse
	err = json.Unmarshal(r, &response)

	return response, err
}

func (c *keyboardClient) DeleteSignal(id int) error {
	c.limiter.Take()
	u := c.generateUrl("signals", url.PathEscape(strconv.Itoa(id)))

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
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received an invalid status code. Expected 200, found %d. Body: %s", resp.StatusCode, body)
	}

	return nil
}

func (c *keyboardClient) DeleteSignalAtZone(zone Zone) error {
	c.limiter.Take()
	u := c.generateUrl("signals", "pid", pid, "zoneid", zone.GetZoneName())

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
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received an invalid status code. Expected 200, found %d. Body: %s", resp.StatusCode, body)
	}

	return nil
}

func (c *keyboardClient) GetSignal(zone Zone) (SignalResponse, error) {
	c.limiter.Take()
	url := c.generateUrl("signals", "pid", pid, "zoneId", zone.GetZoneName())

	result, err := http.Get(url)
	if err != nil {
		return SignalResponse{}, err
	}

	r, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return SignalResponse{}, err
	}

	if result.StatusCode != http.StatusOK {
		return SignalResponse{}, fmt.Errorf("expected 200, got %s", result.Status)
	}

	var response SignalResponse
	err = json.Unmarshal(r, &response)

	return response, err
}

func (c *keyboardClient) ClearAllSignals() {
	var x, y uint
	for x = 0; x < 24; x++ {
		for y = 0; y < 6; y++ {
			z, err := NewXYZone(x, y)
			if err != nil {
				log.Warnf("failed to create zone id. %v", err)
			}
			_ = c.DeleteSignalAtZone(z)
		}
	}
}

func (c *keyboardClient) generateUrl(requestType string, pathArgs ...string) string {
	url := fmt.Sprintf("%s:%d/api/%s/%s", localHost, c.port, apiVersion, requestType)

	if len(pathArgs) != 0 {
		url = fmt.Sprintf("%s/%s", url, strings.Join(pathArgs, "/"))
	}

	return url
}

type CreateSignalRequest struct {
	Name    string    `json:"name"`
	Message string    `json:"message"`
	ZoneId  string    `json:"zoneId"`
	Colour  string    `json:"color"` // should have some colours defined
	Effect  KeyEffect `json:"effect"`
	Pid     string    `json:"pid"` // always DK5QPID?
}

// https://www.daskeyboard.io/api-resources/signal/resource-description/
type SignalResponse struct {
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
