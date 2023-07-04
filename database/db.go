package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/brandonsides/pubblr/activitystreams"
)

type PubblrDatabaseConfig struct {
}

type PubblrDatabase map[string][]json.RawMessage

func NewPubblrDatabase(config PubblrDatabaseConfig) *PubblrDatabase {
	ret := make(PubblrDatabase)
	return &ret
}

func (d *PubblrDatabase) CreatePost(post activitystreams.ObjectIface) (string, error) {
	postType, err := post.Type()
	if err != nil {
		return "", fmt.Errorf("Could not get post type: %w", err)
	}
	postType = strings.ToLower(postType)

	postJson, err := json.Marshal(post)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal retrieved post: %w", err)
	}

	id := strconv.Itoa(len((*d)[postType]))
	(*d)[postType] = append((*d)[postType], postJson)

	return id, nil
}

func (d *PubblrDatabase) GetPostByTypeAndId(typ string, id string) (activitystreams.ObjectIface, error) {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse id: %w", err)
	}

	postsOfType, ok := (*d)[typ]
	if !ok {
		return nil, fmt.Errorf("No posts of type %s", typ)
	}
	if len(postsOfType) <= parsedId {
		return nil, fmt.Errorf("No post of type %s with id %s", typ, id)
	}

	postJson := postsOfType[parsedId]
	var post activitystreams.ObjectIface
	err = activitystreams.DefaultEntityUnmarshaler.Unmarshal(postJson, &post)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal retrieved post: %w", err)
	}

	return post, nil
}
