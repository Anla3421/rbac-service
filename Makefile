# $ make <.PHONY>
origin = if err := r.Run("localhost:5002"); err != nil
alter = if err := r.Run(":5002"); err != nil

.PHONY: prod
prod:
	cp ./makeConfigs/prod/database.json ./configs/database.json
	cat ./main.go | sed -e 's/${origin}/${alter}/g' | tee ./main.go
.PHONY: dev
dev:
	cp ./makeConfigs/dev/database.json ./configs/database.json
	cat ./main.go | sed -e 's/${alter}/${origin}/g' | tee ./main.go