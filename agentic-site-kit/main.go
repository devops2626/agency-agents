package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        query := r.URL.Query().Get("q")
        if query == "" {
            query = "explore"
        }

        // INTENT CLASSIFICATION (Simulated)
        var title, hero, product string
        lower := strings.ToLower(query)

        if strings.Contains(lower, "buy") || strings.Contains(lower, "price") || strings.Contains(lower, "purchase") {
            title = "🛒 Transaction Mode"
            hero = "High-Intent Buyer Detected"
            product = "Showing premium stock with dynamic pricing."
        } else if strings.Contains(lower, "research") || strings.Contains(lower, "compare") || strings.Contains(lower, "specs") {
            title = "🔬 Research Mode"
            hero = "Analytical Deep-Dive"
            product = "Side-by-side specs and white papers."
        } else if strings.Contains(lower, "ai") || strings.Contains(lower, "conference") {
            title = "🧠 Agentic Explorer"
            hero = "You found the AI & Europe conferences"
            product = "Zurich June 2026 | Multi-Agent Systems Track"
        } else {
            title = "🧭 Discovery Mode"
            hero = "Welcome to the Agentic Site"
            product = "Browse our catalog of camping gear and espresso machines."
        }

        // DYNAMIC PAGE COMPOSITION (The "Agentic" UI)
        tmpl := template.Must(template.New("page").Parse(`
        <html>
        <head><title>Agentic Site</title>
        <style>
            body { font-family: system-ui, sans-serif; background: #f0f4f8; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; }
            .card { max-width: 700px; background: white; padding: 40px; border-radius: 24px; box-shadow: 0 20px 60px rgba(0,0,0,0.1); text-align: center; border-top: 6px solid #3b82f6; }
            h1 { margin: 0 0 8px 0; color: #1e293b; }
            .badge { background: #dbeafe; color: #1e40af; padding: 6px 16px; border-radius: 40px; font-size: 14px; display: inline-block; margin-bottom: 20px; }
            .hero { font-size: 24px; font-weight: 600; color: #0f172a; margin: 16px 0; }
            .product { background: #f8fafc; padding: 20px; border-radius: 12px; color: #334155; margin: 20px 0; border: 1px solid #e2e8f0; }
            .query-box { background: #f1f5f9; padding: 12px; border-radius: 40px; font-family: monospace; color: #475569; margin-top: 24px; }
            a { color: #3b82f6; text-decoration: none; }
        </style>
        </head>
        <body>
        <div class="card">
            <h1>{{.Title}}</h1>
            <div class="badge">Audience of One</div>
            <div class="hero">{{.Hero}}</div>
            <div class="product">{{.Product}}</div>
            <div class="query-box">🔍 Intent: "{{.Query}}" | ⚡ Generated in real-time</div>
            <p style="margin-top: 24px; font-size: 14px; color: #94a3b8;">
                <a href="/?q=buy%20coffee">Buy</a> · 
                <a href="/?q=research%20ai">Research</a> · 
                <a href="/?q=explore">Explore</a> · 
                <a href="/?q=europe%20conference">AI Conference</a>
            </p>
        </div>
        </body>
        </html>
        `))
        tmpl.Execute(w, map[string]string{
            "Title":   title,
            "Hero":    hero,
            "Product": product,
            "Query":   query,
        })
    })

    fmt.Println("🚀 Agentic Site running on http://0.0.0.0:8080")
    http.ListenAndServe(":8080", nil)
}
