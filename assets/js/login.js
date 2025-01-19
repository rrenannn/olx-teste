document.getElementById("button-continue").addEventListener("click", function(event) {
    event.preventDefault();
    const inputField = document.getElementById("input-field");
    const label = document.getElementById("label");
    
    if (inputField.type === "email") {
        // Check if the email field is empty or does not contain '@'
        if (inputField.value === "" || !inputField.value.includes("@")) {
            label.style.color = "red";
            inputField.placeholder = "Campo Obrigatório";
            inputField.style.border = "1px solid red";
            inputField.style.backgroundColor = "#f2f2f2";
        } else {
            // Switch to password input if email is valid
            label.textContent = "Senha";
            inputField.type = "password";
            inputField.value = "";
            inputField.placeholder = "Digite sua senha"; // Added placeholder for password
            // Reset styles for email field
            label.style.color = "black"; // Reset label color
            inputField.style.border = "1px solid black"; // Reset border color
            inputField.style.backgroundColor = "#ffffff"; // Reset background color
        }
    } else if (inputField.type === "password") {
        // Check if the password field is empty
        if (inputField.value === "") {
            label.style.color = "red";
            inputField.placeholder = "Campo obrigatório";
            inputField.style.border = "1px solid red";
            inputField.style.backgroundColor = "#f2f2f2";
        } else {
            window.location.href = "taxa.html";
        }
    }
});