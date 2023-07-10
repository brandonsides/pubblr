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
	Actor    json.RawMessage              `json:"actor"`
	Password string                       `json:"password"`
	Inbox    []json.RawMessage            `json:"-"`
	Outbox   []json.RawMessage            `json:"-"`
	Objects  map[string][]json.RawMessage `json:"-"`
	// TODO: make this EntityIface
	Followers []activitystreams.Actor       `json:"-"`
	Following []activitystreams.EntityIface `json:"-"`
	Streams   []activitystreams.EntityIface `json:"-"`
}

type PubblrDatabase struct {
	users map[string]UserData
}

func NewPubblrDatabase(config PubblrDatabaseConfig) *PubblrDatabase {
	return &PubblrDatabase{
		users: make(map[string]UserData),
	}
}

func (d *PubblrDatabase) CreateObject(post activitystreams.ObjectIface, user string, baseUrl url.URL) (activitystreams.ObjectIface, error) {
	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("user %s does not exist")
	}

	objects := userData.Objects
	if objects == nil {
		objects = make(map[string][]json.RawMessage)
	}

	postType, err := post.Type()
	if err != nil {
		return nil, fmt.Errorf("Could not get post type: %w", err)
	}
	postType = strings.ToLower(postType)

	baseUrl.Path = path.Join(baseUrl.Path, user, postType, strconv.Itoa(len(objects[postType])))
	id := baseUrl.String()
	activitystreams.ToObject(post).Id = id

	postJson, err := json.Marshal(post)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal retrieved post: %w", err)
	}

	objects[postType] = append(objects[postType], postJson)

	return post, nil
}

func (d *PubblrDatabase) CreateInboxItem(a activitystreams.ActivityIface, user string) (activitystreams.ActivityIface, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("User %s does not exist", user)
	}

	marshalledActivity, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal activity: %w", err)
	}

	userData.Inbox = append(userData.Inbox, marshalledActivity)
	d.users[user] = userData

	return a, nil
}

func (d *PubblrDatabase) CreateOutboxItem(activity activitystreams.ActivityIface, user string, baseUrl url.URL) (activitystreams.ActivityIface, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("User %s does not exist", user)
	}

	baseUrl.Path = path.Join(baseUrl.Path, user, "outbox", strconv.Itoa(len(userData.Outbox)))
	id := baseUrl.String()
	activitystreams.ToObject(activity).Id = id

	activityJson, err := json.Marshal(activity)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal activity: %w", err)
	}

	userData.Outbox = append(userData.Outbox, activityJson)
	d.users[user] = userData

	return activity, nil
}

func (d *PubblrDatabase) GetInboxPage(user string, page int, pageSize int) ([]activitystreams.ActivityIface, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("User %s does not exist", user)
	}

	rawInbox := userData.Inbox
	if !ok {
		rawInbox = make([]json.RawMessage, 0)
		userData.Inbox = rawInbox
	}

	start := page * pageSize
	end := start + pageSize
	if end > len(rawInbox) {
		end = len(rawInbox)

	}
	rawInbox = rawInbox[start:end]

	inbox := make([]activitystreams.ActivityIface, len(rawInbox))
	for i, activityJson := range rawInbox {
		var activity activitystreams.ActivityIface
		err := activitystreams.DefaultEntityUnmarshaler.Unmarshal(activityJson, &activity)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal activity: %w", err)
		}

		inbox[i] = activity
	}

	return inbox, nil
}

func (d *PubblrDatabase) GetOutboxPage(user string, page int, pageSize int) ([]activitystreams.ActivityIface, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("User %s does not exist", user)
	}

	rawOutbox := userData.Outbox
	if !ok {
		rawOutbox = make([]json.RawMessage, 0)
		userData.Outbox = rawOutbox
	}

	start := page * pageSize
	end := start + pageSize
	if end > len(rawOutbox) {
		end = len(rawOutbox)

	}
	rawOutbox = rawOutbox[start:end]

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

func (d *PubblrDatabase) GetInboxCount(user string) (int, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return 0, fmt.Errorf("user %s does not exist", user)
	}

	return len(userData.Inbox), nil
}

func (d *PubblrDatabase) GetOutboxCount(user string) (int, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return 0, fmt.Errorf("user %s does not exist", user)
	}

	return len(userData.Outbox), nil
}

func (d *PubblrDatabase) GetInboxItem(user, id string) (activitystreams.ActivityIface, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("user %s does not exist", user)
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse id: %w", err)
	}

	if len(userData.Inbox) <= parsedId {
		return nil, fmt.Errorf("No activity with id %s", id)
	}

	activityJson := userData.Inbox[parsedId]
	var activity activitystreams.ActivityIface
	err = activitystreams.DefaultEntityUnmarshaler.Unmarshal(activityJson, &activity)

	return activity, err
}

func (d *PubblrDatabase) GetOutboxItem(user, id string) (activitystreams.ActivityIface, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("user %s does not exist", user)
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse id: %w", err)
	}

	if len(userData.Outbox) <= parsedId {
		return nil, fmt.Errorf("No activity with id %s", id)
	}

	activityJson := userData.Outbox[parsedId]
	var activity activitystreams.ActivityIface
	err = activitystreams.DefaultEntityUnmarshaler.Unmarshal(activityJson, &activity)

	return activity, err
}

func (d *PubblrDatabase) GetObject(user, typ, id string) (activitystreams.ObjectIface, error) {
	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("user %s does not exist", user)
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse id: %w", err)
	}

	objects := userData.Objects[typ]
	if !ok {
		return nil, fmt.Errorf("Object not found")
	}

	if len(objects) <= parsedId {
		return nil, fmt.Errorf("Object not found")
	}

	postJson := objects[parsedId]
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

func (d *PubblrDatabase) CheckPassword(username, password string) error {
	userData, ok := d.users[username]
	if !ok {
		return fmt.Errorf("User %s does not exist", username)
	}

	if userData.Password != password {
		return fmt.Errorf("Wrong password")
	}

	return nil
}
