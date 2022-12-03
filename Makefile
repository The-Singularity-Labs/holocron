#
# Holocron
#
HOLOCRON_NAME?=
HOLOCRON_GATEKEEPER?=
HOLOCRON_TREASURE?=
HOLOCRON_ASCERTAINMENT?=
HOLOCRON_SALT?=foobar
HOLOCRON_OUTDIR?=.build
HOLOCRON_APPDIR?=app


ifeq ($(OS),Windows_NT)
	uname_S := Windows
else
	uname_S := $(shell uname -s)
endif

.PHONY: app

all: app

forge: $(HOLOCRON_OUTDIR)/forge
	$(HOLOCRON_OUTDIR)/forge -n "$(HOLOCRON_NAME)" -g "$(HOLOCRON_GATEKEEPER)" -a "$(HOLOCRON_ASCERTAINMENT)" -t "$(HOLOCRON_TREASURE)" -s "$(HOLOCRON_SALT)" -o "$(HOLOCRON_OUTDIR)"

pick: $(HOLOCRON_OUTDIR)/pick
	$(HOLOCRON_OUTDIR)/pick -name=$(KEY_NAME) -seed="$(KEY_SEED)"

wasm:
	GOOS=js GOARCH=wasm go build -o $(HOLOCRON_APPDIR)/holocron.wasm cmd/wasm/main.go 

app: wasm
	python -m http.server 8000 --directory app

$(HOLOCRON_OUTDIR)/forge:
	go build -o $(HOLOCRON_OUTDIR)/forge cmd/forge/main.go 

$(HOLOCRON_OUTDIR)/pick:
	go build -o $(HOLOCRON_OUTDIR)/pick cmd/pick/main.go 