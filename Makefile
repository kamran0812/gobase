build:
	@go build -o bin/app .

run: build
	@./bin/app

temp:
	templ generate --watch --proxy=http://localhost:3000

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch