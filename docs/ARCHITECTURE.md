# Frontend
## Structure
```
├── components
│   └── [component name]
├── render
├── static
│   ├── css
│   └── js
└── views
    └── [page] 
        ├── handler.go      Define o handler que ira servir o HTML da pagina
        ├── page.templ      Markup da pagina em HTML
        ├── page_templ.go
        └── props.go        Constroi os dados utilizados na pagina
```

