all:
cd espeak && go get -x .

install:
cd espeak && go install -x

fmt:
cd espeak && go fmt .
