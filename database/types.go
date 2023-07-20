package database

import (
	"database/sql"

	"gorm.io/gorm"
)

type dbEntity struct {
	gorm.Model
	Type         sql.NullString // `gorm:"column:entity_type"`
	MediaType    sql.NullString
	PreviewID    *uint
	Preview      *dbEntity  //  `gorm:"foreignKey:PreviewID"`
	AttributedTo []dbEntity `gorm:"many2many:attributions;"`
}

type dbLink struct {
	ID        uint     `gorm:"primaryKey;column:id;"`
	Entity    dbEntity `gorm:"foreignKey:ID;"`
	Href      sql.NullString
	Hreflang  sql.NullString
	MediaType sql.NullString
	Rel       sql.NullString
	Height    sql.NullInt64
	Width     sql.NullInt64
}

type dbObject struct {
	ID             uint     `gorm:"primaryKey;column:id;"`
	Entity         dbEntity `gorm:"foreignKey:ID;"`
	ContextID      *uint
	Context        dbEntity // `gorm:"foreignKey:ContextID;"`
	GeneratorID    *uint
	Generator      dbEntity // `gorm:"foreignKey:GeneratorID;"`
	IconID         *uint
	Icon           dbEntity // `gorm:"foreignKey:IconID;"`
	ImageID        *uint
	Image          dbEntity // `gorm:"foreignKey:ImageID;"`
	LocationID     *uint
	Location       dbEntity // `gorm:"foreignKey:LocationID;"`
	StringUrl      sql.NullString
	LinkUrlID      *uint
	LinkUrl        dbLink // `gorm:"foreignKey:LinkUrlID;"`
	Content        sql.NullString
	Summary        sql.NullString
	StartTime      sql.NullTime
	EndTime        sql.NullTime
	Duration       sql.NullInt64
	Attachment     []dbEntity `gorm:"many2many:attachments;"`
	AudienceEntity []dbEntity `gorm:"many2many:audiences;"`
	Bcc            []dbEntity `gorm:"many2many:bccs;"`
	Bto            []dbEntity `gorm:"many2many:btos;"`
	Cc             []dbEntity `gorm:"many2many:ccs;"`
	Replies        []dbEntity `gorm:"many2many:replies;"`
	InReplyTo      []dbEntity `gorm:"many2many:replies;"`
	Tag            []dbEntity `gorm:"many2many:tags;"`
	To             []dbEntity `gorm:"many2many:tos;"`
}

type dbRelationship struct {
	ID             uint     `gorm:"primaryKey;column:id;"`
	Object         dbObject `gorm:"foreignKey:ID;"`
	SubjectID      *uint
	Subject        dbEntity // `gorm:"foreignKey:SubjectID;"`
	ObjectID       *uint
	Obj            dbEntity `gorm:"foreignKey:ObjectID;"`
	RelationshipID *uint
	Relationship   dbObject // `gorm:"foreignKey:RelationshipID;"`
}

type dbPlace struct {
	ID        uint     `gorm:"primaryKey;column:id;"`
	Object    dbObject `gorm:"foreignKey:ID;"`
	Accuracy  sql.NullFloat64
	Altitude  sql.NullFloat64
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
	Radius    sql.NullFloat64
	Units     sql.NullString
}

type dbProfile struct {
	ID          uint     `gorm:"primaryKey;column:id;"`
	Object      dbObject `gorm:"foreignKey:ID;"`
	DescribesID *uint
	Describes   dbEntity // `gorm:"foreignKey:DescribesID;"`
}

type dbActivity struct {
	ID           uint     `gorm:"primaryKey;column:id;"`
	Object       dbObject `gorm:"foreignKey:ID;"`
	ActorID      *uint
	Actor        dbEntity
	InstrumentID *uint
	Instrument   dbEntity
	OriginID     *uint
	Origin       dbEntity
	TargetID     *uint
	Target       dbEntity
	ResultID     *uint
	Result       dbEntity
}

type dbQuestion struct {
	ID           uint       `gorm:"primaryKey;column:id;"`
	Activity     dbActivity `gorm:"foreignKey:ID;"`
	QuestionType sql.NullString
	Answers      []dbEntity `gorm:"many2many:answers;"`
}

type dbAnswer struct {
	ID         uint     `gorm:"primaryKey;column:id;"`
	Entity     dbEntity `gorm:"foreignKey:ID;"`
	QuestionID *uint
	Question   dbQuestion
}

type dbActor struct {
	ID               uint     `gorm:"primaryKey;column:id;"`
	Entity           dbEntity `gorm:"foreignKey:ID;"`
	PreferreUsername sql.NullString
}

type dbUser struct {
	ID                 uint    `gorm:"primaryKey;column:id;"`
	Actor              dbActor `gorm:"foreignKey:ID;"`
	SaltedPasswordHash string
	Salt               string
	Email              string
}

type dbCollection struct {
	ID      uint     `gorm:"primaryKey;column:id;"`
	Object  dbObject `gorm:"foreignKey:ID;"`
	Ordered bool
	Items   []dbEntity `gorm:"many2many:collection_items;"`
}
