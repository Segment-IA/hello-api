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

# teste de code review


### criando nova feature ###

git checkout develop
git pull origin develop
git checkout -b feature/nova-pasta

# (adicione arquivos)
git add .
git commit -m "feat: adiciona nova pasta"

git push -u origin feature/nova-pasta
# → Crie PR no GitHub: feature/nova-pasta → develop

# Após merge
git checkout develop
git pull
git branch -d feature/nova-pasta
git push origin --delete feature/nova-pasta

✅ Objetivo: fazer merge das mudanças de develop → main
🧱 Etapas completas:
1. ⚙️ Certifique-se de que develop está atualizado
git checkout develop
git pull origin develop

2. 🌿 Crie a branch de release
Substitua 1.4.0 pelo número da versão que deseja lançar:
git checkout -b release/1.4.0

3. 🧪 Faça ajustes finais (se necessário)
Exemplo:
Atualize versão no package.json, pyproject.toml ou similar
Atualize o CHANGELOG.md
Depois, commit:
git add .
git commit -m "chore(release): finaliza release 1.4.0"

4. 🚀 Suba a branch para o GitHub
git push -u origin release/1.4.0

5. 🔁 Crie os Pull Requests no GitHub:
PR: release/1.4.0 → main ✅
(esse é o que efetivamente publica em produção)
PR: release/1.4.0 → develop ⚠️
(caso tenha commits específicos que não estavam na develop)

6. 🏷️ Após merge na main, crie a tag da versão
git checkout main
git pull origin main

git tag -a v1.4.0 -m "Release 1.4.0"
git push origin v1.4.0



### criando um hotfix ###

1. Vá para a branch main e atualize
git checkout main
git pull origin main

2. Crie a branch do hotfix
git checkout -b hotfix/1.3.1

3. Faça as alterações (exemplo)
nano src/api/rota_critica.js
# ou use seu editor preferido
4. Commit e push
git add .
git commit -m "fix: corrige bug em rota crítica"
git push -u origin hotfix/1.3.1

5. 🔁 Crie os dois Pull Requests no GitHub:
✅ PR 1 → produção:
hotfix/1.3.1 → main (obrigatório)
Título: hotfix: 1.3.1
Após aprovação: merge
✅ PR 2 → develop:
hotfix/1.3.1 → develop (obrigatório)
Para que a correção vá para o ciclo de release futuro

6. 🏷️ Após o merge em main, crie a tag da versão
git checkout main
git pull origin main

git tag -a v1.3.1 -m "Hotfix 1.3.1"
git push origin v1.3.1

7. (Opcional) Limpeza
git branch -d hotfix/1.3.1
git push origin --delete hotfix/1.3.1
