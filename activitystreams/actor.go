package activitystreams

type Application struct {
	Object
}

func (a *Application) Type() (string, error) {
	return "Application", nil
}

func (a *Application) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

type Group struct {
	Object
}

func (g *Group) Type() (string, error) {
	return "Group", nil
}

func (g *Group) MarshalJSON() ([]byte, error) {
	return MarshalEntity(g)
}

type Organization struct {
	Object
}

func (o *Organization) Type() (string, error) {
	return "Organization", nil
}

func (o *Organization) MarshalJSON() ([]byte, error) {
	return MarshalEntity(o)
}

type Person struct {
	Object
}

func (p *Person) Type() (string, error) {
	return "Person", nil
}

func (p *Person) MarshalJSON() ([]byte, error) {
	return MarshalEntity(p)
}

type Service struct {
	Object
}

func (s *Service) Type() (string, error) {
	return "Service", nil
}

func (s *Service) MarshalJSON() ([]byte, error) {
	return MarshalEntity(s)
}
