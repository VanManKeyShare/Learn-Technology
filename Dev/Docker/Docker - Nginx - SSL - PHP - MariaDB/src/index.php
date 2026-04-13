<?php

$mysqli = new mysqli("mariadb", "mysql_admin", "mysql_admin_pwd_123", "mysql_db");

if ($mysqli->connect_error) {
    die("CONNECTION FAILED: " . $mysqli->connect_error);
}

echo "NGINX + SSL + PHP + MARIADB";

echo "<br/><br/>";

echo "CONNECTED TO MARIADB SUCCESSFULLY!";

echo "<br/>";

phpinfo();

?>