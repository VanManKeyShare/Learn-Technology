<?php

echo "APACHE - PHP - REDIS - MARIADB";
echo "<br/><br/>";

echo "+ KIỂM TRA KẾT NỐI ĐẾN MARIADB";
echo "<br/><br/>";

$db = getenv('DB_HOST');
$db_name = getenv('DB_NAME');
$db_user = getenv('DB_USER');
$db_password = getenv('DB_PASS');

$mysqli = new mysqli($db, $db_user, $db_password, $db_name);

if ($mysqli->connect_error) {
    die("KHÔNG THỂ KẾT NỐI ĐẾN MARIADB: " . $mysqli->connect_error);
}

echo "KẾT NỐI ĐẾN MARIADB THÀNH CÔNG";
echo "<br/><br/>";

echo "+ KIỂM TRA KẾT NỐI ĐẾN REDIS";
echo "<br/><br/>";

// KIỂM TRA XEM EXTENSION REDIS ĐÃ ĐƯỢC CÀI CHƯA
if (!class_exists('Redis')) {
    die("PHP REDIS EXTENSION CHƯA ĐƯỢC CÀI ĐẶT");
}

$REDIS_HOST = getenv('REDIS_HOST');
$REDIS_PORT = getenv('REDIS_PORT');
$REDIS_TIMEOUT = getenv('REDIS_TIMEOUT');
$REDIS_PASSWORD = getenv('REDIS_PASSWORD');

$redis = new Redis();

try {
    // KẾT NỐI TỚI REDIS SERVER (HOST, PORT, TIMEOUT)
    $redis->connect($REDIS_HOST, $REDIS_PORT, $REDIS_TIMEOUT);
    $redis->auth($REDIS_PASSWORD); // BỎ QUA NẾU KHÔNG CÓ MẬT KHẨU

    // KIỂM TRA PHẢN HỒI TỪ REDIS BẰNG CÁCH GỬI LỆNH PING
    $response = $redis->ping();

    if ($response == "PONG" || $response === true) {
        echo "KẾT NỐI ĐẾN REDIS THÀNH CÔNG";
        echo "<br/><br/>";

        $redis->set('hello_key', 'Hello From Redis');
        $value = $redis->get('hello_key');

        echo "KIỂM TRA GHI VÀ ĐỌC GIÁ TRỊ TỪ REDIS: " . $value;
    }
} catch (Exception $e) {
    echo "KHÔNG THỂ KẾT NỐI ĐẾN REDIS: " . $e->getMessage();
}

echo "<br/><br/>";

phpinfo();

?>