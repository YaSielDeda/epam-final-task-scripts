#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m'

function about() {
	printf 'This script makes two tasks, according to main task reuirements:
	1) it takes all content from Article table of Magazines_db
	and puts it into new file at the /local/files directory
        with name, which represents date of script launch,
	also, it runs 2) at the end.

        2) if number of files in /local/files directory is more than 3,
	script compress and moves files into /local/backup directory.
	
	To use this script for task 1, you must:
        1) have .cnf file in your home directory,
       	which contains login and password of user, who have select permission.
	(if you launch script as root, you need to have .cnf file in root directory)

	2) launch script in sudo\n'

	exit 1
}

function get_data() {
	
	LC_TIME='en_US.UTF-8'
	
	NAME=`date "+%c"`.txt

	CLEAN_NAME=`echo $NAME | tr " " "_"`

	mysql -h "localhost" "magazines_db" < "./.articles_select_all.sql" > "/mnt/lvm/local/files/${CLEAN_NAME}"

	if [ `ls -l /mnt/lvm/local/files/ | grep $CLEAN_NAME 2> /dev/null | wc -l` > 0 ]; then
		printf "${GREEN}file $NAME is successfully created!${NC}\n"
	else
		printf "${RED}error!${NC}\n"
	fi

	backup

	exit 1;
}

function backup() {

	printf "${YELLOW}files in local/files dir: $((`ls -l /mnt/lvm/local/files/ | wc -l` - 1))${NC}\n"

	LC_TIME='en_US.UTF-8'
	
	NAME=$(date "+%c").txt

	CLEAN_NAME=$(echo $NAME | tr " " "_")
	CLEAN_NAME=$(echo $CLEAN_NAME | tr ":" "-")

	if [ $((`ls -l /mnt/lvm/local/files/ | wc -l` - 1 )) > 3 ]; then
		tar zcf "archive_${CLEAN_NAME}.tar.gz" --absolute-names /mnt/lvm/local/files
		mv "archive_${CLEAN_NAME}.tar.gz" /mnt/raid/raid/local/backups
	fi

	if [ $(find /mnt/raid/raid/local/backups -name ${CLEAN_NAME} | wc -l) > 0 ]; then
		printf "${GREEN}backup is successfully created!${NC}\n"
	else
		printf "${RED}something gone wrong!${NC}\n"
	fi
}

while getopts ":a12" opt; do
	case ${opt} in
		1)
		       	get_data
		       	;;
		2) 
			backup
		       	;;
		a)
                        about
                        ;;
               	*)
			echo "Usage: cmd [-a12]"
			exit 1
			;;
	esac
done
