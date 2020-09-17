# DEPLOY BOT


## Installation and USAGE:
TODO: .....


### AUTHENTICATION:
```console
$ export BOT_TOKEN=<yout bot token here>
$ export BOT_WEBHOOK_URL=<your webhook url>
$ export BOT_API_URL=<your bot api url>
```
Or for persistence authentication,open `.bashrc` file and put the lines above on it if your are running `linux` of course.

### Run using Docker
TODO: .....


## ROADMAP
- [X] setup webhook.
- [X] get updates from webhook and parse the json data.
- [X] command handler 
- [ ] setup a bot log file where all logs are gonna live.
- [ ] figure out how to work with deployment (drone api etc ...)
- [ ] introduce services package where you gonna handle deployment and get pods status and logs.
- [ ] .....
