package activitystreams

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/brandonsides/pubblr/util"
)

const (
	LinkTypeLink    = "Link"
	LinkTypeMention = "Mention"
)

type LinkIface interface {
	json.Marshaler
	link() *Link
	Type() string
}

// Marhsal a LinkIface to JSON
// Marshals the implementing type to JSON and adds a "type" field to the JSON
// representation with the value returned by the Type() method.
func MarshalLink(l LinkIface) ([]byte, error) {
	linkMap := make(map[string]interface{})

	LinkIfaceType := reflect.TypeOf((*LinkIface)(nil)).Elem()

	linkType := reflect.TypeOf(l).Elem()
	for fieldIndex := 0; fieldIndex < linkType.NumField(); fieldIndex++ {
		field := linkType.Field(fieldIndex)
		if field.Anonymous && (field.Type == LinkIfaceType || reflect.PointerTo(field.Type).Implements(LinkIfaceType)) {
			fieldInterface := reflect.ValueOf(l).Elem().Field(fieldIndex).Interface()
			if link, ok := fieldInterface.(Link); ok {
				fieldInterface = (rawLink)(link)
			}
			var nestedMap map[string]interface{}
			nestedJson, err := json.Marshal(fieldInterface)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(nestedJson, &nestedMap)
			if err != nil {
				return nil, err
			}

			for k, v := range nestedMap {
				linkMap[k] = v
			}
			continue
		}
		tag := util.FromString(field.Tag.Get("json"))
		if tag.Name == "" {
			tag.Name = strings.ToLower(field.Name[:1]) + field.Name[1:]
		}
		if tag.Name == "-" || tag.OmitEmpty && reflect.ValueOf(l).Elem().Field(fieldIndex).IsZero() {
			continue
		}

		v := reflect.ValueOf(l).Elem().Field(fieldIndex)
		if tag.String {
			linkMap[tag.Name] = v.String()
		} else {
			linkMap[tag.Name] = v.Interface()
		}
	}

	if l != nil {
		linkMap["type"] = l.Type()
	} else {
		linkMap["type"] = "Link"
	}

	return json.Marshal(linkMap)
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

func (l *Link) Type() string {
	return LinkTypeLink
}

func (l *Link) MarshalJSON() ([]byte, error) {
	return MarshalLink(l)
}

type Mention struct {
	Link
}

func (m *Mention) Type() string {
	return LinkTypeMention
}

func (l *Mention) MarshalJSON() ([]byte, error) {
	return MarshalLink(l)
}
