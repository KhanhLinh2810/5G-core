package types

type SNssai struct {
	Sst int    `json:"sst"`
	Sd  string `json:"sd"`
}

type TypeCreateSessionRequest struct {
	Supi         string          `json:"supi"`
	Gpsi         string          `json:"gpsi"`
	PduSessionID int             `json:"pduSessionId"`
	Dnn          string          `json:"dnn"`
	SNssai       SNssai   `json:"sNssai"`
	ServingNHd   string          `json:"servingNHd"`
	AnType       string          `json:"anType"`
}


/* 
supi'^(imsi-[0-9]{5,15}|nai-.+|gci-.+|gli-.+|.+)$'
gpsi '^(msisdn-[0-9]{5,15}|extid-[^@]+@[^@]+|.+)$'
pdu type: integer
  minimum: 0
  maximum: 255

sst:
  type: integer
  minimum: 0
  maximum: 255
		  
sd:
  type: string
  pattern: '^[A-Fa-f0-9]{6}$'

*/
