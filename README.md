# RKL.tv \[Mod Bot\]
## Abstract
This is an early version of the twitch irc bot I use in my streams under [rkl.tv](https://rkl.tv). It is designed to
interact with the streamed games via Twitch chat.

## Support
It would be great if you support me on [twitch](https://twitch.tv/rkl85) with a follow or subscription, the goal is to gather more and more
spare time to work on projects like this.

## Platform
This version only works under a Microsoft Windows OS (tested with win 10 pro).

## Configuration
Ensure that you have the file `%userprofile%\twitch_connector.ini` with your configured content:
```ini
#
# Template for %userprofile%\twitch_connector.ini
#

[twitch]
TwitchIrcUsername = "rkl85"
TwitchIrcAuthentication = "oauth:123456"
TwitchIrcChannel = "rkl85-Channel"

[l4d2]
BoostSeconds = 30
```

## Build instructions
You need the latest [go](https://golang.org/doc/install) version for windows as well as [Mingw64](http://mingw-w64.org/doku.php).
If all is set up well, you can simply run `wire gen && go build` under the source tree. If you wanna run the tests use `go test ./...` under the
source tree. It is important for the tests, that your `steam process` is running (`steam.exe`). We use the steam process
for some basic tests like access process memory and so on. We simply abuse steam here, because to use this bot, it is highly probable,
that steam is running on your machine.

## Supported games and commands

### Left 4 Dead 2
Tested with the official [steam version](https://store.steampowered.com/app/550/Left_4_Dead_2) (non beta) for windows.

#### Commands
* `$help`    shows usage
* `$boost`   give `god mode` to team for a configured period (default 30 seconds).
