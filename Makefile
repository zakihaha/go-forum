export MYSQL_URL ='mysql://root:@tcp(localhost:3306)/fastcampus_forum'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq ${name}

migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down