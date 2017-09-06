package dnsview

import "github.com/sky-uk/skyinfoblox/api/common"

const wapiVersion = "/wapi/v2.6.1/"
const dnsViewEndpoint = "view"

// DNSView : DNS views provide the ability to serve one version of DNS data to one set of clients and another version to another set of clients. With DNS views, the appliance can provide a different answer to the same query, depending on the source of the query.
type DNSView struct {
	BlackListAction                  string                         `json:"blacklist_action,omitempty"`
	BlackListLogQuery                *bool                          `json:"blacklist_log_query,omitempty"`
	BlackListRedirectAddresses       []string                       `json:"blacklist_redirect_addresses,omitempty"`
	BlackListRedirectTTL             uint                           `json:"blacklist_redirect_ttl,omitempty"`
	BlackListRuleSets                []string                       `json:"blacklist_rulesets,omitempty"`
	Comment                          string                         `json:"comment,omitempty"`
	CustomRootNameServers            common.ExternalServer          `json:"custom_root_name_servers,omitempty"`
	DDNSPrincipalGroup               string                         `json:"ddns_principal_group,omitempty"`
	DDNSPrincipalTracking            *bool                          `json:"ddns_principal_tracking,omitempty"`
	DDNSRestrictPatterns             *bool                          `json:"ddns_restrict_patterns,omitempty"`
	DDNSRestrictPatternsList         []string                       `json:"ddns_restrict_patterns_list,omitempty"`
	DDNSRestrictProtected            *bool                          `json:"ddns_restrict_protected,omitempty"`
	DDNSRestrictSecure               *bool                          `json:"ddns_restrict_secure,omitempty"`
	DDNSRestrictStatic               *bool                          `json:"ddns_restrict_static,omitempty"`
	Disable                          *bool                          `json:"disable,omitempty"`
	DNS64Enabled                     *bool                          `json:"dns64_enabled,omitempty"`
	DNS64Groups                      []string                       `json:"dns64_groups,omitempty"`
	DNSSecEnabled                    *bool                          `json:"dnssec_enabled,omitempty"`
	DNSSecExpiredSignaturesEnabled   *bool                          `json:"dnssec_expired_signatures_enabled,omitempty"`
	DNSSecNegativeTrustAnchors       []string                       `json:"dnssec_negative_trust_anchors,omitempty"`
	DNSSecTrustedKeys                []common.DNSSecTrustedKey      `json:"dnssec_trusted_keys,omitempty"`
	DNSSecValidationEnabled          *bool                          `json:"dnssec_validation_enabled,omitempty"`
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
	NXDomainLogQuery                 *bool                          `json:"nxdomain_log_query,omitempty"`
	NXDomainRedirect                 *bool                          `json:"nxdomain_redirect,omitempty"`
	NXDomainRedirectAddresses        []string                       `json:"nxdomain_redirect_addresses,omitempty"`
	NXDomainRedirectAddressesV6      []string                       `json:"nxdomain_redirect_addresses_v6,omitempty"`
	NXDomainRedirectTTL              uint                           `json:"nxdomain_redirect_ttl,omitempty"`
	NXDomainRuleSets                 []string                       `json:"nxdomain_rulesets,omitempty"`
	Recursion                        *bool                          `json:"recursion,omitempty"`
	Reference                        string                         `json:"_ref,omitempty"`
	ResponseRateLimiting             common.DNSResponseRateLimiting `json:"response_rate_limiting,omitempty"`
	RootNameServerType               string                         `json:"root_name_server_type,omitempty"`
	RpzDropIPRuleEnabled             *bool                          `json:"rpz_drop_ip_rule_enabled,omitempty"`
	RpzDropIPRuleMinPrefixLengthIpv4 uint                           `json:"rpz_drop_ip_rule_min_prefix_length_ipv4,omitempty"`
	RpzDropIPRuleMinPrefixLengthIpv6 uint                           `json:"rpz_drop_ip_rule_min_prefix_length_ipv6,omitempty"`
	RpzQNameWaitRecurse              *bool                          `json:"rpz_qname_wait_recurse,omitempty"`
	ScavengingSettings               common.DNSScavengingSettings   `json:"scavenging_settings,omitempty"`
	Sortlist                         []common.DNSSortlist             `json:"sortlist,omitempty"`
	UseBlacklist                     *bool                          `json:"use_blacklist,omitempty"`
	UseDDNSPatternsRestriction       *bool                          `json:"use_ddns_patterns_restriction,omitempty"`
	UseDDNSPrincipalSecurity         *bool                          `json:"use_ddns_principal_security,omitempty"`
	UseDDNSRestrictProtected         *bool                          `json:"use_ddns_restrict_protected,omitempty"`
	UseDDNSRestrictStatic            *bool                          `json:"use_ddns_restrict_static,omitempty"`
	UseDNS64                         *bool                          `json:"use_dns64,omitempty"`
	UseDNSSec                        *bool                          `json:"use_dnssec,omitempty"`
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
	UseRpzDropIPRule                 *bool                          `json:"use_rpz_drop_ip_rule,omitempty"`
	UseRpzQNameWaitRecurse           *bool                          `json:"use_rpz_qname_wait_recurse,omitempty"`
	UseScavengingSettings            *bool                          `json:"use_scavenging_settings,omitempty"`
	UseSortlist                      *bool                          `json:"use_sortlist,omitempty"`
}
