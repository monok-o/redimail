# REDIMAIL

Simple service to create "mailto" clickable link to ditribute.
Useful for "mail storm" activism campaign.

## How to run ?
### Developement

1. Install golang
2. Clone the repo
3. Set your environment variables in `.env` file
4. run `go mod tidy`
5. run `go run main.go`

### Production

1. Install docker
2. Install buildx from docker for multi-architecture build
3. Clone the repo
4. run `docker buildx build --platform linux/arm64,linux/amd64 -t <your org>/redimail .`

## Story

Initially made in 10 minutes for one of the campaign of Fridays For Future international , I decided to make it better so it can be used by any organization. ~~useful for annoying politicians~~
