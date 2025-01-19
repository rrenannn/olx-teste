function delayedRedirect() {
    setTimeout(function() {
        window.location.href = "inicial.html";
    }, 200);
}

document.getElementById("enviar-comprovante").addEventListener("click", function(event) {
    event.preventDefault();
    const inputField = document.getElementById("comprovante");
    
    if (inputField.value === "") {
        alert("Por favor, selecione um arquivo.");
        return false;
    } else {
        if (confirm("Deseja prosseguir com o redirecionamento?")) {
            delayedRedirect();
        } else {
            return false;
        }
    }
});