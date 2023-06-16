package activitystreams

type Application Object

func (a *Application) Type() string {
	return "Application"
}

type Group Object

func (g *Group) Type() string {
	return "Group"
}

type Organization Object

func (o *Organization) Type() string {
	return "Organization"
}

type Person struct {
	Object
}

type rawPerson Person

func (p *Person) Type() string {
	return "Person"
}

func (p *Person) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawPerson)(p))
}

type Service Object

func (s *Service) Type() string {
	return "Service"
}
