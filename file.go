package ArchiverEngine

import "log"

type Archiver struct {
	log                *log.Logger
	EventsHandler      Events
}
