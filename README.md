﻿# TCP-Chat-Golang
 
 <details>
  <summary><strong>‼️ Antes de utilizar</strong></summary>
    Este chat requer que os utilizadores tenham instalado e ativados em seu pc telnet, na maioria dos sistemas linux ela já vem configurada, nos Windows é    necessario sua ativação,[descubra como ativar](https://blog.betrybe.com/tecnologia/comando-telnet)
  <br />
</details>

## O Projeto
Este projeto consiste em uma chat feito em golang, que utiliza o protocolo telnet para as conexões, para inicia-lo você deve rodar o executavel (disponivel para Linux e Windows) esse executar iniciará o servidor do chat, por padrão ele é iniciado na porta 8888. Depois que o servidor estiver iniciado utilize o telnet para conectar, com o comando `telnet localhost (porta)`

### Comandos
- `/nick (seu nick)` - seta seu nick, por padrão ele vem como anon.
- `/join (nome da sala)` - entra em uma sala, se ela não existir será criada.
- `/rooms` - lista todas as salas disponiveis para entrar.
- `/msg` (sua mensagem) - envia uma mensagem na sala.
- `/quit` - desconecta do servidor.
