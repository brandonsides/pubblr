package activitystreams

import (
	"encoding/json"
	"fmt"
)

const (
	LinkTypeLink    = "Link"
	LinkTypeMention = "Mention"
)

type LinkIface interface {
	EntityIface
	link() *Link
}

type Link struct {
	Entity
	Preview  EntityIface `json:"preview,omitempty"`
	Height   *uint64     `json:"height,omitempty"`
	Href     string      `json:"href,omitempty"`
	HrefLang string      `json:"hreflang,omitempty"`
	Rel      []string    `json:"rel,omitempty"`
	Width    *uint64     `json:"width,omitempty"`
}

func (l *Link) link() *Link {
	return l
}

func (l *Link) Type() (string, error) {
	return LinkTypeLink, nil
}

func (l *Link) MarshalJSON() ([]byte, error) {
	link, err := json.Marshal(&l.Entity)
	if err != nil {
		return nil, err
	}

	var linkMap map[string]json.RawMessage
	err = json.Unmarshal(link, &linkMap)
	if err != nil {
		return nil, err
	}

	if l.Preview != nil {
		linkMap["preview"] = []byte(fmt.Sprintf("%q", ToEntity(l.Preview).Id))
	}

	if l.Height != nil {
		linkMap["height"] = []byte(fmt.Sprintf("%d", *l.Height))
	}

	if l.Href != "" {
		linkMap["href"] = []byte(fmt.Sprintf("%q", l.Href))
	}

	if l.HrefLang != "" {
		linkMap["hreflang"] = []byte(fmt.Sprintf("%q", l.HrefLang))
	}

	if len(l.Rel) > 0 {
		rels, err := json.Marshal(l.Rel)
		if err != nil {
			return nil, err
		}
		linkMap["rel"] = rels
	}

	if l.Width != nil {
		linkMap["width"] = []byte(fmt.Sprintf("%d", *l.Width))
	}

	linkMap["type"] = []byte(fmt.Sprintf("%q", LinkTypeLink))

	return json.Marshal(linkMap)
}

func (l *Link) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := l.Entity.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	if preview, ok := objMap["preview"]; ok {
		l.Preview, err = u.UnmarshalEntity(preview)
		if err != nil {
			return err
		}
	}

	if height, ok := objMap["height"]; ok {
		err = json.Unmarshal(height, &l.Height)
		if err != nil {
			return err
		}
	}

	if href, ok := objMap["href"]; ok {
		err = json.Unmarshal(href, &l.Href)
		if err != nil {
			return err
		}
	}

	if hreflang, ok := objMap["hreflang"]; ok {
		err = json.Unmarshal(hreflang, &l.HrefLang)
		if err != nil {
			return err
		}
	}

	if rel, ok := objMap["rel"]; ok {
		err = json.Unmarshal(rel, &l.Rel)
		if err != nil {
			return err
		}
	}

	if width, ok := objMap["width"]; ok {
		err = json.Unmarshal(width, &l.Width)
		if err != nil {
			return err
		}
	}

	return nil
}

type Mention struct {
	Link
}

func (m *Mention) Type() (string, error) {
	return LinkTypeMention, nil
}

func (l *Mention) MarshalJSON() ([]byte, error) {
	mention, err := json.Marshal(&l.Link)
	if err != nil {
		return nil, err
	}

	var mentionMap map[string]json.RawMessage
	err = json.Unmarshal(mention, &mentionMap)
	if err != nil {
		return nil, err
	}

	mentionMap["type"] = []byte(fmt.Sprintf("%q", LinkTypeMention))

	return json.Marshal(mentionMap)
}
