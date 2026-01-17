
// BOOKMARKLET -> ẨN NHỮNG DÒNG KHÔNG CÓ NÚT THÊM TIẾT SỔ ĐẦU BÀI

$('table.table tbody tr.none').each(function(index, row) {
	var $row = $(row);
	var col_hanhdong = $row.find('td[chitiet_id]');
	var add_sodaubai = col_hanhdong.find('a.add_tiet_so_dau_bai');
	if(add_sodaubai.length == 0){
		$row.css("display","none");
	}
});

// BOOKMARKLET -> NHẬP THÔNG TIN NHẬN XÉT VÀ NHẤN NÚT LƯU

$("textarea[name='nhan_xet']").val('Lớp nghiêm túc.');
$("input[name='diem']").val('6');
$("input[name='diem_kl']").val('2');
$("input[name='diem_vs']").val('2');
var IDLABEL_DBHSVTUSODIEMDANH = $("label:contains('Đồng bộ HSV từ Sổ điểm danh:')").attr("for");
var DBHSVTUSODIEMDANH = $("input#" + IDLABEL_DBHSVTUSODIEMDANH);
if(DBHSVTUSODIEMDANH.attr("aria-checked") == "false"){DBHSVTUSODIEMDANH.trigger("click")}
$("span.x-btn-inner:contains('Lưu')").parent().trigger("click");
setTimeout(function(){
	$("span.x-btn-inner:contains('OK')").parent().trigger("click");
},1000);
