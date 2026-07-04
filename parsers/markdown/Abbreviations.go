package markdown

import "sort"

var Abbreviations map[string]string = map[string]string{
	"ACL":   "Access Control List",
	"AES":   "Advanced Encryption Standard",
	"AMSI":  "Anti Malware Scan Interface",
	"AP":    "Access Point",
	"API":   "Application Programming Interface",
	"APT":   "Advanced Persistent Threat",
	"ARP":   "Address Resolution Protocol",
	"ASN":   "Autonomous System Number",

	"BGP":   "Border Gateway Protocol",
	"BOF":   "Beacon Object File",

	"C2":    "Command and Control",
	"CA":    "Certificate Authority",
	"CAPEC": "Common Attack Pattern Enumeration and Classification",
	"CBC":   "Cipher Block Chaining",
	"CERT":  "Computer Emergency Response Team",
	"CI/CD": "Continuous Integration and Continuous Deployment",
	"CIDR":  "Classless Inter-Domain Routing",
	"CISA":  "Cybersecurity and Infrastructure Security Agency",
	"CISO":  "Chief Information Security Officer",
	"CKC":   "Cyber Kill Chain",
	"CNC":   "Command and Control",
	"CoW":   "Copy-on-Write",
	"CPU":   "Central Processing Unit",
	"CSP":   "Content Security Policy",
	"CSPM":  "Cloud Security Posture Management",
	"CSRF":  "Cross-Site Request Forgery",
	"CTF":   "Capture The Flag",
	"CVE":   "Common Vulnerabilities and Exposures",
	"CVSS":  "Common Vulnerability Scoring System",

	"DA":    "Domain Admin",
	"DC":    "Domain Controller",
	"DDoS":  "Distributed Denial of Service",
	"DEP":   "Data Execution Prevention",
	"DLL":   "Dynamic Link Library",
	"DHCP":  "Dynamic Host Configuration Protocol",
	"DMARC": "Domain-based Message Authentication, Reporting, and Conformance",
	"DNS":   "Domain Name System",
	"DNSSEC":"Domain Name System Security Extensions",
	"DoS":   "Denial of Service",
	"DSCP":  "Differentiated Services Code Point",

	"EDR":   "Endpoint Detection and Response",
	"EDRM":  "Electronic Discovery Reference Model",
	"EoP":   "Elevation of Privilege",
	"EPP":   "Endpoint Protection Platform",
	"ESP":   "Encapsulating Security Payload",

	"FQDN":  "Fully Qualified Domain Name",
	"FTP":   "File Transfer Protocol",

	"GRE":   "Generic Routing Encapsulation",
	"GUI":   "Graphical User Interface",
	"GPO":   "Group Policy Object",

	"HTTP":  "Hypertext Transfer Protocol",
	"HTTPS": "Hypertext Transfer Protocol Secure",
	"HMAC":  "Hash-based Message Authentication Code",

	"IAM":   "Identity Access Management",
	"ICMP":  "Internet Control Message Protocol",
	"ICS":   "Intrusion Control System",
	"IDS":   "Intrusion Detection System",
	"IKE":   "Internet Key Exchange",
	"IOC":   "Indicator of Compromise",
	"IOA":   "Indicator of Attack",
	"IP":    "Internet Protocol",
	"IPS":   "Intrusion Prevention System",
	"IPsec": "Internet Protocol Security",
	"IPv4":  "Internet Protocol Version 4",
	"IPv6":  "Internet Protocol Version 6",
	"IR":    "Incident Response",
	"ISP":   "Internet Service Provider",

	"JWT":   "JSON Web Token",

	"LAN":    "Local Area Network",
	"LLDP":   "Link Layer Discovery Protocol",
	"LFI":    "Local File Inclusion",
	"LOLBAS": "Living Off The Land Binaries and Scripts",
	"LOLBIN": "Living Off The Land Binary",
	"LPE":    "Local Privilege Escalation",
	"LSASS":  "Local Security Authority Subsystem Service",

	"MAC":   "Media Access Control",
	"MAN":   "Metropolitan Area Network",
	"MFA":   "Multi-Factor Authentication",
	"MSS":   "Maximum Segment Size",
	"MITM":  "Man-in-the-Middle",
	"MOTW":  "Mark of the Web",

	"NAT":   "Network Address Translation",
	"NDP":   "Neighbor Discovery Protocol",
	"NDR":   "Network Detection and Response",
	"NIC":   "Network Interface Card",
	"NTLM":  "NT LAN Manager",
	"NTP":   "Network Time Protocol",

	"OAST":  "Out-of-Band Application Security Testing",
	"OAuth": "Open Authorization",
	"OCSP":  "Online Certificate Status Protocol",
	"OSCP":  "Offensive Security Certified Professional",
	"OSI":   "Open Systems Interconnection",
	"OSINT": "Open Source Intelligence",
	"OSPF":  "Open Shortest Path First",

	"PAM":   "Privileged Access Management",
	"PAT":   "Port Address Translation",
	"PCI":   "Peripheral Component Interconnect",
	"PKI":   "Public Key Infrastructure",
	"POP3":  "Post Office Protocol Version 3",
	"PoC":   "Proof of Concept",
	"PPTP":  "Point-to-Point Tunneling Protocol",
	"PT":    "Penetration Test",

	"QoS":   "Quality of Service",

	"RARP":  "Reverse Address Resolution Protocol",
	"RAT":   "Remote Access Trojan",
	"RBAC":  "Role-Based Access Control",
	"RCE":   "Remote Code Execution",
	"RDP":   "Remote Desktop Protocol",
	"RFI":   "Remote File Inclusion",
	"RIP":   "Routing Information Protocol",
	"ROP":   "Return-Oriented Programming",
	"RPC":   "Remote Procedure Call",
	"RSA":   "Rivest-Shamir-Adleman",
	"RTT":   "Round-Trip Time",

	"SAML":  "Security Assertion Markup Language",
	"SAST":  "Static Application Security Testing",
	"SBOM":  "Software Bill of Materials",
	"SFTP":  "SSH File Transfer Protocol",
	"SIEM":  "Security Information and Event Management",
	"SMB":   "Server Message Block",
	"SMTP":  "Simple Mail Transfer Protocol",
	"SNMP":  "Simple Network Management Protocol",
	"SOC":   "Security Operations Center",
	"SOAR":  "Security Orchestration, Automation, and Response",
	"SPN":   "Service Principal Name",
	"SQLi":  "SQL Injection",
	"SRP":   "Software Restriction Policies",
	"SSH":   "Secure Shell",
	"SSID":  "Service Set Identifier",
	"SSL":   "Secure Sockets Layer",
	"SSO":   "Single Sign-On",
	"SSRF":  "Server-Side Request Forgery",
	"STP":   "Spanning Tree Protocol",
	"SYN":   "Synchronize",

	"TCP":   "Transmission Control Protocol",
	"TLS":   "Transport Layer Security",
	"TPM":   "Trusted Platform Module",
	"TTL":   "Time To Live",
	"TTP":   "Tactics, Techniques, and Procedures",


	"U2F":   "Universal 2nd Factor",
	"UAC":   "User Account Control",
	"UDP":   "User Datagram Protocol",
	"UPN":   "User Principal Name",
	"URI":   "Uniform Resource Identifier",
	"URL":   "Uniform Resource Locator",
	"USB":   "Universal Serial Bus",

	"VLAN":  "Virtual Local Area Network",
	"VDI":   "Virtual Desktop Infrastructure",
	"VM":    "Virtual Machine",
	"VPN":   "Virtual Private Network",

	"WAF":   "Web Application Firewall",
	"WAN":   "Wide Area Network",
	"WDAC":  "Windows Defender Application Control",
	"WEP":   "Wired Equivalent Privacy",
	"Wi-Fi": "Wireless Fidelity",
	"WLAN":  "Wireless Local Area Network",
	"WMI":   "Windows Management Instrumentation",
	"WPA":   "Wi-Fi Protected Access",
	"WPA2":  "Wi-Fi Protected Access II",
	"WPA3":  "Wi-Fi Protected Access III",
	"WWW":   "World Wide Web",

	"XDR":   "Extended Detection and Response",
	"XML":   "Extensible Markup Language",
	"XSS":   "Cross-Site Scripting",

	"YARA":  "Yet Another Recursive Acronym",

}

var abbreviation_keys []string

func init() {

	abbreviation_keys = make([]string, 0, len(Abbreviations))

	for key, _ := range Abbreviations {
		abbreviation_keys = append(abbreviation_keys, key)
	}

	sort.Slice(abbreviation_keys, func(a int, b int) bool {
		return len(abbreviation_keys[a]) > len(abbreviation_keys[b])
	})

}
