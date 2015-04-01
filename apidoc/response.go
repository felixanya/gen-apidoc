package apidoc

import (
	"net/http"
	"net/http/httptest"
)

func (ad *ApiDefine) SetResponse(resp *httptest.ResponseRecorder, ok interface{}, ng interface{}) {
	if resp.Code == http.StatusOK && ok != nil {
		ad.SetSuccess(ok)
	} else if ng != nil {
		//ad.SetError(ng)
	}
}
