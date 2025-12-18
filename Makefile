.PHONY: install

install:
	go build -o cfget
	mkdir -p ~/.cfget/bin
	mv cfget ~/.cfget/bin
	if [ -f ~/.bashrc ]; then ./writerc.sh; fi
	if [ -f ~/.bash_profile ]; then ./writepr.sh; fi
