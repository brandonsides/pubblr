package activitystreams

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

type ActorEndpoints struct {
	ProxyUrl                   string `json:"proxyUrl,omitempty"`
	OauthAuthorizationEndpoint string `json:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndpoint         string `json:"oauthTokenEndpoint,omitempty"`
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

type Application struct {
	Actor
}

func (a *Application) Type() (string, error) {
	return "Application", nil
}

func (a *Application) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

type Group struct {
	Actor
}

func (g *Group) Type() (string, error) {
	return "Group", nil
}

func (g *Group) MarshalJSON() ([]byte, error) {
	return MarshalEntity(g)
}

type Organization struct {
	Actor
}

func (o *Organization) Type() (string, error) {
	return "Organization", nil
}

func (o *Organization) MarshalJSON() ([]byte, error) {
	return MarshalEntity(o)
}

type Person struct {
	Actor
}

func (p *Person) Type() (string, error) {
	return "Person", nil
}

func (p *Person) MarshalJSON() ([]byte, error) {
	return MarshalEntity(p)
}

type Service struct {
	Actor
}

func (s *Service) Type() (string, error) {
	return "Service", nil
}

func (s *Service) MarshalJSON() ([]byte, error) {
	return MarshalEntity(s)
}
