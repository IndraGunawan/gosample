.PHONY: all

REGISTRY  = registry.bukalapak.io/bukalapak
DDIR      = deploy
ODIR      = $(DDIR)/_output
NOCACHE   = --no-cache
VERSION  ?= $(shell git show -q --format=%h)
SERVICES ?= web
ENV      ?= default
FILE     ?= deployment

export VAR_KUBE_NAMESPACE ?= default
export VAR_CONSUL_PREFIX ?= gosample

dep:
	dep ensure

all:
	consul compile build push deployment

test:
	go test ./...

compile:
	@$(foreach var, $(SERVICES), GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/$(var)/bin/$(var) app/$(var)/main.go;)

$(ODIR):
	mkdir -p $(ODIR)

consul: $(ODIR)
	@wget https://releases.hashicorp.com/envconsul/0.6.2/envconsul_0.6.2_linux_amd64.tgz
	@tar -xf envconsul_0.6.2_linux_amd64.tgz -C $(ODIR)/
	@rm envconsul_0.6.2_linux_amd64.tgz

build:
	@$(foreach var, $(SERVICES), docker build $(NOCACHE) -t $(REGISTRY)/gosample/$(var):$(VERSION) -f ./deploy/$(var)/Dockerfile .;)

push:
	@$(foreach var, $(SERVICES), docker push $(REGISTRY)/gosample/$(var):$(VERSION);)

deployment: $(ODIR)
ifeq ($(ENV),default)
	kubelize deployment -v $(VERSION) $(SERVICES)
else
	kubelize deployment -e $(ENV) -v $(VERSION) $(SERVICES)
endif

$(ENV):
	@$(foreach var, $(SERVICES), kubectl replace -f $(ODIR)/$(var)/$@/deployment.yml;)


checkenv:
ifndef ENV
	$(error ENV must be set.)
endif

deploy: checkenv $(ODIR)
	@$(foreach svc, $(SERVICES), \
		echo deploying "$(svc)" to environment "$(ENV)" && \
		! kubelize genfile --overwrite -c ./ -s $(svc) -e $(ENV) deploy/$(svc)/$(FILE).yml $(ODIR)/$(svc)/ || \
		cat $(ODIR)/$(svc)/$(FILE).yml || \
		kubectl replace -f $(ODIR)/$(svc)/$(FILE).yml || kubectl create -f $(ODIR)/$(svc)/$(FILE).yml ;)
