I don't know go.

### Setup
```bash
# put the discord token into env, and DON'T CHECK THAT SHIT IN.
echo TOKEN_FROM_DISCORD > env

go get -u github.com/bwmarrin/discordgo
go get -u github.com/davecgh/go-spew/spew
```

### Run...?
```bash
make
./goelo -t $(cat env)
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
