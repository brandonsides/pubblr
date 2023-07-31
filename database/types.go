package database

import (
	"database/sql"

	"github.com/brandonsides/pubblr/database/util"
	"gorm.io/gorm"
)

type dbEntity struct {
	gorm.Model
	Type         sql.NullString // `gorm:"column:entity_type"`
	MediaType    sql.NullString
	PreviewID    *uint
	Preview      *dbEntity  //  `gorm:"foreignKey:PreviewID"`
	AttributedTo []dbEntity `gorm:"many2many:attributions;"`
	RestID       *uint
	RestType     sql.NullString
	Rest         interface{}
}

type dbLink struct {
	ID        uint      `gorm:"primaryKey;column:id;"`
	Entity    *dbEntity `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:link;"`
	Href      sql.NullString
	Hreflang  sql.NullString
	MediaType sql.NullString
	Rel       util.StringArray
	Height    *uint64
	Width     *uint64
}

type dbObject struct {
	ID          uint      `gorm:"primaryKey;column:id;"`
	Entity      *dbEntity `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:object;"`
	ContextID   *uint
	Context     *dbEntity // `gorm:"foreignKey:ContextID;"`
	GeneratorID *uint
	Generator   *dbEntity // `gorm:"foreignKey:GeneratorID;"`
	IconID      *uint
	Icon        *dbEntity // `gorm:"foreignKey:IconID;"`
	ImageID     *uint
	Image       *dbEntity // `gorm:"foreignKey:ImageID;"`
	LocationID  *uint
	Location    []dbEntity `gorm:"many2many:locations;"`
	StringUrl   sql.NullString
	LinkUrlID   *uint
	LinkUrl     *dbLink // `gorm:"foreignKey:LinkUrlID;"`
	Content     sql.NullString
	Summary     sql.NullString
	StartTime   sql.NullTime
	EndTime     sql.NullTime
	Duration    sql.NullInt64
	Attachment  []dbEntity `gorm:"many2many:attachments;"`
	Audience    []dbEntity `gorm:"many2many:audiences;"`
	Bcc         []dbEntity `gorm:"many2many:bccs;"`
	Bto         []dbEntity `gorm:"many2many:btos;"`
	Cc          []dbEntity `gorm:"many2many:ccs;"`
	Replies     []dbEntity `gorm:"many2many:replies;"`
	InReplyTo   []dbEntity `gorm:"many2many:replies;"`
	Tag         []dbEntity `gorm:"many2many:tags;"`
	To          []dbEntity `gorm:"many2many:tos;"`
	RestID      *uint
	RestType    sql.NullString
	Rest        interface{}
}

type dbRelationship struct {
	ID             uint      `gorm:"primaryKey;column:id;"`
	Object         *dbObject `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:relationship;"`
	SubjectID      *uint
	Subject        *dbEntity // `gorm:"foreignKey:SubjectID;"`
	ObjectID       *uint
	Obj            *dbEntity `gorm:"foreignKey:ObjectID;"`
	RelationshipID *uint
	Relationship   *dbObject // `gorm:"foreignKey:RelationshipID;"`
}

type dbPlace struct {
	ID        uint      `gorm:"primaryKey;column:id;"`
	Object    *dbObject `gorm:"foreignKey:ID;ploymorphic:Rest;polymorphicValue:place;"`
	Accuracy  sql.NullFloat64
	Altitude  sql.NullFloat64
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
	Radius    sql.NullFloat64
	Units     sql.NullString
}

type dbProfile struct {
	ID          uint      `gorm:"primaryKey;column:id;"`
	Object      *dbObject `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:profile;"`
	DescribesID *uint
	Describes   *dbEntity // `gorm:"foreignKey:DescribesID;"`
}

type dbActivity struct {
	ID           uint      `gorm:"primaryKey;column:id;"`
	Object       *dbObject `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:activity;"`
	ActorID      *uint
	Actor        *dbEntity
	InstrumentID *uint
	Instrument   *dbEntity
	OriginID     *uint
	Origin       *dbEntity
	TargetID     *uint
	Target       *dbEntity
	ResultID     *uint
	Result       *dbEntity
	RestID       *uint
	RestType     sql.NullString
	Rest         interface{}
}

type dbTransitiveActivity struct {
	ID       uint        `gorm:"primaryKey;column:id;"`
	Activity *dbActivity `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:transitive_activity;"`
	ObjectID *uint
	Object   *dbEntity
}

type dbQuestion struct {
	ID           uint        `gorm:"primaryKey;column:id;"`
	Activity     *dbActivity `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:question;"`
	QuestionType sql.NullString
	Answers      []dbEntity `gorm:"many2many:answers;"`
}

type dbActor struct {
	ID               uint      `gorm:"primaryKey;column:id;"`
	Object           *dbObject `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:actor;"`
	PreferreUsername sql.NullString
}

type dbUser struct {
	ID                 uint     `gorm:"primaryKey;column:id;"`
	Actor              *dbActor `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:user;"`
	SaltedPasswordHash string
	Salt               string
	Email              string
}

type dbCollection struct {
	ID      uint      `gorm:"primaryKey;column:id;"`
	Object  *dbObject `gorm:"foreignKey:ID;polymorphic:Rest;polymorphicValue:collection;"`
	Ordered bool
	Items   []dbEntity `gorm:"many2many:collection_items;"`
}
