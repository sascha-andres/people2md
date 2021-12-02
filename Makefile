.DEFAULT_GOAL := help

.PHONY: help

groups: ## create groups.json to have a list of groups
	goobook dump_groups > groups.json

contacts: ## create contacts.json to have a list of contacts
	goobook dump_contacts > contacts.json

help: ## Display help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'