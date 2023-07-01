package types

import (
	"github.com/brandonsides/pubblr/activitystreams"
)

const (
	LinkTypeLink    = "Link"
	LinkTypeMention = "Mention"
)

type LinkIface interface {
	activitystreams.EntityIface
	link() *Link
}

type Link struct {
	activitystreams.Entity
	Preview  activitystreams.EntityIface `json:"preview,omitempty"`
	Height   *uint64                     `json:"height,omitempty"`
	Href     string                      `json:"href,omitempty"`
	HrefLang string                      `json:"hreflang,omitempty"`
	Rel      []string                    `json:"rel,omitempty"`
	Width    *uint64                     `json:"width,omitempty"`
}

func (l *Link) link() *Link {
	return l
}

func (l *Link) Type() (string, error) {
	return LinkTypeLink, nil
}

func (l *Link) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(l)
}

type Mention struct {
	Link
}

func (m *Mention) Type() (string, error) {
	return LinkTypeMention, nil
}

func (l *Mention) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(l)
}
