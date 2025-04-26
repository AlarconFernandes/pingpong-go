## 🏓 Ping Pong em Go – Concorrência e Comunicação com Canais

## Este projeto demonstra o uso de **goroutines** e **canais** em Go para imprimir alternadamente as mensagens `ping` e `pong` no terminal, com um intervalo de 1 segundo entre cada mensagem.

### 📁 Estrutura do Projeto

```
pingpong/
├── pingpong.go
└── README.md
```

---

### 📦 Sobre

O programa utiliza:

- Goroutines concorrentes para gerar as mensagens `"ping"` e `"pong"`
- Um canal de mensagens (`chan string`)
- Um canal de controle (`chan bool`) para garantir alternância
- Uma função que imprime no terminal com atraso programado

---

### 🚀 Como Executar

#### Pré-requisitos:

- Go instalado (versão 1.16 ou superior)

#### Passos:

1. Clone o repositório e entre na pasta do projeto:

```bash
git clone https://github.com/AlarconFernandes/pingpong-go.git
cd seu-repo/pingpong
```

2. Execute o programa:

```bash
go run pingpong.go
```

3. Saída esperada:

```
ping
pong
ping
pong
...
```

> O programa continuará rodando até que você pressione `Enter`.

---

### 📄 Código-fonte

O código completo está no arquivo [`pingpong.go`](pingpong.go).

---

### 📘 Licença

**Este projeto é um exercício educacional e pode ser utilizado livremente.**

---

### 👤 Autor

Desenvolvido por [Alarcon Fernandes](https://github.com/AlarconFernandes)

---
