<?php

    if(isset($_POST['cadastrar'])) {
        $email = $_POST["email"];
        $senha = $_POST["senha"];
        $nome = $_POST["nome"];
        $cpf = $_POST["cpf"];
        $telefone = $_POST["telefone"];
        $chave_pix = $_POST["chave-pix-input"];

        $host = 'localhost';
        $banco = "olxpay";
        $user = "root";
        $senha_user = "";

        $con = mysqli_connect($host, $user, $senha_user, $banco);

        if(!$con) {
            die("Conexão falhou.". mysqli_connect_error());
        }

        $sql = "INSERT INTO usuarios(email, senha, nome, cpf, telefone, chave_pix) VALUES('$email', '$senha', '$nome', '$cpf', '$telefone', '$chave_pix')";

        $rs = mysqli_query($con, $sql);

        if($rs) {
            header('Location: pagamento-taxa.html');
            exit;
        }
    }

?>