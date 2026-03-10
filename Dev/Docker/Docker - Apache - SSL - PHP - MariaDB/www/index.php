<?php

$mysqli = new mysqli("db", "my_user", "my_pwd", "my_db");

if ($mysqli->connect_error) {
    die("CONNECTION FAILED: " . $mysqli->connect_error);
}

echo "CONNECTED TO MARIADB SUCCESSFULLY!";

echo "<br/>";

phpinfo();

?>