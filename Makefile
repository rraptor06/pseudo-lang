##
## EPITECH PROJECT, 2024
## Pseudo Lang Compiler
## File description:
## The main Makefile of our project
##

NAME		=	plc

SRCDIR		=	./src

SRCS		+=	main.go

FLAGS		=

.PHONY: all clean re

all: $(NAME)
	@echo -e "\033[1;33m$(NAME) compiled.\033[0m"

$(NAME):
	@cd $(SRCDIR) && go build $(SRCS)
	@mv $(SRCDIR)/main $(NAME)

clean:
	@make fclean --no-print-directory -C ./template/
	@rm -rf *.log
	@rm -rf $(NAME)
	@echo -e "\033[1;31mPLC Project cleaned.\033[0m"

re:	clean all

run: re
	@echo -e "\033[1;32mRun $(NAME).\033[0m"
	@./$(NAME)

doc:
	@cd src ; google-chrome http://localhost:6060/pkg/pseudo-lang/ ; godoc -http=:6060