# entnum

entviz is an ent extension that provides functions that returns enums.

> ent/todo/todo.go

```go
// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusOngoing   Status = "ongoing"
	StatusCompleted Status = "completed"
	StatusLater     Status = "later"
)
```

> ent/entnum.go

```go
func AllTodoStatus() []todo.Status {
	return []todo.Status{
		todo.StatusOngoing,
		todo.StatusCompleted,
		todo.StatusLater,
	}
}
```
