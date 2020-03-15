#!/bin/bash
for file in ./*.zip ; do
	fname=$(basename "$file")
	#echo "Name is $fname"
done

read -p "Copy the the live-blog package $fname (y/n)? " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Nn]$ ]]
then
	echo "Copy canceled"
	exit 0
fi

echo "Start to upload $fname"
rsync -avz $fname igor@invido.it:/home/igor/app/go/live-blog/zips

echo "That's all folks!"
