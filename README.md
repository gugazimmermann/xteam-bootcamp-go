# X-Team Go Bootcamp

https://www.linkedin.com/in/juliano-silva-de-souza/
https://github.com/julianosouza/go-crash-course

## First Project

Build an HTTP API that's responsible for handling a phone agenda (i know it's obvious, but the nuances on building the project are what really matters).

The http server should contain:
- An endpoint for pushing new contacts
- An endpoint for editing contact information
- An endpoint for deleting a contact
- An endpoint for searching a contact by it's id
- An endpoint for searching contacts by a part of their name
- An endpoint that lists all contacts
- The http service should be configurable through flags on startup (timeouts and port to use)
- Log messages should be written for each endpoint hit, with the response status code and the time it took to fulfill the request
- If an error occurs, a log message should be written with the response status code and a helpful error message, to help an engineer troubleshoot the issue
Service and host metrics should be collected. I suggest using Prometheus (https://prometheus.io/docs/guides/go-application/)
- The application should have a reasonable test coverage, preferably above 70%
- The application should have end-to-end tests (this is a good way to try out the http client)
- The application should contain a buildable Dockerfile (https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e) -- care on this, as using plainly the scratch image might hinder you from making https requests. Not that this will impact our example, but something to always take care into the future
- It would be nice for the application to have some type of storage to persist the data. I'll leave this open, feel free to pick any type of storage you want :slightly_smiling_face:

## Resources

- https://github.com/dariubs/GoBooks#starter-books
- https://peter.bourgon.org/go-for-industrial-programming/
- https://dave.cheney.net/practical-go
- https://dave.cheney.net/high-performance-go
- https://medium.com/golangspec/interfaces-in-go-part-iii-61f5e7c52fb5
- https://golang.org/doc/effective_go.html
- https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
- https://gist.github.com/ghiden/0ff00fe355a24e790512
- https://github.com/golang-standards/project-layout
- https://github.com/golang/go/wiki/CodeReviewComments
