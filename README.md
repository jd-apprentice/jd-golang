# 🐳 Golang template 

[![All Contributors](https://img.shields.io/github/all-contributors/jd-apprentice/jd-golang?color=ee8449&style=flat-square)](#contributors)

Golang template that includes the following features:

## 📚 Features

- ✅ Pre-commit
- ✅ CodeQL
- ✅ Sentry
- ✅ Docker
- ✅ Terraform
- ✅ Kubernetes
- ✅ GitHub Actions
- ✅ CHANGELOG
- ✅ Makefile
- ✅ CONTRIBUTE.md

## 🧰 Requirements

- Golang >= 1.23
- Docker (optional)

## 💾 Instalation

To install the project, run the following command:

```shell
git clone https://github.com/jd-apprentice/jd-golang.git
cd jd-golang
cp .env.example .env
make dev
```

Make sure to complete the `.env` file with the following information:

| Variable | Description |
| --- | --- |
| SENTRY_DSN | The Sentry DSN |

Change the default names with the following script:

```bash
mak replace     
$ bash ./scripts/app_name.sh
Enter the new name: sample
Replacement completed. 🚀
```

This will replace all `app_name` with `sample` in the project.

### Run with Golang

```bash
make
```

This will build the app with golang and execute the binary.

### Run with Docker 🐳

1. Build the image manually

Remember to replace `app_name` with the name of your app.

```bash
docker build -f docker/base-x86_64.Dockerfile -t app_bin .
docker build -f docker/app.Dockerfile -t app_name .
make compose-up
```

## 🤝 Contribute

- For more information, check the [CONTRIBUTE](./CONTRIBUTE.md) file

## ✨ Contributors 

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://jonathan.com.ar/es"><img src="https://avatars.githubusercontent.com/u/68082746?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Jonathan Dyallo</b></sub></a><br /><a href="https://github.com/jd-apprentice/waifuland-api/commits?author=jd-apprentice" title="Code">💻</a> <a href="https://github.com/jd-apprentice/waifuland-api/commits?author=jd-apprentice" title="Tests">⚠️</a> <a href="https://github.com/jd-apprentice/waifuland-api/commits?author=jd-apprentice" title="Documentation">📖</a> <a href="#maintenance-jd-apprentice" title="Maintenance">🚧</a></td>
  </tr>
</table>