run: build
	go run main.go


build: main.js

main.js: src/Main.elm
	elm make src/Main.elm --output=main.js
