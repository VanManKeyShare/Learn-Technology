
// VnEdu -> LỆNH KIỂM TRA ĐÃ NHẬP ĐẦY ĐỦ SỔ ĐẦU BÀI HAY CHƯA

const tuanHoc = '20';
const app_nam_hoc = '2025';
const my_user_id = 'XXXXXX';
const my_token = 'XXXXXX';
const vnedu_hosname = 'https://MMZMCCCUSSITESSGDBINHDUONG.VNEDU.VN';

const Danh_Sach_Ten_Lop_Can_Lay = '7A15|7A16|7A17|7A18|7A19|7A20|7A21|7A22|7A23|7A24|8A14|8A15|8A16|8A17|8A18|8A19|8A20|8A21|8A22|8A23|8A24|8A25|8A26';

const Key_TimKiem_SoDauBaiChuaHoanThanh = "class='add add_tiet_so_dau_bai'";

var Array_Ten_Lop_Can_Lay_Temp = Danh_Sach_Ten_Lop_Can_Lay.split('|');

var Array_ID_Lop_Can_Lay = [];
var Array_Ten_Lop_Can_Lay = [];
var Array_Khoi_Lop_Can_Lay = [];

function RANDOM_BIGINT(min, max) {
  // CHUYỂN ĐỔI ĐẦU VÀO SANG BIGINT ĐỂ ĐẢM BẢO ĐỘ CHÍNH XÁC
  const bMin = BigInt(min);
  const bMax = BigInt(max);

  // TÍNH TOÁN PHẠM VI (RANGE)
  const range = bMax - bMin + 1n;

  // TẠO SỐ NGẪU NHIÊN AN TOÀN BẰNG CÁCH KẾT HỢP NHIỀU PHẦN NHỎ
  // HOẶC DÙNG PHƯƠNG PHÁP NHÂN VỚI MATH.RANDOM() (VỚI ĐỘ CHÍNH XÁC GIỚI HẠN)
  // CÁCH CHUẨN NHẤT CHO SỐ LỚN LÀ DÙNG CRYPTO API:
  const array = new BigUint64Array(1);
  window.crypto.getRandomValues(array);

  // ÉP SỐ NGẪU NHIÊN VÀO KHOẢNG [MIN, MAX]
  return (array[0] % range) + bMin;
}

var _dc = RANDOM_BIGINT(1111111111111,9999999999999);

// LẤY DANH SÁCH LỚP HỌC -> TÊN, MÃ LỚP

fetch(`${vnedu_hosname}/v5/?call=edu.lop_hoc.getLopAndMonChuyen&app_status_lop_hoc=*&nam_hoc=${app_nam_hoc}&_dc=${_dc}&my_user_id=${my_user_id}&app_nam_hoc=${app_nam_hoc}&my_token=${my_token}`, {
  "headers": {
    "accept": "*/*",
    "accept-language": "vi,en-US;q=0.9,en;q=0.8",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"143\", \"Chromium\";v=\"143\", \"Not A(Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"Windows\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "x-requested-with": "XMLHttpRequest"
  },
  "referrer": `${vnedu_hosname}/v5/`,
  "body": "[]",
  "method": "POST",
  "mode": "cors",
  "credentials": "include"
})
.then(response => response.text())
.then(data => {
	var DuLieu = JSON.parse(data);
	for(var i = 0; i < DuLieu.aRoot.length; i++){
	  if(Array_Ten_Lop_Can_Lay_Temp.includes(DuLieu.aRoot[i].ten)){
		Array_ID_Lop_Can_Lay.push(DuLieu.aRoot[i].id);
		Array_Ten_Lop_Can_Lay.push(DuLieu.aRoot[i].ten);
		Array_Khoi_Lop_Can_Lay.push(DuLieu.aRoot[i].khoi);
	  }
	}
	if(Array_ID_Lop_Can_Lay.length > 0 && Array_ID_Lop_Can_Lay.length == Array_Khoi_Lop_Can_Lay.length && Array_ID_Lop_Can_Lay.length == Array_Ten_Lop_Can_Lay.length){
		KIEM_TRA_SO_DAU_BAI(0, Array_ID_Lop_Can_Lay, Array_Khoi_Lop_Can_Lay, Array_Ten_Lop_Can_Lay);
	}
})
.catch(err => console.error('LỖI TẢI DANH SÁCH LỚP: ', err));

// HÀM LẤY SỔ ĐẦU BÀI THEO TUẦN, KHỐI, LỚP

function KIEM_TRA_SO_DAU_BAI(STT, Array_ID_Lop, Array_Khoi_Lop, Array_Ten_Lop){
	if(STT >= Array_ID_Lop.length) { return; }
	var lopHoc = Array_ID_Lop[STT];
	var tenLop = Array_Ten_Lop[STT];
	var khoiHoc = Array_Khoi_Lop[STT];
	console.log(`SỔ ĐẦU BÀI: Tuần ${tuanHoc} - Lớp ${tenLop} - ID Lớp ${lopHoc} -> ĐANG KIỂM TRA`);
	fetch(`${vnedu_hosname}/v5/?load=app.sodaubai.serv.so_dau_bai&app_nam_hoc=${app_nam_hoc}&my_token=${my_token}&tenlop=${tenLop}`, {
	  "headers": {
		"accept": "*/*",
		"accept-language": "vi,en-US;q=0.9,en;q=0.8",
		"content-type": "application/x-www-form-urlencoded; charset=UTF-8",
		"x-requested-with": "XMLHttpRequest"
	  },
	  "referrer": `${vnedu_hosname}/v5/`,
	  "body": `my_token=${my_token}&my_user_id=${my_user_id}&app_nam_hoc=${app_nam_hoc}&winId=5ee7a2b2d953513c0b9ceb729e9381910&capHoc=2&lopHoc=${lopHoc}&khoiHoc=${khoiHoc}&tuanHoc=${tuanHoc}&show_goi_y=1`,
	  "method": "POST",
	  "mode": "cors",
	  "credentials": "include"
	})
	.then(response => response.text())
	.then(data => {
		if(data.indexOf(Key_TimKiem_SoDauBaiChuaHoanThanh) > -1){
			console.log(`%cSỔ ĐẦU BÀI: Tuần ${tuanHoc} - Lớp ${tenLop} - ID Lớp ${lopHoc} -> CÒN THIẾU DỮ LIỆU`,"color:red");
		} else {
			console.log(`%cSỔ ĐẦU BÀI: Tuần ${tuanHoc} - Lớp ${tenLop} - ID Lớp ${lopHoc} -> OK`,"color:blue");
		}
		KIEM_TRA_SO_DAU_BAI(STT + 1, Array_ID_Lop, Array_Khoi_Lop, Array_Ten_Lop);
	})
	.catch(err => console.error('LỖI TẢI SỔ ĐẦU BÀI: ', err));
}
