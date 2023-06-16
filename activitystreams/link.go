package activitystreams

import (
	"encoding/json"

	"github.com/brandonsides/pubblr/util"
)

const (
	LinkTypeLink    = "Link"
	LinkTypeMention = "Mention"
)

type LinkIface interface {
	link() *Link
	Type() string
}

func MarshalLink(l LinkIface) ([]byte, error) {
	var mapped map[string]interface{}
	j, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(j, &mapped)
	if err != nil {
		return nil, err
	}

	objectType := l.Type()
	if objectType != "" {
		mapped["type"] = objectType
	}

	return json.Marshal(mapped)
}

type Link struct {
	Id           string                      `json:"id,omitempty"`
	AttributedTo []util.Either[Object, Link] `json:"attributedTo,omitempty"`
	Preview      *util.Either[Object, Link]  `json:"preview,omitempty"`
	Name         string                      `json:"name,omitempty"`
	Height       *uint64                     `json:"height,omitempty"`
	Href         string                      `json:"href,omitempty"`
	HrefLang     string                      `json:"hreflang,omitempty"`
	MediaType    string                      `json:"mediaType,omitempty"`
	Rel          []string                    `json:"rel,omitempty"`
	Width        *uint64                     `json:"width,omitempty"`
}

type rawLink Link

func (l *Link) link() *Link {
	return l
}

func (l *rawLink) link() *Link {
	return (*Link)(l)
}

func (l *Link) Type() string {
	return LinkTypeLink
}

func (l *rawLink) Type() string {
	return LinkTypeLink
}

func (l Link) MarshalJSON() ([]byte, error) {
	return MarshalLink((*rawLink)(&l))
}

type Mention struct {
	Link
}

type rawMention Mention

func (m *Mention) Type() string {
	return LinkTypeMention
}

func (m *rawMention) Type() string {
	return LinkTypeMention
}

func (l Mention) MarshalJSON() ([]byte, error) {
	return MarshalLink((*rawMention)(&l))
}
