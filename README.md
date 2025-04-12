# 🖼️ Placeholder API

Une API ultra simple et rapide pour générer dynamiquement des images **placeholder** en Go

---

## 🚀 Exemple d'appel

```http request
GET http://localhost:8080/128x64?label=Hello World !&font_size=18&font_color=a255c9&font_style=Comic&background_color=d0d0d0
```

![Exemple de résultat](assets/exemple-result.png)


---

## 📦 Paramètres disponibles

| Paramètre          | Type   | Description                                | Exemple        |
|--------------------|--------|--------------------------------------------|----------------|
| `label`            | string | Texte affiché au centre de l’image         | `Place-holder` |
| `font_size`        | int    | Taille du texte en pixels                  | `18`           |
| `font_color`       | hex    | Couleur du texte en hex sans `#`           | `76ecf5`       |
| `font_style`       | string | Le style de police utilisé                 | `Inter`        |
| `background_color` | hex    | Couleur de fond de l’image en hex sans `#` | `f576f1`       |

---

## 💾 Fonts personnalisées

Par défaut, l’API contient deux polices (Inter et ComicNeuve)
pour ajouter vos propres fonts glissez-les dans le dossier `generator/fonts`
elle sera alors disponible sous le paramètre de requête suivant `?font_style=MyNewFont`
si aucune font n'est trouvé sous ce nom alors Inter sera prise par défaut 
> [!WARNING]
> Le nom du fichier sans l'extention sera utilisé comme valeur
---

## 👷‍♂️ Installation

1) Assurez-vous d'avoir go en version *1.24.2* installé 
2) Clonez le repo [Placeholder-api](https://github.com/LyneQ/Placeholder-API)
```shell
    git clone https://github.com/LyneQ/Placeholder-API
```
3) Télécharger des dependence via
```shell
  go mod download`
```
4) Construire l'app via la commande 
```shell
  go build
```

---

## 📰 À propos

Ce projet utilise des polices de [Google Fonts](https://fonts.google.com/), mises à disposition sous licence open-source.

- **[Inter](https://fonts.google.com/specimen/Inter)**
- **[Comic Neue](https://fonts.google.com/specimen/Comic+Neue)**

---

### Développé par [LyneQ](https://github.com/LyneQ)

