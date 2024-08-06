package mediatr

import (
	"fmt"
	"log/slog"
	"reflect"
)

var (
	commandHandlers = make(map[reflect.Type]interface{})
	queryHandlers   = make(map[reflect.Type]interface{})
	eventHandlers   = make(map[reflect.Type][]interface{})
)

func init() {
	commandHandlers = make(map[reflect.Type]interface{})
	queryHandlers = make(map[reflect.Type]interface{})
	eventHandlers = make(map[reflect.Type][]interface{})
}

// RegisterCommandHandler registers a CommandHandler for a specific type. If a
// type has been registered already, the new handler will overwrite the existing.
// A type can only have a single CommandHandler registered.
func RegisterCommandHandler[T any](handler CommandHandler[T]) {
	commandType := reflect.TypeOf(*new(T))
	commandHandlers[commandType] = handler
}

// RegisterQueryHandler registers a RequestHandler for a specific type. If a
// type has been registered already, the new handler will overwrite the existing.
// A type can only have a single RequestHandler registered.
func RegisterQueryHandler[T any, R any](handler RequestHandler[T, R]) {
	queryType := reflect.TypeOf(*new(T))
	queryHandlers[queryType] = handler
}

// RegisterNotificationHandler registers a NotificationHandler for a specific type.
// A notification type can have multiple handlers registered and each is invoked
// on a separate goroutine when SendNotification is called.
func RegisterNotificationHandler[T any](handler NotificationHandler[T]) {
	eventType := reflect.TypeOf(*new(T))
	eventHandlers[eventType] = append(eventHandlers[eventType], handler)
}

// SendCommand sends a command to the registered CommandHandler. If no handler
// is registered for the command type, an error is returned.
func SendCommand[T any](command T) error {
	commandType := reflect.TypeOf(command)
	handler, ok := commandHandlers[commandType]
	if !ok {
		return fmt.Errorf("no command handler registered command type %s", commandType)
	}

	commandHandler, ok := handler.(CommandHandler[T])
	if !ok {
		return fmt.Errorf("handler type mismatch for command type %s", commandType)
	}
	return commandHandler.Handle(command)
}

// SendRequest sends a request to the registered RequestHandler. If no handler
// is registered for the request type, an error is returned.
func SendRequest[T any, R any](query T) (R, error) {
	var zero R
	queryType := reflect.TypeOf(query)
	handler, ok := queryHandlers[queryType]
	if !ok {
		return zero, fmt.Errorf("no query handler registered for query type %s", queryType)
	}

	queryHandler, ok := handler.(RequestHandler[T, R])
	if !ok {
		return zero, fmt.Errorf("handler type mismatch for query type %s", queryType)
	}
	return queryHandler.Handle(query)
}

// SendNotification sends a notification to all registered NotificationHandlers.
func SendNotification[T any](notification T) {
	notificationType := reflect.TypeOf(notification)
	handlers, ok := eventHandlers[notificationType]
	if !ok {
		logger.Error("No registered NotificationHandler for notification type, notification will be dropped and ignored",
			slog.String("notificationType", notificationType.String()))
		return
	}

	for _, handler := range handlers {
		go func() {
			notificationHandler, ok := handler.(NotificationHandler[T])
			if !ok {
				logger.Error("NotificationHandler type mismatch for notification type, notification will be dropped and ignored",
					slog.String("notificationType", notificationType.String()))
				return
			}
			notificationHandler.Notify(notification)
		}()
	}
}
