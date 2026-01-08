# AUTO LOGIN - AUTO START - WINE APP

> XUBUNTU v15.10 X86
> RELEASE 2026-01-08

## 📜 TẢI VÀ CÀI ĐẶT HỆ ĐIỀU HÀNH -> XUBUNTU v15.10 X86

CẤU HÌNH CÁC THÔNG SỐ CƠ BẢN NHƯ SAU:

```code
USER = vmk
HOSTNAME = MyVPS
-> NHỚ TÍCH CHỌN AUTO LOGIN CHO USER VỪA TẠO Ở TRÊN
```

## 📜 KẾT NỐI ĐẾN VPS SỬ DỤNG NOVNC VÀ CẤU HÌNH CƠ BẢN

- CHANGE SCREEN RESOLUTION TO 1024 x 768
- CHANGE SETTING FOR UBUNTU AUTO DOWNLOAD & INSTALL UPDATE

## 📜 THAY ĐỔI MẬT KHẨU USER ROOT

```bash
sudo -i
passwd
exit
```

## 📜 NÂNG CẤP HỆ THỐNG VÀ CÀI ĐẶT ỨNG DỤNG OPENSSH

```bash
su - root
sudo apt-get update
sudo apt-get install openssh-server openssh-client
reboot
```

## 📜 SỬ DỤNG PUTTY KẾT NỐI ĐẾN VPS

```bash
su - root
```

## 📜 CÀI ĐẶT ỨNG DỤNG FIREFOX, WINE VÀ TEAMVIEWER

```bash
su - root

sudo apt-get update
sudo apt-get install firefox
sudo apt-get install wine

cd /home/vmk
wget http://download.teamviewer.com/download/teamviewer_i386.deb
sudo dpkg -i teamviewer_i386.deb
sudo apt-get install -f
sudo dpkg -i teamviewer_i386.deb

sudo apt-get upgrade
sudo apt-get clean

reboot
```

## 📜 KẾT NỐI ĐẾN VPS SỬ DỤNG NOVNC VÀ CẤU HÌNH TEAMVIEWER, WINE + CÀI ĐẶT MT4

```code
1. RUN TEAMVIEWER AND ACCEPT LICENSE
   - LOGIN ACCOUNT TEAMVIEW
   - PLEASE CHECK "KEEP ME SIGNED IN"
   - ADD SERVER TO LIST PARTNER

2. RUN WINE AND CONFIG IT

3. USE TEAMVIEWER OR WINSCP CONNECT AND SEND FILE TO SERVER
   - INSTALL MT4 TO "C:\MT4_01"
```

## 📜 KẾT NỐI ĐẾN VPS SỬ DỤNG PUTTY

```bash
su - root
```

## 📜 THIẾT LẬP CẤU HÌNH AUTO KHỞI ĐỘNG CÙNG HỆ THỐNG

```bash
mkdir /home/vmk/.config
mkdir /home/vmk/.config/autostart

chmod u=rwx,g=rwx,o=rwx /home/vmk/.config
chmod u=rwx,g=rwx,o=rwx /home/vmk/.config/autostart
```

## 📜 THIẾT LẬP TEAMVIEWER AUTO KHỞI ĐỘNG CÙNG HỆ THỐNG

```bash
nano /home/vmk/.config/autostart/TeamViewer.desktop
```

```code
[Desktop Entry]
Encoding=UTF-8
Name=Teamviewer
Comment=Teamviewer
Type=Application
OnlyShowIn=XFCE;
StartupNotify=false
Terminal=false
Hidden=false
Exec=/usr/bin/teamviewer
```

```bash
chmod u=rwx,g=rwx,o=rwx /home/vmk/.config/autostart/TeamViewer.desktop
```

## 📜 TẠO SCRIPT KHỞI ĐỘNG LẠI ỨNG DỤNG MT4

```bash
nano /home/vmk/MT4.sh
```

```code
#!/bin/sh
pkill terminal.exe &
pkill Terminal.exe &
env DISPLAY=:0 wine "C:\\MT4_01\\Terminal.exe" &
#env DISPLAY=:0 wine "C:\\MT4_02\\Terminal.exe" &
```

```bash
chmod u=rwx,g=rwx,o=rwx /home/vmk/MT4.sh
```

## 📜 THIẾT LẬP MT4 AUTO KHỞI ĐỘNG CÙNG HỆ THỐNG

```bash
nano /home/vmk/.config/autostart/MT4.desktop
```

```code
[Desktop Entry]
Encoding=UTF-8
Name=MT4
Comment=MT4
Type=Application
OnlyShowIn=XFCE;
StartupNotify=false
Terminal=false
Hidden=false
Exec=/home/vmk/MT4.sh
```

```bash
chmod u=rwx,g=rwx,o=rwx /home/vmk/.config/autostart/MT4.desktop
```

## 📜 THIẾT LẬP 'XYZ.exe' AUTO KHỞI ĐỘNG CÙNG HỆ THỐNG

```code
SAO CHÉP ỨNG DỤNG 'XYZ.exe' VÀO THƯ MỤC "/home/vmk/.wine/drive_c/XYZ"
```

```bash
nano /home/vmk/.config/autostart/XYZ.desktop
```

```code
[Desktop Entry]
Encoding=UTF-8
Name=XYZ
Comment=XYZ
Type=Application
OnlyShowIn=XFCE;
StartupNotify=false
Terminal=false
Hidden=false
Exec=env DISPLAY=:0 WINEPREFIX="/home/vmk/.wine" wine C:\\\\XYZ\\\\XYZ.exe
Path=/home/vmk/.wine/dosdevices/c:/XYZ/
```

```bash
chmod u=rwx,g=rwx,o=rwx /home/vmk/.config/autostart/XYZ.desktop
```

## 📜 CẤU HÌNH TASK SCHEDULER AUTO REBOOT VÀO 11:58 VÀ 23:58 HÀNG NGÀY

```bash
nano /etc/crontab
```

```code
58 23 * * * root /sbin/reboot
58 11 * * * root /sbin/reboot
```

```bash
reboot
```
