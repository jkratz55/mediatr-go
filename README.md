# Mediatr

Mediatr is a lightweight zero dependency library implementing the Mediator Pattern. The Mediator pattern promotes lose coupling between components and objects by communicating through a mediator object, rather than directly. This can help reduce or eliminate having tangled web of dependencies as objects/components are coupled to the mediator instead of each other. However, this is not without its tradeoffs. The Mediator pattern can also enable bad development practices, and poor code, structure, design, etc. Additionally, it adds a layer of indirection and abstraction that often isn't needed. 

This project is mostly a toy project to see if I could port MediatR concepts from my Java/Kotlin [library](https://github.com/jkratz55/spring-mediatR). Both are inspired from the C# library MediatR. I would not recommend using this library for building production applications unless you were absolutely sure this pattern would be useful.

## Core Concepts

Mediatr has three core types/concepts

|               |                                                                                                                                                                                                                                                                                          |
|---------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Commands      | Commands represent an intention to perform an action or a change in the system. Commands typically result in state changes. Commands and handled by CommandHandlers                                                                                                                      |
| Requests      | Requests can be used to either retrieve data with no state change or perform an action that changes the state of the system, but also returns data. Generally speaking Requests are much more generalized and used when you need to return data. Requests are handled by RequestHandlers |
| Notifications | Notifications are events/data published to other components of an application to notify them some event has occurred. Notifications are handled by NotificationHandlers.                                                                                                                 |

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/jkratz55/mediatr-go"
)

type CreateUserCommand struct {
	Username string
	Password string
}

type CreateUserCommandHandler struct {
}

func (c CreateUserCommandHandler) Handle(command CreateUserCommand) error {
	fmt.Println("CreateUserCommandHandler")
	return nil
}

type GetUserQuery struct {
	Username string
}

type GetUserQueryHandler struct {
}

func (g GetUserQueryHandler) Handle(query GetUserQuery) (int, error) {
	return 42, nil
}

type UserCreatedEvent struct {
	Username string
}

func main() {
	createUserCommandHandler := &CreateUserCommandHandler{}
	mediatr.RegisterCommandHandler[CreateUserCommand](createUserCommandHandler)

	getUserQueryHandler := &GetUserQueryHandler{}
	mediatr.RegisterQueryHandler[GetUserQuery, int](getUserQueryHandler)

	mediatr.RegisterNotificationHandler[UserCreatedEvent](mediatr.NotificationHandlerFunc[UserCreatedEvent](func(notification UserCreatedEvent) {
		fmt.Println("UserCreatedEvent1")
	}))
	mediatr.RegisterNotificationHandler[UserCreatedEvent](mediatr.NotificationHandlerFunc[UserCreatedEvent](func(notification UserCreatedEvent) {
		fmt.Println("UserCreatedEvent2")
	}))

	mediatr.RegisterNotificationHandler[string](mediatr.NotificationHandlerFunc[string](func(notification string) {
		fmt.Println("UserCreatedEvent2")
	}))

	if err := mediatr.SendCommand(CreateUserCommand{
		Username: "newdude",
		Password: "password",
	}); err != nil {
		panic(err)
	}

	result, err := mediatr.SendRequest[GetUserQuery, int](GetUserQuery{
		Username: "newdude",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	mediatr.SendNotification(UserCreatedEvent{
		Username: "mrnewdude",
	})
	mediatr.SendNotification(42)

	time.Sleep(1 * time.Second)
}
```

Output:

```text
CreateUserCommandHandler
42
UserCreatedEvent2
UserCreatedEvent2
time=2024-08-06T17:01:42.469-04:00 level=ERROR msg="No registered NotificationHandler for notification type, notification will be dropped and ignored" notificationType=int

```