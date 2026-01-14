# MariaDB -> TẠO USER MỚI, CẤP QUYỀN, XÓA USER

## 1. ĐĂNG NHẬP VÀO MariaDB DƯỚI QUYỀN ROOT

```bash
C:\MariaDB\bin>mysql -u root -p
```

## 2. TẠO VÀ CẤP QUYỀN CHO USER 'vmk' QUA LOCALHOST

- TẠO USER 'vmk' VỚI MẬT KHẨU 9999
- CẤP TOÀN BỘ QUYỀN TRÊN TẤT CẢ DATABASE CHO USER 'vmk'
- LÀM MỚI BẢNG PHÂN QUYỀN ĐỂ CÁC THAY ĐỔI CÓ HIỆU LỰC NGAY LẬP TỨC MÀ KHÔNG CẦN KHỞI ĐỘNG LẠI DỊCH VỤ

```bash
CREATE USER 'vmk'@'localhost' IDENTIFIED BY '9999';
GRANT ALL PRIVILEGES ON *.* TO 'vmk'@'localhost' IDENTIFIED BY '9999';
FLUSH PRIVILEGES;
```

## 3. TẠO VÀ CẤP QUYỀN CHO USER 'vmk' QUA IP 127.0.0.1

- TẠO USER 'vmk' VỚI MẬT KHẨU 9999
- CẤP TOÀN BỘ QUYỀN TRÊN TẤT CẢ DATABASE CHO USER 'vmk'
- LÀM MỚI BẢNG PHÂN QUYỀN ĐỂ CÁC THAY ĐỔI CÓ HIỆU LỰC NGAY LẬP TỨC MÀ KHÔNG CẦN KHỞI ĐỘNG LẠI DỊCH VỤ

```bash
CREATE USER 'vmk'@'127.0.0.1' IDENTIFIED BY '9999';
GRANT ALL PRIVILEGES ON *.* TO 'vmk'@'127.0.0.1' IDENTIFIED BY '9999';
FLUSH PRIVILEGES;
```

## 4. XÓA USER

```bash
DROP USER 'vmk'@'localhost';
DROP USER 'vmk'@'127.0.0.1';
```
