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

