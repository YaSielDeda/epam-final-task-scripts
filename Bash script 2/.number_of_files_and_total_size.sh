#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m'

function about() {

	printf 'This script makes two things, according to requirements for script 2.
	1) It sends to root number of files in /local/backups directory is more than X
	2) It sends to root total size of /local/backups directory is more than Y bytes
	3) It runs at the system start, works in detached mode and have protection from
	multiple running\n'

	exit 1

}

function number_of_files() {
	FILESINDIR=$(($(ls -l /mnt/raid/raid/local/backups | wc -l) - 1))

	if [ $FILESINDIR -gt $X ]; then	
#		printf "${YELLOW}X: $X${NC}\n"
#		printf "${YELLOW}Number of files in /mnt/raid/raid/local/backups: $FILESINDIR${NC}\n"

		echo "Number of files in /mnt/raid/raid/local/backups: $FILESINDIR" | mail -s "Num of files" root@db

#		printf "${GREEN}mail has been successfully sended to host!${NC}\n"
#	else
#		printf "${YELLOW}files in backups dir is less than $X${NC}\n"
	fi
}

function total_size() {
	BYTES=$(du -s /mnt/raid/raid/local/backups/ | grep -o '[0-9*]' | tr -d '\n')

	if [ $BYTES -gt $Y ]; then
#		printf "${YELLOW}Y: $X${NC}\n"
#               printf "${YELLOW}Total size of /mnt/raid/raid/local/backups directory: $BYTES${NC}\n"

		echo "Total size of /mnt/raid/raid/local/backups directory: $BYTES" | mail -s "Total size" root@db
#		printf "${GREEN}mail has been successfully sended to host!${NC}\n"
#	else
#		printf "${YELLOW}Total size of dir is less than $Y${NC}\n"
	fi
}

function setup() {
	source /home/dolgov/.conf.sh

	exec 9>/var/run/script2.lock
	if ! flock -n 9; then
#		printf "${RED}another instance is running!${NC}\n"
		exit 1
	fi
}

function execute() {
	number_of_files
	total_size

	rm -f /var/run/script2.lock

	exit 1
}

setup
execute
