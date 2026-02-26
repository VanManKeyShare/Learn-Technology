# FIDDLER CLASSIC - JAVASCRIPT INJECTION INTO HTML DOCUMENTS

### BƯỚC 01: TÌM FUNCTION = ONBEFORERESPONSE(oSession: Session)

### BƯỚC 02: THÊM LỆNH SAU

```JS
var INJECT_JS = "<script>alert('Hello')</script>";
var HOST_LIST = new HostList("*.YOUTUBE.COM");
if (HOST_LIST.ContainsHost(oSession.hostname) && oSession.oResponse.headers.ExistsAndContains("Content-Type","text/html")) {
	oSession.utilDecodeResponse();
	oSession.utilReplaceInResponse('</head>', INJECT_JS + '</head>');
}
```
