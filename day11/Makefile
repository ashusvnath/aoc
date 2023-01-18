DIR := $(shell pwd)
NAME:= $(shell basename $(DIR))

all : clean build tests

.PHONY: clean
clean:
	@echo Cleaning targets
	@rm -f $(NAME).out
	@echo DONE

build : *.go
	@echo Building
	@go build -o $(NAME).out ./...
	@echo DONE

tests : build
	@echo Running against test input
	./$(NAME).out -f test.txt -v
	@echo DONE
