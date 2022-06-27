package users

import "github.com/labstack/echo/v4"

// Handlers
type GetHandler interface {
	Handle(echo.Context) error
}

type PostHandler interface {
	Handle(echo.Context) error
}

type PutHandler interface {
	Handle(echo.Context) error
}

type DeleteHandler interface {
	Handle(echo.Context) error
}

// Requests
type PostRequest interface {
	Validate() error
	Bind(echo.Context) error
	GetName() string
	GetEmail() string
}

type GetRequest interface {
	Bind(echo.Context) error
	Validate() error
	GetId() int
}

type PutRequest interface {
	Validate() error
	Bind(echo.Context) error
	GetID() int
	GetName() string
	GetEmail() string
}

type DeleteRequest interface {
	Bind(echo.Context) error
	Validate() error
	GetId() int
}
