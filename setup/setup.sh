# set env variables
export ECHO_CONTACTS_DB_NAME=echo_contacts_dev # the db name should be same as in the create_db.sql
export DB_USER="root" # change these as per your mysql db
export DB_PASSWORD="" # change these as per your mysql db

echo $ECHO_CONTACTS_DB_NAME
echo $DB_USER
echo $DB_PASSWORD

#Create database
mysql -u $DB_USER -p < ./setup/create_db.sql
if (( $? != 0 )); then
   echo "Database creation failed"
   #exit 1
else
echo "Created table:" $ECHO_CONTACTS_DB_NAME
fi