# JSON server in Go

## Task
Learn about Go's http capabilities by building a simple API.

## About
An API where users can submit ideas for new features for IDEs. The submission is a POST
request with a json body on the form
{ 
    "ide": "<the IDE in question>",
    "idea": "<the suggestion>"
}

It is handel by Go using the struct
struct {
    ID int64,
    IDE string,
    Idea string
}

Yeah, I'm having some fun here ;) 

This struct is also used to communicate with the sqlite3 database, which give the program persistent storage.

### TODO
* Add test to handlers
* Add form validation
* Add TLS support