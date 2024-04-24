package v1

// CreateStubInterfaceRequest
type CreateStubInterfaceRequest struct {
	Url           string `json:"url" valid:"required,stringlength(1|128)"`
	DefRespCode   string `json:"def_resp_code" valid:"required,stringlength(1|16)"`
	DefRespHeader string `json:"def_resp_header,omitempty"`
	DefRespBody   string `json:"def_resp_body" valid:"required"`
	Owner         string `json:"owner" valid:"required,stringlength(1|64)"`
	Desc          string `json:"desc,omitempty"`
}

// CreateStubRuleeRequest
type CreateStubRuleRequest struct {
	Url         string `json:"url" valid:"required,stringlength(1|128)"`
	InterfaceId int32  `json:"interface_id" valid:"required"`
	MatchType   int32  `json:"match_type" valid:"required"`
	MatchRule   string `json:"match_rule" valid:"required,stringlength(1|512)"`
	RespCode    string `json:"resp_code" valid:"required,stringlength(1|16)"`
	RespHeader  string `json:"resp_header,omitempty"`
	RespBody    string `json:"resp_body" valid:"required"`
	DelayTime   int32  `json:"delay_time,omitempty"`
	Desc        string `json:"desc,omitempty"`
}
