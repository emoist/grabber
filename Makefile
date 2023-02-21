#================================
#== GOLANG ENVIRONMENT
#================================
GO := go

install:
	${GO} get .

dev:
	${GO} run main.go

test:
	${GO} test -v

format:
	${GO} fmt ./...
