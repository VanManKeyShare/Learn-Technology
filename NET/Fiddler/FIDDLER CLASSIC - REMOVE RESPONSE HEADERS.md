# FIDDLER CLASSIC - REMOVE RESPONSE HEADERS

### BƯỚC 01: TÌM FUNCTION = ONBEFORERESPONSE(oSession: Session)

### BƯỚC 02: THÊM LỆNH SAU

```JS
var HOST_LIST = new HostList("SITES.GOOGLE.COM");
if (HOST_LIST.ContainsHost(oSession.hostname) && oSession.oResponse.headers.ExistsAndContains("Content-Type","text/html")){
	oSession.oResponse.headers.Remove("Content-Security-Policy");
	oSession.oResponse.headers.Remove("Cross-Origin-Opener-Policy");
	oSession.oResponse.headers.Remove("Referrer-Policy");
	oSession.oResponse.headers.Remove("X-Content-Type-Options");
	oSession.oResponse.headers.Remove("X-Frame-Options");
	oSession.oResponse.headers.Remove("X-XSS-Protection");
}
```