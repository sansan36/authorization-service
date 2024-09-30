install:
	bash install.sh

# generate proto
buf-gen:
	buf generate proto

buf-update:
	buf dep update;

# run
up:
	docker-compose up --build;
	
down:
	docker-compose down

run:
	go run app/main.go

# test
test-all:
	cd repository/; go test -v;
	cd handler/; go test -v;

test-repo:
	cd repository/; go test -v

test-handler:
	cd handler/; go test -v

# Generate mocks with mockgen
# mock-gen:
# 	mockgen -source=repository/role.go -destination=repository/mock/role_mock.go -package=mock;
# 	mockgen -source=repository/user.go -destination=repository/mock/user_mock.go -package=mock;
# 	mockgen -source=core/caller/caller.go -destination=core/caller/caller_mock.go -package=caller;