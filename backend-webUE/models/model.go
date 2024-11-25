package models

// UE Profile
type UeProfile struct {
	// IMSI number of the UE. IMSI = [MCC|MNC|MSISDN] (In total 15 digits)
	Supi string `json:"supi" bson:"supi"`

	PlmnId          PlmnId   `json:"plmnid" bson:"plmnid"`
	ConfiguredSlice []Snssai `json:"configuredSlice" bson:"configuredSlice"`
	DefaultSlice    []Snssai `json:"defaultSlice" bson:"defaultSlice"`

	// SUCI Protection Scheme : 0 for Null-scheme, 1 for Profile A and 2 for Profile B
	ProtectionScheme int `json:"protectionScheme" bson:"protectionScheme"`
	// Home Network Public Key for protecting with SUCI Profile A
	HomeNetworkPublicKey string `json:"homeNetworkPublicKey" bson:"homeNetworkPublicKey"`
	// Home Network Public Key ID for protecting with SUCI Profile A
	HomeNetworkPublicKeyId int `json:"homeNetworkPublicKeyId" bson:"homeNetworkPublicKeyId"`
	//Routing Indicator
	RoutingIndicator string `json:"routingIndicator" bson:"routingIndicator"`
	//Permanent subscription key
	Key string `json:"key" bson:"key"`
	// Operator code (OP or OPC) of the UE
	Op string `json:"op" bson:"op"`
	// This value specifies the OP type and it can be either 'OP' or 'OPC'
	OpType string `json:"opType" bson:"opType"`
	// Authentication Management Field (AMF) value
	Amf string `json:"amf" bson:"amf"`
	// IMEI number of the device. It is used if no SUPI is provided
	Imei string `json:"imei" bson:"imei"`
	// IMEISV number of the device. It is used if no SUPI and IMEI is provided
	Imeisv string `json:"imeiSv" bson:"imeiSv"`
	// List of gNB IP addresses for Radio Link Simulation
	GnbSearchList []string `json:"gnbSearchList" bson:"gnbSearchList"`

	Integrity Integrity `json:"integrity" bson:"integrity"`
	Ciphering Ciphering `json:"ciphering" bson:"ciphering"`

	// UAC Access Identities Configuration
	UacAic struct {
		Mps bool `json:"mps" bson:"mps"`
		Mcs bool `json:"mcs" bson:"mcs"`
	} `json:"uacAic" bson:"uacAic"`

	// UAC Access Control Class
	UacAcc struct {
		NormalClass int  `json:"normalClass" bson:"normalClass"`
		Class11     bool `json:"class11" bson:"class11"`
		Class12     bool `json:"class12" bson:"class12"`
		Class13     bool `json:"class13" bson:"class13"`
		Class14     bool `json:"class14" bson:"class14"`
		Class15     bool `json:"class15" bson:"class15"`
	} `json:"uacAcc" bson:"uacAcc"`

	//Initial PDU sessions to be established
	Sessions []struct {
		Type  string `json:"type" bson:"type"`
		Apn   string `json:"apn" bson:"apn"`
		Slice struct {
			Sst int `json:"sst" bson:"sst"`
			Sd  int `json:"sd" bson:"sd"`
		} `json:"slice" bson:"slice"`
	} `json:"sessions" bson:"sessions"`

	IntegrityMaxRate struct {
		Uplink   string `json:"uplink" bson:"uplink"`
		Downlink string `json:"downlink" bson:"downlink"`
	} `json:"integrityMaxRate" bson:"integrityMaxRate"`

	// Public key of the UE
	PublicKey string `json:"publicKey" bson:"publicKey"`
	// Private key of the UE
	PrivateKey string `json:"privateKey" bson:"privateKey"`
}
type PlmnId struct {
	// Mobile Country Code value of HPLMN
	Mcc string `json:"mcc" bson:"mcc"`
	// Mobile Network Code value of HPLMN (2 or 3 digits)
	Mnc string `json:"mnc" bson:"mnc"`
}

// Integrity algorithms by UE
type Integrity struct {
	IA1 bool `json:"IA1" bson:"IA1"`
	IA2 bool `json:"IA2" bson:"IA2"`
	IA3 bool `json:"IA3" bson:"IA3"`
}

// Ciphering algorithms by UE
type Ciphering struct {
	EA1 bool `json:"EA1" bson:"EA1"`
	EA2 bool `json:"EA2" bson:"EA2"`
	EA3 bool `json:"EA3" bson:"EA3"`
}

type Snssai struct {
	Sst int    `json:"sst" bson:"sst"`
	Sd  string `json:"sd" bson:"sd"`
}
