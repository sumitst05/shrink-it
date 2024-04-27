.PHONY: tailwind-watch
tailwind-watch:
	tailwindcss -i ./web/static/index.css -o ./dist/style.css --watch

.PHONY: templ-watch
templ-watch:
	templ generate -watch -proxy=http://localhost:3000

.PHONY: watch
watch:
	air

.PHONY: dev
dev: 
	make -j3 tailwind-watch templ-watch watch
