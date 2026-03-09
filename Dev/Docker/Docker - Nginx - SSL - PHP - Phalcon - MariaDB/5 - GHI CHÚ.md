# KẾT NỐI CSDL MARIADB KHI VIẾT LỆNH PHP
```code
SỬ DỤNG TÊN CONTAINER CỦA MARIADB TRONG DOCKER LÀM TÊN HOST KHI VIẾT LỆNH PHP
```

# CẤU HÌNH phpMyAdmin

ĐỔI TÊN FILE "CONFIG.EXAMPLE.INC.PHP" THÀNH "CONFIG.INC.PHP"

SAU ĐÓ CẤU HÌNH CƠ BẢN NHƯ SAU:

```php
$cfg['CheckConfigurationPermissions'] = false;
$cfg['blowfish_secret'] = '< 32 KÝ TỰ CHỮ VÀ SỐ >';
$cfg['Servers'][$i]['host'] = '< TÊN CONTAINER MARIADB TRONG DOCKER >';
```