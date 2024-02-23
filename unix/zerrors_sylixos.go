// mkerrors.sh -mcpu=generic -fno-omit-frame-pointer -mstrict-align -ffixed-x18 -Wl,-shared -fPIC -shared -DSYLIXOS -I/home/user/toolchains/libsylixos/SylixOS -I/home/user/toolchains/libsylixos/SylixOS/include -I/home/user/toolchains/libsylixos/SylixOS/include/network
// Code generated by the command above; DO NOT EDIT.

//go:build sylixos

package unix

import "syscall"

const (
	AF_ALG					= 0x26
	AF_APPLETALK				= 0x5
	AF_ATMPVC				= 0x8
	AF_ATMSVC				= 0x14
	AF_AX25					= 0x3
	AF_BLUETOOTH				= 0x1f
	AF_BRIDGE				= 0x7
	AF_CAIF					= 0x25
	AF_CAN					= 0x1d
	AF_DECnet				= 0xc
	AF_ECONET				= 0x13
	AF_IEEE802154				= 0x24
	AF_INET					= 0x2
	AF_INET6				= 0xa
	AF_IPX					= 0x4
	AF_IRDA					= 0x17
	AF_ISDN					= 0x22
	AF_IUCV					= 0x20
	AF_KEY					= 0xf
	AF_LINK					= 0x12
	AF_LLC					= 0x1a
	AF_LOCAL				= 0x1
	AF_MAX					= 0x40
	AF_NETBEUI				= 0xd
	AF_NETROM				= 0x6
	AF_NFC					= 0x27
	AF_PACKET				= 0x11
	AF_PHONET				= 0x23
	AF_PPPOX				= 0x18
	AF_RDS					= 0x15
	AF_ROSE					= 0xb
	AF_ROUTE				= 0x10
	AF_RXRPC				= 0x21
	AF_SECURITY				= 0xe
	AF_SNA					= 0x16
	AF_TIPC					= 0x1e
	AF_UNIX					= 0x1
	AF_UNSPEC				= 0x0
	AF_WANPIPE				= 0x19
	AF_X25					= 0x9
	B0					= 0x0
	B1000000				= 0x10080000
	B110					= 0x30000
	B115200					= 0x10020000
	B1152000				= 0x10090000
	B1200					= 0x90000
	B134					= 0x40000
	B150					= 0x50000
	B1500000				= 0x100a0000
	B1800					= 0xa0000
	B19200					= 0xe0000
	B200					= 0x60000
	B2000000				= 0x100b0000
	B230400					= 0x10030000
	B2400					= 0xb0000
	B2500000				= 0x100c0000
	B300					= 0x70000
	B3000000				= 0x100d0000
	B3500000				= 0x100e0000
	B38400					= 0xf0000
	B4000000				= 0x100f0000
	B460800					= 0x10040000
	B4800					= 0xc0000
	B50					= 0x10000
	B500000					= 0x10050000
	B57600					= 0x10010000
	B576000					= 0x10060000
	B600					= 0x80000
	B75					= 0x20000
	B921600					= 0x10070000
	B9600					= 0xd0000
	BRKINT					= 0x2
	CLOCAL					= 0x1
	CREAD					= 0x2
	CS5					= 0x0
	CS6					= 0x4
	CS7					= 0x8
	CS8					= 0xc
	CSIZE					= 0xc
	CSTOPB					= 0x20
	DT_BLK					= 0x6
	DT_CHR					= 0x2
	DT_DIR					= 0x4
	DT_FIFO					= 0x1
	DT_LNK					= 0xa
	DT_REG					= 0x8
	DT_SOCK					= 0xc
	DT_UNKNOWN				= 0x0
	DT_WHT					= 0xe
	EAP_DEFALLOWREQ				= 0x8
	EAP_DEFREQTIME				= 0x6
	EAP_DEFTIMEOUT				= 0x6
	EAP_SUPPORT				= 0x1
	ECHO					= 0x8
	ECHOCTL					= 0x200
	ECHOE					= 0x10
	ECHOK					= 0x20
	ECHOKE					= 0x800
	ECHONL					= 0x40
	ECHOPRT					= 0x400
	ECP_SUPPORT				= 0x0
	ENSRBADFAMILY				= 0xa8
	ENSRBADNAME				= 0xa7
	ENSRBADQUERY				= 0xa6
	ENSRBADRESP				= 0xa9
	ENSRCNAMELOOP				= 0xb1
	ENSRCONNREFUSED				= 0xaa
	ENSRDESTRUCTION				= 0xaf
	ENSRFILE				= 0xad
	ENSRFORMERR				= 0xa1
	ENSRNODATA				= 0xa0
	ENSRNOMEM				= 0xae
	ENSRNOTFOUND				= 0xa3
	ENSRNOTIMP				= 0xa4
	ENSROF					= 0xac
	ENSROK					= 0x0
	ENSRQUERYDOMAINTOOLONG			= 0xb0
	ENSRREFUSED				= 0xa5
	ENSRSERVFAIL				= 0xa2
	ENSRTIMEOUT				= 0xab
	ETHARP_DEBUG				= 0x0
	ETHARP_STATS				= 0x1
	ETHARP_SUPPORT_STATIC_ENTRIES		= 0x1
	ETHARP_SUPPORT_VLAN			= 0x1
	ETHARP_TABLE_MATCH_NETIF		= 0x1
	FD_CLOEXEC				= 0x1
	FD_SETSIZE				= 0x800
	FIOFLUSH				= 0x2
	FIORFLUSH				= 0x1a
	FIOWFLUSH				= 0x1b
	FLUSHO					= 0x1000
	F_CNVT					= 0xc
	F_DUP2FD				= 0x12
	F_DUP2FD_CLOEXEC			= 0x13
	F_DUPFD					= 0x0
	F_DUPFD_CLOEXEC				= 0x11
	F_GETFD					= 0x1
	F_GETFL					= 0x3
	F_GETLK					= 0x7
	F_GETOWN				= 0x5
	F_GETSIG				= 0xe
	F_LOCK					= 0x1
	F_OK					= 0x0
	F_RDLCK					= 0x1
	F_RGETLK				= 0xa
	F_RSETLK				= 0xb
	F_RSETLKW				= 0xd
	F_SETFD					= 0x2
	F_SETFL					= 0x4
	F_SETLK					= 0x8
	F_SETLKW				= 0x9
	F_SETOWN				= 0x6
	F_SETSIG				= 0xf
	F_TEST					= 0x3
	F_TLOCK					= 0x2
	F_ULOCK					= 0x0
	F_UNLCK					= 0x3
	F_WRLCK					= 0x2
	HUPCL					= 0x10
	ICANON					= 0x2
	ICMP6_STATS				= 0x1
	ICRNL					= 0x100
	IEXTEN					= 0x8000
	IFAN_ARRIVAL				= 0x0
	IFAN_DEPARTURE				= 0x1
	IFF_802_1Q_VLAN				= 0x1
	IFF_ALLMULTI				= 0x800
	IFF_BONDING				= 0x4
	IFF_BROADCAST				= 0x2
	IFF_EBRIDGE				= 0x2
	IFF_LOOPBACK				= 0x100
	IFF_MULTICAST				= 0x80
	IFF_NOARP				= 0x200
	IFF_POINTOPOINT				= 0x4
	IFF_PROMISC				= 0x400
	IFF_RUNNING				= 0x10
	IFF_UP					= 0x1
	IFNAMSIZ				= 0x10
	IFT_1822				= 0x2
	IFT_A12MPPSWITCH			= 0x82
	IFT_AAL2				= 0xbb
	IFT_AAL5				= 0x31
	IFT_ADSL				= 0x5e
	IFT_AFLANE8023				= 0x3b
	IFT_AFLANE8025				= 0x3c
	IFT_ARAP				= 0x58
	IFT_ARCNET				= 0x23
	IFT_ARCNETPLUS				= 0x24
	IFT_ASYNC				= 0x54
	IFT_ATM					= 0x25
	IFT_ATMDXI				= 0x69
	IFT_ATMFUNI				= 0x6a
	IFT_ATMIMA				= 0x6b
	IFT_ATMLOGICAL				= 0x50
	IFT_ATMRADIO				= 0xbd
	IFT_ATMSUBINTERFACE			= 0x86
	IFT_ATMVCIENDPT				= 0xc2
	IFT_ATMVIRTUAL				= 0x95
	IFT_BGPPOLICYACCOUNTING			= 0xa2
	IFT_BRIDGE				= 0xd1
	IFT_BSC					= 0x53
	IFT_CARP				= 0xf8
	IFT_CCTEMUL				= 0x3d
	IFT_CEPT				= 0x13
	IFT_CES					= 0x85
	IFT_CHANNEL				= 0x46
	IFT_CNR					= 0x55
	IFT_COFFEE				= 0x84
	IFT_COMPOSITELINK			= 0x9b
	IFT_DCN					= 0x8d
	IFT_DIGITALPOWERLINE			= 0x8a
	IFT_DIGITALWRAPPEROVERHEADCHANNEL	= 0xba
	IFT_DLSW				= 0x4a
	IFT_DOCSCABLEDOWNSTREAM			= 0x80
	IFT_DOCSCABLEMACLAYER			= 0x7f
	IFT_DOCSCABLEUPSTREAM			= 0x81
	IFT_DS0					= 0x51
	IFT_DS0BUNDLE				= 0x52
	IFT_DS1FDL				= 0xaa
	IFT_DS3					= 0x1e
	IFT_DTM					= 0x8c
	IFT_DVBASILN				= 0xac
	IFT_DVBASIOUT				= 0xad
	IFT_DVBRCCDOWNSTREAM			= 0x93
	IFT_DVBRCCMACLAYER			= 0x92
	IFT_DVBRCCUPSTREAM			= 0x94
	IFT_ENC					= 0xf4
	IFT_EON					= 0x19
	IFT_EPLRS				= 0x57
	IFT_ESCON				= 0x49
	IFT_ETHER				= 0x6
	IFT_FAITH				= 0xf2
	IFT_FAST				= 0x7d
	IFT_FASTETHER				= 0x3e
	IFT_FASTETHERFX				= 0x45
	IFT_FDDI				= 0xf
	IFT_FIBRECHANNEL			= 0x38
	IFT_FRAMERELAYINTERCONNECT		= 0x3a
	IFT_FRAMERELAYMPI			= 0x5c
	IFT_FRDLCIENDPT				= 0xc1
	IFT_FRELAY				= 0x20
	IFT_FRELAYDCE				= 0x2c
	IFT_FRF16MFRBUNDLE			= 0xa3
	IFT_FRFORWARD				= 0x9e
	IFT_G703AT2MB				= 0x43
	IFT_G703AT64K				= 0x42
	IFT_GIF					= 0xf0
	IFT_GIGABITETHERNET			= 0x75
	IFT_GR303IDT				= 0xb2
	IFT_GR303RDT				= 0xb1
	IFT_H323GATEKEEPER			= 0xa4
	IFT_H323PROXY				= 0xa5
	IFT_HDH1822				= 0x3
	IFT_HDLC				= 0x76
	IFT_HDSL2				= 0xa8
	IFT_HIPERLAN2				= 0xb7
	IFT_HIPPI				= 0x2f
	IFT_HIPPIINTERFACE			= 0x39
	IFT_HOSTPAD				= 0x5a
	IFT_HSSI				= 0x2e
	IFT_HY					= 0xe
	IFT_IBM370PARCHAN			= 0x48
	IFT_IDSL				= 0x9a
	IFT_IEEE1394				= 0x90
	IFT_IEEE80211				= 0x47
	IFT_IEEE80212				= 0x37
	IFT_IEEE802154				= 0x103
	IFT_IEEE8023ADLAG			= 0xa1
	IFT_IFGSN				= 0x91
	IFT_IMT					= 0xbe
	IFT_INFINIBAND				= 0xc7
	IFT_INTERLEAVE				= 0x7c
	IFT_IP					= 0x7e
	IFT_IPFORWARD				= 0x8e
	IFT_IPOVERATM				= 0x72
	IFT_IPOVERCDLC				= 0x6d
	IFT_IPOVERCLAW				= 0x6e
	IFT_IPSWITCH				= 0x4e
	IFT_IPXIP				= 0xf9
	IFT_ISDN				= 0x3f
	IFT_ISDNBASIC				= 0x14
	IFT_ISDNPRIMARY				= 0x15
	IFT_ISDNS				= 0x4b
	IFT_ISDNU				= 0x4c
	IFT_ISO88022LLC				= 0x29
	IFT_ISO88023				= 0x7
	IFT_ISO88024				= 0x8
	IFT_ISO88025				= 0x9
	IFT_ISO88025CRFPINT			= 0x62
	IFT_ISO88025DTR				= 0x56
	IFT_ISO88025FIBER			= 0x73
	IFT_ISO88026				= 0xa
	IFT_ISUP				= 0xb3
	IFT_L2VLAN				= 0x87
	IFT_L3IPVLAN				= 0x88
	IFT_L3IPXVLAN				= 0x89
	IFT_LAPB				= 0x10
	IFT_LAPD				= 0x4d
	IFT_LAPF				= 0x77
	IFT_LOCALTALK				= 0x2a
	IFT_LOOP				= 0x18
	IFT_MEDIAMAILOVERIP			= 0x8b
	IFT_MFSIGLINK				= 0xa7
	IFT_MIOX25				= 0x26
	IFT_MODEM				= 0x30
	IFT_MPC					= 0x71
	IFT_MPLS				= 0xa6
	IFT_MPLSTUNNEL				= 0x96
	IFT_MSDSL				= 0x8f
	IFT_MVL					= 0xbf
	IFT_MYRINET				= 0x63
	IFT_NFAS				= 0xaf
	IFT_NSIP				= 0x1b
	IFT_OPTICALCHANNEL			= 0xc3
	IFT_OPTICALTRANSPORT			= 0xc4
	IFT_OTHER				= 0x1
	IFT_P10					= 0xc
	IFT_P80					= 0xd
	IFT_PARA				= 0x22
	IFT_PFLOG				= 0xf6
	IFT_PFSYNC				= 0xf7
	IFT_PLC					= 0xae
	IFT_POS					= 0xab
	IFT_PPP					= 0x17
	IFT_PPPMULTILINKBUNDLE			= 0x6c
	IFT_PROPBWAP2MP				= 0xb8
	IFT_PROPCNLS				= 0x59
	IFT_PROPDOCSWIRELESSDOWNSTREAM		= 0xb5
	IFT_PROPDOCSWIRELESSMACLAYER		= 0xb4
	IFT_PROPDOCSWIRELESSUPSTREAM		= 0xb6
	IFT_PROPMUX				= 0x36
	IFT_PROPVIRTUAL				= 0x35
	IFT_PROPWIRELESSP2P			= 0x9d
	IFT_PTPSERIAL				= 0x16
	IFT_PVC					= 0xf1
	IFT_QLLC				= 0x44
	IFT_RADIOMAC				= 0xbc
	IFT_RADSL				= 0x5f
	IFT_REACHDSL				= 0xc0
	IFT_RFC1483				= 0x9f
	IFT_RS232				= 0x21
	IFT_RSRB				= 0x4f
	IFT_SDLC				= 0x11
	IFT_SDSL				= 0x60
	IFT_SHDSL				= 0xa9
	IFT_SIP					= 0x1f
	IFT_SLIP				= 0x1c
	IFT_SMDSDXI				= 0x2b
	IFT_SMDSICIP				= 0x34
	IFT_SONET				= 0x27
	IFT_SONETOVERHEADCHANNEL		= 0xb9
	IFT_SONETPATH				= 0x32
	IFT_SONETVT				= 0x33
	IFT_SRP					= 0x97
	IFT_SS7SIGLINK				= 0x9c
	IFT_STACKTOSTACK			= 0x6f
	IFT_STARLAN				= 0xb
	IFT_STF					= 0xd7
	IFT_T1					= 0x12
	IFT_TDLC				= 0x74
	IFT_TERMPAD				= 0x5b
	IFT_TR008				= 0xb0
	IFT_TRANSPHDLC				= 0x7b
	IFT_TUNNEL				= 0x83
	IFT_ULTRA				= 0x1d
	IFT_USB					= 0xa0
	IFT_V11					= 0x40
	IFT_V35					= 0x2d
	IFT_V36					= 0x41
	IFT_V37					= 0x78
	IFT_VDSL				= 0x61
	IFT_VIRTUALIPADDRESS			= 0x70
	IFT_VOICEEM				= 0x64
	IFT_VOICEENCAP				= 0x67
	IFT_VOICEFXO				= 0x65
	IFT_VOICEFXS				= 0x66
	IFT_VOICEOVERATM			= 0x98
	IFT_VOICEOVERFRAMERELAY			= 0x99
	IFT_VOICEOVERIP				= 0x68
	IFT_X213				= 0x5d
	IFT_X25					= 0x5
	IFT_X25DDN				= 0x4
	IFT_X25HUNTGROUP			= 0x7a
	IFT_X25MLP				= 0x79
	IFT_X25PLE				= 0x28
	IFT_XETHER				= 0x1a
	IGNBRK					= 0x1
	IGNCR					= 0x80
	IGNPAR					= 0x4
	IMAXBEL					= 0x2000
	INLCR					= 0x40
	INPCK					= 0x10
	IN_CLASSA_HOST				= 0xffffff
	IN_CLASSA_MAX				= 0x80
	IN_CLASSA_NET				= 0xff000000
	IN_CLASSA_NSHIFT			= 0x18
	IN_CLASSB_HOST				= 0xffff
	IN_CLASSB_MAX				= 0x10000
	IN_CLASSB_NET				= 0xffff0000
	IN_CLASSB_NSHIFT			= 0x10
	IN_CLASSC_HOST				= 0xff
	IN_CLASSC_MAX				= 0x0
	IN_CLASSC_NET				= 0xffffff00
	IN_CLASSC_NSHIFT			= 0x8
	IN_CLASSD_HOST				= 0xfffffff
	IN_CLASSD_MAX				= 0x0
	IN_CLASSD_NET				= 0xf0000000
	IN_CLASSD_NSHIFT			= 0x1c
	IN_LOOPBACKNET				= 0x7f
	IPPROTO_ENCAP				= 0x62
	IPPROTO_ICMP				= 0x1
	IPPROTO_ICMPV6				= 0x3a
	IPPROTO_IGMP				= 0x2
	IPPROTO_IP				= 0x0
	IPPROTO_IPIP				= 0x4
	IPPROTO_IPV6				= 0x29
	IPPROTO_PIM				= 0x67
	IPPROTO_RAW				= 0xff
	IPPROTO_RSVP				= 0x2e
	IPPROTO_SCTP				= 0x84
	IPPROTO_TCP				= 0x6
	IPPROTO_UDP				= 0x11
	IPPROTO_UDPLITE				= 0x88
	IPV6_ADD_MEMBERSHIP			= 0xc
	IPV6_CHECKSUM				= 0x7
	IPV6_CUSTOM_SCOPES			= 0x0
	IPV6_DEFHLIM				= 0x40
	IPV6_DROP_MEMBERSHIP			= 0xd
	IPV6_FLOWINFO_MASK			= 0xffffff0f
	IPV6_FLOWLABEL_MASK			= 0xffff0f00
	IPV6_FRAGTTL				= 0x78
	IPV6_FRAG_COPYHEADER			= 0x1
	IPV6_HLIMDEC				= 0x1
	IPV6_HOPLIMIT				= 0x34
	IPV6_JOIN_GROUP				= 0xc
	IPV6_LEAVE_GROUP			= 0xd
	IPV6_MAXHLIM				= 0xff
	IPV6_MAXPACKET				= 0xffff
	IPV6_MINHOPCNT				= 0x63
	IPV6_MMTU				= 0x500
	IPV6_MULTICAST_HOPS			= 0x12
	IPV6_MULTICAST_IF			= 0x11
	IPV6_MULTICAST_LOOP			= 0x13
	IPV6_PKTINFO				= 0x32
	IPV6_REASS_MAXAGE			= 0x3c
	IPV6_RECVHOPLIMIT			= 0x33
	IPV6_RECVPKTINFO			= 0x31
	IPV6_UNICAST_HOPS			= 0x10
	IPV6_V6ONLY				= 0x1b
	IPV6_VERSION				= 0x60
	IPV6_VERSION_MASK			= 0xf0
	IP_ADDR_ANY				= 0x0
	IP_ADDR_BROADCAST			= 0xffffffff
	IP_ADD_MEMBERSHIP			= 0x3
	IP_ADD_SOURCE_MEMBERSHIP		= 0x27
	IP_ANY_TYPE				= 0x0
	IP_BLOCK_SOURCE				= 0x26
	IP_CLASSA_HOST				= 0xffffff
	IP_CLASSA_MAX				= 0x80
	IP_CLASSA_NET				= 0xff000000
	IP_CLASSA_NSHIFT			= 0x18
	IP_CLASSB_HOST				= 0xffff
	IP_CLASSB_MAX				= 0x10000
	IP_CLASSB_NET				= 0xffff0000
	IP_CLASSB_NSHIFT			= 0x10
	IP_CLASSC_HOST				= 0xff
	IP_CLASSC_NET				= 0xffffff00
	IP_CLASSC_NSHIFT			= 0x8
	IP_CLASSD_HOST				= 0xfffffff
	IP_CLASSD_NET				= 0xf0000000
	IP_CLASSD_NSHIFT			= 0x1c
	IP_DEBUG				= 0x0
	IP_DEFAULT_MULTICAST_LOOP		= 0x1
	IP_DEFAULT_MULTICAST_TTL		= 0x1
	IP_DEFAULT_TTL				= 0xff
	IP_DF					= 0x4000
	IP_DROP_MEMBERSHIP			= 0x4
	IP_DROP_SOURCE_MEMBERSHIP		= 0x28
	IP_FORWARD				= 0x1
	IP_FORWARD_ALLOW_TX_ON_RX_NETIF		= 0x0
	IP_FRAG					= 0x1
	IP_HDRINCL				= 0x64
	IP_LOOPBACKNET				= 0x7f
	IP_MAXPACKET				= 0xffff
	IP_MAX_MEMBERSHIPS			= 0x14
	IP_MF					= 0x2000
	IP_MINTTL				= 0x63
	IP_MSFILTER				= 0x29
	IP_MSS					= 0x240
	IP_MULTICAST_IF				= 0x6
	IP_MULTICAST_LOOP			= 0x7
	IP_MULTICAST_TTL			= 0x5
	IP_OFFMASK				= 0x1fff
	IP_OPTIONS				= 0x65
	IP_OPTIONS_ALLOWED			= 0x1
	IP_PKTINFO				= 0x8
	IP_REASSEMBLY				= 0x1
	IP_REASS_DEBUG				= 0x0
	IP_REASS_MAXAGE				= 0xf
	IP_REASS_MAX_PBUFS			= 0x200
	IP_RF					= 0x8000
	IP_RSVP_OFF				= 0x10
	IP_RSVP_ON				= 0xf
	IP_RSVP_VIF_OFF				= 0x12
	IP_RSVP_VIF_ON				= 0x11
	IP_SOF_BROADCAST			= 0x1
	IP_SOF_BROADCAST_RECV			= 0x1
	IP_STATS				= 0x1
	IP_TOS					= 0x1
	IP_TTL					= 0x2
	IP_UNBLOCK_SOURCE			= 0x25
	ISIG					= 0x1
	ISTRIP					= 0x20
	IXANY					= 0x800
	IXOFF					= 0x1000
	IXON					= 0x400
	LOCK_EX					= 0x2
	LOCK_NB					= 0x80
	LOCK_SH					= 0x1
	LOCK_UN					= 0x3
	MADV_DONTNEED				= 0x4
	MADV_NORMAL				= 0x0
	MADV_RANDOM				= 0x1
	MADV_SEQUENTIAL				= 0x2
	MADV_WILLNEED				= 0x3
	MAP_ANON				= 0x8
	MAP_ANONYMOUS			= 0x8
	MAP_ANON_FD				= -0x1
	MAP_CONTIG				= 0x20
	MAP_FILE				= 0x0
	MAP_FIXED				= 0x4
	MAP_NORESERVE			= 0x10
	MAP_PREALLOC			= 0x8000  
	MAP_PRIVATE				= 0x2
	MAP_SHARED				= 0x1
	MCL_CURRENT				= 0x1
	MCL_FUTURE				= 0x2
	MSG_CMSG_CLOEXEC			= 0x40000
	MSG_CTRUNC				= 0x1000
	MSG_DONTWAIT				= 0x8
	MSG_EOR					= 0x8000
	MSG_MORE				= 0x10
	MSG_NOSIGNAL				= 0x4000
	MSG_NOTIFICATION			= 0x80000
	MSG_OOB					= 0x4
	MSG_PEEK				= 0x1
	MSG_TRUNC				= 0x200
	MSG_WAITALL				= 0x2
	MS_ASYNC				= 0x1
	MS_INVALIDATE				= 0x4
	MS_SYNC					= 0x2
	NAME_MAX				= 0x200
	// Maybe NFDBITS is 0x20;
	// typedef ULONG               fd_mask;
	// #define NBBY                8
	// #define NFDBITS             (sizeof(fd_mask) * NBBY) /* bits per mask */
	NFDBITS					= 0x40
	NOFLSH					= 0x80
	OCRNL					= 0x8
	OFDEL					= 0x80
	OFILL					= 0x40
	ONLCR					= 0x4
	ONLRET					= 0x20
	ONOCR					= 0x10
	OPOST					= 0x1
	O_ACCMODE				= 0x3
	O_APPEND				= 0x8
	O_ASYNC					= 0x40
	O_BINARY				= 0x10000
	O_CLOEXEC				= 0x80000
	O_CREAT					= 0x200
	O_DIRECTORY				= 0x40000
	O_DSYNC					= 0x20
	O_EXCL					= 0x800
	O_EXLOCK				= 0x80
	O_FSYNC					= 0x2000
	O_LARGEFILE				= 0x100000
	O_NDELAY				= 0x4000
	O_NOCTTY				= 0x8000
	O_NOFOLLOW				= 0x20000
	O_NONBLOCK				= 0x4000
	O_RDONLY				= 0x0
	O_RDWR					= 0x2
	O_SHLOCK				= 0x10
	O_SYNC					= 0x2000
	O_TRUNC					= 0x400
	O_WRONLY				= 0x1
	PARENB					= 0x40
	PARMRK					= 0x8
	PARODD					= 0x80
	PENDIN					= 0x4000
	PRIO_PGRP				= 0x1
	PRIO_PROCESS				= 0x0
	PRIO_USER				= 0x2
	PROT_EXEC				= 0x4
	PROT_NONE				= 0x0
	PROT_READ				= 0x1
	PROT_WRITE				= 0x2
	PR_FASTHZ				= 0x5
	PR_SLOWHZ				= 0x2
	RLIMIT_AS				= 0x9
	RLIMIT_CORE				= 0x4
	RLIMIT_CPU				= 0x0
	RLIMIT_DATA				= 0x2
	RLIMIT_FSIZE				= 0x1
	RLIMIT_NOFILE				= 0x8
	RLIMIT_STACK				= 0x3
	RLIM_INFINITY				= 0x7fffffff
	RTAX_AUTHOR				= 0x6
	RTAX_BRD				= 0x7
	RTAX_DST				= 0x0
	RTAX_GATEWAY				= 0x1
	RTAX_GENMASK				= 0x3
	RTAX_IFA				= 0x5
	RTAX_IFP				= 0x4
	RTAX_MAX				= 0x8
	RTAX_NETMASK				= 0x2
	RTA_AUTHOR				= 0x40
	RTA_BRD					= 0x80
	RTA_DST					= 0x1
	RTA_GATEWAY				= 0x2
	RTA_GENMASK				= 0x8
	RTA_IFA					= 0x20
	RTA_IFP					= 0x10
	RTA_NETMASK				= 0x4
	RTF_BLACKHOLE				= 0x1000
	RTF_BROADCAST				= 0x400000
	RTF_CLONING				= 0x100
	RTF_DONE				= 0x40
	RTF_DYNAMIC				= 0x10
	RTF_GATEWAY				= 0x2
	RTF_HOST				= 0x4
	RTF_LLINFO				= 0x400
	RTF_LOCAL				= 0x200000
	RTF_MODIFIED				= 0x20
	RTF_MULTICAST				= 0x800000
	RTF_PINNED				= 0x100000
	RTF_PRCLONING				= 0x10000
	RTF_PROTO1				= 0x8000
	RTF_PROTO2				= 0x4000
	RTF_PROTO3				= 0x40000
	RTF_REJECT				= 0x8
	RTF_RUNNING				= 0x80000
	RTF_STATIC				= 0x800
	RTF_UP					= 0x1
	RTF_WASCLONED				= 0x20000
	RTF_XRESOLVE				= 0x200
	RTM_ADD					= 0x1
	RTM_CHANGE				= 0x3
	RTM_DELADDR				= 0xd
	RTM_DELETE				= 0x2
	RTM_DELMADDR				= 0x10
	RTM_GET					= 0x4
	RTM_IEEE80211				= 0x12
	RTM_IFANNOUNCE				= 0x11
	RTM_IFINFO				= 0xe
	RTM_LOCK				= 0x8
	RTM_LOSING				= 0x5
	RTM_MISS				= 0x7
	RTM_NEWADDR				= 0xc
	RTM_NEWMADDR				= 0xf
	RTM_OLDADD				= 0x9
	RTM_OLDDEL				= 0xa
	RTM_REDIRECT				= 0x6
	RTM_RESOLVE				= 0xb
	RTM_VERSION				= 0x5
	RTV_EXPIRE				= 0x4
	RTV_HOPCOUNT				= 0x2
	RTV_MTU					= 0x1
	RTV_RPIPE				= 0x8
	RTV_RTT					= 0x40
	RTV_RTTVAR				= 0x80
	RTV_SPIPE				= 0x10
	RTV_SSTHRESH				= 0x20
	RUSAGE_CHILDREN				= -0x1
	RUSAGE_SELF				= 0x0
	SCM_CRED				= 0x3
	SCM_CREDENTIALS				= 0x2
	SCM_RIGHTS				= 0x1
	SHUT_RD					= 0x0
	SHUT_RDWR				= 0x2
	SHUT_WR					= 0x1
	SIGQUEUE_MAX				= 0x1f4
	SIOCADDMULTI				= 0x8030693c
	SIOCADDRT				= 0x8130720a
	SIOCAIFADDR				= 0x8064691a
	SIOCATMARK				= 0x40087307
	SIOCCHGRT				= 0x8130720c
	SIOCDELMULTI				= 0x8030693d
	SIOCDELRT				= 0x8130720b
	SIOCDEVPRIVATE				= 0x89f0
	SIOCDIFADDR				= 0x80306919
	SIOCDIFADDR6				= 0x80106923
	SIOCG7668AESKEY				= 0x402469d5
	SIOCG7668CTX				= 0x402469d4
	SIOCG7668DSTADDR			= 0x402469d3
	SIOCG802154AESKEY			= 0x402469cb
	SIOCG802154CTX				= 0x402469ca
	SIOCG802154PANID			= 0x402469c8
	SIOCG802154SHRTADDR			= 0x402469c9
	SIOCGETRT				= 0xc130720d
	SIOCGETSGCNT				= 0x89e1
	SIOCGETVIFCNT				= 0x89e0
	SIOCGFWOPT				= 0x40887215
	SIOCGHIWAT				= 0x40087301
	SIOCGIFADDR				= 0xc0306921
	SIOCGIFADDR6				= 0xc0106921
	SIOCGIFAUTOCFG				= 0xc0306946
	SIOCGIFBRDADDR				= 0xc0306923
	SIOCGIFCONF				= 0xc0106914
	SIOCGIFDGWADDR				= 0xc0306910
	SIOCGIFDSTADDR				= 0xc0306922
	SIOCGIFDSTADDR6				= 0xc0106922
	SIOCGIFFLAGS				= 0xc0306911
	SIOCGIFHWADDR				= 0xc0306936
	SIOCGIFINDEX				= 0xc0306933
	SIOCGIFMETRIC				= 0xc0306917
	SIOCGIFMTU				= 0xc0306934
	SIOCGIFNAME				= 0xc0306932
	SIOCGIFNETMASK				= 0xc0306925
	SIOCGIFNETMASK6				= 0xc0106925
	SIOCGIFNUM				= 0x40046914
	SIOCGIFPFLAGS				= 0xc0306942
	SIOCGIFSECREG				= 0xc0306944
	SIOCGIFSTATS				= 0xc0b86950
	SIOCGIFTCPAF				= 0xc030693e
	SIOCGIFTCPWND				= 0xc0306940
	SIOCGIFTYPE				= 0x40306931
	SIOCGLOWAT				= 0x40087303
	SIOCGSIZIFCONF				= 0x4004696a
	SIOCGSIZIFREQ6				= 0x4010696a
	SIOCGTCPMSSADJ				= 0x40047214
	SIOCLSTRT				= 0xc020720e
	SIOCLSTRTM				= 0xc020720f
	SIOCPROTOPRIVATE			= 0x89e0
	SIOCS7668AESKEY				= 0x802469d5
	SIOCS7668CTX				= 0x802469d4
	SIOCS7668DSTADDR			= 0x802469d3
	SIOCS802154AESKEY			= 0x802469cb
	SIOCS802154CTX				= 0x802469ca
	SIOCS802154PANID			= 0x802469c8
	SIOCS802154SHRTADDR			= 0x802469c9
	SIOCSFWOPT				= 0x80887215
	SIOCSHIWAT				= 0x80087300
	SIOCSIFADDR				= 0x8030690c
	SIOCSIFADDR6				= 0x8010690c
	SIOCSIFAUTOCFG				= 0x80306947
	SIOCSIFBRDADDR				= 0x80306913
	SIOCSIFDGWADDR				= 0x8030690f
	SIOCSIFDSTADDR				= 0x8030690e
	SIOCSIFDSTADDR6				= 0x8010690e
	SIOCSIFFLAGS				= 0x80306910
	SIOCSIFHWADDR				= 0x80306937
	SIOCSIFMETRIC				= 0x80306918
	SIOCSIFMTU				= 0x80306935
	SIOCSIFNETMASK				= 0x80306916
	SIOCSIFNETMASK6				= 0x80106916
	SIOCSIFPFLAGS				= 0x80306943
	SIOCSIFSECREG				= 0x80306945
	SIOCSIFTCPAF				= 0x8030693f
	SIOCSIFTCPWND				= 0x80306941
	SIOCSLOWAT				= 0x80087302
	SIOCSTCPMSSADJ				= 0x80047214
	SOCK_CLOEXEC				= 0x10000000
	SOCK_DGRAM				= 0x2
	SOCK_NONBLOCK				= 0x20000000
	SOCK_RAW				= 0x3
	SOCK_SEQPACKET				= 0x5
	SOCK_STREAM				= 0x1
	SOL_ICMPV6				= 0x3a
	SOL_IP					= 0x0
	SOL_IPV6				= 0x29
	SOL_PACKET				= 0x107
	SOL_RAW					= 0xff
	SOL_SCTP				= 0x84
	SOL_SOCKET				= 0xfff
	SOL_TCP					= 0x6
	SOL_UDP					= 0x11
	SOL_UDPLITE				= 0x88
	SOMAXCONN				= 0xff
	SO_ACCEPTCONN				= 0x2
	SO_BINDTODEVICE				= 0x100b
	SO_BROADCAST				= 0x20
	SO_CONTIMEO				= 0x1009
	SO_DEBUG				= 0x1
	SO_DONTLINGER				= -0x81
	SO_DONTROUTE				= 0x10
	SO_ERROR				= 0x1007
	SO_KEEPALIVE				= 0x8
	SO_LINGER				= 0x80
	SO_NO_CHECK				= 0x100a
	SO_OOBINLINE				= 0x100
	SO_PASSCRED				= 0x10
	SO_PEERCRED				= 0x11
	SO_PRIORITY				= 0x4002
	SO_RCVBUF				= 0x1002
	SO_RCVLOWAT				= 0x1004
	SO_RCVTIMEO				= 0x1006
	SO_REUSE				= 0x1
	SO_REUSEADDR				= 0x4
	SO_REUSEPORT				= 0x200
	SO_REUSE_RXTOALL			= 0x1
	SO_SECREGION				= 0x4001
	SO_SNDBUF				= 0x1001
	SO_SNDLOWAT				= 0x1003
	SO_SNDTIMEO				= 0x1005
	SO_TYPE					= 0x1008
	SO_USELOOPBACK				= 0x40
	S_IEXEC					= 0x40
	S_IFBLK					= 0x6000
	S_IFCHR					= 0x2000
	S_IFDIR					= 0x4000
	S_IFIFO					= 0x1000
	S_IFLNK					= 0xa000
	S_IFMT					= 0xf000
	S_IFREG					= 0x8000
	S_IFSOCK				= 0xc000
	S_IREAD					= 0x100
	S_IRGRP					= 0x20
	S_IROTH					= 0x4
	S_IRUSR					= 0x100
	S_IRWXG					= 0x38
	S_IRWXO					= 0x7
	S_IRWXU					= 0x1c0
	S_ISGID					= 0x400
	S_ISUID					= 0x800
	S_ISVTX					= 0x200
	S_IWGRP					= 0x10
	S_IWOTH					= 0x2
	S_IWRITE				= 0x80
	S_IWUSR					= 0x80
	S_IXGRP					= 0x8
	S_IXOTH					= 0x1
	S_IXUSR					= 0x40
	TCSETS					= 0 // equivalent to TCSANOW for tcsetattr
	TCSADRAIN				= 1
	TCSETSW					= 1 // equivalent to TCSADRAIN for tcsetattr
	TCSETSF					= 2 // equivalent to TCSAFLUSH for tcsetattr
	TCGETS					= 3 // not defined in termios.h -- sylixos golang only
	TCIFLUSH				= 0x0
	TCIOFLUSH				= 0x2
	TCOFLUSH				= 0x1
	TCP_CALCULATE_EFF_SEND_MSS		= 0x1
	TCP_CWND_DEBUG				= 0x0
	TCP_DEBUG					= 0x0
	TCP_DEFAULT_LISTEN_BACKLOG		= 0xff
	TCP_DESC					= 0x80
	TCP_FR_DEBUG				= 0x0
	TCP_INPUT_DEBUG				= 0x0
	TCP_ISN_KEY_INTERVAL			= 0x493e0
	TCP_KEEPALIVE				= 0x2
	TCP_KEEPCNT					= 0x5
	TCP_KEEPIDLE				= 0x3
	TCP_KEEPINTVL				= 0x4
	TCP_LISTEN_BACKLOG			= 0x1
	TCP_LISTEN_MULTI			= 0x0
	TCP_MAXRTX				= 0x6
	TCP_MPORTS				= 0x81
	TCP_MSS					= 0x1000
	TCP_NODELAY				= 0x1
	TCP_OOSEQ_MAX_BYTES			= 0x0
	TCP_OOSEQ_MAX_PBUFS			= 0x0
	TCP_OUTPUT_DEBUG			= 0x0
	TCP_OVERSIZE				= 0x1000
	TCP_QLEN_DEBUG				= 0x0
	TCP_QUEUE_OOSEQ				= 0x1
	TCP_RCV_SCALE				= 0x2
	TCP_RST_DEBUG				= 0x0
	TCP_RTO_DEBUG				= 0x0
	TCP_SNDLOWAT				= 0x7fff
	TCP_SNDQUEUELOWAT			= 0x20
	TCP_SND_BUF					= 0xffff
	TCP_SND_QUEUELEN			= 0x40
	TCP_STATE_CLOSED			= 0x0
	TCP_STATE_CLOSE_WAIT		= 0x7
	TCP_STATE_CLOSING			= 0x8
	TCP_STATE_ESTABLISHED		= 0x4
	TCP_STATE_FIN_WAIT_1		= 0x5
	TCP_STATE_FIN_WAIT_2		= 0x6
	TCP_STATE_LAST_ACK			= 0x9
	TCP_STATE_LISTEN			= 0x1
	TCP_STATE_SYN_RCVD			= 0x3
	TCP_STATE_SYN_SENT			= 0x2
	TCP_STATE_TIME_WAIT			= 0xa
	TCP_STATS					= 0x1
	TCP_SYNMAXRTX				= 0x4
	TCP_TTL						= 0xff
	TCP_WND						= 0xffff
	TCP_WND_DEBUG				= 0x0
	TCP_WND_UPDATE_THRESHOLD		= 0x0
	TCSAFLUSH				= 0x2
	TIOCGWINSZ				= 0x40087468
	TIOCSWINSZ				= 0x80087467
	TOSTOP					= 0x100
	VDISCARD				= 0xd
	VEOF					= 0x4
	VEOL					= 0xb
	VEOL2					= 0x10
	VERASE					= 0x2
	VINTR					= 0x0
	VKILL					= 0x3
	VLNEXT					= 0xf
	VMIN					= 0x6
	VQUIT					= 0x1
	VREPRINT				= 0xc
	VSTART					= 0x8
	VSTOP					= 0x9
	VSUSP					= 0xa
	VSWTC					= 0x7
	VT0						= 0x0
	VT1						= 0x4000
	VTDLY					= 0x4000
	VTIME					= 0x5
	VWERASE					= 0xe
	WNOHANG					= 0x1
	WUNTRACED				= 0x2
)

// Errors
const (
	E2BIG					= syscall.Errno( 0x7)
	EACCES					= syscall.Errno( 0xd)
	EADDRINUSE				= syscall.Errno( 0x30)
	EADDRNOTAVAIL			= syscall.Errno( 0x31)
	EADV					= syscall.Errno( 0x70)
	EAFNOSUPPORT			= syscall.Errno( 0x2f)
	EAGAIN					= syscall.Errno( 0xb)
	EALREADY				= syscall.Errno( 0x45)
	EBADE					= syscall.Errno( 0x65)
	EBADF					= syscall.Errno( 0x9)
	EBADFD					= syscall.Errno( 0x77)
	EBADMSG					= syscall.Errno( 0x4d)
	EBADR					= syscall.Errno( 0x66)
	EBADRQC					= syscall.Errno( 0x69)
	EBADSLT					= syscall.Errno( 0x6a)
	EBFONT					= syscall.Errno( 0x6b)
	EBUSY					= syscall.Errno( 0x10)
	ECALLEDINISR			= syscall.Errno( 0x130)
	ECANCELED				= syscall.Errno( 0x48)
	ECHILD					= syscall.Errno( 0xa)
	ECHRNG					= syscall.Errno( 0x5d)
	ECOMM					= syscall.Errno( 0x72)
	ECONNABORTED			= syscall.Errno( 0x35)
	ECONNREFUSED			= syscall.Errno( 0x3d)
	ECONNRESET				= syscall.Errno( 0x36)
	EDEADLK					= syscall.Errno( 0x21)
	EDEADLOCK				= syscall.Errno( 0x21)
	EDESTADDRREQ			= syscall.Errno( 0x28)
	EDOM					= syscall.Errno( 0x25)
	EDOTDOT					= syscall.Errno( 0x74)
	EDQUOT					= syscall.Errno( 0x86)
	EEXIST					= syscall.Errno( 0x11)
	EFAULT					= syscall.Errno( 0xe)
	EFBIG					= syscall.Errno( 0x1b)
	EFORMAT					= syscall.Errno( 0x5b)
	EFTYPE					= syscall.Errno( 0x16)
	EHOSTDOWN				= syscall.Errno( 0x43)
	EHOSTUNREACH			= syscall.Errno( 0x41)
	EIDRM					= syscall.Errno( 0x5c)
	EILSEQ					= syscall.Errno( 0x8a)
	EINPROGRESS				= syscall.Errno( 0x44)
	EINTR					= syscall.Errno( 0x4)
	EINVAL					= syscall.Errno( 0x16)
	EIO						= syscall.Errno( 0x5)
	EISCONN					= syscall.Errno( 0x38)
	EISDIR					= syscall.Errno( 0x15)
	EISNAM					= syscall.Errno( 0x84)
	EL2HLT					= syscall.Errno( 0x64)
	EL2NSYNC				= syscall.Errno( 0x5e)
	EL3HLT					= syscall.Errno( 0x5f)
	EL3RST					= syscall.Errno( 0x60)
	ELIBACC					= syscall.Errno( 0x79)
	ELIBBAD					= syscall.Errno( 0x7a)
	ELIBEXEC				= syscall.Errno( 0x7d)
	ELIBMAX					= syscall.Errno( 0x7c)
	ELIBSCN					= syscall.Errno( 0x7b)
	ELNRNG					= syscall.Errno( 0x61)
	ELOOP					= syscall.Errno( 0x40)
	EMEDIUMTYPE				= syscall.Errno( 0x88)
	EMFILE					= syscall.Errno( 0x18)
	EMLINK					= syscall.Errno( 0x1f)
	EMNOTINITED				= syscall.Errno( 0x16)
	EMSGSIZE				= syscall.Errno( 0x24)
	EMULTIHOP				= syscall.Errno( 0x73)
	ENAMETOOLONG			= syscall.Errno( 0x1a)
	ENAVAIL					= syscall.Errno( 0x83)
	ENETDOWN				= syscall.Errno( 0x3e)
	ENETRESET				= syscall.Errno( 0x34)
	ENETUNREACH				= syscall.Errno( 0x33)
	ENFILE					= syscall.Errno( 0x17)
	ENOANO					= syscall.Errno( 0x68)
	ENOBUFS					= syscall.Errno( 0x37)
	ENOCSI					= syscall.Errno( 0x63)
	ENODATA					= syscall.Errno( 0x4e)
	ENODEV					= syscall.Errno( 0x13)
	ENOENT					= syscall.Errno( 0x2)
	ENOEXEC					= syscall.Errno( 0x8)
	ENOIOCTLCMD				= syscall.Errno( 0x47)
	ENOLCK					= syscall.Errno( 0x22)
	ENOLINK					= syscall.Errno( 0x6f)
	ENOMEDIUM				= syscall.Errno( 0x87)
	ENOMEM					= syscall.Errno( 0xc)
	ENOMSG					= syscall.Errno( 0x50)
	ENONET					= syscall.Errno( 0x6c)
	ENOPKG					= syscall.Errno( 0x6d)
	ENOPROTOOPT				= syscall.Errno( 0x2a)
	ENOSPC					= syscall.Errno( 0x1c)
	ENOSR					= syscall.Errno( 0x4a)
	ENOSTR					= syscall.Errno( 0x4b)
	ENOSYS					= syscall.Errno( 0x47)
	ENOTBLK					= syscall.Errno( 0x42)
	ENOTCONN				= syscall.Errno( 0x39)
	ENOTDIR					= syscall.Errno( 0x14)
	ENOTEMPTY				= syscall.Errno( 0xf)
	ENOTNAM					= syscall.Errno( 0x82)
	ENOTRECOVERABLE			= syscall.Errno( 0x8d)
	ENOTSOCK				= syscall.Errno( 0x32)
	ENOTSUP					= syscall.Errno( 0x23)
	ENOTTY					= syscall.Errno( 0x19)
	ENOTUNIQ				= syscall.Errno( 0x76)
	ENXIO					= syscall.Errno( 0x6)
	EOPNOTSUPP				= syscall.Errno( 0x2d)
	EOVERFLOW				= syscall.Errno( 0x7)
	EOWNERDEAD				= syscall.Errno( 0x8c)
	EPERM					= syscall.Errno( 0x1)
	EPFNOSUPPORT			= syscall.Errno( 0x2e)
	EPIPE					= syscall.Errno( 0x20)
	EPROTO					= syscall.Errno( 0x4c)
	EPROTONOSUPPORT			= syscall.Errno( 0x2b)
	EPROTOTYPE				= syscall.Errno( 0x29)
	ERANGE					= syscall.Errno( 0x26)
	EREMCHG					= syscall.Errno( 0x78)
	EREMOTE					= syscall.Errno( 0x6e)
	EREMOTEIO				= syscall.Errno( 0x85)
	ERESTART				= syscall.Errno( 0x7e)
	EROFS					= syscall.Errno( 0x1e)
	ESHUTDOWN				= syscall.Errno( 0x3a)
	ESOCKTNOSUPPORT			= syscall.Errno( 0x2c)
	ESPIPE					= syscall.Errno( 0x1d)
	ESRCH					= syscall.Errno( 0x3)
	ESRMNT					= syscall.Errno( 0x71)
	ESTALE					= syscall.Errno( 0x81)
	ESTRPIPE				= syscall.Errno( 0x7f)
	ETIME					= syscall.Errno( 0x4f)
	ETIMEDOUT				= syscall.Errno( 0x3c)
	ETOOMANYREFS			= syscall.Errno( 0x3b)
	ETXTBSY					= syscall.Errno( 0x3f)
	EUCLEAN					= syscall.Errno( 0x75)
	EUNATCH					= syscall.Errno( 0x62)
	EUSERS					= syscall.Errno( 0x80)
	EWOULDBLOCK				= syscall.Errno( 0xb)
	EWRPROTECT				= syscall.Errno( 0x5a)
	EXDEV					= syscall.Errno( 0x12)
	EXFULL					= syscall.Errno( 0x67)
	EXIT_FAILURE			= syscall.Errno( 0x1)
	EXIT_SUCCESS			= syscall.Errno( 0x0)
)

// Signals
const (
	SIGABRT					= syscall.Signal( 0x6)
	SIGALRM					= syscall.Signal( 0xe)
	SIGBUS					= syscall.Signal( 0xa)
	SIGCANCEL				= SIGTERM
	SIGCHLD					= syscall.Signal( 0x14)
	SIGCNCL					= syscall.Signal( 0x10)
	SIGCONT					= syscall.Signal( 0x13)
	SIGFPE					= syscall.Signal( 0x8)
	SIGHUP					= syscall.Signal( 0x1)
	SIGILL					= syscall.Signal( 0x4)
	SIGINFO					= syscall.Signal( 0x1d)
	SIGINT					= syscall.Signal( 0x2)
	SIGIO					= syscall.Signal( 0x17)
	SIGKILL					= syscall.Signal( 0x9)
	SIGLOWMEM				= syscall.Signal( 0x2e)
	SIGPIPE					= syscall.Signal( 0xd)
	SIGPOLL					= SIGIO
	SIGPROF					= syscall.Signal( 0x1b)
	SIGPWR					= syscall.Signal( 0x21)
	SIGQUIT					= syscall.Signal( 0x3)
	SIGSEGV					= syscall.Signal( 0xb)
	SIGSTOP					= syscall.Signal( 0x11)
	SIGSYS					= syscall.Signal( 0x22)
	SIGTERM					= syscall.Signal( 0xf)
	SIGTRAP					= syscall.Signal( 0x5)
	SIGTSTP					= syscall.Signal( 0x12)
	SIGTTIN					= syscall.Signal( 0x15)
	SIGTTOU					= syscall.Signal( 0x16)
	SIGUNUSED				= syscall.Signal( 0x7)
	SIGUNUSED2				= syscall.Signal( 0xc)
	SIGURG					= syscall.Signal( 0x23)
	SIGUSR1					= syscall.Signal( 0x1e)
	SIGUSR2					= syscall.Signal( 0x1f)
	SIGVTALRM				= syscall.Signal( 0x1a)
	SIGWINCH				= syscall.Signal( 0x1c)
	SIGXCPU					= syscall.Signal( 0x18)
	SIGXFSZ					= syscall.Signal( 0x19)
)

// Error table
var errorList = [...]struct {
	num  syscall.Errno
	name string
	desc string
}{
	{0, "EXIT_SUCCESS",  "no error"},
	// {1, "EPERM|EXIT_FAILURE",  "not owner"},
	{1, "EPERM",  "not owner"},
	{2, "ENOENT",  "no such file or directory"},
	{3, "ESRCH",  "no such process"},
	{4, "EINTR",  "interrupted system call"},
	{5, "EIO",  "I/O error"},
	{6, "ENXIO",  "no such device or address"},
	{7, "E2BIG|EOVERFLOW",  "arg list too long or overflow"},
	{8, "ENOEXEC",  "exec format error"},
	{9, "EBADF",  "bad file number"},
	{10, "ECHILD",  "no children"},
	{11, "EAGAIN|EWOULDBLOCK",  "no more processes or operation would block"},
	{12, "ENOMEM",  "no enough memory"},
	{13, "EACCES",  "permission denied or can not access"},
	{14, "EFAULT",  "bad address"},
	{15, "ENOTEMPTY",  "directory not empty"},
	{16, "EBUSY",  "mount device busy"},
	{17, "EEXIST",  "file exists"},
	{18, "EXDEV",  "cross-device link"},
	{19, "ENODEV",  "no such device"},
	{20, "ENOTDIR",  "not a directory"},
	{21, "EISDIR",  "is a directory"},
	// {22, "EFTYPE|EINVAL|EMNOTINITED",  "invalid argument or format"},
	{22, "EINVAL",  "invalid argument or format"},
	{23, "ENFILE",  "file table overflow"},
	{24, "EMFILE",  "too many open files"},
	{25, "ENOTTY",  "not a typewriter"},
	{26, "ENAMETOOLONG",  "file name too long"},
	{27, "EFBIG",  "file too large"},
	{28, "ENOSPC",  "no space left on device"},
	{29, "ESPIPE",  "illegal seek"},
	{30, "EROFS",  "read-only file system"},
	{31, "EMLINK",  "too many links"},
	{32, "EPIPE",  "broken pipe"},
	{33, "EDEADLOCK|EDEADLK",  "resource deadlock avoided"},
	{34, "ENOLCK",  "no locks available"},
	{35, "ENOTSUP",  "unsupported value"},
	{36, "EMSGSIZE",  "message size"},
	{37, "EDOM",  "argument too large"},
	{38, "ERANGE",  "result too large"},
	{40, "EDESTADDRREQ",  "destination address required"},
	{41, "EPROTOTYPE",  "protocol wrong type for socket"},
	{42, "ENOPROTOOPT",  "protocol not available"},
	{43, "EPROTONOSUPPORT",  "protocol not supported"},
	{44, "ESOCKTNOSUPPORT",  "socket type not supported"},
	{45, "EOPNOTSUPP",  "operation not supported on socket"},
	{46, "EPFNOSUPPORT",  "protocol family not supported"},
	{47, "EAFNOSUPPORT",  "addr family not supported"},
	{48, "EADDRINUSE",  "address already in use"},
	{49, "EADDRNOTAVAIL",  "can't assign requested address"},
	{50, "ENOTSOCK",  "socket operation on non-socket"},
	{51, "ENETUNREACH",  "network unreachable"},
	{52, "ENETRESET",  "network dropped connection on reset"},
	{53, "ECONNABORTED",  "software caused connection abort"},
	{54, "ECONNRESET",  "connection reset by peer"},
	{55, "ENOBUFS",  "no buffer space available"},
	{56, "EISCONN",  "socket is already connected"},
	{57, "ENOTCONN",  "socket is not connected"},
	{58, "ESHUTDOWN",  "can't send after socket shutdown"},
	{59, "ETOOMANYREFS",  "too many references can't splice"},
	{60, "ETIMEDOUT",  "connection timed out"},
	{61, "ECONNREFUSED",  "connection refused"},
	{62, "ENETDOWN",  "network is down"},
	{63, "ETXTBSY",  "text file busy"},
	{64, "ELOOP",  "too many levels of symbolic links"},
	{65, "EHOSTUNREACH",  "host unreachable"},
	{66, "ENOTBLK",  "block device required"},
	{67, "EHOSTDOWN",  "host is down"},
	{68, "EINPROGRESS",  "operation now in progress"},
	{69, "EALREADY",  "operation already in progress"},
	{71, "ENOSYS|ENOIOCTLCMD",  "function not implemented"},
	{72, "ECANCELED",  "operation canceled"},
	{74, "ENOSR",  "insufficient memory"},
	{75, "ENOSTR",  "STREAMS device required"},
	{76, "EPROTO",  "generic STREAMS error"},
	{77, "EBADMSG",  "invalid STREAMS message"},
	{78, "ENODATA",  "missing expected message data"},
	{79, "ETIME",  "STREAMS timeout occurred"},
	{80, "ENOMSG",  "unexpected message type"},
	{90, "EWRPROTECT",  "write protect"},
	{91, "EFORMAT",  "invalid format"},
	{92, "EIDRM",  "identifier removed"},
	{93, "ECHRNG",  "channel number out of range"},
	{94, "EL2NSYNC",  "level 2 not synchronized"},
	{95, "EL3HLT",  "level 3 halted"},
	{96, "EL3RST",  "level 3 reset"},
	{97, "ELNRNG",  "link number out of range"},
	{98, "EUNATCH",  "protocol driver not attached"},
	{99, "ENOCSI",  "no CSI structure available"},
	{100, "EL2HLT",  "level 2 halted"},
	{101, "EBADE",  "invalid exchange"},
	{102, "EBADR",  "invalid request descriptor"},
	{103, "EXFULL",  "exchange full"},
	{104, "ENOANO",  "no anode"},
	{105, "EBADRQC",  "invalid request code"},
	{106, "EBADSLT",  "invalid slot"},
	{107, "EBFONT",  "bad font file format"},
	{108, "ENONET",  "machine is not on the network"},
	{109, "ENOPKG",  "package not installed"},
	{110, "EREMOTE",  "object is remote"},
	{111, "ENOLINK",  "link has been severed"},
	{112, "EADV",  "advertise error"},
	{113, "ESRMNT",  "srmount error"},
	{114, "ECOMM",  "communication error on send"},
	{115, "EMULTIHOP",  "multihop attempted"},
	{116, "EDOTDOT",  "RFS specific error"},
	{117, "EUCLEAN",  "structure needs cleaning"},
	{118, "ENOTUNIQ",  "name not unique on network"},
	{119, "EBADFD",  "file descriptor in bad state"},
	{120, "EREMCHG",  "remote address changed"},
	{121, "ELIBACC",  "can not access a needed shared library"},
	{122, "ELIBBAD",  "accessing a corrupted shared library"},
	{123, "ELIBSCN",  ".lib section in a.out corrupted"},
	{124, "ELIBMAX",  "attempting to link in too many shared libraries"},
	{125, "ELIBEXEC",  "cannot exec a shared library directly"},
	{126, "ERESTART",  "interrupted system call should be restarted"},
	{127, "ESTRPIPE",  "streams pipe error"},
	{128, "EUSERS",  "too many users"},
	{129, "ESTALE",  "stale NFS file handle"},
	{130, "ENOTNAM",  "not a XENIX named type file"},
	{131, "ENAVAIL",  "no XENIX semaphores available"},
	{132, "EISNAM",  "is a named type file"},
	{133, "EREMOTEIO",  "remote I/O error"},
	{134, "EDQUOT",  "quota exceeded"},
	{135, "ENOMEDIUM",  "no medium found"},
	{136, "EMEDIUMTYPE",  "wrong medium type"},
	{138, "EILSEQ",  "illegal byte sequence"},
	{140, "EOWNERDEAD",  "owner died"},
	{141, "ENOTRECOVERABLE",  "state not recoverable"},
	{304, "ECALLEDINISR",  "kernel in interrupt service mode"},
}

// Signal table
var signalList = [...]struct {
	num  syscall.Signal
	name string
	desc string
}{ 
	{1, "SIGHUP", "hangup"},
	{2, "SIGINT", "interrupt"},
	{3, "SIGQUIT", "quit"},
	{4, "SIGILL", "illegal instruction"},
	{5, "SIGTRAP", "trace/breakpoint trap"},
	{6, "SIGABRT", "aborted"},
	{7, "SIGUNUSED", "unused signal 1"},
	{8, "SIGFPE", "floating point exception"},
	{9, "SIGKILL", "killed"},
	{10, "SIGBUS", "bus error"},
	{11, "SIGSEGV", "segmentation fault"},
	{12, "SIGUNUSED2", "unused signal 2"},
	{13, "SIGPIPE", "broken pipe"},
	{14, "SIGALRM", "alarm clock"},
	{15, "SIGTERM", "terminated"},
	{16, "SIGCNCL", "pthreads cancellation"},
	{17, "SIGSTOP", "stopped (signal)"},
	{18, "SIGTSTP", "stopped"},
	{19, "SIGCONT", "continued"},
	{20, "SIGCHLD", "child exited"},
	{21, "SIGTTIN", "stopped (tty input)"},
	{22, "SIGTTOU", "stopped (tty output)"},
	{23, "SIGIO", "I/O possible"},
	{24, "SIGXCPU", "CPU time limit exceeded"},
	{25, "SIGXFSZ", "file size limit exceeded"},
	{26, "SIGVTALRM", "virtual timer expired"},
	{27, "SIGPROF", "profiling timer expired"},
	{28, "SIGWINCH", "window changed"},
	{29, "SIGINFO", "information request"},
	{30, "SIGUSR1", "user defined signal 1"},
	{31, "SIGUSR2", "user defined signal 2"},
	{33, "SIGPWR", "power failure"},
	{34, "SIGSYS", "bad system call"},
	{35, "SIGURG", "urgent I/O condition"},
	{46, "SIGLOWMEM", "low memory"},
}
