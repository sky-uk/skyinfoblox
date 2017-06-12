if [ $# -lt 1 ]
then
    echo "Usage: delete_network.sh <an network object reference>\n"
    exit
fi
curl -w "\nStatus Code: %{http_code}\n" -k1 -u admin:infoblox -X DELETE https://h1infoblox.devops.int.ovp.bskyb.com/wapi/v2.3.1/$1
