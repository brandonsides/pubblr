package database

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/url"

	"github.com/brandonsides/pubblr/activitystreams"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PubblrDatabaseConfig struct {
	Addr string
	User string
	Pass string
}

type PubblrDatabase struct {
	db gorm.DB
}

func NewPubblrDatabase(config PubblrDatabaseConfig) (*PubblrDatabase, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/pubblr?charset=utf8&parseTime=True&loc=Local", config.User, config.Pass, config.Addr)))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	return &PubblrDatabase{*db}, nil
}

func (d *PubblrDatabase) CreateObject(post activitystreams.ObjectIface, user string, baseUrl url.URL) (activitystreams.ObjectIface, error) {
	/*
		userData, ok := db
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
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) CreateInboxItem(a activitystreams.ActivityIface, user string) (activitystreams.ActivityIface, error) {
	/*
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
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) CreateOutboxItem(activity activitystreams.ActivityIface, user string, baseUrl url.URL) (activitystreams.ActivityIface, error) {
	/*
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
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) GetInboxPage(user string, page int, pageSize int) ([]activitystreams.ActivityIface, error) {
	/*
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
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) GetOutboxPage(user string, page int, pageSize int) ([]activitystreams.ActivityIface, error) {
	/*
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
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) GetInboxCount(user string) (int, error) {
	/*
		if d.users == nil {
			d.users = make(map[string]UserData)
		}

		userData, ok := d.users[user]
		if !ok {
			return 0, fmt.Errorf("user %s does not exist", user)
		}

		return len(userData.Inbox), nil
	*/
	return 0, errors.New("Not implemented")
}

func (d *PubblrDatabase) GetOutboxCount(user string) (int, error) {
	/*
		if d.users == nil {
			d.users = make(map[string]UserData)
		}

		userData, ok := d.users[user]
		if !ok {
			return 0, fmt.Errorf("user %s does not exist", user)
		}

		return len(userData.Outbox), nil
	*/
	return 0, errors.New("Not implemented")
}

func (d *PubblrDatabase) GetInboxItem(user, id string) (activitystreams.ActivityIface, error) {
	/*
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
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) GetOutboxItem(user, id string) (activitystreams.ActivityIface, error) {
	/*
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
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) GetObject(user, typ, id string) (activitystreams.ObjectIface, error) {
	/*
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
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) CreateUser(user activitystreams.ActorIface, email, password string, baseUrl url.URL) (activitystreams.ActorIface, error) {
	actorDbEntity, err := toDBEntity(user)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert actor to db entity: %w", err)
	}

	dbActor, ok := actorDbEntity.Rest.(*dbActor)
	if !ok {
		return nil, fmt.Errorf("Failed to convert actor to db actor")
	}

	salt := GenerateSalt()
	saltedPassword := password + salt
	hasher := sha256.New()
	hasher.Write([]byte(saltedPassword))
	saltedPasswordHash := string(hasher.Sum(nil))

	dbUser := dbUser{
		Actor:              dbActor,
		SaltedPasswordHash: saltedPasswordHash,
		Salt:               salt,
		Email:              email,
	}

	d.db.Save(dbUser)

	return user, nil
}

func (d *PubblrDatabase) GetUser(username string) (activitystreams.ActorIface, error) {
	/*
		userData, ok := d.users[username]
		if !ok {
			return nil, fmt.Errorf("User %s does not exist", username)
		}
		user := userData.Actor
		return user, nil
	*/
	return nil, errors.New("Not implemented")
}

func (d *PubblrDatabase) CheckPassword(username, password string) error {
	/*
		userData, ok := d.users[username]
		if !ok {
			return fmt.Errorf("User %s does not exist", username)
		}

		if userData.Password != password {
			return fmt.Errorf("Wrong password")
		}
	*/

	return errors.New("Not implemented")
}
