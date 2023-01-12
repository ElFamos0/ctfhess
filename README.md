# Challenge Inté

Site web permettant de déployer et d'administrer les challenges inté pour la rentrée 2022 à Telecom-Nancy

Lien en ligne : ctfhess.com


## Setup Frontend

```bash
cd frontend
npm install package-lock.json
npm run serve
```

## Setup Backend

```bash
cd backend
go mod download
go run main.go
```

## Setup Docker

```bash
sudo docker-compose build
sudo docker-compose up
```
