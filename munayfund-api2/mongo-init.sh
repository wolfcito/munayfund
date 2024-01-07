#!/bin/bash
set -e

echo "Executing mongo-init.sh script..."

mongo <<EOF
db = db.getSiblingDB('admin')

db.createUser({
  user: '${ROOT_USERNAME}',
  pwd: '${ROOT_PWD}',
  roles: [{ role: 'root', db: 'admin' }]
});


db = db.getSiblingDB('munayFundDB')

db.createUser({
  user: '${MUNAY_USERNAME}',
  pwd: '${MUNAY_PASSWORD}',
  roles: [{ role: 'readWrite', db: 'munayFundDB' }],
});

db.createCollection('Users')
db.createCollection('Projects')

EOF
