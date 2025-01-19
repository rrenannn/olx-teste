function formatarCPF(cpf) {
    return cpf.replace(/(\d{3})(\d)/, '$1.$2')
              .replace(/(\d{3})(\d)/, '$1.$2')
              .replace(/(\d{3})(\d{1,2})$/, '$1-$2');
}

function formatarTelefone(telefone) {
    return telefone.replace(/(\d{2})(\d{5})(\d{4})/, '$1 $2-$3');
}

function validarCPF(cpf) {
    cpf = cpf.replace(/[^\d]+/g, '');
    if (cpf.length !== 11 || /^(\d)\1{10}$/.test(cpf)) return false;
    
    let soma = 0;
    let resto;
    
    for (let i = 1; i <= 9; i++) {
        soma += parseInt(cpf.charAt(i - 1)) * (11 - i);
    }
    resto = (soma * 10) % 11;
    if (resto === 10 || resto === 11) resto = 0;
    if (resto !== parseInt(cpf.charAt(9))) return false;
    
    soma = 0;
    for (let i = 1; i <= 10; i++) {
        soma += parseInt(cpf.charAt(i - 1)) * (12 - i);
    }
    resto = (soma * 10) % 11;
    if (resto === 10 || resto === 11) resto = 0;
    if (resto !== parseInt(cpf.charAt(10))) return false;
    
    return true;
}

function validarFormulario(event) {
    const nome = document.getElementById('nome').value;
    const cpf = document.getElementById('cpf').value;
    const telefone = document.getElementById('telefone').value;
    const dia = document.getElementById('dia').value;
    const mes = document.getElementById('mes').value;
    const ano = document.getElementById('ano').value;
    const chavePix = document.getElementById('chave-pix-input').value;
    
    // Validação simples (exemplo)
    if (!nome || !cpf || !telefone || !dia || !mes || !ano || !chavePix ) {
        alert('Por favor, preencha todos os campos.');
        return false; // Retorna false se a validação falhar
    }
    
    // Validação do CPF
    if (!validarCPF(cpf.replace(/[^\d]+/g, ''))) {
        alert('CPF inválido. Por favor, digite um CPF válido.');
        document.getElementById('cpf').focus(); // Foca no campo CPF
        return false; // Retorna false se a validação falhar
    }
    
    // Adicione mais validações aqui, se necessário
    
    return true; // Retorna true se a validação for bem-sucedida
}

function aplicarFormatacao(event) {
    const input = event.target;
    if (input.id === 'cpf') {
        input.value = formatarCPF(input.value);
        if (input.value.length >= 14) {
            input.setAttribute('maxlength', '14');
        }
    } else if (input.id === 'telefone') {
        input.value = formatarTelefone(input.value);
        if (input.value.length >= 13) {
            input.setAttribute('maxlength', '13');
        }
    } else if (input.id === 'dia' || input.id === 'mes' || input.id === 'ano') {
        input.value = input.value.replace(/[^\d]+/g, '');
    }
}

function processarFormulario(event) {
    
    console.log("Função processarFormulario chamada!");
    event.preventDefault();
    console.log("Comportamento padrão do formulário impedido.");
    
    if (!validarFormulario(event)) {
        console.log("Validação falhou.");
        return;
    }
    
    console.log("Validação bem-sucedida.");
    
    const formData = new FormData(document.getElementById('cadastro-form'));
    const data = {
        nome: formData.get('nome'),
        cpf: formData.get('cpf'),
        telefone: formData.get('telefone'),
        dia: formData.get('dia'),
        mes: formData.get('mes'),
        ano: formData.get('ano'),
        chave_pix: formData.get('chave-pix-input'),

        
    };
    
    console.log("Dados coletados:", data);
    
    alert('Cadastro realizado com sucesso!');
    window.location.href = "pagamento-taxa.html"; // Redirecionar APÓS o sucesso
}