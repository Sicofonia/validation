build:
	hugo > /dev/null 2>&1
post:
	grep -rl "ChangeMe" archetypes/ | xargs sed -i 's/ChangeMe/${POST_TITLE}/g'
	hugo new posts/${POST_NAME}.md
clean:
	rm -rf ./dist
