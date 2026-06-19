# AI Marketplace (Vue + Go)

Projet **full-stack** :
- **Frontend** : Vue 3 (Vite) + Pinia
- **Backend** : Go (Gin) expose une API REST
- Les “modèles” sont stockés dans un fichier local : **`backend/data/models.json`**

---

## Fonctionnement rapide
1. Au démarrage, le frontend appelle : `GET http://localhost:8080/models`
2. L’UI affiche la liste des modèles (recherche, catégories, section “Trending”).
3. Quand vous publiez / supprimez / téléchargez / ratez un modèle, le frontend appelle les endpoints du backend.
4. Le backend lit/écrit dans `backend/data/models.json` à chaque action.

---

## Prérequis
- **Go** : version compatible avec le `go.mod` (ici Go **1.26.3**)
- **Node.js** : recommandé **Node 20+** (le `package.json` précise `^20.19.0 || >=22.12.0`)

---

## Lancer le projet
Le plus simple : lancer **backend et frontend dans 2 terminaux**.

### 1) Backend (Go)
Dans le terminal :
```bash
cd backend
go run main.go
```
- Le serveur démarre sur : **http://localhost:8080**

### 2) Frontend (Vue/Vite)
Dans un second terminal :
```bash
cd frontend
npm install
npm run dev
```
- Vite démarre en général sur **http://localhost:5173** (ou un autre port si déjà utilisé)

---

## API Backend (Gin)
Le backend expose les endpoints suivants (base : `/models`) :

### 1) Lister les modèles
- **GET** `/models`
- Retour : tableau JSON de modèles

### 2) Publier un modèle
- **POST** `/models`
- Corps JSON :
```json
{
  "name": "...",
  "description": "...",
  "category": "...",
  "accuracy": 90,
  "author": "..."
}
```
- Le backend génère : `id`, `downloads` (0) et `rating` (5.0)

### 3) Supprimer un modèle
- **DELETE** `/models/:id`

### 4) Simuler un téléchargement
- **PATCH** `/models/:id/download`
- Effet : `downloads += 1`

### 5) Noter un modèle
- **PATCH** `/models/:id/rate`
- Effet : `rating += 0.1`

---

## Données
- Fichier de persistance : **`backend/data/models.json`**
- Chaque action (POST/DELETE/PATCH) recharge puis réécrit ce fichier.

---

## Notes
- CORS est autorisé avec `Access-Control-Allow-Origin: *`, donc le frontend peut appeler le backend en local.


