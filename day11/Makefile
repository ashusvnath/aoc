DIR := $(shell pwd)
NAME:= $(shell basename $(DIR))

all : clean build tests

.PHONY: clean
clean:
	@echo Cleaning targets
	@rm -f $(NAME).out
	@echo DONE

unit:
	@echo Ruinning unit tests
	@go test -v
	@echo DONE

build : *.go unit
	@echo Building
	@go build -o $(NAME).out ./...
	@echo DONE

tests : build
	@echo Running against test input
	./$(NAME).out -f test.txt -v
	@echo DONE
