I don't know go.

### Setup
```bash
# put the discord token into env, and DON'T CHECK THAT SHIT IN.
echo TOKEN_FROM_DISCORD > env
chmod 600 env

go get -u github.com/bwmarrin/discordgo
go get -u github.com/davecgh/go-spew/spew
```

The token is put in a file that is specifically ignored by git so that we don't run the risk of commiting an access token, like some kind of hacked lamer. When you invoke the program, you can either specify the token on the command line via `-t TOKEN` or have the program read the token from a file with `-f FILENAME`. The latter is much better for systems that you are sharing (like a colocated linux box that all your friends have shells on), because the token would appear in the process list with `-t`.

### Run...?
```bash
make
./goelo -f env
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
