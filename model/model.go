package model // import "github.com/ipsusila/iotid/mode"

import "github.com/satori/go.uuid"

type Identifiable interface {
	Identifier() uuid.UUID
}

type Command interface {
	Identifiable
}

type Response interface {
	Identifiable
}
