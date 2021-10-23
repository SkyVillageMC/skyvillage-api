migrate:
	go run github.com/prisma/prisma-client-go migrate dev --name init

format:
	go run github.com/prisma/prisma-client-go format

studio:
	go run github.com/prisma/prisma-client-go studio