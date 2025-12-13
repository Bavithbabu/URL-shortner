# URL-shortner

A simple and lightweight URL shortener built in Go — convert long URLs into short, shareable links.

## ✅ What is this

Many URLs online are long, unwieldy, or hard to share.  
URL-shortner takes any valid URL and generates a much shorter version that redirects to the original link. This makes sharing and remembering links easier — useful for messaging, social media, or embedding links in limited-space contexts. :contentReference[oaicite:1]{index=1}

## ⚙️ Tech stack & Structure

- **Language**: Go :contentReference[oaicite:2]{index=2}  
- **Repository layout**:
  - `api/` — contains the HTTP API endpoints to create and resolve shortened links  
  - `db/` — contains database logic (e.g. storing original → short URL mappings)  
  - `Dockerfile` — (optional) includes instructions to containerize the application :contentReference[oaicite:3]{index=3}

## 🚀 How to install & run locally

1. Clone the repository:
   ```bash
   git clone https://github.com/Bavithbabu/URL-shortner.git
   cd URL-shortner
