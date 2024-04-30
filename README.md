<p align="center">
 <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
 <img src="https://img.shields.io/badge/HTMX-005A9C?style=for-the-badge&logo=htmx&logoColor=white" alt="HTMX" />
 <img src="https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white" alt="Tailwind CSS" />
</p>

# ShrinkIt

ShrinkIt is a simple file compression platform written in Go

The UI is written in 🐴 HTMX with [a-h/templ](https://github.com/a-h/templ) and Tailwind CSS

All the compression and shrinking is performed using the Go standard library alongside some very handy and amazing CLI Tools (for now)

### Shrinks:

- 🖼️ **Images**:

  PNG, JPG

- 📄 **Documents**:

  PDF, DOCX

- 🎶 **Audios** (Pending) :

  MP3, WAV

- 🎞️ **Videos** (Pending) :

  MP4, MKV

### Quick Peek:

![swappy-20240430_165715](https://github.com/sumitst05/shrink-it/assets/106669732/5919fcff-a59e-425c-ac2a-bc38c9213422)

![swappy-20240430_165744](https://github.com/sumitst05/shrink-it/assets/106669732/f2b5fb5f-c9e6-4828-9e2b-5d926dd86132)

### Installation:

#### Using Docker:

1. Pull the image with:

```bash
docker pull sumitst05/shrink-it:latest
```

2. Run:

```bash
docker run -p 3000:3000 sumitst05/shrink-it
```

This starts the application and makes it accessible on port 3000.

#### Building from Source:

1. Clone the repository:

```bash
git clone https://github.com/your-repo/shrink-it.git
```

2. Navigate to the project's root:

```bash
cd shrink-it
```

3. Build and run the application:

```bash
make dev
```

### Deployed on:

<a href="https://render.com" style="display: flex; align-items: center; text-decoration: none;">
<img src="https://media.licdn.com/dms/image/D4E0BAQGGDoFoqHtOvA/company-logo_200_200/0/1702595267620/renderco_logo?e=2147483647&v=beta&t=Ywm0UZpTXbiXPopyfCDty8QXSEVz88QWWCwy28qLUyE" width="60" height="60" style="margin-right: 10px;" />

<h2>Render</h2>
</a>
