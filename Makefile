app:
	docker build -t ciao-universe -f _build/app/Dockerfile .

up:
	docker run -p 8080:8080 ciao-universe