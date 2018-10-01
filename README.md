# Fortnite API

A simple to use package for interacting with Fortnite API.

## Install

```
$ go get github.com/jryd/fortnite
```

## API

### FIRST THINGS FIRST

To access the Fortnite API, you need to have an account on Epic Games. After that you need to get 2 headers that your client uses to access Fortnite.

You can get these headers by using a tool to capture incoming and outgoing HTTP traffic. The steps below are for the Fiddler tool that I used.

*   Install & Open Fiddler
*   In Tools -> Options -> HTTPS, Select Capture HTTPS Connects
*   After that start the Epic Games launcher
*   You will see a request with _/account/api/oauth/token_. Click on it -> Click Inspectors to view the headers (you want to grab the long string in the Authorization header, without the basic part - i.e header is basic abcd... you only need the abcd... bit) => **This header is your Client Launcher Token**
*   Launch Fortnite
*   You will see again a request with _/account/api/oauth/token_. Click on it -> Click Inspectors to view the headers (you want to grab the long string in the Authorization header, without the basic part - i.e header is basic abcd... you only need the abcd... bit) => **This header is your Fortnite Client Token**

---

### USAGE

```go
fortniteClient := fortnite.NewClient("email address", "password", "client launcher token", "fortnite client token")
```

---

### METHODS

```go
fortniteClient.Login()

fortniteClient.Lookup("jryd")
fortniteClient.CheckPlayer("jryd")
fortniteClient.GetStatsBR("jryd", "pc")
fortniteClient.GetStatsBRFromID("12345", "pc")
fortniteClient.GetFortniteNews()
fortniteClient.CheckFortniteStatus()
fortniteClient.GetFortnitePVEInfo("en")
fortniteClient.GetStore("en")

fortniteClient.KillSession()
```

More information on the mentods can be found in the [GoDoc](https://godoc.org/github.com/jryd/fortnite).