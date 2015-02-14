# notbot
an IRC notification bot

## Notifiers

Add new sources for IRC notifications by implementing the Notifier interface:

```
type Notifier interface {
        Notifications(*rest.Request) []string
}
```

There are currently three Notifiers: GitHub, GitLab and Sensu, each living under a route of the same name:

```
handler.SetRoutes(
        &rest.Route{"POST", "/github", Post},
        &rest.Route{"POST", "/gitlab", Post},
        &rest.Route{"POST", "/sensu", Post},
)
```

