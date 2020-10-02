I don't know go.

### Setup

Setup the secrets that you will need
```json
{
  "discord": {
    "api_key": "YOUR DISCORD BOT TOKEN"
  },
  "twitter": {
    "screen_name": "",
    "consumer_key": "",
    "consumer_secret": "",
    "access_token": "",
    "access_secret": ""
  }
}
```

Required modules will be installed via `go.mod`. The token is put in a file that is specifically ignored by git so that we don't run the risk of commiting an access token, like some kind of hacked lamer. When you invoke the program, you can either specify the token on the command line via `-t TOKEN` or have the program read the token from a file with `-f FILENAME`. The latter is much better for systems that you are sharing (like a colocated linux box that all your friends have shells on), because the token would appear in the process list with `-t`.

### Run...?
```bash
make
./goelo2
```

### Stripping the binary down to size
Are you insane about binary size? Cool, install and run `upx`. Warning: this shit is slow as fuck.

```bash
brew install upx
go get github.com/pwaller/goupx

make strip
```

### Cross compiling for linux
```bash
make linux
make strip #if you want a smaller binary (you do)
```
