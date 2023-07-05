runPort:
	./Dictionary -port=:8083
runDefault:
	./Dictionary
runEnv:
	./Dictionary -port=.env
builddocker:
	docker build -t wolfxxxz/dictionary_app .
rundocker:
	docker run -it -p 8081:8081 wolfxxxz/dictionary_app
rundockerBash:
	docker run -it -p 8081:8081 wolfxxxz/dictionary_app bash