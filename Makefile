GVM_BINARY=epub-scrapper.exe

build:
	@echo Building App binary...
	set CGO_ENABLED=0 && go build -o ${GVM_BINARY} .
	@echo Done!