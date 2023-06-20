package activitystreams

type Application struct {
	Object
}

func (a *Application) Type() string {
	return "Application"
}

func (a Application) MarshalJSON() ([]byte, error) {
	return MarshalObject(&a)
}

type Group struct {
	Object
}

func (g *Group) Type() string {
	return "Group"
}

func (g Group) MarshalJSON() ([]byte, error) {
	return MarshalObject(&g)
}

type Organization struct {
	Object
}

func (o *Organization) Type() string {
	return "Organization"
}

func (o Organization) MarshalJSON() ([]byte, error) {
	return MarshalObject(&o)
}

type Person struct {
	Object
}

func (p *Person) Type() string {
	return "Person"
}

func (p Person) MarshalJSON() ([]byte, error) {
	return MarshalObject(&p)
}

type Service struct {
	Object
}

func (s *Service) Type() string {
	return "Service"
}

func (s Service) MarshalJSON() ([]byte, error) {
	return MarshalObject(&s)
}
