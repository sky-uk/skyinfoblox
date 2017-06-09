package network

// Network : base DHCP Network object model
// See : https://h1infoblox.devops.int.ovp.bskyb.com/wapidoc/objects/network.html
// Fields name 	| in/out| Descr
// -------------------------------------------------------------------------------
// Ref		|  OUT	| The unique object reference. The name part of it
//	  	|  	| has the following components:
//  	    	|  	| - Address of the network
//  	    	|  	| - CIDR of the network
//  	    	|  	| - Name of the network view
// Network	|  IN	| The network address, in IPv4 Address/CIDR format.
// NetworkView	|  IN	| The name of the network view in which this network resides.
// Comment	|	|  IN   | Object description/notes
type Network struct {
	Ref         string `json:"_ref"`
	Network     string `json:"network"`
	NetworkView string `json:"network_view"`
	Comment     string `json:"comment,omitempty"`
}
