<?php
// error_reporting(0);
$FLAG = $_ENV['FLAG'];
$DB_HOST = $_ENV['DB_HOST'];
$DB_USER = $_ENV['DB_USERNAME'];
$DB_PASS = $_ENV['DB_PASSWORD'];
if (isset($_POST['login'])) {
    if ($_POST['db_host'] === $DB_HOST && $_POST['db_user'] === $DB_USER && $_POST['db_pass'] === $DB_PASS) {
        $is_login = true;
    } else {
        $is_login = false;
    }
}
?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Welcome to phpmyadmin handmade version</title>
</head>
<body>
    <h1>Welcome to phpmyadmin handmade version</h1>
    
    <?php if (isset($is_login)): ?>
        <?php if ($is_login === true): ?>
            <h2>Login ! => <?php echo $FLAG; ?></h2>
        <?php else: ?>
            <h2>Can't login</h2>
        <?php endif ?>
    <?php endif ?>

    <form method="post">
        Database Host: <input type="text" name="db_host"/><br>
        Database User: <input type="text" name="db_user"/><br>
        Database Pass: <input type="text" name="db_pass"/><br>
        <button type="submit" name='login'>Login</button>
    </form>
</body>
</html>