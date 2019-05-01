all: plugin app

LuaPlugin: LuaPlugin.go
	go build --buildmode=plugin $<
HTTPPlugin:HTTPPlugin.go
	go build --buildmode=plugin $<

plugin:LuaPlugin HTTPPlugin

app: app.go
	go build -o app app.go
clean:
	rm app *.so
