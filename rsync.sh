#!/bin/sh
rsync -avz --exclude '*.swp' --delete $1 $2
