# 🩺 IA-Saúde

API REST para gerenciamento de pacientes e consultas médicas, desenvolvida em **Go (Golang)** com **Gin** e **GORM**.
O projeto foi desenvolvido como estudo em arquitetura de software, boas práticas de CRUD e integração com banco de dados relacional.

---

## 🚀 Tecnologias

- **Go (Golang)**  
- **Gin** → framework HTTP  
- **GORM** → ORM para PostgreSQL  
- **PostgreSQL** → banco de dados relacional  
- **Docker** → ambiente isolado para banco e aplicação

---

## ⚙️ Configuração do projeto

### Pré-requisitos
- Go 1.22+  
- PostgreSQL 14+ (se não usar Docker)  
- Docker + Docker Compose (opcional)

### Clonar o repositório
```bash
git clone https://github.com/Luysdev/ia-saude-2.git
cd ia-saude-2
```

## 📦 Rodando com Docker

### docker-compose.yml
```yaml
version: "3.9"

services:
  db:
    image: postgres:14
    container_name: saude-db
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: 123
      POSTGRES_DB: saude
    ports:
      - "5432:5432"
    volumes:
      - db_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 5s
      retries: 5

  app:
    build: .
    container_name: saude-app
    depends_on:
      db:
        condition: service_healthy
    environment:
      DB_HOST: db
      DB_USER: postgres
      DB_PASSWORD: 123
      DB_NAME: saude
      DB_PORT: 5432
    ports:
      - "8080:8080"
    command: go run cmd/main.go

volumes:
  db_data:
```

### Rodar com Docker
```bash
docker compose up --build
```
API disponível em `http://localhost:8080`

---

## 📌 Rotas da API

| Método | Rota            | Descrição          | Body JSON (exemplo)                                              |
| ------ | --------------- | ------------------ | ---------------------------------------------------------------- |
| POST   | `/paciente/`    | Criar paciente     | `{ "nome": "João da Silva", "idade": 45, "cpf": "12345678900" }` |
| GET    | `/paciente/`    | Listar todos       | -                                                                |
| GET    | `/paciente/:id` | Buscar por ID      | -                                                                |
| PUT    | `/paciente/:id` | Atualizar paciente | `{ "nome": "João da Silva Jr", "idade": 46 }`                    |
| DELETE | `/paciente/:id` | Deletar paciente   | -                                                                |

| Método | Rota            | Descrição          | Body JSON (exemplo)                                                                |
| ------ | --------------- | ------------------ | ---------------------------------------------------------------------------------- |
| POST   | `/consulta/`    | Criar consulta     | `{ "paciente_id": 1, "especialidade": "Cardiologia", "sintomas": "Dor no peito" }` |
| GET    | `/consulta/`    | Listar todas       | -                                                                                  |
| GET    | `/consulta/:id` | Buscar por ID      | -                                                                                  |
| PUT    | `/consulta/:id` | Atualizar consulta | `{ "diagnostico": "Angina", "prescricao": "Medicamento X" }`                       |
| DELETE | `/consulta/:id` | Deletar consulta   | -                                                                                  |

| Método | Rota             | Descrição           | Body JSON (exemplo)                                                                            |
| ------ | ---------------- | ------------------- | ---------------------------------------------------------------------------------------------- |
| POST   | `/historico/`    | Criar histórico     | `{ "paciente_id": 1, "condicao": "Hipertensão", "descricao": "Pressão alta", "status": true }` |
| GET    | `/historico/`    | Listar todos        | -                                                                                              |
| GET    | `/historico/:id` | Buscar por ID       | -                                                                                              |
| PUT    | `/historico/:id` | Atualizar histórico | `{ "descricao": "Pressão controlada" }`                                                        |
| DELETE | `/historico/:id` | Deletar histórico   | -                                                                                              |

| Método | Rota          | Descrição        | Body JSON (exemplo)                                                        |
| ------ | ------------- | ---------------- | -------------------------------------------------------------------------- |
| POST   | `/medico/`    | Criar médico     | `{ "nome": "Dr. Carlos", "crm": "12345", "especialidade": "Cardiologia" }` |
| GET    | `/medico/`    | Listar todos     | -                                                                          |
| GET    | `/medico/:id` | Buscar por ID    | -                                                                          |
| PUT    | `/medico/:id` | Atualizar médico | `{ "telefone": "999999999" }`                                              |
| DELETE | `/medico/:id` | Deletar médico   | -                                                                          |

| Método | Rota         | Descrição       | Body JSON (exemplo)                                                                   |
| ------ | ------------ | --------------- | ------------------------------------------------------------------------------------- |
| POST   | `/exame/`    | Criar exame     | `{ "paciente_id": 1, "tipo": "Sangue", "data": "2025-10-01", "resultado": "Normal" }` |
| GET    | `/exame/`    | Listar todos    | -                                                                                     |
| GET    | `/exame/:id` | Buscar por ID   | -                                                                                     |
| PUT    | `/exame/:id` | Atualizar exame | `{ "resultado": "Levemente alterado" }`                                               |
| DELETE | `/exame/:id` | Deletar exame   | -                                                                                     |

| Método | Rota                    | Descrição                                     | Body JSON (exemplo) |
| ------ | ----------------------- | --------------------------------------------- | ------------------- |
| GET    | `/analyse/:id` | Gera um resumo do paciente usando IA (Gemini) | -                   |


## 📂 Estrutura do projeto
```
.
├── cmd/
│   └── main.go            # inicialização da API
├── internal/
│   ├── db/                # conexão com o banco
│   ├── models/            # structs do domínio (Paciente, Consulta)
│   ├── repository/        # acesso ao banco
│   ├── service/           # regras de negócio
│   ├── handlers/          # controladores HTTP
│   └── routes/            # definição de rotas
└── README.md
```

---

## 📈 Próximos Passos
- Implementar autenticação JWT para médicos/pacientes
- Melhorar validações de dados (CPF, datas de consulta)
- Adicionar testes automatizados mais robustos
- Configurar CI/CD com GitHub Actions

