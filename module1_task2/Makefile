TITLE = ${POST_TITLE}

build: ## Builds the Awesome Inc. website for publishing
	hugo > /dev/null 2>&1
post:  ## Creates a new post, takes parameters POST_NAME and POST_TITLE
	grep -rl "ChangeMe" archetypes/ | xargs sed -i 's/ChangeMe/$(TITLE)/g'
	hugo new posts/${POST_NAME}.md
	grep -rl "$(TITLE)" archetypes/ | xargs sed -i 's/$(TITLE)/ChangeMe/g'
clean: ## Deletes the existing built website
	rm -rf ./dist
help:  ## Shows this help usage
	@echo "build: Builds the Awesome Inc. website for publishing"
	@echo "post: Creates a new post, takes parameters POST_NAME and POST_TITLE"
	@echo "clean: Deletes"
	@echo "help: Shows this help usage"
