# Counter API
This repository includes consumer driven contract test for provider, unit test and counter api.

## Consumer API
- https://github.com/azumber/basic-counter

## Goal
We aim to verify the consumer test created by the project given url, in this repository. While doing this verification, we use the credentials of the Pact-Broker that published consumer test.

## Endpoints
```go
We use GET method call for all endpoints.
app.Get("/display", DisplayHandler)
app.Get("/increase", IncreaseHandler)
app.Get("/decrease", DecreaseHandler)
```
## Responses
```javascript
// /display
{
    counter: 0
}
// /increase
{
    counter: 1
}
// /decrease
{
    counter: -1
}
```