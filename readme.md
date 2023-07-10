# Pubblr

Pubblr is an open-source, federated blogging platform built on the ActivityPub protocol.  This repo houses the Pubblr
server; the front-end will be housed separately.

The server is currently in **pre-Alpha** - that is, it is not yet functional.  The roadmap below details what features
are present, and what remains to be completed before the project is in a minimally viable state.

# Alpha Roadmap

These are the components that need to be minimally functional before Alpha.  They may not be working perfectly; once
Alpha phase is reached, focus will shift to testing, bug fixes and incremental improvements.

- [x] ActivityStreams type hierarchy
- [x] User registration
- [x] User authentication
- [ ] User authorization
- [x] Outboxes
- [x] Inboxes
- [ ] Delivery
    - [x] Local Delivery
    - [ ] Remote Delivery
- [x] Object Retrieval
- [ ] Activity Processing
    - [ ] Client-to-Server
        - [x] Create
        - [ ] Update
        - [ ] Delete
        - [ ] Follow
        - [ ] Add
        - [ ] Remove
        - [ ] Like
        - [ ] Block
        - [ ] Undo
    - [ ] Server-to-Server
        - [ ] Create
        - [ ] Update
        - [ ] Delete
        - [ ] Follow
        - [ ] Accept
        - [ ] Reject
        - [ ] Add
        - [ ] Remove
        - [ ] Like
        - [ ] Announce
        - [ ] Undo
- [ ] Media Uploading