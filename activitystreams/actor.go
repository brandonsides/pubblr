package activitystreams

import (
	"encoding/json"
	"fmt"
)

type Actor struct {
	Object
	Inbox             CollectionIface `json:"inbox,omitempty"`
	Outbox            CollectionIface `json:"outbox,omitempty"`
	Following         CollectionIface `json:"following,omitempty"`
	Followers         CollectionIface `json:"followers,omitempty"`
	Liked             CollectionIface `json:"liked,omitempty"`
	Streams           CollectionIface `json:"streams,omitempty"`
	PreferredUsername string          `json:"preferredUsername,omitempty"`
	Endpoints         *ActorEndpoints `json:"endpoints,omitempty"`
}

func (a *Actor) MarshalJSON() ([]byte, error) {
	actor, err := json.Marshal(a.Object)
	if err != nil {
		return nil, err
	}

	var actorMap map[string]json.RawMessage
	err = json.Unmarshal(actor, &actorMap)
	if err != nil {
		return nil, err
	}

	if a.Inbox != nil {
		actorMap["inbox"] = []byte(fmt.Sprintf("%q", ToEntity(a.Inbox).Id))
	}

	if a.Outbox != nil {
		actorMap["outbox"] = []byte(fmt.Sprintf("%q", ToEntity(a.Outbox).Id))
	}

	if a.Following != nil {
		actorMap["following"] = []byte(fmt.Sprintf("%q", ToEntity(a.Following).Id))
	}

	if a.Followers != nil {
		actorMap["followers"] = []byte(fmt.Sprintf("%q", ToEntity(a.Followers).Id))
	}

	if a.Liked != nil {
		actorMap["liked"] = []byte(fmt.Sprintf("%q", ToEntity(a.Liked).Id))
	}

	if a.Streams != nil {
		actorMap["streams"] = []byte(fmt.Sprintf("%q", ToEntity(a.Streams).Id))
	}

	if a.PreferredUsername != "" {
		actorMap["preferredUsername"] = []byte(fmt.Sprintf("%q", a.PreferredUsername))
	}

	if a.Endpoints != nil {
		endpoints, err := json.Marshal(a.Endpoints)
		if err != nil {
			return nil, err
		}
		actorMap["endpoints"] = endpoints
	}

	return json.Marshal(actorMap)
}

type ActorEndpoints struct {
	ProxyUrl                   string `json:"proxyUrl,omitempty"`
	OauthAuthorizationEndpoint string `json:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndpoint         string `json:"oauthTokenEndpoint,omitempty"`
	UploadMedia                string `json:"uploadMedia,omitempty"`
	ProvideClientKey           bool   `json:"provideClientKey,omitempty"`
	SignClientKey              bool   `json:"signClientKey,omitempty"`
}

type ActorIface interface {
	ObjectIface
	actor() *Actor
	Type() (string, error)
}

func (a *Actor) actor() *Actor {
	return a
}

func ToActor(a ActorIface) *Actor {
	return a.actor()
}

type Application struct {
	Actor
}

func (a *Application) Type() (string, error) {
	return "Application", nil
}

func (a *Application) MarshalJSON() ([]byte, error) {
	application, err := json.Marshal(a.Actor)
	if err != nil {
		return nil, err
	}

	var applicationMap map[string]json.RawMessage
	err = json.Unmarshal(application, &applicationMap)
	if err != nil {
		return nil, err
	}

	applicationMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Application"))

	return json.Marshal(applicationMap)
}

type Group struct {
	Actor
}

func (g *Group) Type() (string, error) {
	return "Group", nil
}

func (g *Group) MarshalJSON() ([]byte, error) {
	group, err := json.Marshal(g.Actor)
	if err != nil {
		return nil, err
	}

	var groupMap map[string]json.RawMessage
	err = json.Unmarshal(group, &groupMap)
	if err != nil {
		return nil, err
	}

	groupMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Group"))

	return json.Marshal(groupMap)
}

type Organization struct {
	Actor
}

func (o *Organization) Type() (string, error) {
	return "Organization", nil
}

func (o *Organization) MarshalJSON() ([]byte, error) {
	organization, err := json.Marshal(o.Actor)
	if err != nil {
		return nil, err
	}

	var organizationMap map[string]json.RawMessage
	err = json.Unmarshal(organization, &organizationMap)
	if err != nil {
		return nil, err
	}

	organizationMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Organization"))

	return json.Marshal(organizationMap)
}

type Person struct {
	Actor
}

func (p *Person) Type() (string, error) {
	return "Person", nil
}

func (p *Person) MarshalJSON() ([]byte, error) {
	person, err := json.Marshal(p.Actor)
	if err != nil {
		return nil, err
	}

	var personMap map[string]json.RawMessage
	err = json.Unmarshal(person, &personMap)
	if err != nil {
		return nil, err
	}

	personMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Person"))

	return json.Marshal(personMap)
}

type Service struct {
	Actor
}

func (s *Service) Type() (string, error) {
	return "Service", nil
}

func (s *Service) MarshalJSON() ([]byte, error) {
	service, err := json.Marshal(s.Actor)
	if err != nil {
		return nil, err
	}

	var serviceMap map[string]json.RawMessage
	err = json.Unmarshal(service, &serviceMap)
	if err != nil {
		return nil, err
	}

	serviceMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Service"))

	return json.Marshal(serviceMap)
}
