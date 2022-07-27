db-migrate:
	docker exec go-mysql sh -c "cat /migrations/jira.sql | mysql --user=root --password=root"
