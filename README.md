# Hello API – Segment IA

Uma API em Go simples que retorna `"Hello, World"` na rota raiz `/`.

## 🛠️ Como executar

```bash
go run main.go
```

Acesse [http://localhost:8080](http://localhost:8080)

## ✅ Testes

```bash
go test ./...
```

## 📦 Estrutura

```
.
├── main.go               # Início do servidor
├── handler/hello.go      # Handler principal da API
├── handler/hello_test.go # Teste unitário
```

## 🔄 GitHub Actions

Este projeto já possui CI configurado para rodar os testes em cada `push` e `pull request`.

#teste de code review

