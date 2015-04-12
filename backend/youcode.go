package youcode

import (
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"

	"appengine/datastore"
)

const (
	webClientID         = "replace this with your Web client ID"
	apiExplorerClientID = endpoints.APIExplorerClientID
)

var (
	scopes  = []string{endpoints.EmailScope}
	clients = []string{webClientID, apiExplorerClientID}
)

type Channel struct {
	YouTubeID string
	Name      string
}

type Channels struct {
	Items []Channel
}

func init() {
	_, err := endpoints.RegisterService(ChannelsAPI{}, "channels", "v1", "channels API", true)
	if err != nil {
		panic(err)
	}
	endpoints.HandleHTTP()
}

type ChannelsAPI struct{}

func (ChannelsAPI) List(c endpoints.Context) (*Channels, error) {
	channels := []Channel{}

	_, err := datastore.NewQuery("Channel").GetAll(c, &channels)
	if err != nil {
		return nil, endpoints.NewInternalServerError("get channels: %v", err)
	}

	return &Channels{channels}, nil
}

func (ChannelsAPI) Add(c endpoints.Context, ch *Channel) (*Channel, error) {
	u, err := endpoints.CurrentUser(c, scopes, nil, clients)
	if err != nil {
		c.Errorf("auth: %v", err)
		return nil, endpoints.NewUnauthorizedError("authorization required")
	}

	if !strings.HasSuffix(u.Email, "gmail.com") {
		return nil, endpoints.NewUnauthorizedError("authorization refused")
	}

	if ch.Name == "" || ch.YouTubeID == "" {
		return nil, endpoints.NewBadRequestError("empty channel")
	}

	k := datastore.NewKey(c, "Channel", ch.YouTubeID, 0, nil)
	_, err = datastore.Put(c, k, ch)
	if err != nil {
		return nil, endpoints.NewInternalServerError("save channel: %v", err)
	}

	return ch, endpoints.NewAPIError("created", "created", http.StatusCreated)
}
