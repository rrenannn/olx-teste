function copyToClipboard() {
    const inputElement = document.getElementById('pix-code');
    inputElement.select();
    document.execCommand('copy');
    alert('Código PIX copiado para a área de transferência!');
}

function delayedRedirect() {
    setTimeout(function() {
        window.location.href = "comprovante.html";
    }, 200);
}



window.onload = function() {
    generateQRCode();
};