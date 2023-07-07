package database

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/brandonsides/pubblr/activitystreams"
)

type PubblrDatabaseConfig struct {
}

type UserData struct {
	Actor    json.RawMessage `json:"actor"`
	Password string          `json:"password"`
}

type PubblrDatabase struct {
	posts    map[string]map[string][]json.RawMessage
	users    map[string]UserData
	outboxes map[string][]json.RawMessage
}

func NewPubblrDatabase(config PubblrDatabaseConfig) *PubblrDatabase {
	return &PubblrDatabase{
		posts:    make(map[string]map[string][]json.RawMessage),
		users:    make(map[string]UserData),
		outboxes: make(map[string][]json.RawMessage),
	}
}

func (d *PubblrDatabase) CreateObject(post activitystreams.ObjectIface, user string, baseUrl url.URL) (activitystreams.ObjectIface, error) {
	if d.posts == nil {
		d.posts = make(map[string]map[string][]json.RawMessage)
	}

	postType, err := post.Type()
	if err != nil {
		return nil, fmt.Errorf("Could not get post type: %w", err)
	}
	postType = strings.ToLower(postType)

	postsByUser, ok := d.posts[user]
	if !ok {
		postsByUser = make(map[string][]json.RawMessage)
		d.posts[user] = postsByUser
	}

	postsOfType, ok := postsByUser[postType]
	if !ok {
		postsOfType = make([]json.RawMessage, 0)
		postsByUser[postType] = postsOfType
	}

	baseUrl.Path = path.Join(baseUrl.Path, user, postType, strconv.Itoa(len(postsOfType)))
	id := baseUrl.String()
	activitystreams.ToObject(post).Id = id

	postJson, err := json.Marshal(post)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal retrieved post: %w", err)
	}

	postsOfType = append(postsOfType, postJson)
	postsByUser[postType] = postsOfType

	return post, nil
}

func (d *PubblrDatabase) CreateActivity(activity activitystreams.ActivityIface, user string, baseUrl url.URL) (activitystreams.ActivityIface, error) {
	if d.outboxes == nil {
		d.outboxes = make(map[string][]json.RawMessage)
	}

	outbox, ok := d.outboxes[user]
	if !ok {
		outbox = make([]json.RawMessage, 0)
		d.outboxes[user] = outbox
	}

	baseUrl.Path = path.Join(baseUrl.Path, user, "outbox", strconv.Itoa(len(outbox)))
	id := baseUrl.String()
	activitystreams.ToObject(activity).Id = id

	activityJson, err := json.Marshal(activity)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal activity: %w", err)
	}

	outbox = append(outbox, activityJson)
	d.outboxes[user] = outbox

	return activity, nil
}

func (d *PubblrDatabase) GetOutbox(user string) ([]activitystreams.ActivityIface, error) {
	if d.outboxes == nil {
		d.outboxes = make(map[string][]json.RawMessage)
	}

	rawOutbox, ok := d.outboxes[user]
	if !ok {
		rawOutbox = make([]json.RawMessage, 0)
		d.outboxes[user] = rawOutbox
	}

	outbox := make([]activitystreams.ActivityIface, len(rawOutbox))
	for i, activityJson := range rawOutbox {
		var activity activitystreams.ActivityIface
		err := activitystreams.DefaultEntityUnmarshaler.Unmarshal(activityJson, &activity)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal activity: %w", err)
		}

		outbox[i] = activity
	}

	return outbox, nil
}

func (d *PubblrDatabase) GetActivity(user, id string) (activitystreams.ActivityIface, error) {
	if d.outboxes == nil {
		d.outboxes = make(map[string][]json.RawMessage)
	}

	outbox, ok := d.outboxes[user]
	if !ok {
		outbox = make([]json.RawMessage, 0)
		d.outboxes[user] = outbox
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse id: %w", err)
	}

	if len(outbox) <= parsedId {
		return nil, fmt.Errorf("No activity with id %s", id)
	}

	activityJson := outbox[parsedId]
	var activity activitystreams.ActivityIface
	err = activitystreams.DefaultEntityUnmarshaler.Unmarshal(activityJson, &activity)

	return activity, err
}

func (d *PubblrDatabase) GetPost(user, typ, id string) (activitystreams.ObjectIface, error) {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse id: %w", err)
	}

	postsByUser, ok := d.posts[user]
	if !ok {
		return nil, fmt.Errorf("No posts by user %s", user)
	}

	postsOfType, ok := postsByUser[typ]
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

func (d *PubblrDatabase) CreateUser(user activitystreams.ActorIface, username, password string, baseUrl url.URL) (activitystreams.ActorIface, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	_, ok := d.users[username]
	if ok {
		return nil, fmt.Errorf("User %s already exists", username)
	}

	baseUrl.Path = path.Join(baseUrl.Path, username)
	id := baseUrl.String()

	activitystreams.ToObject(user).Id = id

	bytes, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	userdata := d.users[username]
	userdata.Actor = bytes
	userdata.Password = password
	d.users[username] = userdata
	return user, nil
}

func (d *PubblrDatabase) GetUser(username string) (activitystreams.ActorIface, error) {
	userData, ok := d.users[username]
	if !ok {
		return nil, fmt.Errorf("User %s does not exist", username)
	}
	userJson := userData.Actor

	var user activitystreams.ActorIface
	err := activitystreams.DefaultEntityUnmarshaler.Unmarshal(userJson, &user)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal retrieved user: %w", err)
	}

	return user, nil
}
