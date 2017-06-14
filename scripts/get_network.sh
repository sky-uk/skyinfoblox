#!/usr/bin/env bash
if [ $# -lt 1 ]
then
    echo "Usage: get_network.sh <a valid obj ref>\n";
    exit
fi
curl -w "\nStatus Code: %{http_code}\n" -k1 -u admin:infoblox -X GET https://h1infoblox.devops.int.ovp.bskyb.com/wapi/v2.3.1/$1 \
-d _return_fields=network,network_view,comment,authority,disable,discover_now_status,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,\
ipv4addr,lease_scavenge_time,\
netmask,network_container,\
options,recycle_leases,update_dns_on_lease_renewal,use_authority,use_blackout_setting,\
use_discovery_basic_polling_settings,use_discovery_basic_polling_settings,\
use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,\
use_enable_ifmap_publishing,use_ignore_dhcp_option_list_request,\
use_ignore_id,use_ipam_email_addresses,use_ipam_threshold_settings,\
use_lease_scavenge_time,use_logic_filter_rules,use_nextserver,\
use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,\
use_update_dns_on_lease_renewal,use_zone_associations,zone_associations
