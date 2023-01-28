DIR := $(shell pwd)
NAME:= $(shell basename $(DIR))

all : clean build tests

.PHONY: clean
clean:
	@echo Cleaning targets
	@rm -f $(NAME).out
	go clean -x
	@echo DONE

unit:
	@echo Running unit tests
	@go test -tags testing -v ./...
	@echo DONE

build : *.go unit
	@echo Building
	@go build -o $(NAME).out $(NAME)
	@echo DONE

tests : build
	@echo Running against test input
	./$(NAME).out -f test.txt -v
	@echo DONE
