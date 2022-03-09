pip install virtualenv
# echo "1"
# # exec bash
# exec $SHELL -l;

# sleep 2
# echo "2"
rm -rf ./myenv
sudo apt install libpq-dev -y
bash -c """
virtualenv -p python3 ./myenv
. ./myenv/bin/activate
pip3 install psycopg2

"""