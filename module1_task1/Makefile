TITLE = ${POST_TITLE}
build:
	hugo > /dev/null 2>&1
post:
	grep -rl "ChangeMe" archetypes/ | xargs sed -i 's/ChangeMe/$(TITLE)/g'
	hugo new posts/${POST_NAME}.md
	grep -rl "$(TITLE)" archetypes/ | xargs sed -i 's/$(TITLE)/ChangeMe/g'
clean:
	rm -rf ./dist
