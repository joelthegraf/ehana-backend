.PHONY := hooks

hooks:
	chmod 755 githooks/*
	cp githooks/* .git/hooks
