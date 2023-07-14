package database

import (
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
	Actor    activitystreams.ActorIface               `json:"actor"`
	Password string                                   `json:"password"`
	Inbox    []activitystreams.ActivityIface          `json:"-"`
	Outbox   []activitystreams.ActivityIface          `json:"-"`
	Objects  map[string][]activitystreams.ObjectIface `json:"-"`
	// TODO: make this EntityIface
	Followers []activitystreams.ActorIface  `json:"-"`
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
		objects = make(map[string][]activitystreams.ObjectIface)
	}

	postType, err := post.Type()
	if err != nil {
		return nil, fmt.Errorf("Could not get post type: %w", err)
	}
	postType = strings.ToLower(postType)

	baseUrl.Path = path.Join(baseUrl.Path, user, postType, strconv.Itoa(len(objects[postType])))
	id := baseUrl.String()
	activitystreams.ToObject(post).Id = id

	objects[postType] = append(objects[postType], post)

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

	userData.Inbox = append(userData.Inbox, a)
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

	userData.Outbox = append(userData.Outbox, activity)
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

	inbox := userData.Inbox
	if !ok {
		inbox = make([]activitystreams.ActivityIface, 0)
		userData.Inbox = inbox
	}

	start := page * pageSize
	end := start + pageSize
	if end > len(inbox) {
		end = len(inbox)

	}
	inboxPage := inbox[start:end]

	return inboxPage, nil
}

func (d *PubblrDatabase) GetOutboxPage(user string, page int, pageSize int) ([]activitystreams.ActivityIface, error) {
	if d.users == nil {
		d.users = make(map[string]UserData)
	}

	userData, ok := d.users[user]
	if !ok {
		return nil, fmt.Errorf("User %s does not exist", user)
	}

	outbox := userData.Outbox
	if !ok {
		outbox = make([]activitystreams.ActivityIface, 0)
		userData.Outbox = outbox
	}

	start := page * pageSize
	end := start + pageSize
	if end > len(outbox) {
		end = len(outbox)

	}
	outboxPage := outbox[start:end]

	return outboxPage, nil
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

	activity := userData.Inbox[parsedId]
	return activity, nil
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

	activity := userData.Outbox[parsedId]
	return activity, nil
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

	post := objects[parsedId]
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

	userdata := d.users[username]
	userdata.Actor = user
	userdata.Password = password
	d.users[username] = userdata
	return user, nil
}

func (d *PubblrDatabase) GetUser(username string) (activitystreams.ActorIface, error) {
	userData, ok := d.users[username]
	if !ok {
		return nil, fmt.Errorf("User %s does not exist", username)
	}
	user := userData.Actor
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
