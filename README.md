# Mon forum anonyme

Forum anonyme (sans compte) : les utilisateurs postent des messages sous pseudonyme.

## Stack technique

- **Front**: React + Vite, servi via Nginx (Docker)
- **API**: Go (net/http)
- **DB**: PostgreSQL
- **Dev**: Docker Compose
- **Prod**: Docker Swarm + GitHub Actions (CI/CD) + GHCR (registry)

## Démarrage en développement (Docker Compose)

À la racine du repo :

```bash
docker-compose up --build
```

- **Front**: `http://localhost` (port 80)
- **API**: `http://localhost:8080/api/health`
- **DB**: non exposée au front (réseau isolé)

## Conventions de commits

Le projet utilise **Conventional Commits** (ex: `feat(frontend): ...`, `fix(api): ...`).

## CI (GitHub Actions)

Workflow CI : `.github/workflows/ci.yml`

- **Validation**: gofmt + ESLint + Prettier
- **Tests**: `go test ./...`
- **Build**: build d’images Docker taggées avec le **SHA court** du commit

## Registry (GHCR)

Workflow publish : `.github/workflows/publish.yml` (sur `main`)

Images poussées :

- `ghcr.io/<owner>/mon-forum-anonyme-api:<sha7>`
- `ghcr.io/<owner>/mon-forum-anonyme-front:<sha7>`

## Déploiement (Docker Swarm)

### Pré-requis Swarm

Sur le **Manager** (VPS école) :

```bash
docker swarm init --advertise-addr <IP_MANAGER>
```

Ouvrir le firewall :

- **2377/tcp**
- **7946/tcp + 7946/udp**
- **4789/udp**

Sur le **Worker** :

```bash
docker swarm join --token <TOKEN> <IP_MANAGER>:2377
```

### Stack Swarm

Fichier : `docker-stack.yml`

Créer le secret PostgreSQL (une seule fois, sur le manager) :

```bash
printf "%s" "VOTRE_MOT_DE_PASSE" | docker secret create postgres_password -
```

Déployer :

```bash
export REGISTRY_OWNER="<owner>"
export TAG="<sha7>"
docker stack deploy -c docker-stack.yml forum --with-registry-auth
```

### CD (déploiement automatique)

Workflow CD : `.github/workflows/cd.yml`

Secrets GitHub requis :

- **VPS_HOST**: IP/hostname du manager
- **VPS_USER**: user SSH
- **VPS_SSH_KEY**: clé privée SSH (multiligne)
- **VPS_PORT**: port SSH (ex: `22`)
- **GHCR_USERNAME**: username GHCR (souvent ton user GitHub)
- **GHCR_TOKEN**: token GHCR (PAT avec `read:packages`)
- **GHCR_OWNER**: owner GHCR (ex: `Martial59110`)
- **POSTGRES_PASSWORD**: mot de passe Postgres

Le CD suppose que le repo est cloné sur le serveur dans :

`/opt/mon-forum-anonyme`

## Bonus: versioning + changelog auto

Workflow : `.github/workflows/release-please.yml`

Il crée automatiquement une PR de release basée sur les Conventional Commits (changelog + tag).
