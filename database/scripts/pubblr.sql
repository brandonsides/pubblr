CREATE DATABASE pubblr;


TYPE activity_streams_type AS ENUM (
    'Article',
    'Audio',
    'Document',
    'Event',
    'Image',
    'Note',
    'Page',
    'Place',
    'Profile',
    'Relationship',
    'Tombstone',
    'Video',
    'Collection',
    'OrderedCollection',
    'CollectionPage',
    'OrderedCollectionPage',
    'Accept',
    'Add',
    'Announce',
    'Arrive',
    'Block',
    'Create',
    'Delete',
    'Dislike',
    'Flag',
    'Follow',
    'Ignore',
    'Invite',
    'Join',
    'Leave',
    'Like',
    'Listen',
    'Move',
    'Offer',
    'Question',
    'Reject',
    'Read',
    'Remove',
    'TentativeReject',
    'TentativeAccept',
    'Travel',
    'Undo',
    'Update',
    'View',
    'Application',
    'Group',
    'Organization',
    'Person',
    'Service',
    'Link',
    'Mention',
);

CREATE TABLE pubblr.entities (
    id SERIAL PRIMARY KEY,
    entity_type activity_streams_type,
    entity_name VARCHAR(255) NOT NULL,
    media_type VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.attributions (
    id SERIAL PRIMARY KEY,
    entity INT FOREIGN KEY REFERENCES pubblr.entities(id),
    attributedTo INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.objects (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.entities(id),
    context INT FOREIGN KEY REFERENCES pubblr.entities(id),
    generator INT FOREIGN KEY REFERENCES pubblr.entities(id),
    icon INT FOREIGN KEY REFERENCES pubblr.entities(id),
    img INT FOREIGN KEY REFERENCES pubblr.entities(id),
    loc INT FOREIGN KEY REFERENCES pubblr.entities(id),
    preview INT FOREIGN KEY REFERENCES pubblr.entities(id),
    stringUrl VARCHAR(255),
    linkUrl INT FOREIGN KEY REFERENCES pubblr.links(id),
    content LONGTEXT,
    summary LONGTEXT,
    startTime TIMESTAMP,
    endTime TIMESTAMP,
    duration INT,
);

CREATE TABLE pubblr.attachments (
    id SERIAL PRIMARY KEY
    obj INT FOREIGN KEY REFERENCES pubblr.objects(id),
    attachment INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.audiences (
    id SERIAL PRIMARY KEY
    obj INT FOREIGN KEY REFERENCES pubblr.objects(id),
    audience INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.bccs (
    id SERIAL PRIMARY KEY
    obj INT FOREIGN KEY REFERENCES pubblr.objects(id),
    bcc INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.btos (
    id SERIAL PRIMARY KEY
    obj INT FOREIGN KEY REFERENCES pubblr.objects(id),
    bto INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.ccs (
    id SERIAL PRIMARY KEY
    obj INT FOREIGN KEY REFERENCES pubblr.objects(id),
    cc INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.replies (
    id SERIAL PRIMARY KEY
    entity INT FOREIGN KEY REFERENCES pubblr.entities(id),
    reply INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.tags (
    id SERIAL PRIMARY KEY
    entity INT FOREIGN KEY REFERENCES pubblr.entities(id),
    tag INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.tos (
    id SERIAL PRIMARY KEY
    entity INT FOREIGN KEY REFERENCES pubblr.entities(id),
    to INT FOREIGN KEY REFERENCES pubblr.entities(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE pubblr.relationships (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.objects(id),
    subj INT FOREIGN KEY REFERENCES pubblr.entities(id),
    obj INT FOREIGN KEY REFERENCES pubblr.entities(id),
    relationship INT FOREIGN KEY REFERENCES pubblr.objects(id)
);

CREATE TABLE pubblr.places (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.objects(id),
    accuracy REAL,
    altitude REAL,
    latitude REAL,
    longitude REAL,
    radius REAL,
    units VARCHAR(255)
);

CREATE TABLE pubblr.profiles (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.objects(id),
    describes INT FOREIGN KEY REFERENCES pubblr.objects(id),
);

CREATE TABLE pubblr.activities (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.objects(id),
    actor INT FOREIGN KEY REFERENCES pubblr.entities(id),
    obj INT FOREIGN KEY REFERENCES pubblr.entities(id),
    trg INT FOREIGN KEY REFERENCES pubblr.entities(id),
    result INT FOREIGN KEY REFERENCES pubblr.entities(id),
    origin INT FOREIGN KEY REFERENCES pubblr.entities(id),
    instrument INT FOREIGN KEY REFERENCES pubblr.entities(id),
);

type question_type as ENUM (
    'closed',
    'single-choice',
    'multi-choice',
);

CREATE TABLE pubblr.questions (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.activities(id),
    question_type question_type
);

CREATE TABLE pubblr.answers (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.entities(id),
    question INT FOREIGN KEY REFERENCES pubblr.questions(id),
);

CREATE TABLE pubblr.actors (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.objects(id),
    preferredUsername VARCHAR(255),
);

CREATE TABLE pubblr.users (
    id int PRIMARY KEY FOREIGN KEY REFERENCES pubblr.actors(id),
    salted_password_hash VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
);