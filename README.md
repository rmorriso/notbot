# notbot
an IRC notification bot

---

notbot accepts POST requests from clients that wish to send notifications to an IRC channel.

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

## TODO

All the client to specify the target channel in the POST request.

