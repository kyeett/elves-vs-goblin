# Elves vs Goblins

A command line game.

```sh
docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
```

## Features

- Nothing, so far

## Todo

- server: handle multiple users
- server: multiple users testcase
- server: clean up usage according to plan
- server: send only state needed or timeout
- client: clean up usage according to plan
- client: Show multiple users

* Change the handling of buffer/writes in views

* Logrus for logging
* Logrus over NATS
* In game chat
* Twitch support
* Admin console
* Everything else

## Done

- ~~Navigation (WASD)~~
- ~~Propagate world state~~
- ~~rename xxx.NewXXX to xxx.New~~
- ~~Connect to NATS~~
- ~~Transmit player state~~
- ~~Draw world~~
- ~~Add a player type~~
- ~~Draw player~~
- ~~bug: ctrl+c in client breaks terminal layout~~

## Bugs

- Server sends empty player ID first time

## References
