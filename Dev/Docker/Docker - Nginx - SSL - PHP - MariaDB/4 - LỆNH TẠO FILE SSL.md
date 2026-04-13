# CÀI ĐẶT OPENSSL

- ĐỊA CHỈ TẢI OPENSSL: [https://slproweb.com/download/Win64OpenSSL-3_6_1.msi](https://slproweb.com/download/Win64OpenSSL-3_6_1.msi)
- SAU KHI CÀI ĐẶT THÌ MỞ FILE "SYSDM.CPL"
- CHỌN THẺ "ADVANCED" NHẤN NÚT "ENVIRONMENT VARIABLES"
- CHỌN "PATH" TRONG DANH SÁCH "SYSTEM VARIABLES" NHẤN NÚT "EDIT"
- NHẤN NÚT "NEW" VÀ THÊM "C:\Program Files\OpenSSL-Win64\bin" VÀO DANH SÁCH
- MỞ TERMINAL CHẠY LỆNH "openssl version" ĐỂ KIỂM TRA

## LỆNH TẠO KEY FILE SSL

MỞ TERMINAL CHẠY LỆNH SAU

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout vmk.key -out vmk.crt
```
