package dnsview

import "github.com/sky-uk/skyinfoblox/api/common"

const wapiVersion = "/wapi/v2.6.1/"
const dnsViewEndpoint = "view"

// DnsView : DNS views provide the ability to serve one version of DNS data to one set of clients and another version to another set of clients. With DNS views, the appliance can provide a different answer to the same query, depending on the source of the query.
type DNSView struct {
	BlackListAction                  string                         `json:"blacklist_action,omitempty"`
	BlackListLogQuery                *bool                          `json:"blacklist_log_query,omitempty"`
	BlackListRedirectAddresses       []string                       `json:"blacklist_redirect_addresses,omitempty"`
	BlackListRedirectTTL             uint                           `json:"blacklist_redirect_ttl,omitempty"`
	BlackListRuleSets                []string                       `json:"blacklist_rulesets,omitempty"`
	Comment                          string                         `json:"comment,omitempty"`
	CustomRootNameServers            common.ExternalServer          `json:"custom_root_name_servers,omitempty"`
	DdnsPrincipalGroup               string                         `json:"ddns_principal_group,omitempty"`
	DdnsPrincipalTracking            *bool                          `json:"ddns_principal_tracking,omitempty"`
	DdnsRestrictPatterns             *bool                          `json:"ddns_restrict_patterns,omitempty"`
	DdnsRestrictPatternsList         []string                       `json:"ddns_restrict_patterns_list,omitempty"`
	DdnsRestrictProtected            *bool                          `json:"ddns_restrict_protected,omitempty"`
	DdnsRestrictSecure               *bool                          `json:"ddns_restrict_secure,omitempty"`
	DdnsRestrictStatic               *bool                          `json:"ddns_restrict_static,omitempty"`
	Disable                          *bool                          `json:"disable,omitempty"`
	Dns64Enabled                     *bool                          `json:"dns64_enabled,omitempty"`
	Dns64Groups                      []string                       `json:"dns64_groups,omitempty"`
	DnsSecEnabled                    *bool                          `json:"dnssec_enabled,omitempty"`
	DnsSecExpiredSignaturesEnabled   *bool                          `json:"dnssec_expired_signatures_enabled,omitempty"`
	DnsSecNegativeTrustAnchors       []string                       `json:"dnssec_negative_trust_anchors,omitempty"`
	DnsSecTrustedKeys                []common.DnsSecTrustedKey      `json:"dnssec_trusted_keys,omitempty"`
	DnsSecValidationEnabled          *bool                          `json:"dnssec_validation_enabled,omitempty"`
	EnableBlacklist                  *bool                          `json:"enable_blacklist,omitempty"`
	EnableFixedRRsetOrderFQDN        *bool                          `json:"enable_fixed_rrset_order_fqdns,omitempty"`
	EnableMatchRecursiveOnly         *bool                          `json:"enable_match_recursive_only,omitempty"`
	FilterAAAA                       string                         `json:"filter_aaaa,omitempty"`
	FilterAAAAList                   []common.AddressAC             `json:"filter_aaaa_list,omitempty"`
	FilterRRsetOrderFQDNs            common.FixedRRSetOrderFQDN     `json:"fixed_rrset_order_fqdns,omitempty"`
	ForwardOnly                      *bool                          `json:"forward_only,omitempty"`
	Forwarders                       []string                       `json:"forwarders,omitempty"`
	IsDefault                        *bool                          `json:"is_default,omitempty"`
	LameTTL                          uint                           `json:"lame_ttl,omitempty"`
	MatchClients                     []interface{}                  `json:"match_clients,omitempty"`
	MatchDestinations                []interface{}                  `json:"match_destinations,omitempty"`
	MaxCacheTTL                      uint                           `json:"max_cache_ttl,omitempty"`
	MaxNCacheTTL                     uint                           `json:"max_ncache_ttl,omitempty"`
	Name                             string                         `json:"name,omitempty"`
	NetworkView                      string                         `json:"network_view,omitempty"`
	NotifyDelay                      uint                           `json:"notify_delay,omitempty"`
	NxDomainLogQuery                 *bool                          `json:"nxdomain_log_query,omitempty"`
	NxDomainRedirect                 *bool                          `json:"nxdomain_redirect,omitempty"`
	NxDomainRedirectAddresses        []string                       `json:"nxdomain_redirect_addresses,omitempty"`
	NxDomainRedirectAddressesV6      []string                       `json:"nxdomain_redirect_addresses_v6,omitempty"`
	NxDomainRedirectTTL              uint                           `json:"nxdomain_redirect_ttl,omitempty"`
	NxDomainRuleSets                 []string                       `json:"nxdomain_rulesets,omitempty"`
	Recursion                        *bool                          `json:"recursion,omitempty"`
	Reference                        string                         `json:"_ref,omitempty"`
	ResponseRateLimiting             common.DnsResponseRateLimiting `json:"response_rate_limiting,omitempty"`
	RootNameServerType               string                         `json:"root_name_server_type,omitempty"`
	RpzDropIpRuleEnabled             *bool                          `json:"rpz_drop_ip_rule_enabled,omitempty"`
	RpzDropIpRuleMinPrefixLengthIpv4 uint                           `json:"rpz_drop_ip_rule_min_prefix_length_ipv4,omitempty"`
	RpzDropIpRuleMinPrefixLengthIpv6 uint                           `json:"rpz_drop_ip_rule_min_prefix_length_ipv6,omitempty"`
	RpzQNameWaitRecurse              *bool                          `json:"rpz_qname_wait_recurse,omitempty"`
	ScavengingSettings               common.DnsScavengingSettings   `json:"scavenging_settings,omitempty"`
	Sortlist                         common.DnsSortlist             `json:"sortlist,omitempty"`
	UseBlacklist                     *bool                          `json:"use_blacklist,omitempty"`
	UseDdnsPatternsRestriction       *bool                          `json:"use_ddns_patterns_restriction,omitempty"`
	UseDdnsPrincipalSecurity         *bool                          `json:"use_ddns_principal_security,omitempty"`
	UseDdnsRestrictProtected         *bool                          `json:"use_ddns_restrict_protected,omitempty"`
	UseDdnsRestrictStatic            *bool                          `json:"use_ddns_restrict_static,omitempty"`
	UseDns64                         *bool                          `json:"use_dns64,omitempty"`
	UseDnsSec                        *bool                          `json:"use_dnssec,omitempty"`
	UseFilterAAAA                    *bool                          `json:"use_filter_aaaa,omitempty"`
	UseFixedRRsetOrderFQDNs          *bool                          `json:"use_fixed_rrset_order_fqdns,omitempty"`
	UseForwarders                    *bool                          `json:"use_forwarders,omitempty"`
	UseLameTTL                       *bool                          `json:"use_lame_ttl,omitempty"`
	UseMaxCacheTTL                   *bool                          `json:"use_max_cache_ttl,omitempty"`
	UseMaxNCacheTTL                  *bool                          `json:"use_max_ncache_ttl,omitempty"`
	UseNXDomainRedirect              *bool                          `json:"use_nxdomain_redirect,omitempty"`
	UseRecursion                     *bool                          `json:"use_recursion,omitempty"`
	UseResponseRateLimiting          *bool                          `json:"use_response_rate_limiting,omitempty"`
	UseRootNameServer                *bool                          `json:"use_root_name_server,omitempty"`
	UseRpzDropIpRule                 *bool                          `json:"use_rpz_drop_ip_rule,omitempty"`
	UseRpzQNameWaitRecurse           *bool                          `json:"use_rpz_qname_wait_recurse,omitempty"`
	UseScavengingSettings            *bool                          `json:"use_scavenging_settings,omitempty"`
	UseSortlist                      *bool                          `json:"use_sortlist,omitempty"`
}
