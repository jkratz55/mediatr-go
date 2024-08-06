package mediatr

// CommandHandler is a type that handles a particular type of command. A
// CommandHandler is meant to execute a command (action) and returns an
// error if the CommandHandler doesn't handle the command successfully.
type CommandHandler[T any] interface {
	Handle(command T) error
}

type CommandHandlerFunc[T any] func(command T) error

func (c CommandHandlerFunc[T]) Handle(command T) error {
	return c(command)
}

// RequestHandler is a type that handles a particular type of request. A
// RequestHandler is meant to retrieve data and return it to the caller.
type RequestHandler[T any, R any] interface {
	Handle(query T) (R, error)
}

type RequestHandlerFunc[T any, R any] func(query T) (R, error)

func (q RequestHandlerFunc[T, R]) Handle(query T) (R, error) {
	return q(query)
}

// NotificationHandler is a type that handles a particular type of notification.
// A NotificationHandler is meant to handle a notification (event) and doesn't
// return anything.
type NotificationHandler[T any] interface {
	Notify(notification T)
}

type NotificationHandlerFunc[T any] func(notification T)

func (n NotificationHandlerFunc[T]) Notify(notification T) {
	n(notification)
}
