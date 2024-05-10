smtp: main.go
	go build -o $@ $^

.PHONY: clean
clean:
	$(RM) smtp

